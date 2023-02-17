// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: proto/todotask/v1/todotask.proto

package __todotask

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

type TodoTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Date        uint64 `protobuf:"varint,4,opt,name=date,proto3" json:"date,omitempty"`
	IsCompleted bool   `protobuf:"varint,5,opt,name=is_completed,json=isCompleted,proto3" json:"is_completed,omitempty"`
}

func (x *TodoTask) Reset() {
	*x = TodoTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todotask_v1_todotask_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoTask) ProtoMessage() {}

func (x *TodoTask) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todotask_v1_todotask_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoTask.ProtoReflect.Descriptor instead.
func (*TodoTask) Descriptor() ([]byte, []int) {
	return file_proto_todotask_v1_todotask_proto_rawDescGZIP(), []int{0}
}

func (x *TodoTask) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TodoTask) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *TodoTask) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *TodoTask) GetDate() uint64 {
	if x != nil {
		return x.Date
	}
	return 0
}

func (x *TodoTask) GetIsCompleted() bool {
	if x != nil {
		return x.IsCompleted
	}
	return false
}

type SaveTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task *TodoTask `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
}

func (x *SaveTaskRequest) Reset() {
	*x = SaveTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todotask_v1_todotask_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveTaskRequest) ProtoMessage() {}

func (x *SaveTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todotask_v1_todotask_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveTaskRequest.ProtoReflect.Descriptor instead.
func (*SaveTaskRequest) Descriptor() ([]byte, []int) {
	return file_proto_todotask_v1_todotask_proto_rawDescGZIP(), []int{1}
}

func (x *SaveTaskRequest) GetTask() *TodoTask {
	if x != nil {
		return x.Task
	}
	return nil
}

type SaveTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId uint32 `protobuf:"varint,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
}

func (x *SaveTaskResponse) Reset() {
	*x = SaveTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todotask_v1_todotask_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveTaskResponse) ProtoMessage() {}

func (x *SaveTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todotask_v1_todotask_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveTaskResponse.ProtoReflect.Descriptor instead.
func (*SaveTaskResponse) Descriptor() ([]byte, []int) {
	return file_proto_todotask_v1_todotask_proto_rawDescGZIP(), []int{2}
}

func (x *SaveTaskResponse) GetTaskId() uint32 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

type TasksListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title    string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	IsDone   bool   `protobuf:"varint,2,opt,name=is_done,json=isDone,proto3" json:"is_done,omitempty"`
	IsUndone bool   `protobuf:"varint,3,opt,name=is_undone,json=isUndone,proto3" json:"is_undone,omitempty"`
}

func (x *TasksListRequest) Reset() {
	*x = TasksListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todotask_v1_todotask_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TasksListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TasksListRequest) ProtoMessage() {}

func (x *TasksListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todotask_v1_todotask_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TasksListRequest.ProtoReflect.Descriptor instead.
func (*TasksListRequest) Descriptor() ([]byte, []int) {
	return file_proto_todotask_v1_todotask_proto_rawDescGZIP(), []int{3}
}

func (x *TasksListRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *TasksListRequest) GetIsDone() bool {
	if x != nil {
		return x.IsDone
	}
	return false
}

func (x *TasksListRequest) GetIsUndone() bool {
	if x != nil {
		return x.IsUndone
	}
	return false
}

type TasksListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tasks []*TodoTask `protobuf:"bytes,2,rep,name=tasks,proto3" json:"tasks,omitempty"`
}

func (x *TasksListResponse) Reset() {
	*x = TasksListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todotask_v1_todotask_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TasksListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TasksListResponse) ProtoMessage() {}

func (x *TasksListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todotask_v1_todotask_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TasksListResponse.ProtoReflect.Descriptor instead.
func (*TasksListResponse) Descriptor() ([]byte, []int) {
	return file_proto_todotask_v1_todotask_proto_rawDescGZIP(), []int{4}
}

func (x *TasksListResponse) GetTasks() []*TodoTask {
	if x != nil {
		return x.Tasks
	}
	return nil
}

type DeleteTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId uint32 `protobuf:"varint,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
}

func (x *DeleteTaskRequest) Reset() {
	*x = DeleteTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todotask_v1_todotask_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTaskRequest) ProtoMessage() {}

func (x *DeleteTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todotask_v1_todotask_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTaskRequest.ProtoReflect.Descriptor instead.
func (*DeleteTaskRequest) Descriptor() ([]byte, []int) {
	return file_proto_todotask_v1_todotask_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteTaskRequest) GetTaskId() uint32 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

type DeleteTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteTaskResponse) Reset() {
	*x = DeleteTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todotask_v1_todotask_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTaskResponse) ProtoMessage() {}

func (x *DeleteTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todotask_v1_todotask_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTaskResponse.ProtoReflect.Descriptor instead.
func (*DeleteTaskResponse) Descriptor() ([]byte, []int) {
	return file_proto_todotask_v1_todotask_proto_rawDescGZIP(), []int{6}
}

type UpdateTaskStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId uint32 `protobuf:"varint,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	IsDone bool   `protobuf:"varint,2,opt,name=is_done,json=isDone,proto3" json:"is_done,omitempty"`
}

