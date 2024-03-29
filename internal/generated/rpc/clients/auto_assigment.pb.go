// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: proto/auto_assigment.proto

package clients

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Agent_Statuses int32

const (
	Agent_ACTIVE        Agent_Statuses = 0
	Agent_REQ_INACTIVE  Agent_Statuses = 1
	Agent_INACTIVE      Agent_Statuses = 2
	Agent_REQ_BREAK     Agent_Statuses = 3
	Agent_BREAK         Agent_Statuses = 4
	Agent_FORCE_MAJEURE Agent_Statuses = 5
	Agent_CHAT          Agent_Statuses = 6
	Agent_LETTER        Agent_Statuses = 7
)

// Enum value maps for Agent_Statuses.
var (
	Agent_Statuses_name = map[int32]string{
		0: "ACTIVE",
		1: "REQ_INACTIVE",
		2: "INACTIVE",
		3: "REQ_BREAK",
		4: "BREAK",
		5: "FORCE_MAJEURE",
		6: "CHAT",
		7: "LETTER",
	}
	Agent_Statuses_value = map[string]int32{
		"ACTIVE":        0,
		"REQ_INACTIVE":  1,
		"INACTIVE":      2,
		"REQ_BREAK":     3,
		"BREAK":         4,
		"FORCE_MAJEURE": 5,
		"CHAT":          6,
		"LETTER":        7,
	}
)

func (x Agent_Statuses) Enum() *Agent_Statuses {
	p := new(Agent_Statuses)
	*p = x
	return p
}

func (x Agent_Statuses) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Agent_Statuses) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_auto_assigment_proto_enumTypes[0].Descriptor()
}

func (Agent_Statuses) Type() protoreflect.EnumType {
	return &file_proto_auto_assigment_proto_enumTypes[0]
}

func (x Agent_Statuses) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Agent_Statuses.Descriptor instead.
func (Agent_Statuses) EnumDescriptor() ([]byte, []int) {
	return file_proto_auto_assigment_proto_rawDescGZIP(), []int{2, 0}
}

type ChangeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *Agent `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *ChangeRequest) Reset() {
	*x = ChangeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auto_assigment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeRequest) ProtoMessage() {}

func (x *ChangeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auto_assigment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeRequest.ProtoReflect.Descriptor instead.
func (*ChangeRequest) Descriptor() ([]byte, []int) {
	return file_proto_auto_assigment_proto_rawDescGZIP(), []int{0}
}

func (x *ChangeRequest) GetUser() *Agent {
	if x != nil {
		return x.User
	}
	return nil
}

type UserChangeStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserChangeStatusResponse) Reset() {
	*x = UserChangeStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auto_assigment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserChangeStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserChangeStatusResponse) ProtoMessage() {}

func (x *UserChangeStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auto_assigment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserChangeStatusResponse.ProtoReflect.Descriptor instead.
func (*UserChangeStatusResponse) Descriptor() ([]byte, []int) {
	return file_proto_auto_assigment_proto_rawDescGZIP(), []int{1}
}

type Agent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login  string         `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Status Agent_Statuses `protobuf:"varint,2,opt,name=status,proto3,enum=clients.Agent_Statuses" json:"status,omitempty"`
}

func (x *Agent) Reset() {
	*x = Agent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auto_assigment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Agent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Agent) ProtoMessage() {}

func (x *Agent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auto_assigment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Agent.ProtoReflect.Descriptor instead.
func (*Agent) Descriptor() ([]byte, []int) {
	return file_proto_auto_assigment_proto_rawDescGZIP(), []int{2}
}

func (x *Agent) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *Agent) GetStatus() Agent_Statuses {
	if x != nil {
		return x.Status
	}
	return Agent_ACTIVE
}

var File_proto_auto_assigment_proto protoreflect.FileDescriptor

var file_proto_auto_assigment_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x61, 0x73, 0x73,
	0x69, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x33, 0x0a, 0x0d, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x1a, 0x0a, 0x18, 0x55, 0x73,
	0x65, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xc9, 0x01, 0x0a, 0x05, 0x41, 0x67, 0x65, 0x6e, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x2f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x79, 0x0a, 0x08, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x65, 0x73, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x00, 0x12,
	0x10, 0x0a, 0x0c, 0x52, 0x45, 0x51, 0x5f, 0x49, 0x4e, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10,
	0x01, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x4e, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x02, 0x12,
	0x0d, 0x0a, 0x09, 0x52, 0x45, 0x51, 0x5f, 0x42, 0x52, 0x45, 0x41, 0x4b, 0x10, 0x03, 0x12, 0x09,
	0x0a, 0x05, 0x42, 0x52, 0x45, 0x41, 0x4b, 0x10, 0x04, 0x12, 0x11, 0x0a, 0x0d, 0x46, 0x4f, 0x52,
	0x43, 0x45, 0x5f, 0x4d, 0x41, 0x4a, 0x45, 0x55, 0x52, 0x45, 0x10, 0x05, 0x12, 0x08, 0x0a, 0x04,
	0x43, 0x48, 0x41, 0x54, 0x10, 0x06, 0x12, 0x0a, 0x0a, 0x06, 0x4c, 0x45, 0x54, 0x54, 0x45, 0x52,
	0x10, 0x07, 0x32, 0x5e, 0x0a, 0x0d, 0x41, 0x75, 0x74, 0x6f, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x4d, 0x0a, 0x10, 0x75, 0x73, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x21, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x20, 0x5a, 0x1e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_auto_assigment_proto_rawDescOnce sync.Once
	file_proto_auto_assigment_proto_rawDescData = file_proto_auto_assigment_proto_rawDesc
)

func file_proto_auto_assigment_proto_rawDescGZIP() []byte {
	file_proto_auto_assigment_proto_rawDescOnce.Do(func() {
		file_proto_auto_assigment_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_auto_assigment_proto_rawDescData)
	})
	return file_proto_auto_assigment_proto_rawDescData
}

var file_proto_auto_assigment_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_auto_assigment_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_auto_assigment_proto_goTypes = []interface{}{
	(Agent_Statuses)(0),              // 0: clients.Agent.Statuses
	(*ChangeRequest)(nil),            // 1: clients.ChangeRequest
	(*UserChangeStatusResponse)(nil), // 2: clients.UserChangeStatusResponse
	(*Agent)(nil),                    // 3: clients.Agent
}
var file_proto_auto_assigment_proto_depIdxs = []int32{
	3, // 0: clients.ChangeRequest.user:type_name -> clients.Agent
	0, // 1: clients.Agent.status:type_name -> clients.Agent.Statuses
	1, // 2: clients.AutoAssigment.userChangeStatus:input_type -> clients.ChangeRequest
	2, // 3: clients.AutoAssigment.userChangeStatus:output_type -> clients.UserChangeStatusResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_auto_assigment_proto_init() }
func file_proto_auto_assigment_proto_init() {
	if File_proto_auto_assigment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_auto_assigment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_auto_assigment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserChangeStatusResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_auto_assigment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Agent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_auto_assigment_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_auto_assigment_proto_goTypes,
		DependencyIndexes: file_proto_auto_assigment_proto_depIdxs,
		EnumInfos:         file_proto_auto_assigment_proto_enumTypes,
		MessageInfos:      file_proto_auto_assigment_proto_msgTypes,
	}.Build()
	File_proto_auto_assigment_proto = out.File
	file_proto_auto_assigment_proto_rawDesc = nil
	file_proto_auto_assigment_proto_goTypes = nil
	file_proto_auto_assigment_proto_depIdxs = nil
}
