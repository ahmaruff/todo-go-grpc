// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: todo.proto

package models

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type TodoItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title     string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	DoneAt    *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=done_at,json=doneAt,proto3" json:"done_at,omitempty"`
	IsDone    *bool                  `protobuf:"varint,5,opt,name=is_done,json=isDone,proto3,oneof" json:"is_done,omitempty"`
}

func (x *TodoItem) Reset() {
	*x = TodoItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoItem) ProtoMessage() {}

func (x *TodoItem) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoItem.ProtoReflect.Descriptor instead.
func (*TodoItem) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{0}
}

func (x *TodoItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TodoItem) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *TodoItem) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *TodoItem) GetDoneAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DoneAt
	}
	return nil
}

func (x *TodoItem) GetIsDone() bool {
	if x != nil && x.IsDone != nil {
		return *x.IsDone
	}
	return false
}

type TodoItemId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TodoItemId) Reset() {
	*x = TodoItemId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoItemId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoItemId) ProtoMessage() {}

func (x *TodoItemId) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoItemId.ProtoReflect.Descriptor instead.
func (*TodoItemId) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{1}
}

func (x *TodoItemId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type TodoItems struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*TodoItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *TodoItems) Reset() {
	*x = TodoItems{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoItems) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoItems) ProtoMessage() {}

func (x *TodoItems) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoItems.ProtoReflect.Descriptor instead.
func (*TodoItems) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{2}
}

func (x *TodoItems) GetItems() []*TodoItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type CreateTodoItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *CreateTodoItemRequest) Reset() {
	*x = CreateTodoItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTodoItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTodoItemRequest) ProtoMessage() {}

func (x *CreateTodoItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTodoItemRequest.ProtoReflect.Descriptor instead.
func (*CreateTodoItemRequest) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{3}
}

func (x *CreateTodoItemRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type UpdateTodoItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *UpdateTodoItemRequest) Reset() {
	*x = UpdateTodoItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTodoItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTodoItemRequest) ProtoMessage() {}

func (x *UpdateTodoItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTodoItemRequest.ProtoReflect.Descriptor instead.
func (*UpdateTodoItemRequest) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateTodoItemRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateTodoItemRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type MakeTodoItemDoneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DoneAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=done_at,json=doneAt,proto3" json:"done_at,omitempty"`
}

func (x *MakeTodoItemDoneRequest) Reset() {
	*x = MakeTodoItemDoneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MakeTodoItemDoneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MakeTodoItemDoneRequest) ProtoMessage() {}

func (x *MakeTodoItemDoneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MakeTodoItemDoneRequest.ProtoReflect.Descriptor instead.
func (*MakeTodoItemDoneRequest) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{5}
}

func (x *MakeTodoItemDoneRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MakeTodoItemDoneRequest) GetDoneAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DoneAt
	}
	return nil
}

var File_todo_proto protoreflect.FileDescriptor

var file_todo_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xca, 0x01, 0x0a, 0x08, 0x54,
	0x6f, 0x64, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x39, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x33, 0x0a, 0x07, 0x64, 0x6f, 0x6e, 0x65,
	0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x06, 0x64, 0x6f, 0x6e, 0x65, 0x41, 0x74, 0x12, 0x1c, 0x0a,
	0x07, 0x69, 0x73, 0x5f, 0x64, 0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00,
	0x52, 0x06, 0x69, 0x73, 0x44, 0x6f, 0x6e, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f,
	0x69, 0x73, 0x5f, 0x64, 0x6f, 0x6e, 0x65, 0x22, 0x1c, 0x0a, 0x0a, 0x54, 0x6f, 0x64, 0x6f, 0x49,
	0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2c, 0x0a, 0x09, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x12, 0x1f, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x09, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x22, 0x2d, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64,
	0x6f, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x22, 0x3d, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x22, 0x5e, 0x0a, 0x17, 0x4d, 0x61, 0x6b, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x74, 0x65,
	0x6d, 0x44, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x33, 0x0a, 0x07,
	0x64, 0x6f, 0x6e, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x06, 0x64, 0x6f, 0x6e, 0x65, 0x41,
	0x74, 0x32, 0xbc, 0x02, 0x0a, 0x04, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x25, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0b, 0x2e, 0x54, 0x6f, 0x64, 0x6f,
	0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x1a, 0x09, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x74, 0x65,
	0x6d, 0x12, 0x33, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0a, 0x2e, 0x54, 0x6f, 0x64,
	0x6f, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x33, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x6f, 0x64, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x16, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x09, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x33, 0x0a, 0x0e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x16, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x74, 0x65, 0x6d,
	0x12, 0x35, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x0b, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x37, 0x0a, 0x10, 0x4d, 0x61, 0x6b, 0x65, 0x54,
	0x6f, 0x64, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x44, 0x6f, 0x6e, 0x65, 0x12, 0x0b, 0x2e, 0x54, 0x6f,
	0x64, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x42, 0x17, 0x5a, 0x15, 0x74, 0x6f, 0x64, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x6f,
	0x64, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_todo_proto_rawDescOnce sync.Once
	file_todo_proto_rawDescData = file_todo_proto_rawDesc
)

