// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: proto/todotask/v1/todotask.proto

package __todotask

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type TasksCompletedFilter int32

const (
	TasksCompletedFilter_TASKS_COMPLETED_FILTER_ALL    TasksCompletedFilter = 0
	TasksCompletedFilter_TASKS_COMPLETED_FILTER_DONE   TasksCompletedFilter = 1
	TasksCompletedFilter_TASKS_COMPLETED_FILTER_UNDONE TasksCompletedFilter = 2
)

// Enum value maps for TasksCompletedFilter.
var (
	TasksCompletedFilter_name = map[int32]string{
		0: "TASKS_COMPLETED_FILTER_ALL",
		1: "TASKS_COMPLETED_FILTER_DONE",
		2: "TASKS_COMPLETED_FILTER_UNDONE",
	}
	TasksCompletedFilter_value = map[string]int32{
		"TASKS_COMPLETED_FILTER_ALL":    0,
		"TASKS_COMPLETED_FILTER_DONE":   1,
		"TASKS_COMPLETED_FILTER_UNDONE": 2,
	}
)

func (x TasksCompletedFilter) Enum() *TasksCompletedFilter {
	p := new(TasksCompletedFilter)
	*p = x
	return p
}

func (x TasksCompletedFilter) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TasksCompletedFilter) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_todotask_v1_todotask_proto_enumTypes[0].Descriptor()
}

func (TasksCompletedFilter) Type() protoreflect.EnumType {
	return &file_proto_todotask_v1_todotask_proto_enumTypes[0]
}

func (x TasksCompletedFilter) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TasksCompletedFilter.Descriptor instead.
func (TasksCompletedFilter) EnumDescriptor() ([]byte, []int) {
	return file_proto_todotask_v1_todotask_proto_rawDescGZIP(), []int{0}
}

type ResponseCode int32

const (
	ResponseCode_RESPONSE_CODE_UNSPECIFIED ResponseCode = 0
	ResponseCode_RESPONSE_CODE_SUCCESS     ResponseCode = 1
	ResponseCode_RESPONSE_CODE_NOT_FOUND   ResponseCode = 2
	ResponseCode_RESPONSE_CODE_FAILED      ResponseCode = 3
)

// Enum value maps for ResponseCode.
var (
	ResponseCode_name = map[int32]string{
		0: "RESPONSE_CODE_UNSPECIFIED",
		1: "RESPONSE_CODE_SUCCESS",
		2: "RESPONSE_CODE_NOT_FOUND",
		3: "RESPONSE_CODE_FAILED",
	}
	ResponseCode_value = map[string]int32{
		"RESPONSE_CODE_UNSPECIFIED": 0,
		"RESPONSE_CODE_SUCCESS":     1,
		"RESPONSE_CODE_NOT_FOUND":   2,
		"RESPONSE_CODE_FAILED":      3,
	}
)

func (x ResponseCode) Enum() *ResponseCode {
	p := new(ResponseCode)
	*p = x
	return p
}

func (x ResponseCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResponseCode) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_todotask_v1_todotask_proto_enumTypes[1].Descriptor()
}

func (ResponseCode) Type() protoreflect.EnumType {
	return &file_proto_todotask_v1_todotask_proto_enumTypes[1]
}

func (x ResponseCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResponseCode.Descriptor instead.
func (ResponseCode) EnumDescriptor() ([]byte, []int) {
	return file_proto_todotask_v1_todotask_proto_rawDescGZIP(), []int{1}
}

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

type CreateTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task *TodoTask `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
}

func (x *CreateTaskRequest) Reset() {
	*x = CreateTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todotask_v1_todotask_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaskRequest) ProtoMessage() {}

func (x *CreateTaskRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CreateTaskRequest.ProtoReflect.Descriptor instead.
func (*CreateTaskRequest) Descriptor() ([]byte, []int) {
	return file_proto_todotask_v1_todotask_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTaskRequest) GetTask() *TodoTask {
	if x != nil {
		return x.Task
	}
	return nil
}

type CreateTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId uint32 `protobuf:"varint,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
}

func (x *CreateTaskResponse) Reset() {
	*x = CreateTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todotask_v1_todotask_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaskResponse) ProtoMessage() {}

func (x *CreateTaskResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CreateTaskResponse.ProtoReflect.Descriptor instead.
func (*CreateTaskResponse) Descriptor() ([]byte, []int) {
	return file_proto_todotask_v1_todotask_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTaskResponse) GetTaskId() uint32 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

type TasksListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string               `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	IsCompleted TasksCompletedFilter `protobuf:"varint,2,opt,name=is_completed,json=isCompleted,proto3,enum=proto.todotask.v1.TasksCompletedFilter" json:"is_completed,omitempty"`
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

func (x *TasksListRequest) GetIsCompleted() TasksCompletedFilter {
	if x != nil {
		return x.IsCompleted
	}
	return TasksCompletedFilter_TASKS_COMPLETED_FILTER_ALL
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

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
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

func (x *DeleteTaskRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code ResponseCode `protobuf:"varint,1,opt,name=code,proto3,enum=proto.todotask.v1.ResponseCode" json:"code,omitempty"`
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

func (x *DeleteTaskResponse) GetCode() ResponseCode {
	if x != nil {
		return x.Code
	}
	return ResponseCode_RESPONSE_CODE_UNSPECIFIED
}

type UpdateTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task *TodoTask `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
}

