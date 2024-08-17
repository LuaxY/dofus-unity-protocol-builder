package main

import (
	"encoding/binary"
	"fmt"
	"github.com/fatih/color"
	"go-xp-dofus-unity-proto-builder/src/protocol"
	connectionMessage "go-xp-dofus-unity-proto-builder/src/protocol/connection/message"
	gameMessage "go-xp-dofus-unity-proto-builder/src/protocol/game/message"
	"google.golang.org/protobuf/reflect/protoreflect"
	"math/rand/v2"
	"net"
	"reflect"
	"strconv"
	"sync"
)

type Direction int

const (
	ClientToServer Direction = iota
	ServerToClient
)

var serverProxyMap = make(map[string]*Proxy)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	connectionProxy := NewProxy(protocol.ConnectionServer, 5555, "dofus2-co-beta.ankama-games.com:5555", &connectionMessage.Message{})
	connectionProxy.ClientMessageFunc = processConnectionClientMessage
	connectionProxy.ServerMessageFunc = processConnectionServerMessage

	go func() {
		defer wg.Done()
		err := connectionProxy.Start()
		if err != nil {
			panic(err)
		}
		connectionProxy.Close()
	}()

	wg.Wait()
}

type ProcessMessageFunc func(client *Client, message protoreflect.ProtoMessage) ([]protoreflect.ProtoMessage, error)

type Client struct {
	clientConn             net.Conn
	serverConn             net.Conn
	bufferConnectionClient []byte
	bufferConnectionServer []byte
}

type Proxy struct {
	serverType  protocol.ServerType
	port        int32
	target      string
	listener    net.Listener
	messageType protoreflect.ProtoMessage

	ServerMessageFunc ProcessMessageFunc
	ClientMessageFunc ProcessMessageFunc
}

func NewProxy(serverType protocol.ServerType, port int32, target string, messageType protoreflect.ProtoMessage) *Proxy {
	p := &Proxy{
		serverType:  serverType,
		port:        port,
		target:      target,
		messageType: messageType,
	}

	p.ClientMessageFunc = p.passThrough
	p.ServerMessageFunc = p.passThrough

	return p
}

func (p *Proxy) Start() error {
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(int(p.port)))
	if err != nil {
		return err
	}

	color.Yellow("[%s] listening on port %d targeting %s", p.serverType, p.port, p.target)

	p.listener = listener

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}

		color.Green("[%s] new connection from %s", p.serverType, conn.RemoteAddr().String())

		go p.handleConnection(conn)
	}
}

func (p *Proxy) Close() error {
	return p.listener.Close()
}

func (p *Proxy) passThrough(client *Client, message protoreflect.ProtoMessage) ([]protoreflect.ProtoMessage, error) {
	return []protoreflect.ProtoMessage{message}, nil
}

func (p *Proxy) handleConnection(clientConn net.Conn) {
	// Connect to the server
	serverConn, err := net.Dial("tcp", p.target)
	if err != nil {
		color.Red("[%s] error connecting to %s: %v", p.serverType, p.target, err)
		clientConn.Close()
		return
	}

	defer clientConn.Close()
	defer serverConn.Close()

	client := &Client{
		clientConn: clientConn,
		serverConn: serverConn,

		bufferConnectionClient: make([]byte, 0),
		bufferConnectionServer: make([]byte, 0),
	}

	handleConn := func(client *Client, direction Direction) {
		var partialBuffer []byte

		var from net.Conn
		var to net.Conn
		var processMessageFunc ProcessMessageFunc
		var directionStr string
		var colorFunc func(format string, a ...interface{})

		switch direction {
		case ClientToServer:
			from = client.clientConn
			to = client.serverConn
			processMessageFunc = p.ClientMessageFunc
			directionStr = "C->S"
			colorFunc = color.Cyan
		case ServerToClient:
			from = client.serverConn
			to = client.clientConn
			processMessageFunc = p.ServerMessageFunc
			directionStr = "S->C"
			colorFunc = color.Blue
		}

		for {
			var readBuffer [2048]byte

			n, err := from.Read(readBuffer[:])
			if err != nil {
				fmt.Println("Error reading", err)
				client.clientConn.Close()
				return
			}

			partialBuffer = append(partialBuffer, readBuffer[:n]...)

			//fmt.Println("Received", n, "bytes")
			//fmt.Println(hex.EncodeToString(readBuffer[:n]))

			for {
				size, sizeLength := binary.Uvarint(partialBuffer[:])

				//fmt.Println("Size", size)
				//fmt.Println("SizeLength", sizeLength)

				// Incomplete message
				if size == 0 || int(size) > len(partialBuffer)-sizeLength {
					break
				}

				payload := partialBuffer[sizeLength : sizeLength+int(size)]
				//fmt.Println(sizeLength, size, hex.EncodeToString(payload[:]))

				// instantiate the message type using reflection
				message := reflect.New(reflect.TypeOf(p.messageType).Elem()).Interface().(protoreflect.ProtoMessage)
				if err := protocol.DecodeMessage(payload, message); err != nil {
					color.Red("[%s] [%s] Error decoding message: %v", p.serverType, directionStr, err)
				}

				messages, err := processMessageFunc(client, message)
				if err != nil {
					color.Red("[%s] [%s] Error processing message: %v", p.serverType, directionStr, err)
				}

				for _, message := range messages {
					colorFunc("[%s] [%s] %s\n", p.serverType, directionStr, protocol.PrettyPrintMessage(message))

					encodedMessage, err := protocol.EncodeMessage(message)
					if err != nil {
						color.Red("[%s] [%s] Error encoding message: %v", p.serverType, directionStr, err)
					}

					_, err = to.Write(encodedMessage)
					if err != nil {
						color.Red("[%s] [%s] Error writing to %s: %v", p.serverType, directionStr, to.RemoteAddr().String(), err)
					}
				}

				partialBuffer = partialBuffer[sizeLength+int(size):]
			}
		}
	}

	wg := sync.WaitGroup{}
	wg.Add(2)

	go handleConn(client, ClientToServer)
	go handleConn(client, ServerToClient)

	wg.Wait()
}

