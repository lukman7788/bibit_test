// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: movie.proto

package movie_grpc

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

type MovieRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImdbID string `protobuf:"bytes,1,opt,name=ImdbID,proto3" json:"ImdbID,omitempty"`
}

func (x *MovieRequest) Reset() {
	*x = MovieRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieRequest) ProtoMessage() {}

func (x *MovieRequest) ProtoReflect() protoreflect.Message {
	mi := &file_movie_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieRequest.ProtoReflect.Descriptor instead.
func (*MovieRequest) Descriptor() ([]byte, []int) {
	return file_movie_proto_rawDescGZIP(), []int{0}
}

func (x *MovieRequest) GetImdbID() string {
	if x != nil {
		return x.ImdbID
	}
	return ""
}

type SearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination int64  `protobuf:"varint,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	Searchword string `protobuf:"bytes,2,opt,name=searchword,proto3" json:"searchword,omitempty"`
}

func (x *SearchRequest) Reset() {
	*x = SearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRequest) ProtoMessage() {}

func (x *SearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_movie_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRequest.ProtoReflect.Descriptor instead.
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return file_movie_proto_rawDescGZIP(), []int{1}
}

func (x *SearchRequest) GetPagination() int64 {
	if x != nil {
		return x.Pagination
	}
	return 0
}

func (x *SearchRequest) GetSearchword() string {
	if x != nil {
		return x.Searchword
	}
	return ""
}

type SearchData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Search      []*ShortMovie `protobuf:"bytes,1,rep,name=Search,proto3" json:"Search,omitempty"`
	TotalResult string        `protobuf:"bytes,2,opt,name=totalResult,proto3" json:"totalResult,omitempty"`
	Response    string        `protobuf:"bytes,3,opt,name=Response,proto3" json:"Response,omitempty"`
}

func (x *SearchData) Reset() {
	*x = SearchData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchData) ProtoMessage() {}

func (x *SearchData) ProtoReflect() protoreflect.Message {
	mi := &file_movie_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchData.ProtoReflect.Descriptor instead.
func (*SearchData) Descriptor() ([]byte, []int) {
	return file_movie_proto_rawDescGZIP(), []int{2}
}

func (x *SearchData) GetSearch() []*ShortMovie {
	if x != nil {
		return x.Search
	}
	return nil
}

func (x *SearchData) GetTotalResult() string {
	if x != nil {
		return x.TotalResult
	}
	return ""
}

func (x *SearchData) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type ShortMovie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImdbID string `protobuf:"bytes,1,opt,name=ImdbID,proto3" json:"ImdbID,omitempty"`
	Title  string `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Year   string `protobuf:"bytes,3,opt,name=Year,proto3" json:"Year,omitempty"`
	Type   string `protobuf:"bytes,4,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (x *ShortMovie) Reset() {
	*x = ShortMovie{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortMovie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortMovie) ProtoMessage() {}

func (x *ShortMovie) ProtoReflect() protoreflect.Message {
	mi := &file_movie_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortMovie.ProtoReflect.Descriptor instead.
func (*ShortMovie) Descriptor() ([]byte, []int) {
	return file_movie_proto_rawDescGZIP(), []int{3}
}

func (x *ShortMovie) GetImdbID() string {
	if x != nil {
		return x.ImdbID
	}
	return ""
}

func (x *ShortMovie) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ShortMovie) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *ShortMovie) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type MovieData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImdbID   string    `protobuf:"bytes,1,opt,name=ImdbID,proto3" json:"ImdbID,omitempty"`
	Title    string    `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Year     string    `protobuf:"bytes,3,opt,name=Year,proto3" json:"Year,omitempty"`
	Type     string    `protobuf:"bytes,4,opt,name=Type,proto3" json:"Type,omitempty"`
	Poster   string    `protobuf:"bytes,5,opt,name=Poster,proto3" json:"Poster,omitempty"`
	Rated    string    `protobuf:"bytes,6,opt,name=Rated,proto3" json:"Rated,omitempty"`
	Released string    `protobuf:"bytes,7,opt,name=Released,proto3" json:"Released,omitempty"`
	Runtime  string    `protobuf:"bytes,8,opt,name=Runtime,proto3" json:"Runtime,omitempty"`
	Genre    string    `protobuf:"bytes,9,opt,name=Genre,proto3" json:"Genre,omitempty"`
	Director string    `protobuf:"bytes,10,opt,name=Director,proto3" json:"Director,omitempty"`
	Writer   string    `protobuf:"bytes,11,opt,name=Writer,proto3" json:"Writer,omitempty"`
	Actors   string    `protobuf:"bytes,12,opt,name=Actors,proto3" json:"Actors,omitempty"`
	Ratings  []*Rating `protobuf:"bytes,13,rep,name=Ratings,proto3" json:"Ratings,omitempty"`
}

