// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: proto/v1/bigfile.proto

package bigfilev1

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

type FileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName    string `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	ChunkNumber int32  `protobuf:"varint,2,opt,name=chunk_number,json=chunkNumber,proto3" json:"chunk_number,omitempty"`
}

func (x *FileRequest) Reset() {
	*x = FileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_bigfile_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileRequest) ProtoMessage() {}

func (x *FileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_bigfile_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileRequest.ProtoReflect.Descriptor instead.
func (*FileRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1_bigfile_proto_rawDescGZIP(), []int{0}
}

func (x *FileRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *FileRequest) GetChunkNumber() int32 {
	if x != nil {
		return x.ChunkNumber
	}
	return 0
}

type FileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string              `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Stat  *FileResponse_Stat  `protobuf:"bytes,2,opt,name=stat,proto3" json:"stat,omitempty"`
	Chunk *FileResponse_Chunk `protobuf:"bytes,3,opt,name=chunk,proto3" json:"chunk,omitempty"`
}

func (x *FileResponse) Reset() {
	*x = FileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_bigfile_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileResponse) ProtoMessage() {}

func (x *FileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_bigfile_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileResponse.ProtoReflect.Descriptor instead.
func (*FileResponse) Descriptor() ([]byte, []int) {
	return file_proto_v1_bigfile_proto_rawDescGZIP(), []int{1}
}

func (x *FileResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FileResponse) GetStat() *FileResponse_Stat {
	if x != nil {
		return x.Stat
	}
	return nil
}

func (x *FileResponse) GetChunk() *FileResponse_Chunk {
	if x != nil {
		return x.Chunk
	}
	return nil
}

type FileResponse_Stat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LastModified *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=last_modified,json=lastModified,proto3" json:"last_modified,omitempty"`
	TotalSize    int32                  `protobuf:"varint,2,opt,name=total_size,json=totalSize,proto3" json:"total_size,omitempty"`
}

func (x *FileResponse_Stat) Reset() {
	*x = FileResponse_Stat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_bigfile_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileResponse_Stat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileResponse_Stat) ProtoMessage() {}

func (x *FileResponse_Stat) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_bigfile_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileResponse_Stat.ProtoReflect.Descriptor instead.
func (*FileResponse_Stat) Descriptor() ([]byte, []int) {
	return file_proto_v1_bigfile_proto_rawDescGZIP(), []int{1, 0}
}

func (x *FileResponse_Stat) GetLastModified() *timestamppb.Timestamp {
	if x != nil {
		return x.LastModified
	}
	return nil
}

func (x *FileResponse_Stat) GetTotalSize() int32 {
	if x != nil {
		return x.TotalSize
	}
	return 0
}

