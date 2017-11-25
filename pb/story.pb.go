// Code generated by protoc-gen-go. DO NOT EDIT.
// source: story.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Story struct {
	Start int32  `protobuf:"varint,1,opt,name=start" json:"start,omitempty"`
	End   int32  `protobuf:"varint,2,opt,name=end" json:"end,omitempty"`
	Doc   string `protobuf:"bytes,3,opt,name=doc" json:"doc,omitempty"`
}

func (m *Story) Reset()                    { *m = Story{} }
func (m *Story) String() string            { return proto.CompactTextString(m) }
func (*Story) ProtoMessage()               {}
func (*Story) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *Story) GetStart() int32 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *Story) GetEnd() int32 {
	if m != nil {
		return m.End
	}
	return 0
}

func (m *Story) GetDoc() string {
	if m != nil {
		return m.Doc
	}
	return ""
}

func init() {
	proto.RegisterType((*Story)(nil), "pb.Story")
}

func init() { proto.RegisterFile("story.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 97 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x2e, 0xc9, 0x2f,
	0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x72, 0xe4, 0x62, 0x0d,
	0x06, 0x09, 0x09, 0x89, 0x70, 0xb1, 0x16, 0x97, 0x24, 0x16, 0x95, 0x48, 0x30, 0x2a, 0x30, 0x6a,
	0xb0, 0x06, 0x41, 0x38, 0x42, 0x02, 0x5c, 0xcc, 0xa9, 0x79, 0x29, 0x12, 0x4c, 0x60, 0x31, 0x10,
	0x13, 0x24, 0x92, 0x92, 0x9f, 0x2c, 0xc1, 0xac, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0x62, 0x26, 0xb1,
	0x81, 0x4d, 0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xeb, 0x5a, 0x66, 0xb3, 0x5c, 0x00, 0x00,
	0x00,
}
