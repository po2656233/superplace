// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.7.0
// source: proto.proto

package superProto

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type I32 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *I32) Reset() {
	*x = I32{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *I32) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*I32) ProtoMessage() {}

func (x *I32) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use I32.ProtoReflect.Descriptor instead.
func (*I32) Descriptor() ([]byte, []int) {
	return file_proto_proto_rawDescGZIP(), []int{0}
}

func (x *I32) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

// member data
type Member struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId   string            `protobuf:"bytes,1,opt,name=nodeId,proto3" json:"nodeId,omitempty"`                                                                                             // node id
	NodeType string            `protobuf:"bytes,2,opt,name=nodeType,proto3" json:"nodeType,omitempty"`                                                                                         // node type
	Address  string            `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`                                                                                           // rpc ip address
	Settings map[string]string `protobuf:"bytes,4,rep,name=settings,proto3" json:"settings,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // node settings data
}

func (x *Member) Reset() {
	*x = Member{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Member) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Member) ProtoMessage() {}

func (x *Member) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Member.ProtoReflect.Descriptor instead.
func (*Member) Descriptor() ([]byte, []int) {
	return file_proto_proto_rawDescGZIP(), []int{1}
}

func (x *Member) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *Member) GetNodeType() string {
	if x != nil {
		return x.NodeType
	}
	return ""
}

func (x *Member) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Member) GetSettings() map[string]string {
	if x != nil {
		return x.Settings
	}
	return nil
}

// member list data
type MemberList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*Member `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *MemberList) Reset() {
	*x = MemberList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberList) ProtoMessage() {}

func (x *MemberList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberList.ProtoReflect.Descriptor instead.
func (*MemberList) Descriptor() ([]byte, []int) {
	return file_proto_proto_rawDescGZIP(), []int{2}
}

func (x *MemberList) GetList() []*Member {
	if x != nil {
		return x.List
	}
	return nil
}

// cross node response data
type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"` // message code
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`  // message data
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_proto_proto_rawDescGZIP(), []int{3}
}

func (x *Response) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Response) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type ClusterPacket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BuildTime  int64    `protobuf:"varint,1,opt,name=buildTime,proto3" json:"buildTime,omitempty"`
	SourcePath string   `protobuf:"bytes,2,opt,name=sourcePath,proto3" json:"sourcePath,omitempty"`
	TargetPath string   `protobuf:"bytes,3,opt,name=targetPath,proto3" json:"targetPath,omitempty"`
	FuncName   string   `protobuf:"bytes,4,opt,name=funcName,proto3" json:"funcName,omitempty"`
	ArgBytes   []byte   `protobuf:"bytes,5,opt,name=argBytes,proto3" json:"argBytes,omitempty"`
	Session    *Session `protobuf:"bytes,6,opt,name=session,proto3" json:"session,omitempty"`
}

func (x *ClusterPacket) Reset() {
	*x = ClusterPacket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterPacket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterPacket) ProtoMessage() {}

func (x *ClusterPacket) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterPacket.ProtoReflect.Descriptor instead.
func (*ClusterPacket) Descriptor() ([]byte, []int) {
	return file_proto_proto_rawDescGZIP(), []int{4}
}

func (x *ClusterPacket) GetBuildTime() int64 {
	if x != nil {
		return x.BuildTime
	}
	return 0
}

func (x *ClusterPacket) GetSourcePath() string {
	if x != nil {
		return x.SourcePath
	}
	return ""
}

func (x *ClusterPacket) GetTargetPath() string {
	if x != nil {
		return x.TargetPath
	}
	return ""
}

func (x *ClusterPacket) GetFuncName() string {
	if x != nil {
		return x.FuncName
	}
	return ""
}

func (x *ClusterPacket) GetArgBytes() []byte {
	if x != nil {
		return x.ArgBytes
	}
	return nil
}

func (x *ClusterPacket) GetSession() *Session {
	if x != nil {
		return x.Session
	}
	return nil
}

type Session struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sid       string            `protobuf:"bytes,1,opt,name=sid,proto3" json:"sid,omitempty"`                                                                                           // session unique id
	Uid       int64             `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`                                                                                          // user id
	AgentPath string            `protobuf:"bytes,3,opt,name=agentPath,proto3" json:"agentPath,omitempty"`                                                                               // frontend actor agent path
	Ip        string            `protobuf:"bytes,4,opt,name=ip,proto3" json:"ip,omitempty"`                                                                                             // ip address
	Mid       uint32            `protobuf:"varint,5,opt,name=mid,proto3" json:"mid,omitempty"`                                                                                          // message id build by client
	Data      map[string]string `protobuf:"bytes,7,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // extend data
}

func (x *Session) Reset() {
	*x = Session{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Session) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session) ProtoMessage() {}

func (x *Session) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Session.ProtoReflect.Descriptor instead.
func (*Session) Descriptor() ([]byte, []int) {
	return file_proto_proto_rawDescGZIP(), []int{5}
}

