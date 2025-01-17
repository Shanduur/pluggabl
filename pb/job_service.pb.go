// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.6.1
// source: job_service.proto

package pb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type JobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//	*JobRequest_Job
	//	*JobRequest_ChunkData
	Data isJobRequest_Data `protobuf_oneof:"data"`
}

func (x *JobRequest) Reset() {
	*x = JobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobRequest) ProtoMessage() {}

func (x *JobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_job_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobRequest.ProtoReflect.Descriptor instead.
func (*JobRequest) Descriptor() ([]byte, []int) {
	return file_job_service_proto_rawDescGZIP(), []int{0}
}

func (m *JobRequest) GetData() isJobRequest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *JobRequest) GetJob() *Job {
	if x, ok := x.GetData().(*JobRequest_Job); ok {
		return x.Job
	}
	return nil
}

func (x *JobRequest) GetChunkData() *Chunk {
	if x, ok := x.GetData().(*JobRequest_ChunkData); ok {
		return x.ChunkData
	}
	return nil
}

type isJobRequest_Data interface {
	isJobRequest_Data()
}

type JobRequest_Job struct {
	Job *Job `protobuf:"bytes,1,opt,name=job,proto3,oneof"`
}

type JobRequest_ChunkData struct {
	ChunkData *Chunk `protobuf:"bytes,2,opt,name=chunk_data,json=chunkData,proto3,oneof"`
}

func (*JobRequest_Job) isJobRequest_Data() {}

func (*JobRequest_ChunkData) isJobRequest_Data() {}

type InternalJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Job *InternalJob `protobuf:"bytes,1,opt,name=job,proto3" json:"job,omitempty"`
}

func (x *InternalJobRequest) Reset() {
	*x = InternalJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InternalJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InternalJobRequest) ProtoMessage() {}

func (x *InternalJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_job_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InternalJobRequest.ProtoReflect.Descriptor instead.
func (*InternalJobRequest) Descriptor() ([]byte, []int) {
	return file_job_service_proto_rawDescGZIP(), []int{1}
}

func (x *InternalJobRequest) GetJob() *InternalJob {
	if x != nil {
		return x.Job
	}
	return nil
}

type JobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//	*JobResponse_FileInfo
	//	*JobResponse_Response
	//	*JobResponse_ChunkData
	Data isJobResponse_Data `protobuf_oneof:"data"`
}

func (x *JobResponse) Reset() {
	*x = JobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobResponse) ProtoMessage() {}

func (x *JobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_job_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobResponse.ProtoReflect.Descriptor instead.
func (*JobResponse) Descriptor() ([]byte, []int) {
	return file_job_service_proto_rawDescGZIP(), []int{2}
}

func (m *JobResponse) GetData() isJobResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *JobResponse) GetFileInfo() *FileInfo {
	if x, ok := x.GetData().(*JobResponse_FileInfo); ok {
		return x.FileInfo
	}
	return nil
}

func (x *JobResponse) GetResponse() *Response {
	if x, ok := x.GetData().(*JobResponse_Response); ok {
		return x.Response
	}
	return nil
}

func (x *JobResponse) GetChunkData() *Chunk {
	if x, ok := x.GetData().(*JobResponse_ChunkData); ok {
		return x.ChunkData
	}
	return nil
}

type isJobResponse_Data interface {
	isJobResponse_Data()
}

type JobResponse_FileInfo struct {
	FileInfo *FileInfo `protobuf:"bytes,1,opt,name=file_info,json=fileInfo,proto3,oneof"`
}

type JobResponse_Response struct {
	Response *Response `protobuf:"bytes,2,opt,name=response,proto3,oneof"`
}

type JobResponse_ChunkData struct {
	ChunkData *Chunk `protobuf:"bytes,3,opt,name=chunk_data,json=chunkData,proto3,oneof"`
}

func (*JobResponse_FileInfo) isJobResponse_Data() {}

func (*JobResponse_Response) isJobResponse_Data() {}

func (*JobResponse_ChunkData) isJobResponse_Data() {}

type InternalJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileId   []int64   `protobuf:"varint,1,rep,packed,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	Response *Response `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *InternalJobResponse) Reset() {
	*x = InternalJobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InternalJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InternalJobResponse) ProtoMessage() {}

