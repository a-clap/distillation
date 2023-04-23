// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: pkg/distillation/distillationproto/ds18b20.proto

package distillationproto

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

type DSConfigs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Configs []*DSConfig `protobuf:"bytes,1,rep,name=configs,proto3" json:"configs,omitempty"`
}

func (x *DSConfigs) Reset() {
	*x = DSConfigs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_distillation_distillationproto_ds18b20_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DSConfigs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DSConfigs) ProtoMessage() {}

func (x *DSConfigs) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_distillation_distillationproto_ds18b20_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DSConfigs.ProtoReflect.Descriptor instead.
func (*DSConfigs) Descriptor() ([]byte, []int) {
	return file_pkg_distillation_distillationproto_ds18b20_proto_rawDescGZIP(), []int{0}
}

func (x *DSConfigs) GetConfigs() []*DSConfig {
	if x != nil {
		return x.Configs
	}
	return nil
}

type DSConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID           string  `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name         string  `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Correction   float32 `protobuf:"fixed32,3,opt,name=Correction,proto3" json:"Correction,omitempty"`
	Resolution   int32   `protobuf:"varint,4,opt,name=Resolution,proto3" json:"Resolution,omitempty"`
	PollInterval int32   `protobuf:"varint,5,opt,name=PollInterval,proto3" json:"PollInterval,omitempty"`
	Samples      uint32  `protobuf:"varint,6,opt,name=Samples,proto3" json:"Samples,omitempty"`
	Enabled      bool    `protobuf:"varint,7,opt,name=Enabled,proto3" json:"Enabled,omitempty"`
}

func (x *DSConfig) Reset() {
	*x = DSConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_distillation_distillationproto_ds18b20_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DSConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DSConfig) ProtoMessage() {}

func (x *DSConfig) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_distillation_distillationproto_ds18b20_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DSConfig.ProtoReflect.Descriptor instead.
func (*DSConfig) Descriptor() ([]byte, []int) {
	return file_pkg_distillation_distillationproto_ds18b20_proto_rawDescGZIP(), []int{1}
}

func (x *DSConfig) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *DSConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DSConfig) GetCorrection() float32 {
	if x != nil {
		return x.Correction
	}
	return 0
}

func (x *DSConfig) GetResolution() int32 {
	if x != nil {
		return x.Resolution
	}
	return 0
}

func (x *DSConfig) GetPollInterval() int32 {
	if x != nil {
		return x.PollInterval
	}
	return 0
}

func (x *DSConfig) GetSamples() uint32 {
	if x != nil {
		return x.Samples
	}
	return 0
}

func (x *DSConfig) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

type DSTemperatures struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Temps []*DSTemperature `protobuf:"bytes,1,rep,name=temps,proto3" json:"temps,omitempty"`
}

func (x *DSTemperatures) Reset() {
	*x = DSTemperatures{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_distillation_distillationproto_ds18b20_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DSTemperatures) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DSTemperatures) ProtoMessage() {}

func (x *DSTemperatures) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_distillation_distillationproto_ds18b20_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DSTemperatures.ProtoReflect.Descriptor instead.
func (*DSTemperatures) Descriptor() ([]byte, []int) {
	return file_pkg_distillation_distillationproto_ds18b20_proto_rawDescGZIP(), []int{2}
}

func (x *DSTemperatures) GetTemps() []*DSTemperature {
	if x != nil {
		return x.Temps
	}
	return nil
}

