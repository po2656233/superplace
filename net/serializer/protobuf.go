package cherrySerializer

import (
	cerr "github.com/po2656233/superplace/logger/error"
	"google.golang.org/protobuf/proto"
)

// Protobuf implements the serialize.Protobuf interface
type Protobuf struct{}

// NewProtobuf NewSerializer returns a new Protobuf.
func NewProtobuf() *Protobuf {
	return &Protobuf{}
}

// Marshal returns the protobuf encoding of v.
func (p *Protobuf) Marshal(v interface{}) ([]byte, error) {
	if data, ok := v.([]byte); ok {
		return data, nil
	}

	pb, ok := v.(proto.Message)
	if !ok {
		return nil, cerr.ProtobufWrongValueType
	}
	return proto.Marshal(pb)
}

// Unmarshal parses the protobuf-encoded data and stores the result
// in the value pointed to by v.
func (p *Protobuf) Unmarshal(data []byte, v interface{}) error {
	pb, ok := v.(proto.Message)
	if !ok {
		return cerr.ProtobufWrongValueType
	}
	return proto.Unmarshal(data, pb)
}

// Name returns the name of the serializer.
func (p *Protobuf) Name() string {
	return "protobuf"
}
