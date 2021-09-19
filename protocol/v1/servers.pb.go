// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: v1/servers.proto

package serversv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A enumerated list of possible types of an event
type Event_Type int32

const (
	Event_UNKNOWN Event_Type = 0
	Event_REQUEST Event_Type = 1
	Event_EVENT   Event_Type = 2
)

// Enum value maps for Event_Type.
var (
	Event_Type_name = map[int32]string{
		0: "UNKNOWN",
		1: "REQUEST",
		2: "EVENT",
	}
	Event_Type_value = map[string]int32{
		"UNKNOWN": 0,
		"REQUEST": 1,
		"EVENT":   2,
	}
)

func (x Event_Type) Enum() *Event_Type {
	p := new(Event_Type)
	*p = x
	return p
}

func (x Event_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Event_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_servers_proto_enumTypes[0].Descriptor()
}

func (Event_Type) Type() protoreflect.EnumType {
	return &file_v1_servers_proto_enumTypes[0]
}

func (x Event_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Event_Type.Descriptor instead.
func (Event_Type) EnumDescriptor() ([]byte, []int) {
	return file_v1_servers_proto_rawDescGZIP(), []int{6, 0}
}

// A enumerated list of possible actions on an event
type Event_Action int32

const (
	Event_NONE   Event_Action = 0
	Event_CREATE Event_Action = 1
	Event_UPDATE Event_Action = 2
	Event_DELETE Event_Action = 3
)

// Enum value maps for Event_Action.
var (
	Event_Action_name = map[int32]string{
		0: "NONE",
		1: "CREATE",
		2: "UPDATE",
		3: "DELETE",
	}
	Event_Action_value = map[string]int32{
		"NONE":   0,
		"CREATE": 1,
		"UPDATE": 2,
		"DELETE": 3,
	}
)

func (x Event_Action) Enum() *Event_Action {
	p := new(Event_Action)
	*p = x
	return p
}

func (x Event_Action) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Event_Action) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_servers_proto_enumTypes[1].Descriptor()
}

func (Event_Action) Type() protoreflect.EnumType {
	return &file_v1_servers_proto_enumTypes[1]
}

func (x Event_Action) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Event_Action.Descriptor instead.
func (Event_Action) EnumDescriptor() ([]byte, []int) {
	return file_v1_servers_proto_rawDescGZIP(), []int{6, 1}
}

// A enumerated list of possible status of servers
type Server_Status int32

const (
	Server_UNKNOWN             Server_Status = 0
	Server_PENDING             Server_Status = 1
	Server_READY               Server_Status = 2
	Server_MARKED_FOR_DELETION Server_Status = 3
)

// Enum value maps for Server_Status.
var (
	Server_Status_name = map[int32]string{
		0: "UNKNOWN",
		1: "PENDING",
		2: "READY",
		3: "MARKED_FOR_DELETION",
	}
	Server_Status_value = map[string]int32{
		"UNKNOWN":             0,
		"PENDING":             1,
		"READY":               2,
		"MARKED_FOR_DELETION": 3,
	}
)

func (x Server_Status) Enum() *Server_Status {
	p := new(Server_Status)
	*p = x
	return p
}

func (x Server_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Server_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_servers_proto_enumTypes[2].Descriptor()
}

func (Server_Status) Type() protoreflect.EnumType {
	return &file_v1_servers_proto_enumTypes[2]
}

func (x Server_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Server_Status.Descriptor instead.
func (Server_Status) EnumDescriptor() ([]byte, []int) {
	return file_v1_servers_proto_rawDescGZIP(), []int{7, 0}
}

// ChangeResponse for requests that results on a change on datastores
type ChangeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// a straightforward message on the result of the request
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ChangeResponse) Reset() {
	*x = ChangeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_servers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeResponse) ProtoMessage() {}