func processConnectionClientMessage(client *Client, msg protoreflect.ProtoMessage) ([]protoreflect.ProtoMessage, error) {
	message := msg.(*connectionMessage.Message)

	switch message.Content.(type) {
	case *connectionMessage.Message_Request:
		//request := message.GetRequest()
	}

	return []protoreflect.ProtoMessage{message}, nil
}

func processConnectionServerMessage(client *Client, msg protoreflect.ProtoMessage) ([]protoreflect.ProtoMessage, error) {
	message := msg.(*connectionMessage.Message)

	switch message.Content.(type) {
	case *connectionMessage.Message_Response:
		response := message.GetResponse()
		switch response.Content.(type) {
		case *connectionMessage.Response_SelectServer:
			selectServer := response.GetSelectServer()
			switch selectServer.Result.(type) {
			case *connectionMessage.SelectServerResponse_Success_:
				success := selectServer.GetSuccess()

				gameProxy, found := serverProxyMap[success.GetHost()]
				if !found {
					// generate random port
					port := rand.Int32N(65535-49152) + 49152

					// Start the game proxy
					gameProxy = NewProxy(protocol.GameServer, port, fmt.Sprintf("%s:%d", success.GetHost(), success.GetPorts()[0]), &gameMessage.Message{})
					gameProxy.ClientMessageFunc = processGameClientMessage
					gameProxy.ServerMessageFunc = processGameServerMessage

					go func() {
						err := gameProxy.Start()
						if err != nil {
							panic(err)
						}
						gameProxy.Close()
					}()

					serverProxyMap[success.GetHost()] = gameProxy
				}

				success.Host = "localhost"
				success.Ports[0] = gameProxy.port
			}
		case *connectionMessage.Response_Identification:
			identificationResponse := response.GetIdentification()
			switch identificationResponse.Result.(type) {
			case *connectionMessage.IdentificationResponse_Success_:
				success := identificationResponse.GetSuccess()
				success.Rights.ShowForceAccount = true
				success.Rights.ShowConsole = true
				success.Rights.UnlimitedAccess = true
				success.Rights.InfiniteSubscription = true
				success.Rights.Report = true
			}
		}

	case *connectionMessage.Message_Event:
		//event := message.GetEvent()
	}

	return []protoreflect.ProtoMessage{message}, nil
}

func processGameClientMessage(client *Client, msg protoreflect.ProtoMessage) ([]protoreflect.ProtoMessage, error) {
	message := msg.(*gameMessage.Message)

	//subMessage, err := protocol.DecodeSubMessage(message)
	//if err != nil {
	//	return nil, err
	//}

	switch message.Content.(type) {
	case *gameMessage.Message_Request:
		//request := message.GetRequest()
	}

	return []protoreflect.ProtoMessage{message}, nil
}

func processGameServerMessage(client *Client, msg protoreflect.ProtoMessage) ([]protoreflect.ProtoMessage, error) {
	message := msg.(*gameMessage.Message)

	//subMessage, err := protocol.DecodeSubMessage(message)
	//if err != nil {
	//	return nil, err
	//}

	switch message.Content.(type) {
	case *gameMessage.Message_Response:
		//response := message.GetResponse()
	case *gameMessage.Message_Event:
		//event := message.GetEvent()
	}

	return []protoreflect.ProtoMessage{message}, nil
}
