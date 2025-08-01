// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: apptest/tenant.proto

package apptest

import (
	_ "github.com/protobuf-orm/protobuf-orm/ormpb"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Tenant struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Id          []byte                 `protobuf:"bytes,1,opt,name=id"`
	xxx_hidden_Alias       string                 `protobuf:"bytes,4,opt,name=alias"`
	xxx_hidden_Name        string                 `protobuf:"bytes,5,opt,name=name"`
	xxx_hidden_Labels      map[string]string      `protobuf:"bytes,7,rep,name=labels" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	xxx_hidden_DateCreated *timestamppb.Timestamp `protobuf:"bytes,15,opt,name=date_created,json=dateCreated"`
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *Tenant) Reset() {
	*x = Tenant{}
	mi := &file_apptest_tenant_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Tenant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tenant) ProtoMessage() {}

func (x *Tenant) ProtoReflect() protoreflect.Message {
	mi := &file_apptest_tenant_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Tenant) GetId() []byte {
	if x != nil {
		return x.xxx_hidden_Id
	}
	return nil
}

func (x *Tenant) GetAlias() string {
	if x != nil {
		return x.xxx_hidden_Alias
	}
	return ""
}

func (x *Tenant) GetName() string {
	if x != nil {
		return x.xxx_hidden_Name
	}
	return ""
}

func (x *Tenant) GetLabels() map[string]string {
	if x != nil {
		return x.xxx_hidden_Labels
	}
	return nil
}

func (x *Tenant) GetDateCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.xxx_hidden_DateCreated
	}
	return nil
}

func (x *Tenant) SetId(v []byte) {
	if v == nil {
		v = []byte{}
	}
	x.xxx_hidden_Id = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 5)
}

func (x *Tenant) SetAlias(v string) {
	x.xxx_hidden_Alias = v
}

func (x *Tenant) SetName(v string) {
	x.xxx_hidden_Name = v
}

func (x *Tenant) SetLabels(v map[string]string) {
	x.xxx_hidden_Labels = v
}

func (x *Tenant) SetDateCreated(v *timestamppb.Timestamp) {
	x.xxx_hidden_DateCreated = v
}

func (x *Tenant) HasId() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *Tenant) HasDateCreated() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_DateCreated != nil
}

func (x *Tenant) ClearId() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_Id = nil
}

func (x *Tenant) ClearDateCreated() {
	x.xxx_hidden_DateCreated = nil
}

type Tenant_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Id          []byte
	Alias       string
	Name        string
	Labels      map[string]string
	DateCreated *timestamppb.Timestamp
}

func (b0 Tenant_builder) Build() *Tenant {
	m0 := &Tenant{}
	b, x := &b0, m0
	_, _ = b, x
	if b.Id != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 5)
		x.xxx_hidden_Id = b.Id
	}
	x.xxx_hidden_Alias = b.Alias
	x.xxx_hidden_Name = b.Name
	x.xxx_hidden_Labels = b.Labels
	x.xxx_hidden_DateCreated = b.DateCreated
	return m0
}

var File_apptest_tenant_proto protoreflect.FileDescriptor

const file_apptest_tenant_proto_rawDesc = "" +
	"\n" +
	"\x14apptest/tenant.proto\x12\aapptest\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\torm.proto\"\xac\x02\n" +
	"\x06Tenant\x12\x18\n" +
	"\x02id\x18\x01 \x01(\fB\b\xea\x82\x16\x04\x10@(\x01R\x02id\x12\"\n" +
	"\x05alias\x18\x04 \x01(\tB\f\xea\x82\x16\x03\x82\x01\x00\xaa\x01\x02\b\x02R\x05alias\x12 \n" +
	"\x04name\x18\x05 \x01(\tB\f\xea\x82\x16\x03\x82\x01\x00\xaa\x01\x02\b\x02R\x04name\x123\n" +
	"\x06labels\x18\a \x03(\v2\x1b.apptest.Tenant.LabelsEntryR\x06labels\x12H\n" +
	"\fdate_created\x18\x0f \x01(\v2\x1a.google.protobuf.TimestampB\t\xea\x82\x16\x05@\x01\x82\x01\x00R\vdateCreated\x1a9\n" +
	"\vLabelsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01:\b\xca\xfc\x15\x04\x12\x02\x10\x01BAZ?github.com/protobuf-orm/protoc-gen-orm-service/internal/apptestb\beditionsp\xe8\a"

var file_apptest_tenant_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_apptest_tenant_proto_goTypes = []any{
	(*Tenant)(nil),                // 0: apptest.Tenant
	nil,                           // 1: apptest.Tenant.LabelsEntry
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_apptest_tenant_proto_depIdxs = []int32{
	1, // 0: apptest.Tenant.labels:type_name -> apptest.Tenant.LabelsEntry
	2, // 1: apptest.Tenant.date_created:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_apptest_tenant_proto_init() }
func file_apptest_tenant_proto_init() {
	if File_apptest_tenant_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_apptest_tenant_proto_rawDesc), len(file_apptest_tenant_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apptest_tenant_proto_goTypes,
		DependencyIndexes: file_apptest_tenant_proto_depIdxs,
		MessageInfos:      file_apptest_tenant_proto_msgTypes,
	}.Build()
	File_apptest_tenant_proto = out.File
	file_apptest_tenant_proto_goTypes = nil
	file_apptest_tenant_proto_depIdxs = nil
}