func (x *ChangeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_servers_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeResponse.ProtoReflect.Descriptor instead.
func (*ChangeResponse) Descriptor() ([]byte, []int) {
	return file_v1_servers_proto_rawDescGZIP(), []int{0}
}

func (x *ChangeResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// A standard message for methods that expects an single resource unique identifier
type ResourceIdentity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique name of the resource
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ResourceIdentity) Reset() {
	*x = ResourceIdentity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_servers_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceIdentity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceIdentity) ProtoMessage() {}

func (x *ResourceIdentity) ProtoReflect() protoreflect.Message {
	mi := &file_v1_servers_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceIdentity.ProtoReflect.Descriptor instead.
func (*ResourceIdentity) Descriptor() ([]byte, []int) {
	return file_v1_servers_proto_rawDescGZIP(), []int{1}
}

func (x *ResourceIdentity) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// A representation of a message to state a resource was changed
type ResourceChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The message to state the resource change
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ResourceChange) Reset() {
	*x = ResourceChange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_servers_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceChange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceChange) ProtoMessage() {}

func (x *ResourceChange) ProtoReflect() protoreflect.Message {
	mi := &file_v1_servers_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceChange.ProtoReflect.Descriptor instead.
func (*ResourceChange) Descriptor() ([]byte, []int) {
	return file_v1_servers_proto_rawDescGZIP(), []int{2}
}

func (x *ResourceChange) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// A representation of a resource list filter
type ListFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The maximum number of resources to return. The service may return fewer than
	// this value.
	// If unspecified, at most 10 resources will be returned.
	// The maximum value is 100; values above 100 result on failure.
	PageSize int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// A page token, received from a previous `List` call.
	// Provide this to retrieve the subsequent page.
	//
	// When paginating, all other parameters provided to `List` must match
	// the call that provided the page token.
	PageNumber int32 `protobuf:"varint,2,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	// Filters directly related to resource type being queried
	Filters map[string]string `protobuf:"bytes,3,rep,name=filters,proto3" json:"filters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ListFilter) Reset() {
	*x = ListFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_servers_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFilter) ProtoMessage() {}

func (x *ListFilter) ProtoReflect() protoreflect.Message {
	mi := &file_v1_servers_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFilter.ProtoReflect.Descriptor instead.
func (*ListFilter) Descriptor() ([]byte, []int) {
	return file_v1_servers_proto_rawDescGZIP(), []int{3}
}

func (x *ListFilter) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListFilter) GetPageNumber() int32 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *ListFilter) GetFilters() map[string]string {
	if x != nil {
		return x.Filters
	}
	return nil
}

// A representation of a resource patch
type Patch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique name of the resource
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The map of changes of the resource
	Changes map[string]string `protobuf:"bytes,2,rep,name=changes,proto3" json:"changes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Patch) Reset() {
	*x = Patch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_servers_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Patch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Patch) ProtoMessage() {}

func (x *Patch) ProtoReflect() protoreflect.Message {
	mi := &file_v1_servers_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Patch.ProtoReflect.Descriptor instead.
func (*Patch) Descriptor() ([]byte, []int) {
	return file_v1_servers_proto_rawDescGZIP(), []int{4}
}

func (x *Patch) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Patch) GetChanges() map[string]string {
	if x != nil {
		return x.Changes
	}
	return nil
}

// A representation of a message to state a error has occurred
type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The message to state the error cause
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_servers_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_v1_servers_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_v1_servers_proto_rawDescGZIP(), []int{5}
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// A standard message for events streamed on Watch methods
type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique name of the resource
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The type of the event
	Type Event_Type `protobuf:"varint,2,opt,name=type,proto3,enum=me.vniche.collective.v1.Event_Type" json:"type,omitempty"`
	// The action of the event
	Action Event_Action `protobuf:"varint,3,opt,name=action,proto3,enum=me.vniche.collective.v1.Event_Action" json:"action,omitempty"`
	// The resource of the event
	Resource []byte `protobuf:"bytes,4,opt,name=resource,proto3" json:"resource,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_servers_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_v1_servers_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_v1_servers_proto_rawDescGZIP(), []int{6}
}

func (x *Event) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Event) GetType() Event_Type {
	if x != nil {
		return x.Type
	}
	return Event_UNKNOWN
}

