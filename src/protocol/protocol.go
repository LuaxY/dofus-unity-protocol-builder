package protocol

import (
	"encoding/binary"
	"fmt"
	gameMessage "go-xp-dofus-unity-proto-builder/src/protocol/game/message"
	"go-xp-dofus-unity-proto-builder/src/resolver"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
)

var customResolver = resolver.NewCustomResolver()

var marshaller = protojson.MarshalOptions{
	Indent:   "  ",
	Resolver: customResolver,
}

var unmarshaller = proto.UnmarshalOptions{
	Resolver: customResolver,
}

const typeUrlPrefix = "type.ankama.com/"

type ServerType string

const (
	ConnectionServer ServerType = "Connection"
	GameServer       ServerType = "Game"
)

type Direction string

const (
	ClientToServer Direction = "ClientToServer"
	ServerToClient Direction = "ServerToClient"
)

func PrettyPrintMessage(message protoreflect.ProtoMessage) string {
	jsonBytes, err := marshaller.Marshal(message)
	if err != nil {
		return "" // TODO: handle error
	}

	return string(jsonBytes)
}

func EncodeMessage(message protoreflect.ProtoMessage) ([]byte, error) {
	data, err := proto.Marshal(message)
	if err != nil {
		fmt.Println("Error marshalling message:", err)
	}

	varIntBuffer := make([]byte, binary.MaxVarintLen64)
	sizeLength := binary.PutUvarint(varIntBuffer, uint64(len(data)))
	buffer := append(varIntBuffer[:sizeLength], data...)

	return buffer, nil
}

func DecodeMessage(data []byte, message proto.Message) error {
	if err := unmarshaller.Unmarshal(data, message); err != nil {
		return fmt.Errorf("error unmarshalling message: %w", err)
	}

	return nil
}

func DecodeSubMessage(message *gameMessage.Message) (protoreflect.ProtoMessage, error) {
	request := message.GetRequest()
	typeUrl := request.GetContent().GetTypeUrl()
	anyMsg := request.GetContent()

	messageType, err := customResolver.FindMessageByURL(typeUrl)
	if err != nil {
		return nil, err
	}

	specificMsg := messageType.New().Interface()
	if err = anyMsg.UnmarshalTo(specificMsg); err != nil {
		return nil, err
	}

	return specificMsg, nil
}

func EncodeGameRequest(request protoreflect.ProtoMessage) (*gameMessage.Message, error) {
	marshalMsg, err := proto.Marshal(request)
	if err != nil {
		return nil, err
	}

	anyMsg := &anypb.Any{
		TypeUrl: typeUrlPrefix + string(request.ProtoReflect().Descriptor().FullName()),
		Value:   marshalMsg,
	}

	return &gameMessage.Message{
		Content: &gameMessage.Message_Request{
			Request: &gameMessage.Request{
				Content: anyMsg,
			},
		},
	}, nil
}

func EncodeGameResponse(response protoreflect.ProtoMessage) (*gameMessage.Message, error) {
	marshalMsg, err := proto.Marshal(response)
	if err != nil {
		return nil, err
	}

	anyMsg := &anypb.Any{
		TypeUrl: typeUrlPrefix + string(response.ProtoReflect().Descriptor().FullName()),
		Value:   marshalMsg,
	}

	return &gameMessage.Message{
		Content: &gameMessage.Message_Response{
			Response: &gameMessage.Response{
				Content: anyMsg,
			},
		},
	}, nil
}

func EncodeGameEvent(event protoreflect.ProtoMessage) (*gameMessage.Message, error) {
	marshalMsg, err := proto.Marshal(event)
	if err != nil {
		return nil, err
	}

	anyMsg := &anypb.Any{
		TypeUrl: typeUrlPrefix + string(event.ProtoReflect().Descriptor().FullName()),
		Value:   marshalMsg,
	}

	return &gameMessage.Message{
		Content: &gameMessage.Message_Event{
			Event: &gameMessage.Event{
				Content: anyMsg,
			},
		},
	}, nil
}