func (x *UpdateTaskStatusRequest) Reset() {
	*x = UpdateTaskStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todotask_v1_todotask_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTaskStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTaskStatusRequest) ProtoMessage() {}

func (x *UpdateTaskStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todotask_v1_todotask_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTaskStatusRequest.ProtoReflect.Descriptor instead.
func (*UpdateTaskStatusRequest) Descriptor() ([]byte, []int) {
	return file_proto_todotask_v1_todotask_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateTaskStatusRequest) GetTaskId() uint32 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

func (x *UpdateTaskStatusRequest) GetIsDone() bool {
	if x != nil {
		return x.IsDone
	}
	return false
}

type UpdateTaskStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateTaskStatusResponse) Reset() {
	*x = UpdateTaskStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todotask_v1_todotask_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTaskStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTaskStatusResponse) ProtoMessage() {}

func (x *UpdateTaskStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todotask_v1_todotask_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTaskStatusResponse.ProtoReflect.Descriptor instead.
func (*UpdateTaskStatusResponse) Descriptor() ([]byte, []int) {
	return file_proto_todotask_v1_todotask_proto_rawDescGZIP(), []int{8}
}

var File_proto_todotask_v1_todotask_proto protoreflect.FileDescriptor

var file_proto_todotask_v1_todotask_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x74, 0x61, 0x73, 0x6b,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x74, 0x61,
	0x73, 0x6b, 0x2e, 0x76, 0x31, 0x22, 0x89, 0x01, 0x0a, 0x08, 0x54, 0x6f, 0x64, 0x6f, 0x54, 0x61,
	0x73, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x22, 0x42, 0x0a, 0x0f, 0x53, 0x61, 0x76, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x74,
	0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x04, 0x74, 0x61, 0x73, 0x6b, 0x22, 0x2b, 0x0a, 0x10, 0x53, 0x61, 0x76, 0x65, 0x54, 0x61, 0x73,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73,
	0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b,
	0x49, 0x64, 0x22, 0x5e, 0x0a, 0x10, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x17, 0x0a, 0x07,
	0x69, 0x73, 0x5f, 0x64, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69,
	0x73, 0x44, 0x6f, 0x6e, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x75, 0x6e, 0x64, 0x6f,
	0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x55, 0x6e, 0x64, 0x6f,
	0x6e, 0x65, 0x22, 0x46, 0x0a, 0x11, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74,
	0x6f, 0x64, 0x6f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x22, 0x2c, 0x0a, 0x11, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4b,
	0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73,
	0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b,
	0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x64, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x44, 0x6f, 0x6e, 0x65, 0x22, 0x1a, 0x0a, 0x18, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x8e, 0x03, 0x0a, 0x0f, 0x54, 0x6f, 0x64, 0x6f,
	0x54, 0x61, 0x73, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x08, 0x53,
	0x61, 0x76, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x74, 0x6f, 0x64, 0x6f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x76, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x61, 0x76, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x58, 0x0a, 0x09, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x74, 0x61, 0x73, 0x6b,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x6f, 0x64,
	0x6f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5b, 0x0a, 0x0a,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x24, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x74, 0x61, 0x73,
	0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6d, 0x0a, 0x10, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x97, 0x01, 0x0a, 0x15, 0x63, 0x6f, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x74, 0x61, 0x73, 0x6b, 0x2e,
	0x76, 0x31, 0x42, 0x0d, 0x54, 0x6f, 0x64, 0x6f, 0x74, 0x61, 0x73, 0x6b, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x09, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x74, 0x61, 0x73, 0x6b, 0xa2, 0x02,
	0x03, 0x50, 0x54, 0x58, 0xaa, 0x02, 0x11, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x64,
	0x6f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x11, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x5c, 0x54, 0x6f, 0x64, 0x6f, 0x74, 0x61, 0x73, 0x6b, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1d, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x54, 0x6f, 0x64, 0x6f, 0x74, 0x61, 0x73, 0x6b, 0x5c, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x13, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x3a, 0x54, 0x6f, 0x64, 0x6f, 0x74, 0x61, 0x73, 0x6b, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_todotask_v1_todotask_proto_rawDescOnce sync.Once
	file_proto_todotask_v1_todotask_proto_rawDescData = file_proto_todotask_v1_todotask_proto_rawDesc
)