func (x *MovieData) Reset() {
	*x = MovieData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieData) ProtoMessage() {}

func (x *MovieData) ProtoReflect() protoreflect.Message {
	mi := &file_movie_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieData.ProtoReflect.Descriptor instead.
func (*MovieData) Descriptor() ([]byte, []int) {
	return file_movie_proto_rawDescGZIP(), []int{4}
}

func (x *MovieData) GetImdbID() string {
	if x != nil {
		return x.ImdbID
	}
	return ""
}

func (x *MovieData) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MovieData) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *MovieData) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *MovieData) GetPoster() string {
	if x != nil {
		return x.Poster
	}
	return ""
}

func (x *MovieData) GetRated() string {
	if x != nil {
		return x.Rated
	}
	return ""
}

func (x *MovieData) GetReleased() string {
	if x != nil {
		return x.Released
	}
	return ""
}

func (x *MovieData) GetRuntime() string {
	if x != nil {
		return x.Runtime
	}
	return ""
}

func (x *MovieData) GetGenre() string {
	if x != nil {
		return x.Genre
	}
	return ""
}

func (x *MovieData) GetDirector() string {
	if x != nil {
		return x.Director
	}
	return ""
}

func (x *MovieData) GetWriter() string {
	if x != nil {
		return x.Writer
	}
	return ""
}

func (x *MovieData) GetActors() string {
	if x != nil {
		return x.Actors
	}
	return ""
}

func (x *MovieData) GetRatings() []*Rating {
	if x != nil {
		return x.Ratings
	}
	return nil
}

