// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: proto/chat.proto

package chat

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RequestType int32

const (
	RequestType_MESSAGE RequestType = 0 // Default
	RequestType_JOIN    RequestType = 1
	RequestType_LEAVE   RequestType = 2
)

// Enum value maps for RequestType.
var (
	RequestType_name = map[int32]string{
		0: "MESSAGE",
		1: "JOIN",
		2: "LEAVE",
	}
	RequestType_value = map[string]int32{
		"MESSAGE": 0,
		"JOIN":    1,
		"LEAVE":   2,
	}
)

func (x RequestType) Enum() *RequestType {
	p := new(RequestType)
	*p = x
	return p
}

func (x RequestType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RequestType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_chat_proto_enumTypes[0].Descriptor()
}

func (RequestType) Type() protoreflect.EnumType {
	return &file_proto_chat_proto_enumTypes[0]
}

func (x RequestType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RequestType.Descriptor instead.
func (RequestType) EnumDescriptor() ([]byte, []int) {
	return file_proto_chat_proto_rawDescGZIP(), []int{0}
}

type MessageRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Sender        string                 `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Channel       string                 `protobuf:"bytes,3,opt,name=channel,proto3" json:"channel,omitempty"`
	Type          RequestType            `protobuf:"varint,4,opt,name=type,proto3,enum=chat.RequestType" json:"type,omitempty"` // NEW: join/leave/message
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MessageRequest) Reset() {
	*x = MessageRequest{}
	mi := &file_proto_chat_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageRequest) ProtoMessage() {}

func (x *MessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chat_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageRequest.ProtoReflect.Descriptor instead.
func (*MessageRequest) Descriptor() ([]byte, []int) {
	return file_proto_chat_proto_rawDescGZIP(), []int{0}
}

func (x *MessageRequest) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *MessageRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *MessageRequest) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *MessageRequest) GetType() RequestType {
	if x != nil {
		return x.Type
	}
	return RequestType_MESSAGE
}

type MessageResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Sender        string                 `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Timestamp     int64                  `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Channel       string                 `protobuf:"bytes,4,opt,name=channel,proto3" json:"channel,omitempty"` // NEW: clients may want to know which channel it came from
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MessageResponse) Reset() {
	*x = MessageResponse{}
	mi := &file_proto_chat_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageResponse) ProtoMessage() {}

func (x *MessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chat_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageResponse.ProtoReflect.Descriptor instead.
func (*MessageResponse) Descriptor() ([]byte, []int) {
	return file_proto_chat_proto_rawDescGZIP(), []int{1}
}

func (x *MessageResponse) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *MessageResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *MessageResponse) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *MessageResponse) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

var File_proto_chat_proto protoreflect.FileDescriptor

const file_proto_chat_proto_rawDesc = "" +
	"\n" +
	"\x10proto/chat.proto\x12\x04chat\"\x83\x01\n" +
	"\x0eMessageRequest\x12\x16\n" +
	"\x06sender\x18\x01 \x01(\tR\x06sender\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\x12\x18\n" +
	"\achannel\x18\x03 \x01(\tR\achannel\x12%\n" +
	"\x04type\x18\x04 \x01(\x0e2\x11.chat.RequestTypeR\x04type\"{\n" +
	"\x0fMessageResponse\x12\x16\n" +
	"\x06sender\x18\x01 \x01(\tR\x06sender\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\x12\x1c\n" +
	"\ttimestamp\x18\x03 \x01(\x03R\ttimestamp\x12\x18\n" +
	"\achannel\x18\x04 \x01(\tR\achannel*/\n" +
	"\vRequestType\x12\v\n" +
	"\aMESSAGE\x10\x00\x12\b\n" +
	"\x04JOIN\x10\x01\x12\t\n" +
	"\x05LEAVE\x10\x022L\n" +
	"\vChatService\x12=\n" +
	"\n" +
	"ChatStream\x12\x14.chat.MessageRequest\x1a\x15.chat.MessageResponse(\x010\x01Bo\n" +
	"\bcom.chatB\tChatProtoP\x01Z(github.com/pgibb96/MessageApp/proto;chat\xa2\x02\x03CXX\xaa\x02\x04Chat\xca\x02\x04Chat\xe2\x02\x10Chat\\GPBMetadata\xea\x02\x04Chatb\x06proto3"

var (
	file_proto_chat_proto_rawDescOnce sync.Once
	file_proto_chat_proto_rawDescData []byte
)

func file_proto_chat_proto_rawDescGZIP() []byte {
	file_proto_chat_proto_rawDescOnce.Do(func() {
		file_proto_chat_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_chat_proto_rawDesc), len(file_proto_chat_proto_rawDesc)))
	})
	return file_proto_chat_proto_rawDescData
}

var file_proto_chat_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_chat_proto_goTypes = []any{
	(RequestType)(0),        // 0: chat.RequestType
	(*MessageRequest)(nil),  // 1: chat.MessageRequest
	(*MessageResponse)(nil), // 2: chat.MessageResponse
}
var file_proto_chat_proto_depIdxs = []int32{
	0, // 0: chat.MessageRequest.type:type_name -> chat.RequestType
	1, // 1: chat.ChatService.ChatStream:input_type -> chat.MessageRequest
	2, // 2: chat.ChatService.ChatStream:output_type -> chat.MessageResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_chat_proto_init() }
func file_proto_chat_proto_init() {
	if File_proto_chat_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_chat_proto_rawDesc), len(file_proto_chat_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_chat_proto_goTypes,
		DependencyIndexes: file_proto_chat_proto_depIdxs,
		EnumInfos:         file_proto_chat_proto_enumTypes,
		MessageInfos:      file_proto_chat_proto_msgTypes,
	}.Build()
	File_proto_chat_proto = out.File
	file_proto_chat_proto_goTypes = nil
	file_proto_chat_proto_depIdxs = nil
}