func file_proto_todotask_v1_todotask_proto_rawDescGZIP() []byte {
	file_proto_todotask_v1_todotask_proto_rawDescOnce.Do(func() {
		file_proto_todotask_v1_todotask_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_todotask_v1_todotask_proto_rawDescData)
	})
	return file_proto_todotask_v1_todotask_proto_rawDescData
}

var file_proto_todotask_v1_todotask_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_todotask_v1_todotask_proto_goTypes = []interface{}{
	(*TodoTask)(nil),                 // 0: proto.todotask.v1.TodoTask
	(*SaveTaskRequest)(nil),          // 1: proto.todotask.v1.SaveTaskRequest
	(*SaveTaskResponse)(nil),         // 2: proto.todotask.v1.SaveTaskResponse
	(*TasksListRequest)(nil),         // 3: proto.todotask.v1.TasksListRequest
	(*TasksListResponse)(nil),        // 4: proto.todotask.v1.TasksListResponse
	(*DeleteTaskRequest)(nil),        // 5: proto.todotask.v1.DeleteTaskRequest
	(*DeleteTaskResponse)(nil),       // 6: proto.todotask.v1.DeleteTaskResponse
	(*UpdateTaskStatusRequest)(nil),  // 7: proto.todotask.v1.UpdateTaskStatusRequest
	(*UpdateTaskStatusResponse)(nil), // 8: proto.todotask.v1.UpdateTaskStatusResponse
}
var file_proto_todotask_v1_todotask_proto_depIdxs = []int32{
	0, // 0: proto.todotask.v1.SaveTaskRequest.task:type_name -> proto.todotask.v1.TodoTask
	0, // 1: proto.todotask.v1.TasksListResponse.tasks:type_name -> proto.todotask.v1.TodoTask
	1, // 2: proto.todotask.v1.TodoTaskService.SaveTask:input_type -> proto.todotask.v1.SaveTaskRequest
	3, // 3: proto.todotask.v1.TodoTaskService.TasksList:input_type -> proto.todotask.v1.TasksListRequest
	5, // 4: proto.todotask.v1.TodoTaskService.DeleteTask:input_type -> proto.todotask.v1.DeleteTaskRequest
	7, // 5: proto.todotask.v1.TodoTaskService.UpdateTaskStatus:input_type -> proto.todotask.v1.UpdateTaskStatusRequest
	2, // 6: proto.todotask.v1.TodoTaskService.SaveTask:output_type -> proto.todotask.v1.SaveTaskResponse
	4, // 7: proto.todotask.v1.TodoTaskService.TasksList:output_type -> proto.todotask.v1.TasksListResponse
	6, // 8: proto.todotask.v1.TodoTaskService.DeleteTask:output_type -> proto.todotask.v1.DeleteTaskResponse
	8, // 9: proto.todotask.v1.TodoTaskService.UpdateTaskStatus:output_type -> proto.todotask.v1.UpdateTaskStatusResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_todotask_v1_todotask_proto_init() }
func file_proto_todotask_v1_todotask_proto_init() {
	if File_proto_todotask_v1_todotask_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_todotask_v1_todotask_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoTask); i {
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
		file_proto_todotask_v1_todotask_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveTaskRequest); i {
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
		file_proto_todotask_v1_todotask_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveTaskResponse); i {
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
		file_proto_todotask_v1_todotask_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TasksListRequest); i {
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
		file_proto_todotask_v1_todotask_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TasksListResponse); i {
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
		file_proto_todotask_v1_todotask_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTaskRequest); i {
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
		file_proto_todotask_v1_todotask_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTaskResponse); i {
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
		file_proto_todotask_v1_todotask_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTaskStatusRequest); i {
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
		file_proto_todotask_v1_todotask_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTaskStatusResponse); i {
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
			RawDescriptor: file_proto_todotask_v1_todotask_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_todotask_v1_todotask_proto_goTypes,
		DependencyIndexes: file_proto_todotask_v1_todotask_proto_depIdxs,
		MessageInfos:      file_proto_todotask_v1_todotask_proto_msgTypes,
	}.Build()
	File_proto_todotask_v1_todotask_proto = out.File
	file_proto_todotask_v1_todotask_proto_rawDesc = nil
	file_proto_todotask_v1_todotask_proto_goTypes = nil
	file_proto_todotask_v1_todotask_proto_depIdxs = nil
}