type FileResponse_Chunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChunkNumber int32  `protobuf:"varint,1,opt,name=chunk_number,json=chunkNumber,proto3" json:"chunk_number,omitempty"`
	ChunkSize   int32  `protobuf:"varint,2,opt,name=chunk_size,json=chunkSize,proto3" json:"chunk_size,omitempty"`
	TotalChunks int32  `protobuf:"varint,3,opt,name=total_chunks,json=totalChunks,proto3" json:"total_chunks,omitempty"`
	End         bool   `protobuf:"varint,4,opt,name=end,proto3" json:"end,omitempty"`
	Data        []byte `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *FileResponse_Chunk) Reset() {
	*x = FileResponse_Chunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_bigfile_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileResponse_Chunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileResponse_Chunk) ProtoMessage() {}

func (x *FileResponse_Chunk) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_bigfile_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileResponse_Chunk.ProtoReflect.Descriptor instead.
func (*FileResponse_Chunk) Descriptor() ([]byte, []int) {
	return file_proto_v1_bigfile_proto_rawDescGZIP(), []int{1, 1}
}

func (x *FileResponse_Chunk) GetChunkNumber() int32 {
	if x != nil {
		return x.ChunkNumber
	}
	return 0
}

func (x *FileResponse_Chunk) GetChunkSize() int32 {
	if x != nil {
		return x.ChunkSize
	}
	return 0
}

func (x *FileResponse_Chunk) GetTotalChunks() int32 {
	if x != nil {
		return x.TotalChunks
	}
	return 0
}

func (x *FileResponse_Chunk) GetEnd() bool {
	if x != nil {
		return x.End
	}
	return false
}

func (x *FileResponse_Chunk) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_proto_v1_bigfile_proto protoreflect.FileDescriptor

var file_proto_v1_bigfile_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x69, 0x67, 0x66, 0x69,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x62, 0x69, 0x67, 0x66, 0x69, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4d, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x22, 0x88, 0x03, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x73, 0x74, 0x61,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x62, 0x69, 0x67, 0x66, 0x69, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x52, 0x04, 0x73, 0x74, 0x61, 0x74, 0x12, 0x34, 0x0a, 0x05,
	0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x62, 0x69,
	0x67, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x05, 0x63, 0x68, 0x75,
	0x6e, 0x6b, 0x1a, 0x66, 0x0a, 0x04, 0x53, 0x74, 0x61, 0x74, 0x12, 0x3f, 0x0a, 0x0d, 0x6c, 0x61,
	0x73, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x6c,
	0x61, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x69, 0x7a, 0x65, 0x1a, 0x92, 0x01, 0x0a, 0x05, 0x43,
	0x68, 0x75, 0x6e, 0x6b, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x68, 0x75, 0x6e,
	0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x75, 0x6e, 0x6b,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x68, 0x75,
	0x6e, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32,
	0x57, 0x0a, 0x13, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c,
	0x65, 0x12, 0x17, 0x2e, 0x62, 0x69, 0x67, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x62, 0x69, 0x67,
	0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x9e, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d,
	0x2e, 0x62, 0x69, 0x67, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x42, 0x69, 0x67,
	0x66, 0x69, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x35, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x70, 0x65, 0x6c, 0x69, 0x63, 0x69, 0x61,
	0x72, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x62, 0x69, 0x67, 0x66, 0x69, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x3b, 0x62, 0x69, 0x67, 0x66, 0x69, 0x6c, 0x65,
	0x76, 0x31, 0xa2, 0x02, 0x03, 0x42, 0x58, 0x58, 0xaa, 0x02, 0x0a, 0x42, 0x69, 0x67, 0x66, 0x69,
	0x6c, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0a, 0x42, 0x69, 0x67, 0x66, 0x69, 0x6c, 0x65, 0x5c,
	0x56, 0x31, 0xe2, 0x02, 0x16, 0x42, 0x69, 0x67, 0x66, 0x69, 0x6c, 0x65, 0x5c, 0x56, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0b, 0x42, 0x69,
	0x67, 0x66, 0x69, 0x6c, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_v1_bigfile_proto_rawDescOnce sync.Once
	file_proto_v1_bigfile_proto_rawDescData = file_proto_v1_bigfile_proto_rawDesc
)

func file_proto_v1_bigfile_proto_rawDescGZIP() []byte {
	file_proto_v1_bigfile_proto_rawDescOnce.Do(func() {
		file_proto_v1_bigfile_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_v1_bigfile_proto_rawDescData)
	})
	return file_proto_v1_bigfile_proto_rawDescData
}

var file_proto_v1_bigfile_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_v1_bigfile_proto_goTypes = []any{
	(*FileRequest)(nil),           // 0: bigfile.v1.FileRequest
	(*FileResponse)(nil),          // 1: bigfile.v1.FileResponse
	(*FileResponse_Stat)(nil),     // 2: bigfile.v1.FileResponse.Stat
	(*FileResponse_Chunk)(nil),    // 3: bigfile.v1.FileResponse.Chunk
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_proto_v1_bigfile_proto_depIdxs = []int32{
	2, // 0: bigfile.v1.FileResponse.stat:type_name -> bigfile.v1.FileResponse.Stat
	3, // 1: bigfile.v1.FileResponse.chunk:type_name -> bigfile.v1.FileResponse.Chunk
	4, // 2: bigfile.v1.FileResponse.Stat.last_modified:type_name -> google.protobuf.Timestamp
	0, // 3: bigfile.v1.FileTransferService.GetFile:input_type -> bigfile.v1.FileRequest
	1, // 4: bigfile.v1.FileTransferService.GetFile:output_type -> bigfile.v1.FileResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_v1_bigfile_proto_init() }
func file_proto_v1_bigfile_proto_init() {
	if File_proto_v1_bigfile_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_v1_bigfile_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*FileRequest); i {
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
		file_proto_v1_bigfile_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*FileResponse); i {
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
		file_proto_v1_bigfile_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*FileResponse_Stat); i {
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
		file_proto_v1_bigfile_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*FileResponse_Chunk); i {
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
			RawDescriptor: file_proto_v1_bigfile_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_v1_bigfile_proto_goTypes,
		DependencyIndexes: file_proto_v1_bigfile_proto_depIdxs,
		MessageInfos:      file_proto_v1_bigfile_proto_msgTypes,
	}.Build()
	File_proto_v1_bigfile_proto = out.File
	file_proto_v1_bigfile_proto_rawDesc = nil
	file_proto_v1_bigfile_proto_goTypes = nil
	file_proto_v1_bigfile_proto_depIdxs = nil
}