func file_todo_proto_rawDescGZIP() []byte {
	file_todo_proto_rawDescOnce.Do(func() {
		file_todo_proto_rawDescData = protoimpl.X.CompressGZIP(file_todo_proto_rawDescData)
	})
	return file_todo_proto_rawDescData
}

var file_todo_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_todo_proto_goTypes = []interface{}{
	(*TodoItem)(nil),                // 0: TodoItem
	(*TodoItemId)(nil),              // 1: TodoItemId
	(*TodoItems)(nil),               // 2: TodoItems
	(*CreateTodoItemRequest)(nil),   // 3: CreateTodoItemRequest
	(*UpdateTodoItemRequest)(nil),   // 4: UpdateTodoItemRequest
	(*MakeTodoItemDoneRequest)(nil), // 5: MakeTodoItemDoneRequest
	(*timestamppb.Timestamp)(nil),   // 6: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),           // 7: google.protobuf.Empty
}
var file_todo_proto_depIdxs = []int32{
	6,  // 0: TodoItem.created_at:type_name -> google.protobuf.Timestamp
	6,  // 1: TodoItem.done_at:type_name -> google.protobuf.Timestamp
	0,  // 2: TodoItems.items:type_name -> TodoItem
	6,  // 3: MakeTodoItemDoneRequest.done_at:type_name -> google.protobuf.Timestamp
	1,  // 4: Todo.GetTodoItem:input_type -> TodoItemId
	7,  // 5: Todo.ListTodoItems:input_type -> google.protobuf.Empty
	3,  // 6: Todo.CreateTodoItem:input_type -> CreateTodoItemRequest
	4,  // 7: Todo.UpdateTodoItem:input_type -> UpdateTodoItemRequest
	1,  // 8: Todo.DeleteTodoItem:input_type -> TodoItemId
	1,  // 9: Todo.MakeTodoItemDone:input_type -> TodoItemId
	0,  // 10: Todo.GetTodoItem:output_type -> TodoItem
	2,  // 11: Todo.ListTodoItems:output_type -> TodoItems
	0,  // 12: Todo.CreateTodoItem:output_type -> TodoItem
	0,  // 13: Todo.UpdateTodoItem:output_type -> TodoItem
	7,  // 14: Todo.DeleteTodoItem:output_type -> google.protobuf.Empty
	7,  // 15: Todo.MakeTodoItemDone:output_type -> google.protobuf.Empty
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_todo_proto_init() }
func file_todo_proto_init() {
	if File_todo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_todo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoItem); i {
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
		file_todo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoItemId); i {
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
		file_todo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoItems); i {
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
		file_todo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTodoItemRequest); i {
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
		file_todo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTodoItemRequest); i {
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
		file_todo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MakeTodoItemDoneRequest); i {
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
	file_todo_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_todo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_todo_proto_goTypes,
		DependencyIndexes: file_todo_proto_depIdxs,
		MessageInfos:      file_todo_proto_msgTypes,
	}.Build()
	File_todo_proto = out.File
	file_todo_proto_rawDesc = nil
	file_todo_proto_goTypes = nil
	file_todo_proto_depIdxs = nil
}
