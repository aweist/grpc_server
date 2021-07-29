// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: blog/blog.proto

package blog

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

type BlogPost_Type int32

const (
	BlogPost_TEXT    BlogPost_Type = 0
	BlogPost_PICTURE BlogPost_Type = 1
	BlogPost_VIDEO   BlogPost_Type = 2
)

// Enum value maps for BlogPost_Type.
var (
	BlogPost_Type_name = map[int32]string{
		0: "TEXT",
		1: "PICTURE",
		2: "VIDEO",
	}
	BlogPost_Type_value = map[string]int32{
		"TEXT":    0,
		"PICTURE": 1,
		"VIDEO":   2,
	}
)

func (x BlogPost_Type) Enum() *BlogPost_Type {
	p := new(BlogPost_Type)
	*p = x
	return p
}

func (x BlogPost_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BlogPost_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_blog_blog_proto_enumTypes[0].Descriptor()
}

func (BlogPost_Type) Type() protoreflect.EnumType {
	return &file_blog_blog_proto_enumTypes[0]
}

func (x BlogPost_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BlogPost_Type.Descriptor instead.
func (BlogPost_Type) EnumDescriptor() ([]byte, []int) {
	return file_blog_blog_proto_rawDescGZIP(), []int{0, 0}
}

type BlogPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string        `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Body  string        `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	Id    int32         `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	Type  BlogPost_Type `protobuf:"varint,4,opt,name=type,proto3,enum=blog.BlogPost_Type" json:"type,omitempty"`
}

func (x *BlogPost) Reset() {
	*x = BlogPost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_blog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlogPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlogPost) ProtoMessage() {}

func (x *BlogPost) ProtoReflect() protoreflect.Message {
	mi := &file_blog_blog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlogPost.ProtoReflect.Descriptor instead.
func (*BlogPost) Descriptor() ([]byte, []int) {
	return file_blog_blog_proto_rawDescGZIP(), []int{0}
}

func (x *BlogPost) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *BlogPost) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *BlogPost) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BlogPost) GetType() BlogPost_Type {
	if x != nil {
		return x.Type
	}
	return BlogPost_TEXT
}

type BlogSearch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
}

func (x *BlogSearch) Reset() {
	*x = BlogSearch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_blog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlogSearch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlogSearch) ProtoMessage() {}

func (x *BlogSearch) ProtoReflect() protoreflect.Message {
	mi := &file_blog_blog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlogSearch.ProtoReflect.Descriptor instead.
func (*BlogSearch) Descriptor() ([]byte, []int) {
	return file_blog_blog_proto_rawDescGZIP(), []int{1}
}

func (x *BlogSearch) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

type BlogSearchResults struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blogposts []*BlogPost `protobuf:"bytes,1,rep,name=blogposts,proto3" json:"blogposts,omitempty"`
}

func (x *BlogSearchResults) Reset() {
	*x = BlogSearchResults{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_blog_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlogSearchResults) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlogSearchResults) ProtoMessage() {}

func (x *BlogSearchResults) ProtoReflect() protoreflect.Message {
	mi := &file_blog_blog_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlogSearchResults.ProtoReflect.Descriptor instead.
func (*BlogSearchResults) Descriptor() ([]byte, []int) {
	return file_blog_blog_proto_rawDescGZIP(), []int{2}
}

func (x *BlogSearchResults) GetBlogposts() []*BlogPost {
	if x != nil {
		return x.Blogposts
	}
	return nil
}

type BlogPostById struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *BlogPostById) Reset() {
	*x = BlogPostById{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_blog_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlogPostById) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlogPostById) ProtoMessage() {}

func (x *BlogPostById) ProtoReflect() protoreflect.Message {
	mi := &file_blog_blog_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlogPostById.ProtoReflect.Descriptor instead.
func (*BlogPostById) Descriptor() ([]byte, []int) {
	return file_blog_blog_proto_rawDescGZIP(), []int{3}
}

func (x *BlogPostById) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_blog_blog_proto protoreflect.FileDescriptor

var file_blog_blog_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x62, 0x6c, 0x6f, 0x67, 0x22, 0x97, 0x01, 0x0a, 0x08, 0x42, 0x6c, 0x6f, 0x67,
	0x50, 0x6f, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x27,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x62,
	0x6c, 0x6f, 0x67, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x50, 0x6f, 0x73, 0x74, 0x2e, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x28, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x08, 0x0a, 0x04, 0x54, 0x45, 0x58, 0x54, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x49, 0x43,
	0x54, 0x55, 0x52, 0x45, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x56, 0x49, 0x44, 0x45, 0x4f, 0x10,
	0x02, 0x22, 0x22, 0x0a, 0x0a, 0x42, 0x6c, 0x6f, 0x67, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12,
	0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x22, 0x41, 0x0a, 0x11, 0x42, 0x6c, 0x6f, 0x67, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x2c, 0x0a, 0x09, 0x62, 0x6c,
	0x6f, 0x67, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x09, 0x62,
	0x6c, 0x6f, 0x67, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x22, 0x1e, 0x0a, 0x0c, 0x42, 0x6c, 0x6f, 0x67,
	0x50, 0x6f, 0x73, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x32, 0x77, 0x0a, 0x0b, 0x42, 0x6c, 0x6f, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x42, 0x6c, 0x6f, 0x67, 0x12, 0x10, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x42, 0x6c, 0x6f,
	0x67, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x1a, 0x17, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x42,
	0x6c, 0x6f, 0x67, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73,
	0x22, 0x00, 0x12, 0x2d, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x2e,
	0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x50, 0x6f, 0x73, 0x74, 0x42, 0x79, 0x49,
	0x64, 0x1a, 0x0e, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x50, 0x6f, 0x73,
	0x74, 0x42, 0x1d, 0x5a, 0x1b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x61, 0x77, 0x65, 0x69, 0x73, 0x74, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x62, 0x6c, 0x6f, 0x67,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blog_blog_proto_rawDescOnce sync.Once
	file_blog_blog_proto_rawDescData = file_blog_blog_proto_rawDesc
)

func file_blog_blog_proto_rawDescGZIP() []byte {
	file_blog_blog_proto_rawDescOnce.Do(func() {
		file_blog_blog_proto_rawDescData = protoimpl.X.CompressGZIP(file_blog_blog_proto_rawDescData)
	})
	return file_blog_blog_proto_rawDescData
}

var file_blog_blog_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_blog_blog_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_blog_blog_proto_goTypes = []interface{}{
	(BlogPost_Type)(0),        // 0: blog.BlogPost.Type
	(*BlogPost)(nil),          // 1: blog.BlogPost
	(*BlogSearch)(nil),        // 2: blog.BlogSearch
	(*BlogSearchResults)(nil), // 3: blog.BlogSearchResults
	(*BlogPostById)(nil),      // 4: blog.BlogPostById
}
var file_blog_blog_proto_depIdxs = []int32{
	0, // 0: blog.BlogPost.type:type_name -> blog.BlogPost.Type
	1, // 1: blog.BlogSearchResults.blogposts:type_name -> blog.BlogPost
	2, // 2: blog.BlogService.SearchBlog:input_type -> blog.BlogSearch
	4, // 3: blog.BlogService.GetPost:input_type -> blog.BlogPostById
	3, // 4: blog.BlogService.SearchBlog:output_type -> blog.BlogSearchResults
	1, // 5: blog.BlogService.GetPost:output_type -> blog.BlogPost
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_blog_blog_proto_init() }
func file_blog_blog_proto_init() {
	if File_blog_blog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_blog_blog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlogPost); i {
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
		file_blog_blog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlogSearch); i {
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
		file_blog_blog_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlogSearchResults); i {
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
		file_blog_blog_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlogPostById); i {
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
			RawDescriptor: file_blog_blog_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_blog_blog_proto_goTypes,
		DependencyIndexes: file_blog_blog_proto_depIdxs,
		EnumInfos:         file_blog_blog_proto_enumTypes,
		MessageInfos:      file_blog_blog_proto_msgTypes,
	}.Build()
	File_blog_blog_proto = out.File
	file_blog_blog_proto_rawDesc = nil
	file_blog_blog_proto_goTypes = nil
	file_blog_blog_proto_depIdxs = nil
}