func (x *InternalJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_job_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InternalJobResponse.ProtoReflect.Descriptor instead.
func (*InternalJobResponse) Descriptor() ([]byte, []int) {
	return file_job_service_proto_rawDescGZIP(), []int{3}
}

func (x *InternalJobResponse) GetFileId() []int64 {
	if x != nil {
		return x.FileId
	}
	return nil
}

func (x *InternalJobResponse) GetResponse() *Response {
	if x != nil {
		return x.Response
	}
	return nil
}

var File_job_service_proto protoreflect.FileDescriptor

var file_job_service_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x12, 0x70, 0x6c, 0x75, 0x67, 0x67, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x1a, 0x11, 0x6a, 0x6f, 0x62, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x13, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x7d, 0x0a, 0x0a, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a,
	0x03, 0x6a, 0x6f, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x6c, 0x75,
	0x67, 0x67, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e,
	0x4a, 0x6f, 0x62, 0x48, 0x00, 0x52, 0x03, 0x6a, 0x6f, 0x62, 0x12, 0x3a, 0x0a, 0x0a, 0x63, 0x68,
	0x75, 0x6e, 0x6b, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x70, 0x6c, 0x75, 0x67, 0x67, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x48, 0x00, 0x52, 0x09, 0x63, 0x68, 0x75,
	0x6e, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x47,
	0x0a, 0x12, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x67, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4a,
	0x6f, 0x62, 0x52, 0x03, 0x6a, 0x6f, 0x62, 0x22, 0xca, 0x01, 0x0a, 0x0b, 0x4a, 0x6f, 0x62, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x6c, 0x75,
	0x67, 0x67, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e,
	0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x67, 0x61, 0x62,
	0x6c, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3a, 0x0a, 0x0a, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x67, 0x61, 0x62, 0x6c, 0x65,
	0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x48,
	0x00, 0x52, 0x09, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x42, 0x06, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x68, 0x0a, 0x13, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x06, 0x66, 0x69,
	0x6c, 0x65, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x67, 0x61, 0x62,
	0x6c, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x60,
	0x0a, 0x0a, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x52, 0x0a, 0x09,
	0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x4a, 0x6f, 0x62, 0x12, 0x1e, 0x2e, 0x70, 0x6c, 0x75, 0x67,
	0x67, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x4a,
	0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x6c, 0x75, 0x67,
	0x67, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x4a,
	0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01,
	0x32, 0x74, 0x0a, 0x12, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4a, 0x6f, 0x62, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5e, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74,
	0x4a, 0x6f, 0x62, 0x12, 0x26, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x67, 0x61, 0x62, 0x6c, 0x65, 0x2e,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x70, 0x6c,
	0x75, 0x67, 0x67, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_job_service_proto_rawDescOnce sync.Once
	file_job_service_proto_rawDescData = file_job_service_proto_rawDesc
)

func file_job_service_proto_rawDescGZIP() []byte {
	file_job_service_proto_rawDescOnce.Do(func() {
		file_job_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_job_service_proto_rawDescData)
	})
	return file_job_service_proto_rawDescData
}

var file_job_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_job_service_proto_goTypes = []interface{}{
	(*JobRequest)(nil),          // 0: dsipe.transfer.JobRequest
	(*InternalJobRequest)(nil),  // 1: dsipe.transfer.InternalJobRequest
	(*JobResponse)(nil),         // 2: dsipe.transfer.JobResponse
	(*InternalJobResponse)(nil), // 3: dsipe.transfer.InternalJobResponse
	(*Job)(nil),                 // 4: dsipe.transfer.Job
	(*Chunk)(nil),               // 5: dsipe.transfer.Chunk
	(*InternalJob)(nil),         // 6: dsipe.transfer.InternalJob
	(*FileInfo)(nil),            // 7: dsipe.transfer.FileInfo
	(*Response)(nil),            // 8: dsipe.transfer.Response
}
var file_job_service_proto_depIdxs = []int32{
	4, // 0: dsipe.transfer.JobRequest.job:type_name -> dsipe.transfer.Job
	5, // 1: dsipe.transfer.JobRequest.chunk_data:type_name -> dsipe.transfer.Chunk
	6, // 2: dsipe.transfer.InternalJobRequest.job:type_name -> dsipe.transfer.InternalJob
	7, // 3: dsipe.transfer.JobResponse.file_info:type_name -> dsipe.transfer.FileInfo
	8, // 4: dsipe.transfer.JobResponse.response:type_name -> dsipe.transfer.Response
	5, // 5: dsipe.transfer.JobResponse.chunk_data:type_name -> dsipe.transfer.Chunk
	8, // 6: dsipe.transfer.InternalJobResponse.response:type_name -> dsipe.transfer.Response
	0, // 7: dsipe.transfer.JobService.SubmitJob:input_type -> dsipe.transfer.JobRequest
	1, // 8: dsipe.transfer.InternalJobService.SubmitJob:input_type -> dsipe.transfer.InternalJobRequest
	2, // 9: dsipe.transfer.JobService.SubmitJob:output_type -> dsipe.transfer.JobResponse
	3, // 10: dsipe.transfer.InternalJobService.SubmitJob:output_type -> dsipe.transfer.InternalJobResponse
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_job_service_proto_init() }
func file_job_service_proto_init() {
	if File_job_service_proto != nil {
		return
	}
	file_job_message_proto_init()
	file_response_message_proto_init()
	file_chunk_message_proto_init()
	file_fileInfo_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_job_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobRequest); i {
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
		file_job_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InternalJobRequest); i {
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
		file_job_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobResponse); i {
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
		file_job_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InternalJobResponse); i {
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
	file_job_service_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*JobRequest_Job)(nil),
		(*JobRequest_ChunkData)(nil),
	}
	file_job_service_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*JobResponse_FileInfo)(nil),
		(*JobResponse_Response)(nil),
		(*JobResponse_ChunkData)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_job_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_job_service_proto_goTypes,
		DependencyIndexes: file_job_service_proto_depIdxs,
		MessageInfos:      file_job_service_proto_msgTypes,
	}.Build()
	File_job_service_proto = out.File
	file_job_service_proto_rawDesc = nil
	file_job_service_proto_goTypes = nil
	file_job_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// JobServiceClient is the client API for JobService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type JobServiceClient interface {
	SubmitJob(ctx context.Context, opts ...grpc.CallOption) (JobService_SubmitJobClient, error)
}

type jobServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewJobServiceClient(cc grpc.ClientConnInterface) JobServiceClient {
	return &jobServiceClient{cc}
}

func (c *jobServiceClient) SubmitJob(ctx context.Context, opts ...grpc.CallOption) (JobService_SubmitJobClient, error) {
	stream, err := c.cc.NewStream(ctx, &_JobService_serviceDesc.Streams[0], "/dsipe.transfer.JobService/SubmitJob", opts...)
	if err != nil {
		return nil, err
	}
	x := &jobServiceSubmitJobClient{stream}
	return x, nil
}

type JobService_SubmitJobClient interface {
	Send(*JobRequest) error
	Recv() (*JobResponse, error)
	grpc.ClientStream
}

type jobServiceSubmitJobClient struct {
	grpc.ClientStream
}

func (x *jobServiceSubmitJobClient) Send(m *JobRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *jobServiceSubmitJobClient) Recv() (*JobResponse, error) {
	m := new(JobResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// JobServiceServer is the server API for JobService service.
type JobServiceServer interface {
	SubmitJob(JobService_SubmitJobServer) error
}

// UnimplementedJobServiceServer can be embedded to have forward compatible implementations.
type UnimplementedJobServiceServer struct {
}

func (*UnimplementedJobServiceServer) SubmitJob(JobService_SubmitJobServer) error {
	return status.Errorf(codes.Unimplemented, "method SubmitJob not implemented")
}

func RegisterJobServiceServer(s *grpc.Server, srv JobServiceServer) {
	s.RegisterService(&_JobService_serviceDesc, srv)
}

func _JobService_SubmitJob_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(JobServiceServer).SubmitJob(&jobServiceSubmitJobServer{stream})
}

type JobService_SubmitJobServer interface {
	Send(*JobResponse) error
	Recv() (*JobRequest, error)
	grpc.ServerStream
}

type jobServiceSubmitJobServer struct {
	grpc.ServerStream
}

func (x *jobServiceSubmitJobServer) Send(m *JobResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *jobServiceSubmitJobServer) Recv() (*JobRequest, error) {
	m := new(JobRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _JobService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dsipe.transfer.JobService",
	HandlerType: (*JobServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubmitJob",
			Handler:       _JobService_SubmitJob_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "job_service.proto",
}

// InternalJobServiceClient is the client API for InternalJobService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InternalJobServiceClient interface {
	SubmitJob(ctx context.Context, in *InternalJobRequest, opts ...grpc.CallOption) (*InternalJobResponse, error)
}

type internalJobServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInternalJobServiceClient(cc grpc.ClientConnInterface) InternalJobServiceClient {
	return &internalJobServiceClient{cc}
}

func (c *internalJobServiceClient) SubmitJob(ctx context.Context, in *InternalJobRequest, opts ...grpc.CallOption) (*InternalJobResponse, error) {
	out := new(InternalJobResponse)
	err := c.cc.Invoke(ctx, "/dsipe.transfer.InternalJobService/SubmitJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InternalJobServiceServer is the server API for InternalJobService service.
type InternalJobServiceServer interface {
	SubmitJob(context.Context, *InternalJobRequest) (*InternalJobResponse, error)
}

// UnimplementedInternalJobServiceServer can be embedded to have forward compatible implementations.
type UnimplementedInternalJobServiceServer struct {
}

func (*UnimplementedInternalJobServiceServer) SubmitJob(context.Context, *InternalJobRequest) (*InternalJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitJob not implemented")
}

func RegisterInternalJobServiceServer(s *grpc.Server, srv InternalJobServiceServer) {
	s.RegisterService(&_InternalJobService_serviceDesc, srv)
}

func _InternalJobService_SubmitJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InternalJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalJobServiceServer).SubmitJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dsipe.transfer.InternalJobService/SubmitJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalJobServiceServer).SubmitJob(ctx, req.(*InternalJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _InternalJobService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dsipe.transfer.InternalJobService",
	HandlerType: (*InternalJobServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitJob",
			Handler:    _InternalJobService_SubmitJob_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "job_service.proto",
}