func (x *Session) GetSid() string {
	if x != nil {
		return x.Sid
	}
	return ""
}

func (x *Session) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *Session) GetAgentPath() string {
	if x != nil {
		return x.AgentPath
	}
	return ""
}

func (x *Session) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *Session) GetMid() uint32 {
	if x != nil {
		return x.Mid
	}
	return 0
}

func (x *Session) GetData() map[string]string {
	if x != nil {
		return x.Data
	}
	return nil
}

type PomeloResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sid  string `protobuf:"bytes,1,opt,name=sid,proto3" json:"sid,omitempty"`
	Mid  uint32 `protobuf:"varint,2,opt,name=mid,proto3" json:"mid,omitempty"`
	Data []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Code int32  `protobuf:"varint,4,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *PomeloResponse) Reset() {
	*x = PomeloResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PomeloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PomeloResponse) ProtoMessage() {}

func (x *PomeloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PomeloResponse.ProtoReflect.Descriptor instead.
func (*PomeloResponse) Descriptor() ([]byte, []int) {
	return file_proto_proto_rawDescGZIP(), []int{6}
}

func (x *PomeloResponse) GetSid() string {
	if x != nil {
		return x.Sid
	}
	return ""
}

func (x *PomeloResponse) GetMid() uint32 {
	if x != nil {
		return x.Mid
	}
	return 0
}

func (x *PomeloResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *PomeloResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type PomeloPush struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sid   string `protobuf:"bytes,1,opt,name=sid,proto3" json:"sid,omitempty"`
	Route string `protobuf:"bytes,2,opt,name=route,proto3" json:"route,omitempty"`
	Data  []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *PomeloPush) Reset() {
	*x = PomeloPush{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PomeloPush) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PomeloPush) ProtoMessage() {}

func (x *PomeloPush) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PomeloPush.ProtoReflect.Descriptor instead.
func (*PomeloPush) Descriptor() ([]byte, []int) {
	return file_proto_proto_rawDescGZIP(), []int{7}
}

func (x *PomeloPush) GetSid() string {
	if x != nil {
		return x.Sid
	}
	return ""
}

func (x *PomeloPush) GetRoute() string {
	if x != nil {
		return x.Route
	}
	return ""
}

func (x *PomeloPush) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type PomeloKick struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sid    string `protobuf:"bytes,1,opt,name=sid,proto3" json:"sid,omitempty"`
	Uid    int64  `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Reason []byte `protobuf:"bytes,3,opt,name=reason,proto3" json:"reason,omitempty"`
	Close  bool   `protobuf:"varint,4,opt,name=close,proto3" json:"close,omitempty"`
}

func (x *PomeloKick) Reset() {
	*x = PomeloKick{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PomeloKick) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PomeloKick) ProtoMessage() {}

func (x *PomeloKick) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PomeloKick.ProtoReflect.Descriptor instead.
func (*PomeloKick) Descriptor() ([]byte, []int) {
	return file_proto_proto_rawDescGZIP(), []int{8}
}

func (x *PomeloKick) GetSid() string {
	if x != nil {
		return x.Sid
	}
	return ""
}

func (x *PomeloKick) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *PomeloKick) GetReason() []byte {
	if x != nil {
		return x.Reason
	}
	return nil
}

func (x *PomeloKick) GetClose() bool {
	if x != nil {
		return x.Close
	}
	return false
}

type PomeloBroadcastPush struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UidList []int64 `protobuf:"varint,1,rep,packed,name=uidList,proto3" json:"uidList,omitempty"` // broadcast the uid list
	AllUID  bool    `protobuf:"varint,2,opt,name=allUID,proto3" json:"allUID,omitempty"`          // broadcast all uid
	Route   string  `protobuf:"bytes,3,opt,name=route,proto3" json:"route,omitempty"`             // route
	Data    []byte  `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`               // data
}

func (x *PomeloBroadcastPush) Reset() {
	*x = PomeloBroadcastPush{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PomeloBroadcastPush) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PomeloBroadcastPush) ProtoMessage() {}

func (x *PomeloBroadcastPush) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PomeloBroadcastPush.ProtoReflect.Descriptor instead.
func (*PomeloBroadcastPush) Descriptor() ([]byte, []int) {
	return file_proto_proto_rawDescGZIP(), []int{9}
}

func (x *PomeloBroadcastPush) GetUidList() []int64 {
	if x != nil {
		return x.UidList
	}
	return nil
}

func (x *PomeloBroadcastPush) GetAllUID() bool {
	if x != nil {
		return x.AllUID
	}
	return false
}

func (x *PomeloBroadcastPush) GetRoute() string {
	if x != nil {
		return x.Route
	}
	return ""
}

func (x *PomeloBroadcastPush) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_proto_proto protoreflect.FileDescriptor