type DSTemperature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          string  `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Temperature float32 `protobuf:"fixed32,2,opt,name=Temperature,proto3" json:"Temperature,omitempty"`
}

func (x *DSTemperature) Reset() {
	*x = DSTemperature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_distillation_distillationproto_ds18b20_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DSTemperature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DSTemperature) ProtoMessage() {}

func (x *DSTemperature) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_distillation_distillationproto_ds18b20_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DSTemperature.ProtoReflect.Descriptor instead.
func (*DSTemperature) Descriptor() ([]byte, []int) {
	return file_pkg_distillation_distillationproto_ds18b20_proto_rawDescGZIP(), []int{3}
}

func (x *DSTemperature) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *DSTemperature) GetTemperature() float32 {
	if x != nil {
		return x.Temperature
	}
	return 0
}

var File_pkg_distillation_distillationproto_ds18b20_proto protoreflect.FileDescriptor

var file_pkg_distillation_distillationproto_ds18b20_proto_rawDesc = []byte{
	0x0a, 0x30, 0x70, 0x6b, 0x67, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x73, 0x31, 0x38, 0x62, 0x32, 0x30, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x11, 0x64, 0x69, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x42, 0x0a, 0x09, 0x44, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12,
	0x35, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x22, 0xc6, 0x01, 0x0a, 0x08, 0x44, 0x53, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x6f, 0x72, 0x72, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x43, 0x6f, 0x72,
	0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x6f, 0x6c,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x52, 0x65, 0x73,
	0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x50, 0x6f, 0x6c, 0x6c, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x50,
	0x6f, 0x6c, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x53,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x53, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22,
	0x48, 0x0a, 0x0e, 0x44, 0x53, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x73, 0x12, 0x36, 0x0a, 0x05, 0x74, 0x65, 0x6d, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x53, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x52, 0x05, 0x74, 0x65, 0x6d, 0x70, 0x73, 0x22, 0x41, 0x0a, 0x0d, 0x44, 0x53, 0x54,
	0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x54, 0x65,
	0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x0b, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x32, 0xe2, 0x01, 0x0a,
	0x02, 0x44, 0x53, 0x12, 0x3f, 0x0a, 0x05, 0x44, 0x53, 0x47, 0x65, 0x74, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1c, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x73, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0b, 0x44, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x65, 0x12, 0x1b, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x1a, 0x1b, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x00, 0x12,
	0x50, 0x0a, 0x11, 0x44, 0x53, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x21, 0x2e, 0x64,
	0x69, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x44, 0x53, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x22,
	0x00, 0x42, 0x45, 0x50, 0x01, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x2d, 0x63, 0x6c, 0x61, 0x70, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x69, 0x6c, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x69, 0x6c,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_distillation_distillationproto_ds18b20_proto_rawDescOnce sync.Once
	file_pkg_distillation_distillationproto_ds18b20_proto_rawDescData = file_pkg_distillation_distillationproto_ds18b20_proto_rawDesc
)

func file_pkg_distillation_distillationproto_ds18b20_proto_rawDescGZIP() []byte {
	file_pkg_distillation_distillationproto_ds18b20_proto_rawDescOnce.Do(func() {
		file_pkg_distillation_distillationproto_ds18b20_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_distillation_distillationproto_ds18b20_proto_rawDescData)
	})
	return file_pkg_distillation_distillationproto_ds18b20_proto_rawDescData
}

var file_pkg_distillation_distillationproto_ds18b20_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_distillation_distillationproto_ds18b20_proto_goTypes = []interface{}{
	(*DSConfigs)(nil),      // 0: distillationproto.DSConfigs
	(*DSConfig)(nil),       // 1: distillationproto.DSConfig
	(*DSTemperatures)(nil), // 2: distillationproto.DSTemperatures
	(*DSTemperature)(nil),  // 3: distillationproto.DSTemperature
	(*empty.Empty)(nil),    // 4: google.protobuf.Empty
}
var file_pkg_distillation_distillationproto_ds18b20_proto_depIdxs = []int32{
	1, // 0: distillationproto.DSConfigs.configs:type_name -> distillationproto.DSConfig
	3, // 1: distillationproto.DSTemperatures.temps:type_name -> distillationproto.DSTemperature
	4, // 2: distillationproto.DS.DSGet:input_type -> google.protobuf.Empty
	1, // 3: distillationproto.DS.DSConfigure:input_type -> distillationproto.DSConfig
	4, // 4: distillationproto.DS.DSGetTemperatures:input_type -> google.protobuf.Empty
	0, // 5: distillationproto.DS.DSGet:output_type -> distillationproto.DSConfigs
	1, // 6: distillationproto.DS.DSConfigure:output_type -> distillationproto.DSConfig
	2, // 7: distillationproto.DS.DSGetTemperatures:output_type -> distillationproto.DSTemperatures
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_distillation_distillationproto_ds18b20_proto_init() }
func file_pkg_distillation_distillationproto_ds18b20_proto_init() {
	if File_pkg_distillation_distillationproto_ds18b20_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_distillation_distillationproto_ds18b20_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DSConfigs); i {
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
		file_pkg_distillation_distillationproto_ds18b20_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DSConfig); i {
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
		file_pkg_distillation_distillationproto_ds18b20_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DSTemperatures); i {
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
		file_pkg_distillation_distillationproto_ds18b20_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DSTemperature); i {
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
			RawDescriptor: file_pkg_distillation_distillationproto_ds18b20_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_distillation_distillationproto_ds18b20_proto_goTypes,
		DependencyIndexes: file_pkg_distillation_distillationproto_ds18b20_proto_depIdxs,
		MessageInfos:      file_pkg_distillation_distillationproto_ds18b20_proto_msgTypes,
	}.Build()
	File_pkg_distillation_distillationproto_ds18b20_proto = out.File
	file_pkg_distillation_distillationproto_ds18b20_proto_rawDesc = nil
	file_pkg_distillation_distillationproto_ds18b20_proto_goTypes = nil
	file_pkg_distillation_distillationproto_ds18b20_proto_depIdxs = nil
}
