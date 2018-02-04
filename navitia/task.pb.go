// Code generated by protoc-gen-go. DO NOT EDIT.
// source: task.proto

package pbnavitia

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Action int32

const (
	Action_RELOAD        Action = 0
	Action_HEARTBEAT     Action = 3
	Action_LOAD_REALTIME Action = 4
)

var Action_name = map[int32]string{
	0: "RELOAD",
	3: "HEARTBEAT",
	4: "LOAD_REALTIME",
}
var Action_value = map[string]int32{
	"RELOAD":        0,
	"HEARTBEAT":     3,
	"LOAD_REALTIME": 4,
}

func (x Action) Enum() *Action {
	p := new(Action)
	*p = x
	return p
}
func (x Action) String() string {
	return proto.EnumName(Action_name, int32(x))
}
func (x *Action) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Action_value, data, "Action")
	if err != nil {
		return err
	}
	*x = Action(value)
	return nil
}
func (Action) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

type Task struct {
	Action           *Action       `protobuf:"varint,1,req,name=action,enum=pbnavitia.Action" json:"action,omitempty"`
	LoadRealtime     *LoadRealtime `protobuf:"bytes,2,opt,name=load_realtime" json:"load_realtime,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *Task) Reset()                    { *m = Task{} }
func (m *Task) String() string            { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()               {}
func (*Task) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *Task) GetAction() Action {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return Action_RELOAD
}

func (m *Task) GetLoadRealtime() *LoadRealtime {
	if m != nil {
		return m.LoadRealtime
	}
	return nil
}

type LoadRealtime struct {
	QueueName        *string  `protobuf:"bytes,1,req,name=queue_name" json:"queue_name,omitempty"`
	Contributors     []string `protobuf:"bytes,2,rep,name=contributors" json:"contributors,omitempty"`
	BeginDate        *string  `protobuf:"bytes,3,opt,name=begin_date" json:"begin_date,omitempty"`
	EndDate          *string  `protobuf:"bytes,4,opt,name=end_date" json:"end_date,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *LoadRealtime) Reset()                    { *m = LoadRealtime{} }
func (m *LoadRealtime) String() string            { return proto.CompactTextString(m) }
func (*LoadRealtime) ProtoMessage()               {}
func (*LoadRealtime) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *LoadRealtime) GetQueueName() string {
	if m != nil && m.QueueName != nil {
		return *m.QueueName
	}
	return ""
}

func (m *LoadRealtime) GetContributors() []string {
	if m != nil {
		return m.Contributors
	}
	return nil
}

func (m *LoadRealtime) GetBeginDate() string {
	if m != nil && m.BeginDate != nil {
		return *m.BeginDate
	}
	return ""
}

func (m *LoadRealtime) GetEndDate() string {
	if m != nil && m.EndDate != nil {
		return *m.EndDate
	}
	return ""
}

func init() {
	proto.RegisterType((*Task)(nil), "pbnavitia.Task")
	proto.RegisterType((*LoadRealtime)(nil), "pbnavitia.LoadRealtime")
	proto.RegisterEnum("pbnavitia.Action", Action_name, Action_value)
}

func init() { proto.RegisterFile("task.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8e, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0x47, 0xcd, 0x1f, 0x82, 0x3b, 0x36, 0x92, 0x0e, 0x82, 0x39, 0xc6, 0x9e, 0x82, 0x87, 0x1c,
	0x7a, 0xf0, 0xbe, 0xe2, 0x82, 0x42, 0x44, 0x08, 0xb9, 0x78, 0x31, 0x4c, 0x9a, 0x45, 0x96, 0xb6,
	0xbb, 0x75, 0x33, 0xf1, 0xf3, 0x4b, 0x6a, 0x29, 0x3d, 0xce, 0x7b, 0x0f, 0xe6, 0x07, 0xc0, 0x34,
	0x6e, 0xab, 0x83, 0x77, 0xec, 0x50, 0x1c, 0x7a, 0x4b, 0xbf, 0x86, 0x0d, 0xad, 0x3e, 0x21, 0x6e,
	0x69, 0xdc, 0xe2, 0x03, 0x24, 0xb4, 0x61, 0xe3, 0x6c, 0x1e, 0x14, 0x61, 0x79, 0xbb, 0x5e, 0x56,
	0xe7, 0xa6, 0x92, 0x47, 0x81, 0x15, 0xa4, 0x3b, 0x47, 0x43, 0xe7, 0x35, 0xed, 0xd8, 0xec, 0x75,
	0x1e, 0x16, 0x41, 0x79, 0xb3, 0xbe, 0xbf, 0x28, 0x6b, 0x47, 0x43, 0x73, 0xd2, 0xab, 0x2f, 0x58,
	0x5c, 0xde, 0x88, 0x00, 0x3f, 0x93, 0x9e, 0x74, 0x67, 0x69, 0xaf, 0x8f, 0x6f, 0x04, 0xde, 0xc1,
	0x62, 0xe3, 0x2c, 0x7b, 0xd3, 0x4f, 0xec, 0xfc, 0x98, 0x87, 0x45, 0x54, 0x8a, 0xb9, 0xec, 0xf5,
	0xb7, 0xb1, 0xdd, 0x40, 0xac, 0xf3, 0xa8, 0x08, 0x4a, 0x81, 0x19, 0x5c, 0x6b, 0x3b, 0xfc, 0x93,
	0x78, 0x26, 0x8f, 0x4f, 0x90, 0x9c, 0x96, 0x01, 0x24, 0x8d, 0xaa, 0x3f, 0xe4, 0x4b, 0x76, 0x85,
	0x29, 0x88, 0x57, 0x25, 0x9b, 0xf6, 0x59, 0xc9, 0x36, 0x8b, 0x70, 0x09, 0xe9, 0x2c, 0xba, 0x46,
	0xc9, 0xba, 0x7d, 0x7b, 0x57, 0x59, 0xfc, 0x17, 0x00, 0x00, 0xff, 0xff, 0x92, 0x1a, 0x28, 0x6a,
	0x0a, 0x01, 0x00, 0x00,
}