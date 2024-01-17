// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: entry.proto

package index

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

type InscriptionEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Charms            uint32  `protobuf:"varint,1,opt,name=charms,proto3" json:"charms,omitempty"`
	Fee               uint64  `protobuf:"varint,2,opt,name=fee,proto3" json:"fee,omitempty"`
	Height            uint64  `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	Id                []byte  `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
	InscriptionNumber int64   `protobuf:"varint,5,opt,name=inscription_number,json=inscriptionNumber,proto3" json:"inscription_number,omitempty"`
	Sat               *uint64 `protobuf:"varint,6,opt,name=sat,proto3,oneof" json:"sat,omitempty"`
	SequenceNumber    int64   `protobuf:"varint,7,opt,name=sequence_number,json=sequenceNumber,proto3" json:"sequence_number,omitempty"`
	Timestamp         int64   `protobuf:"varint,8,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *InscriptionEntry) Reset() {
	*x = InscriptionEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InscriptionEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InscriptionEntry) ProtoMessage() {}

func (x *InscriptionEntry) ProtoReflect() protoreflect.Message {
	mi := &file_entry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InscriptionEntry.ProtoReflect.Descriptor instead.
func (*InscriptionEntry) Descriptor() ([]byte, []int) {
	return file_entry_proto_rawDescGZIP(), []int{0}
}

func (x *InscriptionEntry) GetCharms() uint32 {
	if x != nil {
		return x.Charms
	}
	return 0
}

func (x *InscriptionEntry) GetFee() uint64 {
	if x != nil {
		return x.Fee
	}
	return 0
}

func (x *InscriptionEntry) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *InscriptionEntry) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *InscriptionEntry) GetInscriptionNumber() int64 {
	if x != nil {
		return x.InscriptionNumber
	}
	return 0
}

func (x *InscriptionEntry) GetSat() uint64 {
	if x != nil && x.Sat != nil {
		return *x.Sat
	}
	return 0
}

func (x *InscriptionEntry) GetSequenceNumber() int64 {
	if x != nil {
		return x.SequenceNumber
	}
	return 0
}

func (x *InscriptionEntry) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

var File_entry_proto protoreflect.FileDescriptor

var file_entry_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x22, 0xf9, 0x01, 0x0a, 0x10, 0x49, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x68, 0x61,
	0x72, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x63, 0x68, 0x61, 0x72, 0x6d,
	0x73, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x65, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03,
	0x66, 0x65, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2d, 0x0a, 0x12, 0x69,
	0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x69, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x15, 0x0a, 0x03, 0x73, 0x61,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x03, 0x73, 0x61, 0x74, 0x88, 0x01,
	0x01, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x73, 0x65, 0x71, 0x75,
	0x65, 0x6e, 0x63, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x73, 0x61, 0x74,
	0x42, 0x08, 0x5a, 0x06, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_entry_proto_rawDescOnce sync.Once
	file_entry_proto_rawDescData = file_entry_proto_rawDesc
)

func file_entry_proto_rawDescGZIP() []byte {
	file_entry_proto_rawDescOnce.Do(func() {
		file_entry_proto_rawDescData = protoimpl.X.CompressGZIP(file_entry_proto_rawDescData)
	})
	return file_entry_proto_rawDescData
}

var file_entry_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_entry_proto_goTypes = []interface{}{
	(*InscriptionEntry)(nil), // 0: index.InscriptionEntry
}
var file_entry_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_entry_proto_init() }
func file_entry_proto_init() {
	if File_entry_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_entry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InscriptionEntry); i {
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
	file_entry_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_entry_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_entry_proto_goTypes,
		DependencyIndexes: file_entry_proto_depIdxs,
		MessageInfos:      file_entry_proto_msgTypes,
	}.Build()
	File_entry_proto = out.File
	file_entry_proto_rawDesc = nil
	file_entry_proto_goTypes = nil
	file_entry_proto_depIdxs = nil
}
