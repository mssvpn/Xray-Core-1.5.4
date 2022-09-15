// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: proxy/trojan/config.proto

package trojan

import (
	protocol "github.com/xtls/xray-core/common/protocol"
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

type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Password string `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	Flow     string `protobuf:"bytes,2,opt,name=flow,proto3" json:"flow,omitempty"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_trojan_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_trojan_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_proxy_trojan_config_proto_rawDescGZIP(), []int{0}
}

func (x *Account) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Account) GetFlow() string {
	if x != nil {
		return x.Flow
	}
	return ""
}

type Fallback struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Alpn string `protobuf:"bytes,2,opt,name=alpn,proto3" json:"alpn,omitempty"`
	Path string `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Type string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Dest string `protobuf:"bytes,5,opt,name=dest,proto3" json:"dest,omitempty"`
	Xver uint64 `protobuf:"varint,6,opt,name=xver,proto3" json:"xver,omitempty"`
}

func (x *Fallback) Reset() {
	*x = Fallback{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_trojan_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Fallback) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fallback) ProtoMessage() {}

func (x *Fallback) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_trojan_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fallback.ProtoReflect.Descriptor instead.
func (*Fallback) Descriptor() ([]byte, []int) {
	return file_proxy_trojan_config_proto_rawDescGZIP(), []int{1}
}

func (x *Fallback) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Fallback) GetAlpn() string {
	if x != nil {
		return x.Alpn
	}
	return ""
}

func (x *Fallback) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Fallback) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Fallback) GetDest() string {
	if x != nil {
		return x.Dest
	}
	return ""
}

func (x *Fallback) GetXver() uint64 {
	if x != nil {
		return x.Xver
	}
	return 0
}

type ClientConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Server []*protocol.ServerEndpoint `protobuf:"bytes,1,rep,name=server,proto3" json:"server,omitempty"`
}

func (x *ClientConfig) Reset() {
	*x = ClientConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_trojan_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientConfig) ProtoMessage() {}

func (x *ClientConfig) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_trojan_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientConfig.ProtoReflect.Descriptor instead.
func (*ClientConfig) Descriptor() ([]byte, []int) {
	return file_proxy_trojan_config_proto_rawDescGZIP(), []int{2}
}

func (x *ClientConfig) GetServer() []*protocol.ServerEndpoint {
	if x != nil {
		return x.Server
	}
	return nil
}

type ServerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users     []*protocol.User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	Fallbacks []*Fallback      `protobuf:"bytes,3,rep,name=fallbacks,proto3" json:"fallbacks,omitempty"`
}

func (x *ServerConfig) Reset() {
	*x = ServerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_trojan_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerConfig) ProtoMessage() {}

func (x *ServerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_trojan_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerConfig.ProtoReflect.Descriptor instead.
func (*ServerConfig) Descriptor() ([]byte, []int) {
	return file_proxy_trojan_config_proto_rawDescGZIP(), []int{3}
}

func (x *ServerConfig) GetUsers() []*protocol.User {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *ServerConfig) GetFallbacks() []*Fallback {
	if x != nil {
		return x.Fallbacks
	}
	return nil
}

var File_proxy_trojan_config_proto protoreflect.FileDescriptor

var file_proxy_trojan_config_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x74, 0x72, 0x6f, 0x6a, 0x61, 0x6e, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x78, 0x72, 0x61,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x74, 0x72, 0x6f, 0x6a, 0x61, 0x6e, 0x1a, 0x1a,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x39, 0x0a,
	0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x6c, 0x6f, 0x77, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x66, 0x6c, 0x6f, 0x77, 0x22, 0x82, 0x01, 0x0a, 0x08, 0x46, 0x61, 0x6c,
	0x6c, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x6c, 0x70,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x6c, 0x70, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x78, 0x76, 0x65,
	0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x78, 0x76, 0x65, 0x72, 0x22, 0x4c, 0x0a,
	0x0c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3c, 0x0a,
	0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e,
	0x78, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x45, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x7b, 0x0a, 0x0c, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x30, 0x0a, 0x05, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x78, 0x72, 0x61,
	0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x39, 0x0a,
	0x09, 0x66, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x74, 0x72,
	0x6f, 0x6a, 0x61, 0x6e, 0x2e, 0x46, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x09, 0x66,
	0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x73, 0x42, 0x55, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e,
	0x78, 0x72, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x74, 0x72, 0x6f, 0x6a, 0x61,
	0x6e, 0x50, 0x01, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x78, 0x74, 0x6c, 0x73, 0x2f, 0x78, 0x72, 0x61, 0x79, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2f, 0x74, 0x72, 0x6f, 0x6a, 0x61, 0x6e, 0xaa, 0x02, 0x11, 0x58, 0x72,
	0x61, 0x79, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x54, 0x72, 0x6f, 0x6a, 0x61, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proxy_trojan_config_proto_rawDescOnce sync.Once
	file_proxy_trojan_config_proto_rawDescData = file_proxy_trojan_config_proto_rawDesc
)

func file_proxy_trojan_config_proto_rawDescGZIP() []byte {
	file_proxy_trojan_config_proto_rawDescOnce.Do(func() {
		file_proxy_trojan_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_proxy_trojan_config_proto_rawDescData)
	})
	return file_proxy_trojan_config_proto_rawDescData
}

var file_proxy_trojan_config_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proxy_trojan_config_proto_goTypes = []interface{}{
	(*Account)(nil),                 // 0: xray.proxy.trojan.Account
	(*Fallback)(nil),                // 1: xray.proxy.trojan.Fallback
	(*ClientConfig)(nil),            // 2: xray.proxy.trojan.ClientConfig
	(*ServerConfig)(nil),            // 3: xray.proxy.trojan.ServerConfig
	(*protocol.ServerEndpoint)(nil), // 4: xray.common.protocol.ServerEndpoint
	(*protocol.User)(nil),           // 5: xray.common.protocol.User
}
var file_proxy_trojan_config_proto_depIdxs = []int32{
	4, // 0: xray.proxy.trojan.ClientConfig.server:type_name -> xray.common.protocol.ServerEndpoint
	5, // 1: xray.proxy.trojan.ServerConfig.users:type_name -> xray.common.protocol.User
	1, // 2: xray.proxy.trojan.ServerConfig.fallbacks:type_name -> xray.proxy.trojan.Fallback
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proxy_trojan_config_proto_init() }
func file_proxy_trojan_config_proto_init() {
	if File_proxy_trojan_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proxy_trojan_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account); i {
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
		file_proxy_trojan_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Fallback); i {
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
		file_proxy_trojan_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientConfig); i {
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
		file_proxy_trojan_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerConfig); i {
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
			RawDescriptor: file_proxy_trojan_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proxy_trojan_config_proto_goTypes,
		DependencyIndexes: file_proxy_trojan_config_proto_depIdxs,
		MessageInfos:      file_proxy_trojan_config_proto_msgTypes,
	}.Build()
	File_proxy_trojan_config_proto = out.File
	file_proxy_trojan_config_proto_rawDesc = nil
	file_proxy_trojan_config_proto_goTypes = nil
	file_proxy_trojan_config_proto_depIdxs = nil
}