type Rating struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source string `protobuf:"bytes,1,opt,name=Source,proto3" json:"Source,omitempty"`
	Value  string `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *Rating) Reset() {
	*x = Rating{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rating) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rating) ProtoMessage() {}

func (x *Rating) ProtoReflect() protoreflect.Message {
	mi := &file_movie_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rating.ProtoReflect.Descriptor instead.
func (*Rating) Descriptor() ([]byte, []int) {
	return file_movie_proto_rawDescGZIP(), []int{5}
}

func (x *Rating) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Rating) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_movie_proto protoreflect.FileDescriptor

var file_movie_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6d,
	0x6f, 0x76, 0x69, 0x65, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x22, 0x26, 0x0a, 0x0c, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x6d, 0x64,
	0x62, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x49, 0x6d, 0x64, 0x62, 0x49,
	0x44, 0x22, 0x4f, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x77, 0x6f,
	0x72, 0x64, 0x22, 0x7a, 0x0a, 0x0a, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x2e, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x68,
	0x6f, 0x72, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x12, 0x20, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x62,
	0x0a, 0x0a, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x49, 0x6d, 0x64, 0x62, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x49, 0x6d,
	0x64, 0x62, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x59, 0x65,
	0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x59, 0x65, 0x61, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79,
	0x70, 0x65, 0x22, 0xd5, 0x02, 0x0a, 0x09, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x16, 0x0a, 0x06, 0x49, 0x6d, 0x64, 0x62, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x49, 0x6d, 0x64, 0x62, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x59, 0x65, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x59, 0x65,
	0x61, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x12, 0x14,
	0x0a, 0x05, 0x52, 0x61, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x52,
	0x61, 0x74, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x47, 0x65,
	0x6e, 0x72, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x47, 0x65, 0x6e, 0x72, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06,
	0x57, 0x72, 0x69, 0x74, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x57, 0x72,
	0x69, 0x74, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x2c, 0x0a, 0x07,
	0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x6d, 0x6f, 0x76, 0x69, 0x65, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x52, 0x07, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x36, 0x0a, 0x06, 0x52, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x32, 0x87, 0x01, 0x0a, 0x0c, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x48, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x72, 0x12, 0x3a, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x18,
	0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x6f, 0x76, 0x69,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65,
	0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x3b, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x19, 0x2e, 0x6d, 0x6f, 0x76, 0x69,
	0x65, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x5f, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x44, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_movie_proto_rawDescOnce sync.Once
	file_movie_proto_rawDescData = file_movie_proto_rawDesc
)

func file_movie_proto_rawDescGZIP() []byte {
	file_movie_proto_rawDescOnce.Do(func() {
		file_movie_proto_rawDescData = protoimpl.X.CompressGZIP(file_movie_proto_rawDescData)
	})
	return file_movie_proto_rawDescData
}

var file_movie_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_movie_proto_goTypes = []interface{}{
	(*MovieRequest)(nil),  // 0: movie_grpc.MovieRequest
	(*SearchRequest)(nil), // 1: movie_grpc.SearchRequest
	(*SearchData)(nil),    // 2: movie_grpc.SearchData
	(*ShortMovie)(nil),    // 3: movie_grpc.ShortMovie
	(*MovieData)(nil),     // 4: movie_grpc.MovieData
	(*Rating)(nil),        // 5: movie_grpc.Rating
}
var file_movie_proto_depIdxs = []int32{
	3, // 0: movie_grpc.SearchData.Search:type_name -> movie_grpc.ShortMovie
	5, // 1: movie_grpc.MovieData.Ratings:type_name -> movie_grpc.Rating
	0, // 2: movie_grpc.MovieHandler.GetByID:input_type -> movie_grpc.MovieRequest
	1, // 3: movie_grpc.MovieHandler.Search:input_type -> movie_grpc.SearchRequest
	4, // 4: movie_grpc.MovieHandler.GetByID:output_type -> movie_grpc.MovieData
	2, // 5: movie_grpc.MovieHandler.Search:output_type -> movie_grpc.SearchData
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_movie_proto_init() }
func file_movie_proto_init() {
	if File_movie_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_movie_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieRequest); i {
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
		file_movie_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchRequest); i {
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
		file_movie_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchData); i {
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
		file_movie_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortMovie); i {
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
		file_movie_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieData); i {
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
		file_movie_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rating); i {
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
			RawDescriptor: file_movie_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_movie_proto_goTypes,
		DependencyIndexes: file_movie_proto_depIdxs,
		MessageInfos:      file_movie_proto_msgTypes,
	}.Build()
	File_movie_proto = out.File
	file_movie_proto_rawDesc = nil
	file_movie_proto_goTypes = nil
	file_movie_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MovieHandlerClient is the client API for MovieHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MovieHandlerClient interface {
	GetByID(ctx context.Context, in *MovieRequest, opts ...grpc.CallOption) (*MovieData, error)
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchData, error)
}

type movieHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewMovieHandlerClient(cc grpc.ClientConnInterface) MovieHandlerClient {
	return &movieHandlerClient{cc}
}

func (c *movieHandlerClient) GetByID(ctx context.Context, in *MovieRequest, opts ...grpc.CallOption) (*MovieData, error) {
	out := new(MovieData)
	err := c.cc.Invoke(ctx, "/movie_grpc.MovieHandler/GetByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieHandlerClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchData, error) {
	out := new(SearchData)
	err := c.cc.Invoke(ctx, "/movie_grpc.MovieHandler/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MovieHandlerServer is the server API for MovieHandler service.
type MovieHandlerServer interface {
	GetByID(context.Context, *MovieRequest) (*MovieData, error)
	Search(context.Context, *SearchRequest) (*SearchData, error)
}

// UnimplementedMovieHandlerServer can be embedded to have forward compatible implementations.
type UnimplementedMovieHandlerServer struct {
}

func (*UnimplementedMovieHandlerServer) GetByID(context.Context, *MovieRequest) (*MovieData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (*UnimplementedMovieHandlerServer) Search(context.Context, *SearchRequest) (*SearchData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}

func RegisterMovieHandlerServer(s *grpc.Server, srv MovieHandlerServer) {
	s.RegisterService(&_MovieHandler_serviceDesc, srv)
}

func _MovieHandler_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieHandlerServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/movie_grpc.MovieHandler/GetByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieHandlerServer).GetByID(ctx, req.(*MovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieHandler_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieHandlerServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/movie_grpc.MovieHandler/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieHandlerServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MovieHandler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "movie_grpc.MovieHandler",
	HandlerType: (*MovieHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetByID",
			Handler:    _MovieHandler_GetByID_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _MovieHandler_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "movie.proto",
}