var file_proto_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63,
	0x68, 0x65, 0x72, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1b, 0x0a, 0x03, 0x49, 0x33,
	0x32, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xd2, 0x01, 0x0a, 0x06, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x6f,
	0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f,
	0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x3d, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x68, 0x65, 0x72, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x1a,
	0x3b, 0x0a, 0x0d, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x35, 0x0a, 0x0a,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x04, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x68, 0x65, 0x72, 0x72,
	0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x04, 0x6c,
	0x69, 0x73, 0x74, 0x22, 0x32, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xd5, 0x01, 0x0a, 0x0d, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x50, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x50, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x72, 0x67, 0x42, 0x79, 0x74, 0x65, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x61, 0x72, 0x67, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12,
	0x2e, 0x0a, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x63, 0x68, 0x65, 0x72, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22,
	0xda, 0x01, 0x0a, 0x07, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x73,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x69, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x74, 0x68, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6d, 0x69, 0x64, 0x12,
	0x32, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x63, 0x68, 0x65, 0x72, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x1a, 0x37, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x5c, 0x0a, 0x0e,
	0x50, 0x6f, 0x6d, 0x65, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x73, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x69, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6d,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x48, 0x0a, 0x0a, 0x50, 0x6f,
	0x6d, 0x65, 0x6c, 0x6f, 0x50, 0x75, 0x73, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x5e, 0x0a, 0x0a, 0x50, 0x6f, 0x6d, 0x65, 0x6c, 0x6f, 0x4b, 0x69,
	0x63, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x73, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x63,
	0x6c, 0x6f, 0x73, 0x65, 0x22, 0x71, 0x0a, 0x13, 0x50, 0x6f, 0x6d, 0x65, 0x6c, 0x6f, 0x42, 0x72,
	0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x50, 0x75, 0x73, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x75,
	0x69, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x07, 0x75, 0x69,
	0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6c, 0x6c, 0x55, 0x49, 0x44, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x6c, 0x6c, 0x55, 0x49, 0x44, 0x12, 0x14, 0x0a,
	0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x27, 0x5a, 0x25, 0x73, 0x75, 0x70, 0x65, 0x72,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x2f, 0x6e, 0x65, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x73, 0x75, 0x70, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_proto_rawDescOnce sync.Once
	file_proto_proto_rawDescData = file_proto_proto_rawDesc
)

func file_proto_proto_rawDescGZIP() []byte {
	file_proto_proto_rawDescOnce.Do(func() {
		file_proto_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_proto_rawDescData)
	})
	return file_proto_proto_rawDescData
}

var file_proto_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_proto_proto_goTypes = []interface{}{
	(*I32)(nil),                 // 0: cherryProto.I32
	(*Member)(nil),              // 1: cherryProto.Member
	(*MemberList)(nil),          // 2: cherryProto.MemberList
	(*Response)(nil),            // 3: cherryProto.Response
	(*ClusterPacket)(nil),       // 4: cherryProto.ClusterPacket
	(*Session)(nil),             // 5: cherryProto.Session
	(*PomeloResponse)(nil),      // 6: cherryProto.PomeloResponse
	(*PomeloPush)(nil),          // 7: cherryProto.PomeloPush
	(*PomeloKick)(nil),          // 8: cherryProto.PomeloKick
	(*PomeloBroadcastPush)(nil), // 9: cherryProto.PomeloBroadcastPush
	nil,                         // 10: cherryProto.Member.SettingsEntry
	nil,                         // 11: cherryProto.Session.DataEntry
}
var file_proto_proto_depIdxs = []int32{
	10, // 0: cherryProto.Member.settings:type_name -> cherryProto.Member.SettingsEntry
	1,  // 1: cherryProto.MemberList.list:type_name -> cherryProto.Member
	5,  // 2: cherryProto.ClusterPacket.session:type_name -> cherryProto.Session
	11, // 3: cherryProto.Session.data:type_name -> cherryProto.Session.DataEntry
	4,  // [4:4] is the sub-list for method output_type
	4,  // [4:4] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_proto_proto_init() }
func file_proto_proto_init() {
	if File_proto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*I32); i {
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
		file_proto_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Member); i {
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
		file_proto_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberList); i {
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
		file_proto_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_proto_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterPacket); i {
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
		file_proto_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Session); i {
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
		file_proto_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PomeloResponse); i {
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
		file_proto_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PomeloPush); i {
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
		file_proto_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PomeloKick); i {
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
		file_proto_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PomeloBroadcastPush); i {
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
			RawDescriptor: file_proto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_proto_goTypes,
		DependencyIndexes: file_proto_proto_depIdxs,
		MessageInfos:      file_proto_proto_msgTypes,
	}.Build()
	File_proto_proto = out.File
	file_proto_proto_rawDesc = nil
	file_proto_proto_goTypes = nil
	file_proto_proto_depIdxs = nil
}
