// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: editor.proto

package api

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

type NewFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	Content  string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *NewFileRequest) Reset() {
	*x = NewFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_editor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewFileRequest) ProtoMessage() {}

func (x *NewFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_editor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewFileRequest.ProtoReflect.Descriptor instead.
func (*NewFileRequest) Descriptor() ([]byte, []int) {
	return file_editor_proto_rawDescGZIP(), []int{0}
}

func (x *NewFileRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *NewFileRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type NewFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *NewFileResponse) Reset() {
	*x = NewFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_editor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewFileResponse) ProtoMessage() {}

func (x *NewFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_editor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewFileResponse.ProtoReflect.Descriptor instead.
func (*NewFileResponse) Descriptor() ([]byte, []int) {
	return file_editor_proto_rawDescGZIP(), []int{1}
}

func (x *NewFileResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type ReadFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
}

func (x *ReadFileRequest) Reset() {
	*x = ReadFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_editor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadFileRequest) ProtoMessage() {}

func (x *ReadFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_editor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadFileRequest.ProtoReflect.Descriptor instead.
func (*ReadFileRequest) Descriptor() ([]byte, []int) {
	return file_editor_proto_rawDescGZIP(), []int{2}
}

func (x *ReadFileRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

type ReadFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *ReadFileResponse) Reset() {
	*x = ReadFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_editor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadFileResponse) ProtoMessage() {}

func (x *ReadFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_editor_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadFileResponse.ProtoReflect.Descriptor instead.
func (*ReadFileResponse) Descriptor() ([]byte, []int) {
	return file_editor_proto_rawDescGZIP(), []int{3}
}

func (x *ReadFileResponse) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type SaveFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	Content  []byte `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *SaveFileRequest) Reset() {
	*x = SaveFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_editor_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveFileRequest) ProtoMessage() {}

func (x *SaveFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_editor_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveFileRequest.ProtoReflect.Descriptor instead.
func (*SaveFileRequest) Descriptor() ([]byte, []int) {
	return file_editor_proto_rawDescGZIP(), []int{4}
}

func (x *SaveFileRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *SaveFileRequest) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type SaveFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *SaveFileResponse) Reset() {
	*x = SaveFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_editor_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveFileResponse) ProtoMessage() {}

func (x *SaveFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_editor_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveFileResponse.ProtoReflect.Descriptor instead.
func (*SaveFileResponse) Descriptor() ([]byte, []int) {
	return file_editor_proto_rawDescGZIP(), []int{5}
}

func (x *SaveFileResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type FindTextRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename   string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	SearchText string `protobuf:"bytes,2,opt,name=search_text,json=searchText,proto3" json:"search_text,omitempty"`
}

func (x *FindTextRequest) Reset() {
	*x = FindTextRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_editor_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindTextRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindTextRequest) ProtoMessage() {}

func (x *FindTextRequest) ProtoReflect() protoreflect.Message {
	mi := &file_editor_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindTextRequest.ProtoReflect.Descriptor instead.
func (*FindTextRequest) Descriptor() ([]byte, []int) {
	return file_editor_proto_rawDescGZIP(), []int{6}
}

func (x *FindTextRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *FindTextRequest) GetSearchText() string {
	if x != nil {
		return x.SearchText
	}
	return ""
}

type FindTextResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Lines []string `protobuf:"bytes,2,rep,name=lines,proto3" json:"lines,omitempty"`
}

func (x *FindTextResponse) Reset() {
	*x = FindTextResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_editor_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindTextResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindTextResponse) ProtoMessage() {}

func (x *FindTextResponse) ProtoReflect() protoreflect.Message {
	mi := &file_editor_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindTextResponse.ProtoReflect.Descriptor instead.
func (*FindTextResponse) Descriptor() ([]byte, []int) {
	return file_editor_proto_rawDescGZIP(), []int{7}
}

func (x *FindTextResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *FindTextResponse) GetLines() []string {
	if x != nil {
		return x.Lines
	}
	return nil
}

type ReadAllFilesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content      []*FileContent `protobuf:"bytes,1,rep,name=content,proto3" json:"content,omitempty"`
	ResponseTime float32        `protobuf:"fixed32,2,opt,name=response_time,json=responseTime,proto3" json:"response_time,omitempty"`
}

func (x *ReadAllFilesResponse) Reset() {
	*x = ReadAllFilesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_editor_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadAllFilesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadAllFilesResponse) ProtoMessage() {}

func (x *ReadAllFilesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_editor_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadAllFilesResponse.ProtoReflect.Descriptor instead.
func (*ReadAllFilesResponse) Descriptor() ([]byte, []int) {
	return file_editor_proto_rawDescGZIP(), []int{8}
}

func (x *ReadAllFilesResponse) GetContent() []*FileContent {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *ReadAllFilesResponse) GetResponseTime() float32 {
	if x != nil {
		return x.ResponseTime
	}
	return 0
}

type FindAndReplaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename    string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	FindText    string `protobuf:"bytes,2,opt,name=find_text,json=findText,proto3" json:"find_text,omitempty"`
	ReplaceText string `protobuf:"bytes,3,opt,name=replace_text,json=replaceText,proto3" json:"replace_text,omitempty"`
}

func (x *FindAndReplaceRequest) Reset() {
	*x = FindAndReplaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_editor_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAndReplaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAndReplaceRequest) ProtoMessage() {}

func (x *FindAndReplaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_editor_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAndReplaceRequest.ProtoReflect.Descriptor instead.
func (*FindAndReplaceRequest) Descriptor() ([]byte, []int) {
	return file_editor_proto_rawDescGZIP(), []int{9}
}

func (x *FindAndReplaceRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *FindAndReplaceRequest) GetFindText() string {
	if x != nil {
		return x.FindText
	}
	return ""
}

func (x *FindAndReplaceRequest) GetReplaceText() string {
	if x != nil {
		return x.ReplaceText
	}
	return ""
}

type FindAndReplaceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count     int64                      `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Positions []*FindAndReplacePositions `protobuf:"bytes,2,rep,name=positions,proto3" json:"positions,omitempty"`
}

func (x *FindAndReplaceResponse) Reset() {
	*x = FindAndReplaceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_editor_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAndReplaceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAndReplaceResponse) ProtoMessage() {}

func (x *FindAndReplaceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_editor_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAndReplaceResponse.ProtoReflect.Descriptor instead.
func (*FindAndReplaceResponse) Descriptor() ([]byte, []int) {
	return file_editor_proto_rawDescGZIP(), []int{10}
}

func (x *FindAndReplaceResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *FindAndReplaceResponse) GetPositions() []*FindAndReplacePositions {
	if x != nil {
		return x.Positions
	}
	return nil
}

type DeleteTextRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename      string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	StartPosition int32  `protobuf:"varint,2,opt,name=start_position,json=startPosition,proto3" json:"start_position,omitempty"`
	Length        int32  `protobuf:"varint,3,opt,name=length,proto3" json:"length,omitempty"`
}

func (x *DeleteTextRequest) Reset() {
	*x = DeleteTextRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_editor_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTextRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTextRequest) ProtoMessage() {}

func (x *DeleteTextRequest) ProtoReflect() protoreflect.Message {
	mi := &file_editor_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTextRequest.ProtoReflect.Descriptor instead.
func (*DeleteTextRequest) Descriptor() ([]byte, []int) {
	return file_editor_proto_rawDescGZIP(), []int{11}
}

func (x *DeleteTextRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *DeleteTextRequest) GetStartPosition() int32 {
	if x != nil {
		return x.StartPosition
	}
	return 0
}

func (x *DeleteTextRequest) GetLength() int32 {
	if x != nil {
		return x.Length
	}
	return 0
}

type DeleteTextResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteTextResponse) Reset() {
	*x = DeleteTextResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_editor_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTextResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTextResponse) ProtoMessage() {}

func (x *DeleteTextResponse) ProtoReflect() protoreflect.Message {
	mi := &file_editor_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTextResponse.ProtoReflect.Descriptor instead.
func (*DeleteTextResponse) Descriptor() ([]byte, []int) {
	return file_editor_proto_rawDescGZIP(), []int{12}
}

func (x *DeleteTextResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_editor_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_editor_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_editor_proto_rawDescGZIP(), []int{13}
}

type FileContent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *FileContent) Reset() {
	*x = FileContent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_editor_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileContent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileContent) ProtoMessage() {}

func (x *FileContent) ProtoReflect() protoreflect.Message {
	mi := &file_editor_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileContent.ProtoReflect.Descriptor instead.
func (*FileContent) Descriptor() ([]byte, []int) {
	return file_editor_proto_rawDescGZIP(), []int{14}
}

func (x *FileContent) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FileContent) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type FindAndReplacePositions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Row    int64 `protobuf:"varint,1,opt,name=row,proto3" json:"row,omitempty"`
	Column int64 `protobuf:"varint,2,opt,name=column,proto3" json:"column,omitempty"`
}

func (x *FindAndReplacePositions) Reset() {
	*x = FindAndReplacePositions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_editor_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAndReplacePositions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAndReplacePositions) ProtoMessage() {}

func (x *FindAndReplacePositions) ProtoReflect() protoreflect.Message {
	mi := &file_editor_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAndReplacePositions.ProtoReflect.Descriptor instead.
func (*FindAndReplacePositions) Descriptor() ([]byte, []int) {
	return file_editor_proto_rawDescGZIP(), []int{15}
}

func (x *FindAndReplacePositions) GetRow() int64 {
	if x != nil {
		return x.Row
	}
	return 0
}

func (x *FindAndReplacePositions) GetColumn() int64 {
	if x != nil {
		return x.Column
	}
	return 0
}

var File_editor_proto protoreflect.FileDescriptor

var file_editor_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x65, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x65, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x22, 0x46, 0x0a, 0x0e, 0x4e, 0x65, 0x77, 0x46, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x2d,
	0x0a, 0x0f, 0x4e, 0x65, 0x77, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2d, 0x0a,
	0x0f, 0x52, 0x65, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2c, 0x0a, 0x10,
	0x52, 0x65, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x47, 0x0a, 0x0f, 0x53, 0x61,
	0x76, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x22, 0x2e, 0x0a, 0x10, 0x53, 0x61, 0x76, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x4e, 0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x64, 0x54, 0x65, 0x78, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x74, 0x65, 0x78,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54,
	0x65, 0x78, 0x74, 0x22, 0x3e, 0x0a, 0x10, 0x46, 0x69, 0x6e, 0x64, 0x54, 0x65, 0x78, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x69,
	0x6e, 0x65, 0x73, 0x22, 0x6a, 0x0a, 0x14, 0x52, 0x65, 0x61, 0x64, 0x41, 0x6c, 0x6c, 0x46, 0x69,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x65,
	0x64, 0x69, 0x74, 0x6f, 0x72, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22,
	0x73, 0x0a, 0x15, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6e, 0x64, 0x5f, 0x74, 0x65, 0x78,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6e, 0x64, 0x54, 0x65, 0x78,
	0x74, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x5f, 0x74, 0x65, 0x78,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x54, 0x65, 0x78, 0x74, 0x22, 0x6d, 0x0a, 0x16, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6e, 0x64, 0x52,
	0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3d, 0x0a, 0x09, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x65, 0x64, 0x69, 0x74, 0x6f, 0x72,
	0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x09, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x22, 0x6e, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x78,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6c,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6c, 0x65, 0x6e,
	0x67, 0x74, 0x68, 0x22, 0x2e, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x78,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x3b, 0x0a, 0x0b,
	0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x43, 0x0a, 0x17, 0x46, 0x69, 0x6e,
	0x64, 0x41, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x6f, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x03, 0x72, 0x6f, 0x77, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x32, 0xd8,
	0x03, 0x0a, 0x0a, 0x54, 0x65, 0x78, 0x74, 0x45, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x12, 0x3a, 0x0a,
	0x07, 0x4e, 0x65, 0x77, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x16, 0x2e, 0x65, 0x64, 0x69, 0x74, 0x6f,
	0x72, 0x2e, 0x4e, 0x65, 0x77, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x65, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x2e, 0x4e, 0x65, 0x77, 0x46, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x08, 0x52, 0x65, 0x61,
	0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x17, 0x2e, 0x65, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x2e, 0x52,
	0x65, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x65, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0c, 0x52, 0x65, 0x61, 0x64,
	0x41, 0x6c, 0x6c, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x0d, 0x2e, 0x65, 0x64, 0x69, 0x74, 0x6f,
	0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1c, 0x2e, 0x65, 0x64, 0x69, 0x74, 0x6f, 0x72,
	0x2e, 0x52, 0x65, 0x61, 0x64, 0x41, 0x6c, 0x6c, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x08, 0x53, 0x61, 0x76, 0x65, 0x46, 0x69, 0x6c,
	0x65, 0x12, 0x17, 0x2e, 0x65, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x65, 0x64, 0x69,
	0x74, 0x6f, 0x72, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x08, 0x46, 0x69, 0x6e, 0x64, 0x54, 0x65, 0x78, 0x74,
	0x12, 0x17, 0x2e, 0x65, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x54, 0x65,
	0x78, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x65, 0x64, 0x69, 0x74,
	0x6f, 0x72, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6e, 0x64, 0x52, 0x65,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x12, 0x1d, 0x2e, 0x65, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x2e, 0x46,
	0x69, 0x6e, 0x64, 0x41, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x65, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x2e, 0x46, 0x69,
	0x6e, 0x64, 0x41, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65,
	0x78, 0x74, 0x12, 0x19, 0x2e, 0x65, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x65, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x78,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x3b,
	0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_editor_proto_rawDescOnce sync.Once
	file_editor_proto_rawDescData = file_editor_proto_rawDesc
)

func file_editor_proto_rawDescGZIP() []byte {
	file_editor_proto_rawDescOnce.Do(func() {
		file_editor_proto_rawDescData = protoimpl.X.CompressGZIP(file_editor_proto_rawDescData)
	})
	return file_editor_proto_rawDescData
}

var file_editor_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_editor_proto_goTypes = []any{
	(*NewFileRequest)(nil),          // 0: editor.NewFileRequest
	(*NewFileResponse)(nil),         // 1: editor.NewFileResponse
	(*ReadFileRequest)(nil),         // 2: editor.ReadFileRequest
	(*ReadFileResponse)(nil),        // 3: editor.ReadFileResponse
	(*SaveFileRequest)(nil),         // 4: editor.SaveFileRequest
	(*SaveFileResponse)(nil),        // 5: editor.SaveFileResponse
	(*FindTextRequest)(nil),         // 6: editor.FindTextRequest
	(*FindTextResponse)(nil),        // 7: editor.FindTextResponse
	(*ReadAllFilesResponse)(nil),    // 8: editor.ReadAllFilesResponse
	(*FindAndReplaceRequest)(nil),   // 9: editor.FindAndReplaceRequest
	(*FindAndReplaceResponse)(nil),  // 10: editor.FindAndReplaceResponse
	(*DeleteTextRequest)(nil),       // 11: editor.DeleteTextRequest
	(*DeleteTextResponse)(nil),      // 12: editor.DeleteTextResponse
	(*Empty)(nil),                   // 13: editor.Empty
	(*FileContent)(nil),             // 14: editor.FileContent
	(*FindAndReplacePositions)(nil), // 15: editor.FindAndReplacePositions
}
var file_editor_proto_depIdxs = []int32{
	14, // 0: editor.ReadAllFilesResponse.content:type_name -> editor.FileContent
	15, // 1: editor.FindAndReplaceResponse.positions:type_name -> editor.FindAndReplacePositions
	0,  // 2: editor.TextEditor.NewFile:input_type -> editor.NewFileRequest
	2,  // 3: editor.TextEditor.ReadFile:input_type -> editor.ReadFileRequest
	13, // 4: editor.TextEditor.ReadAllFiles:input_type -> editor.Empty
	4,  // 5: editor.TextEditor.SaveFile:input_type -> editor.SaveFileRequest
	6,  // 6: editor.TextEditor.FindText:input_type -> editor.FindTextRequest
	9,  // 7: editor.TextEditor.FindAndReplace:input_type -> editor.FindAndReplaceRequest
	11, // 8: editor.TextEditor.DeleteText:input_type -> editor.DeleteTextRequest
	1,  // 9: editor.TextEditor.NewFile:output_type -> editor.NewFileResponse
	3,  // 10: editor.TextEditor.ReadFile:output_type -> editor.ReadFileResponse
	8,  // 11: editor.TextEditor.ReadAllFiles:output_type -> editor.ReadAllFilesResponse
	5,  // 12: editor.TextEditor.SaveFile:output_type -> editor.SaveFileResponse
	7,  // 13: editor.TextEditor.FindText:output_type -> editor.FindTextResponse
	10, // 14: editor.TextEditor.FindAndReplace:output_type -> editor.FindAndReplaceResponse
	12, // 15: editor.TextEditor.DeleteText:output_type -> editor.DeleteTextResponse
	9,  // [9:16] is the sub-list for method output_type
	2,  // [2:9] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_editor_proto_init() }
func file_editor_proto_init() {
	if File_editor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_editor_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*NewFileRequest); i {
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
		file_editor_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*NewFileResponse); i {
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
		file_editor_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ReadFileRequest); i {
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
		file_editor_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ReadFileResponse); i {
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
		file_editor_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*SaveFileRequest); i {
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
		file_editor_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*SaveFileResponse); i {
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
		file_editor_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*FindTextRequest); i {
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
		file_editor_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*FindTextResponse); i {
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
		file_editor_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*ReadAllFilesResponse); i {
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
		file_editor_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*FindAndReplaceRequest); i {
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
		file_editor_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*FindAndReplaceResponse); i {
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
		file_editor_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteTextRequest); i {
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
		file_editor_proto_msgTypes[12].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteTextResponse); i {
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
		file_editor_proto_msgTypes[13].Exporter = func(v any, i int) any {
			switch v := v.(*Empty); i {
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
		file_editor_proto_msgTypes[14].Exporter = func(v any, i int) any {
			switch v := v.(*FileContent); i {
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
		file_editor_proto_msgTypes[15].Exporter = func(v any, i int) any {
			switch v := v.(*FindAndReplacePositions); i {
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
			RawDescriptor: file_editor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_editor_proto_goTypes,
		DependencyIndexes: file_editor_proto_depIdxs,
		MessageInfos:      file_editor_proto_msgTypes,
	}.Build()
	File_editor_proto = out.File
	file_editor_proto_rawDesc = nil
	file_editor_proto_goTypes = nil
	file_editor_proto_depIdxs = nil
}