func (x *Event) GetAction() Event_Action {
	if x != nil {
		return x.Action
	}
	return Event_NONE
}

func (x *Event) GetResource() []byte {
	if x != nil {
		return x.Resource
	}
	return nil
}

// A represention of an server
type Server struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the server
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The brief description of the server
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// The status of the server
	Status Server_Status `protobuf:"varint,3,opt,name=status,proto3,enum=me.vniche.collective.v1.Server_Status" json:"status,omitempty"`
	// The creation timestamp of the server
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// The update timestamp of the server
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Server) Reset() {
	*x = Server{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_servers_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server) ProtoMessage() {}

func (x *Server) ProtoReflect() protoreflect.Message {
	mi := &file_v1_servers_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server.ProtoReflect.Descriptor instead.
func (*Server) Descriptor() ([]byte, []int) {
	return file_v1_servers_proto_rawDescGZIP(), []int{7}
}

func (x *Server) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Server) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Server) GetStatus() Server_Status {
	if x != nil {
		return x.Status
	}
	return Server_UNKNOWN
}

func (x *Server) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Server) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type PaginatedServers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of servers of the current page
	Servers []*Server `protobuf:"bytes,1,rep,name=servers,proto3" json:"servers,omitempty"`
	// The estimated total number of servers
	Total int32 `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *PaginatedServers) Reset() {
	*x = PaginatedServers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_servers_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaginatedServers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginatedServers) ProtoMessage() {}

func (x *PaginatedServers) ProtoReflect() protoreflect.Message {
	mi := &file_v1_servers_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginatedServers.ProtoReflect.Descriptor instead.
func (*PaginatedServers) Descriptor() ([]byte, []int) {
	return file_v1_servers_proto_rawDescGZIP(), []int{8}
}

func (x *PaginatedServers) GetServers() []*Server {
	if x != nil {
		return x.Servers
	}
	return nil
}

func (x *PaginatedServers) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_v1_servers_proto protoreflect.FileDescriptor

var file_v1_servers_proto_rawDesc = []byte{
	0x0a, 0x10, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x17, 0x6d, 0x65, 0x2e, 0x76, 0x6e, 0x69, 0x63, 0x68, 0x65, 0x2e, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2a, 0x0a, 0x0e,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x26, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x2a, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xd2, 0x01, 0x0a,
	0x0a, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70,
	0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x4a, 0x0a, 0x07, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x6d, 0x65, 0x2e,
	0x76, 0x6e, 0x69, 0x63, 0x68, 0x65, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x73, 0x1a, 0x3a, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x9e, 0x01, 0x0a, 0x05, 0x50, 0x61, 0x74, 0x63, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x45, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2b, 0x2e, 0x6d, 0x65, 0x2e, 0x76, 0x6e, 0x69, 0x63, 0x68, 0x65, 0x2e, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x74, 0x63, 0x68,
	0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x1a, 0x3a, 0x0a, 0x0c, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x21, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x94, 0x02, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x23, 0x2e, 0x6d, 0x65, 0x2e, 0x76, 0x6e, 0x69, 0x63, 0x68, 0x65, 0x2e, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x3d, 0x0a, 0x06,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x6d,
	0x65, 0x2e, 0x76, 0x6e, 0x69, 0x63, 0x68, 0x65, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x2b, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x56, 0x45,
	0x4e, 0x54, 0x10, 0x02, 0x22, 0x36, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x08,
	0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x52, 0x45, 0x41,
	0x54, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x02,
	0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x03, 0x22, 0xbc, 0x02, 0x0a,
	0x06, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3e, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e,
	0x6d, 0x65, 0x2e, 0x76, 0x6e, 0x69, 0x63, 0x68, 0x65, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0x46, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a,
	0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45,
	0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x45, 0x41, 0x44, 0x59,
	0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x4d, 0x41, 0x52, 0x4b, 0x45, 0x44, 0x5f, 0x46, 0x4f, 0x52,
	0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x03, 0x22, 0x63, 0x0a, 0x10, 0x50,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x12,
	0x39, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x6d, 0x65, 0x2e, 0x76, 0x6e, 0x69, 0x63, 0x68, 0x65, 0x2e, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x32, 0x8f, 0x04, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x12, 0x51, 0x0a, 0x03,
	0x47, 0x65, 0x74, 0x12, 0x29, 0x2e, 0x6d, 0x65, 0x2e, 0x76, 0x6e, 0x69, 0x63, 0x68, 0x65, 0x2e,
	0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x1a, 0x1f,
	0x2e, 0x6d, 0x65, 0x2e, 0x76, 0x6e, 0x69, 0x63, 0x68, 0x65, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12,
	0x56, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x23, 0x2e, 0x6d, 0x65, 0x2e, 0x76, 0x6e, 0x69,
	0x63, 0x68, 0x65, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x29, 0x2e, 0x6d,
	0x65, 0x2e, 0x76, 0x6e, 0x69, 0x63, 0x68, 0x65, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x64,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x12, 0x52, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x1f, 0x2e, 0x6d, 0x65, 0x2e, 0x76, 0x6e, 0x69, 0x63, 0x68, 0x65, 0x2e, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x1a, 0x27, 0x2e, 0x6d, 0x65, 0x2e, 0x76, 0x6e, 0x69, 0x63, 0x68, 0x65, 0x2e, 0x63,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x51, 0x0a, 0x06, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x6d, 0x65, 0x2e, 0x76, 0x6e, 0x69, 0x63, 0x68,
	0x65, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x61, 0x74, 0x63, 0x68, 0x1a, 0x27, 0x2e, 0x6d, 0x65, 0x2e, 0x76, 0x6e, 0x69, 0x63, 0x68,
	0x65, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x5c,
	0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x29, 0x2e, 0x6d, 0x65, 0x2e, 0x76, 0x6e,
	0x69, 0x63, 0x68, 0x65, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x1a, 0x27, 0x2e, 0x6d, 0x65, 0x2e, 0x76, 0x6e, 0x69, 0x63, 0x68, 0x65, 0x2e,
	0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x54, 0x0a, 0x05,
	0x57, 0x61, 0x74, 0x63, 0x68, 0x12, 0x29, 0x2e, 0x6d, 0x65, 0x2e, 0x76, 0x6e, 0x69, 0x63, 0x68,
	0x65, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x1a, 0x1e, 0x2e, 0x6d, 0x65, 0x2e, 0x76, 0x6e, 0x69, 0x63, 0x68, 0x65, 0x2e, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x30, 0x01, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x76, 0x6e, 0x69, 0x63, 0x68, 0x65, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x31, 0x3b, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_servers_proto_rawDescOnce sync.Once
	file_v1_servers_proto_rawDescData = file_v1_servers_proto_rawDesc
)

func file_v1_servers_proto_rawDescGZIP() []byte {
	file_v1_servers_proto_rawDescOnce.Do(func() {
		file_v1_servers_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_servers_proto_rawDescData)
	})
	return file_v1_servers_proto_rawDescData
}

var file_v1_servers_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_v1_servers_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_v1_servers_proto_goTypes = []interface{}{
	(Event_Type)(0),               // 0: me.vniche.collective.v1.Event.Type
	(Event_Action)(0),             // 1: me.vniche.collective.v1.Event.Action
	(Server_Status)(0),            // 2: me.vniche.collective.v1.Server.Status
	(*ChangeResponse)(nil),        // 3: me.vniche.collective.v1.ChangeResponse
	(*ResourceIdentity)(nil),      // 4: me.vniche.collective.v1.ResourceIdentity
	(*ResourceChange)(nil),        // 5: me.vniche.collective.v1.ResourceChange
	(*ListFilter)(nil),            // 6: me.vniche.collective.v1.ListFilter
	(*Patch)(nil),                 // 7: me.vniche.collective.v1.Patch
	(*Error)(nil),                 // 8: me.vniche.collective.v1.Error
	(*Event)(nil),                 // 9: me.vniche.collective.v1.Event
	(*Server)(nil),                // 10: me.vniche.collective.v1.Server
	(*PaginatedServers)(nil),      // 11: me.vniche.collective.v1.PaginatedServers
	nil,                           // 12: me.vniche.collective.v1.ListFilter.FiltersEntry
	nil,                           // 13: me.vniche.collective.v1.Patch.ChangesEntry
	(*timestamppb.Timestamp)(nil), // 14: google.protobuf.Timestamp
}
var file_v1_servers_proto_depIdxs = []int32{
	12, // 0: me.vniche.collective.v1.ListFilter.filters:type_name -> me.vniche.collective.v1.ListFilter.FiltersEntry
	13, // 1: me.vniche.collective.v1.Patch.changes:type_name -> me.vniche.collective.v1.Patch.ChangesEntry
	0,  // 2: me.vniche.collective.v1.Event.type:type_name -> me.vniche.collective.v1.Event.Type
	1,  // 3: me.vniche.collective.v1.Event.action:type_name -> me.vniche.collective.v1.Event.Action
	2,  // 4: me.vniche.collective.v1.Server.status:type_name -> me.vniche.collective.v1.Server.Status
	14, // 5: me.vniche.collective.v1.Server.created_at:type_name -> google.protobuf.Timestamp
	14, // 6: me.vniche.collective.v1.Server.updated_at:type_name -> google.protobuf.Timestamp
	10, // 7: me.vniche.collective.v1.PaginatedServers.servers:type_name -> me.vniche.collective.v1.Server
	4,  // 8: me.vniche.collective.v1.Servers.Get:input_type -> me.vniche.collective.v1.ResourceIdentity
	6,  // 9: me.vniche.collective.v1.Servers.List:input_type -> me.vniche.collective.v1.ListFilter
	10, // 10: me.vniche.collective.v1.Servers.Create:input_type -> me.vniche.collective.v1.Server
	7,  // 11: me.vniche.collective.v1.Servers.Update:input_type -> me.vniche.collective.v1.Patch
	4,  // 12: me.vniche.collective.v1.Servers.Delete:input_type -> me.vniche.collective.v1.ResourceIdentity
	4,  // 13: me.vniche.collective.v1.Servers.Watch:input_type -> me.vniche.collective.v1.ResourceIdentity
	10, // 14: me.vniche.collective.v1.Servers.Get:output_type -> me.vniche.collective.v1.Server
	11, // 15: me.vniche.collective.v1.Servers.List:output_type -> me.vniche.collective.v1.PaginatedServers
	5,  // 16: me.vniche.collective.v1.Servers.Create:output_type -> me.vniche.collective.v1.ResourceChange
	5,  // 17: me.vniche.collective.v1.Servers.Update:output_type -> me.vniche.collective.v1.ResourceChange
	5,  // 18: me.vniche.collective.v1.Servers.Delete:output_type -> me.vniche.collective.v1.ResourceChange
	9,  // 19: me.vniche.collective.v1.Servers.Watch:output_type -> me.vniche.collective.v1.Event
	14, // [14:20] is the sub-list for method output_type
	8,  // [8:14] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_v1_servers_proto_init() }
func file_v1_servers_proto_init() {
	if File_v1_servers_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_servers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeResponse); i {
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
		file_v1_servers_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceIdentity); i {
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
		file_v1_servers_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceChange); i {
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
		file_v1_servers_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFilter); i {
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
		file_v1_servers_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Patch); i {
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
		file_v1_servers_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
		file_v1_servers_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_v1_servers_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server); i {
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
		file_v1_servers_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaginatedServers); i {
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
			RawDescriptor: file_v1_servers_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_servers_proto_goTypes,
		DependencyIndexes: file_v1_servers_proto_depIdxs,
		EnumInfos:         file_v1_servers_proto_enumTypes,
		MessageInfos:      file_v1_servers_proto_msgTypes,
	}.Build()
	File_v1_servers_proto = out.File
	file_v1_servers_proto_rawDesc = nil
	file_v1_servers_proto_goTypes = nil
	file_v1_servers_proto_depIdxs = nil
}
