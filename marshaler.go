package protocol

import (
	"github.com/lynx-go/x/encoding/json"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type Marshaler interface {
	Name() string
	UseBytes() bool
	Marshal(data any) ([]byte, error)
	Unmarshal(data []byte, out any) error
}

var ProtobufMarshaler = &protobufMarshaler{}

var ProtoJSONMarshaler = &protoJsonMarshaler{}

var JSONMarshaler = &jsonMarshaler{}

var Marshalers = []Marshaler{
	ProtobufMarshaler,
	ProtoJSONMarshaler,
}

type jsonMarshaler struct{}

func (msh *jsonMarshaler) Name() string {
	return "json"
}

func (msh *jsonMarshaler) UseBytes() bool {
	return false
}

func (msh *jsonMarshaler) Marshal(data any) ([]byte, error) {
	return json.Marshal(data)
}

func (msh *jsonMarshaler) Unmarshal(data []byte, out any) error {
	return json.Unmarshal(data, out)
}

var _ Marshaler = new(jsonMarshaler)

type protobufMarshaler struct{}

func (m *protobufMarshaler) UseBytes() bool {
	return true
}

func (m *protobufMarshaler) Name() string {
	return "protobuf"
}

func (m *protobufMarshaler) Marshal(data any) ([]byte, error) {
	return proto.Marshal(data.(proto.Message))
}

func (m *protobufMarshaler) Unmarshal(data []byte, out any) error {
	return proto.Unmarshal(data, out.(proto.Message))
}

var _ Marshaler = new(protobufMarshaler)

type protoJsonMarshaler struct{}

func (m *protoJsonMarshaler) UseBytes() bool {
	return false
}

func (m *protoJsonMarshaler) Name() string {
	return "json"
}

func (m *protoJsonMarshaler) Marshal(data any) ([]byte, error) {
	return protojson.Marshal(data.(proto.Message))
}

func (m *protoJsonMarshaler) Unmarshal(data []byte, out any) error {
	return protojson.Unmarshal(data, out.(proto.Message))
}

var _ Marshaler = new(protoJsonMarshaler)