func (x *UpdateTaskRequest) Reset() {
	*x = UpdateTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todotask_v1_todotask_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTaskRequest) ProtoMessage() {}

func (x *UpdateTaskRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UpdateTaskRequest.ProtoReflect.Descriptor instead.
func (*UpdateTaskRequest) Descriptor() ([]byte, []int) {
	return file_proto_todotask_v1_todotask_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateTaskRequest) GetTask() *TodoTask {
	if x != nil {
		return x.Task
	}
	return nil
}

type UpdateTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code ResponseCode `protobuf:"varint,1,opt,name=code,proto3,enum=proto.todotask.v1.ResponseCode" json:"code,omitempty"`
}

func (x *UpdateTaskResponse) Reset() {
	*x = UpdateTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todotask_v1_todotask_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTaskResponse) ProtoMessage() {}

func (x *UpdateTaskResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UpdateTaskResponse.ProtoReflect.Descriptor instead.
func (*UpdateTaskResponse) Descriptor() ([]byte, []int) {
	return file_proto_todotask_v1_todotask_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateTaskResponse) GetCode() ResponseCode {
	if x != nil {
		return x.Code
	}
	return ResponseCode_RESPONSE_CODE_UNSPECIFIED
}

var File_proto_todotask_v1_todotask_proto protoreflect.FileDescriptor

var file_proto_todotask_v1_todotask_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x74, 0x61, 0x73, 0x6b,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x74, 0x61,
	0x73, 0x6b, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x89, 0x01, 0x0a, 0x08, 0x54, 0x6f, 0x64, 0x6f, 0x54, 0x61, 0x73, 0x6b,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x69, 0x73, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x22,
	0x44, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x74,
	0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x04, 0x74, 0x61, 0x73, 0x6b, 0x22, 0x2d, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x74,
	0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x74, 0x61,
	0x73, 0x6b, 0x49, 0x64, 0x22, 0x74, 0x0a, 0x10, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x4a,
	0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x6f, 0x64,
	0x6f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x0b, 0x69,
	0x73, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x22, 0x46, 0x0a, 0x11, 0x54, 0x61,
	0x73, 0x6b, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x31, 0x0a, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x74, 0x61, 0x73, 0x6b, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x05, 0x74, 0x61, 0x73,
	0x6b, 0x73, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x49, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x22, 0x44, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x6f,
	0x64, 0x6f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x22, 0x49, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x2a, 0x7a, 0x0a, 0x14, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x1a, 0x54,
	0x41, 0x53, 0x4b, 0x53, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x5f, 0x46,
	0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x41, 0x4c, 0x4c, 0x10, 0x00, 0x12, 0x1f, 0x0a, 0x1b, 0x54,
	0x41, 0x53, 0x4b, 0x53, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x5f, 0x46,
	0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x44, 0x4f, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x21, 0x0a, 0x1d,
	0x54, 0x41, 0x53, 0x4b, 0x53, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x5f,
	0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x55, 0x4e, 0x44, 0x4f, 0x4e, 0x45, 0x10, 0x02, 0x2a,
	0x7f, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x1d, 0x0a, 0x19, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x19,
	0x0a, 0x15, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f,
	0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x52, 0x45, 0x53,
	0x50, 0x4f, 0x4e, 0x53, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46,
	0x4f, 0x55, 0x4e, 0x44, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e,
	0x53, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x03,
	0x32, 0x88, 0x04, 0x0a, 0x0f, 0x54, 0x6f, 0x64, 0x6f, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x7d, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x12, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x74,
	0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x3a, 0x01, 0x2a, 0x22, 0x17, 0x2f, 0x76, 0x31, 0x2f,
	0x74, 0x6f, 0x64, 0x6f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x74,
	0x61, 0x73, 0x6b, 0x12, 0x76, 0x0a, 0x09, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x74, 0x61, 0x73,
	0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x6f,
	0x64, 0x6f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x74, 0x61, 0x73,
	0x6b, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x7f, 0x0a, 0x0a, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x25, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x74, 0x61, 0x73, 0x6b,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x2a, 0x1c,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x7d, 0x0a, 0x0a,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x24, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x74, 0x61, 0x73,
	0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x3a,
	0x01, 0x2a, 0x1a, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x74, 0x61, 0x73, 0x6b,
	0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x74, 0x61, 0x73, 0x6b, 0x42, 0x97, 0x01, 0x0a, 0x15,
	0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x74, 0x61,
	0x73, 0x6b, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x54, 0x6f, 0x64, 0x6f, 0x74, 0x61, 0x73, 0x6b, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x09, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x74, 0x61, 0x73,
	0x6b, 0xa2, 0x02, 0x03, 0x50, 0x54, 0x58, 0xaa, 0x02, 0x11, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x54, 0x6f, 0x64, 0x6f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x11, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x5c, 0x54, 0x6f, 0x64, 0x6f, 0x74, 0x61, 0x73, 0x6b, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x1d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x54, 0x6f, 0x64, 0x6f, 0x74, 0x61, 0x73, 0x6b,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x13, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x3a, 0x54, 0x6f, 0x64, 0x6f, 0x74, 0x61, 0x73,
	0x6b, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_proto_todotask_v1_todotask_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_todotask_v1_todotask_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_todotask_v1_todotask_proto_goTypes = []interface{}{
	(TasksCompletedFilter)(0),  // 0: proto.todotask.v1.TasksCompletedFilter
	(ResponseCode)(0),          // 1: proto.todotask.v1.ResponseCode
	(*TodoTask)(nil),           // 2: proto.todotask.v1.TodoTask
	(*CreateTaskRequest)(nil),  // 3: proto.todotask.v1.CreateTaskRequest
	(*CreateTaskResponse)(nil), // 4: proto.todotask.v1.CreateTaskResponse
	(*TasksListRequest)(nil),   // 5: proto.todotask.v1.TasksListRequest
	(*TasksListResponse)(nil),  // 6: proto.todotask.v1.TasksListResponse
	(*DeleteTaskRequest)(nil),  // 7: proto.todotask.v1.DeleteTaskRequest
	(*DeleteTaskResponse)(nil), // 8: proto.todotask.v1.DeleteTaskResponse
	(*UpdateTaskRequest)(nil),  // 9: proto.todotask.v1.UpdateTaskRequest
	(*UpdateTaskResponse)(nil), // 10: proto.todotask.v1.UpdateTaskResponse
}
var file_proto_todotask_v1_todotask_proto_depIdxs = []int32{
	2,  // 0: proto.todotask.v1.CreateTaskRequest.task:type_name -> proto.todotask.v1.TodoTask
	0,  // 1: proto.todotask.v1.TasksListRequest.is_completed:type_name -> proto.todotask.v1.TasksCompletedFilter
	2,  // 2: proto.todotask.v1.TasksListResponse.tasks:type_name -> proto.todotask.v1.TodoTask
	1,  // 3: proto.todotask.v1.DeleteTaskResponse.code:type_name -> proto.todotask.v1.ResponseCode
	2,  // 4: proto.todotask.v1.UpdateTaskRequest.task:type_name -> proto.todotask.v1.TodoTask
	1,  // 5: proto.todotask.v1.UpdateTaskResponse.code:type_name -> proto.todotask.v1.ResponseCode
	3,  // 6: proto.todotask.v1.TodoTaskService.CreateTask:input_type -> proto.todotask.v1.CreateTaskRequest
	5,  // 7: proto.todotask.v1.TodoTaskService.TasksList:input_type -> proto.todotask.v1.TasksListRequest
	7,  // 8: proto.todotask.v1.TodoTaskService.DeleteTask:input_type -> proto.todotask.v1.DeleteTaskRequest
	9,  // 9: proto.todotask.v1.TodoTaskService.UpdateTask:input_type -> proto.todotask.v1.UpdateTaskRequest
	4,  // 10: proto.todotask.v1.TodoTaskService.CreateTask:output_type -> proto.todotask.v1.CreateTaskResponse
	6,  // 11: proto.todotask.v1.TodoTaskService.TasksList:output_type -> proto.todotask.v1.TasksListResponse
	8,  // 12: proto.todotask.v1.TodoTaskService.DeleteTask:output_type -> proto.todotask.v1.DeleteTaskResponse
	10, // 13: proto.todotask.v1.TodoTaskService.UpdateTask:output_type -> proto.todotask.v1.UpdateTaskResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
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
			switch v := v.(*CreateTaskRequest); i {
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
			switch v := v.(*CreateTaskResponse); i {
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
			switch v := v.(*UpdateTaskRequest); i {
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
			switch v := v.(*UpdateTaskResponse); i {
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
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_todotask_v1_todotask_proto_goTypes,
		DependencyIndexes: file_proto_todotask_v1_todotask_proto_depIdxs,
		EnumInfos:         file_proto_todotask_v1_todotask_proto_enumTypes,
		MessageInfos:      file_proto_todotask_v1_todotask_proto_msgTypes,
	}.Build()
	File_proto_todotask_v1_todotask_proto = out.File
	file_proto_todotask_v1_todotask_proto_rawDesc = nil
	file_proto_todotask_v1_todotask_proto_goTypes = nil
	file_proto_todotask_v1_todotask_proto_depIdxs = nil
}
