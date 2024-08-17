package main

import (
	"encoding/binary"
	"fmt"
	"github.com/fatih/color"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"go-xp-dofus-unity-proto-builder/src/protocol"
	connectionMessage "go-xp-dofus-unity-proto-builder/src/protocol/connection/message"
	gameMessage "go-xp-dofus-unity-proto-builder/src/protocol/game/message"
	"google.golang.org/protobuf/proto"
	"log"
)

var bufferMap = make(map[string][]byte)

var connectionIPList = []string{
	"34.249.146.47",
	"54.74.22.247",
}

func main() {
	bufferMap[string(protocol.ConnectionServer)+string(protocol.ClientToServer)] = make([]byte, 0)
	bufferMap[string(protocol.ConnectionServer)+string(protocol.ServerToClient)] = make([]byte, 0)
	bufferMap[string(protocol.GameServer)+string(protocol.ClientToServer)] = make([]byte, 0)
	bufferMap[string(protocol.GameServer)+string(protocol.ServerToClient)] = make([]byte, 0)

	handle, err := pcap.OpenLive("en0", 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		handleEthernetPacket(handle, packet)
	}
}

func handleEthernetPacket(handle *pcap.Handle, packet gopacket.Packet) {
	ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
	if ethernetLayer != nil {
		ethernetPacket, _ := ethernetLayer.(*layers.Ethernet)
		switch ethernetPacket.EthernetType {
		case layers.EthernetTypeIPv4:
			ipv4Handler(handle, ethernetPacket, packet)
		}
	}
}

func ipv4Handler(handle *pcap.Handle, ethernetPacket *layers.Ethernet, packet gopacket.Packet) {
	ipv4Layer := packet.Layer(layers.LayerTypeIPv4)
	if ipv4Layer != nil {
		ipv4Packet, _ := ipv4Layer.(*layers.IPv4)
		handleTransportProtocol(handle, ipv4Packet, ethernetPacket)
	}
}

func handleTransportProtocol(handle *pcap.Handle, ipv4Packet *layers.IPv4, ethernetPacket *layers.Ethernet) {
	switch ipv4Packet.Protocol {
	case layers.IPProtocolTCP:
		handleTCPPacket(handle, ipv4Packet, ethernetPacket)
	}
}

func handleTCPPacket(handle *pcap.Handle, ipv4Packet *layers.IPv4, ethernetPacket *layers.Ethernet) {
	tcpLayer := gopacket.NewPacket(ipv4Packet.Payload, layers.LayerTypeTCP, gopacket.Default).Layer(layers.LayerTypeTCP)
	if tcpLayer != nil {
		tcpPacket, _ := tcpLayer.(*layers.TCP)
		handleDofusPacket(handle, ipv4Packet, tcpPacket, ethernetPacket)
	}
}

func handleDofusPacket(handle *pcap.Handle, ipv4Packet *layers.IPv4, tcpPacket *layers.TCP, ethernetPacket *layers.Ethernet) {
	if tcpPacket.DstPort != 5555 && tcpPacket.SrcPort != 5555 {
		return
	}

	if len(tcpPacket.Payload) == 0 {
		return
	}

	var serverType protocol.ServerType
	var message proto.Message

	if contains(connectionIPList, ipv4Packet.SrcIP.String()) || contains(connectionIPList, ipv4Packet.DstIP.String()) {
		serverType = protocol.ConnectionServer
		message = &connectionMessage.Message{}
	} else {
		serverType = protocol.GameServer
		message = &gameMessage.Message{}
	}

	direction := protocol.ClientToServer
	if tcpPacket.SrcPort == 5555 {
		direction = protocol.ServerToClient
	}

	bufferId := string(serverType) + string(direction)
	partialBuffer := bufferMap[bufferId]
	partialBuffer = append(partialBuffer, tcpPacket.Payload...)

	for {
		size, sizeLength := binary.Uvarint(partialBuffer[:])

		// Incomplete message
		if size == 0 || int(size) > len(partialBuffer)-sizeLength {
			break
		}

		payload := partialBuffer[sizeLength : sizeLength+int(size)]

		if err := protocol.DecodeMessage(payload, message); err != nil {
			fmt.Printf("WRONG PROTO PAYLOAD READ %x\n", payload)
		}

		msgTyp := "UNK"
		msgName := "Unknown"

		switch serverType {
		case protocol.ConnectionServer:
			cm := message.(*connectionMessage.Message)
			switch cm.Content.(type) {
			case *connectionMessage.Message_Request:
				msgTyp = "REQ"
				request := cm.GetRequest()
				msgName = request.String()
			case *connectionMessage.Message_Response:
				msgTyp = "RES"
				response := cm.GetResponse()
				msgName = response.String()
			case *connectionMessage.Message_Event:
				msgTyp = "EVT"
				event := cm.GetEvent()
				msgName = event.String()
			}
		case protocol.GameServer:
			gm := message.(*gameMessage.Message)
			switch gm.Content.(type) {
			case *gameMessage.Message_Request:
				msgTyp = "REQ"
				request := gm.GetRequest()
				msgName = request.GetContent().GetTypeUrl()
			case *gameMessage.Message_Response:
				msgTyp = "RES"
				response := gm.GetResponse()
				msgName = response.GetContent().GetTypeUrl()
			case *gameMessage.Message_Event:
				msgTyp = "EVT"
				event := gm.GetEvent()
				msgName = event.GetContent().GetTypeUrl()
			}
		}

		messageDetails := protocol.PrettyPrintMessage(message)

		switch direction {
		case protocol.ClientToServer:
			color.Cyan("[C->S] [%s] %s\n", msgTyp, msgName)
			color.Cyan(messageDetails)
		case protocol.ServerToClient:
			color.Blue("[S->C] [%s] %s\n", msgTyp, msgName)
			color.Blue(messageDetails)
		}

		partialBuffer = partialBuffer[sizeLength+int(size):]
	}

	bufferMap[bufferId] = partialBuffer
}

func contains[T comparable](s []T, e T) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
