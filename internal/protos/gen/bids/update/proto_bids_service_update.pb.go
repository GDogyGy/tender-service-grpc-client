// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: internal/protos/proto/bids/proto_bids_service_update.proto

package update

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BidEditV1 struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Description   string                 `protobuf:"bytes,2,opt,name=Description,proto3" json:"Description,omitempty"`
	Name          string                 `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	Responsible   string                 `protobuf:"bytes,4,opt,name=Responsible,proto3" json:"Responsible,omitempty"`
	Status        string                 `protobuf:"bytes,5,opt,name=Status,proto3" json:"Status,omitempty"`
	TenderId      string                 `protobuf:"bytes,6,opt,name=TenderId,proto3" json:"TenderId,omitempty"`
	Version       int32                  `protobuf:"varint,7,opt,name=Version,proto3" json:"Version,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BidEditV1) Reset() {
	*x = BidEditV1{}
	mi := &file_internal_protos_proto_bids_proto_bids_service_update_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BidEditV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidEditV1) ProtoMessage() {}

func (x *BidEditV1) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protos_proto_bids_proto_bids_service_update_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidEditV1.ProtoReflect.Descriptor instead.
func (*BidEditV1) Descriptor() ([]byte, []int) {
	return file_internal_protos_proto_bids_proto_bids_service_update_proto_rawDescGZIP(), []int{0}
}

func (x *BidEditV1) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BidEditV1) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *BidEditV1) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BidEditV1) GetResponsible() string {
	if x != nil {
		return x.Responsible
	}
	return ""
}

func (x *BidEditV1) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *BidEditV1) GetTenderId() string {
	if x != nil {
		return x.TenderId
	}
	return ""
}

func (x *BidEditV1) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

type BidsRequestEditV1 struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	BidID         string                 `protobuf:"bytes,2,opt,name=BidID,proto3" json:"BidID,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	Name          string                 `protobuf:"bytes,4,opt,name=Name,proto3" json:"Name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BidsRequestEditV1) Reset() {
	*x = BidsRequestEditV1{}
	mi := &file_internal_protos_proto_bids_proto_bids_service_update_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BidsRequestEditV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidsRequestEditV1) ProtoMessage() {}

func (x *BidsRequestEditV1) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protos_proto_bids_proto_bids_service_update_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidsRequestEditV1.ProtoReflect.Descriptor instead.
func (*BidsRequestEditV1) Descriptor() ([]byte, []int) {
	return file_internal_protos_proto_bids_proto_bids_service_update_proto_rawDescGZIP(), []int{1}
}

func (x *BidsRequestEditV1) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *BidsRequestEditV1) GetBidID() string {
	if x != nil {
		return x.BidID
	}
	return ""
}

func (x *BidsRequestEditV1) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *BidsRequestEditV1) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type BidsRequestRollbackV1 struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	BidID         string                 `protobuf:"bytes,2,opt,name=BidID,proto3" json:"BidID,omitempty"`
	Version       string                 `protobuf:"bytes,3,opt,name=Version,proto3" json:"Version,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BidsRequestRollbackV1) Reset() {
	*x = BidsRequestRollbackV1{}
	mi := &file_internal_protos_proto_bids_proto_bids_service_update_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BidsRequestRollbackV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidsRequestRollbackV1) ProtoMessage() {}

func (x *BidsRequestRollbackV1) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protos_proto_bids_proto_bids_service_update_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidsRequestRollbackV1.ProtoReflect.Descriptor instead.
func (*BidsRequestRollbackV1) Descriptor() ([]byte, []int) {
	return file_internal_protos_proto_bids_proto_bids_service_update_proto_rawDescGZIP(), []int{2}
}

func (x *BidsRequestRollbackV1) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *BidsRequestRollbackV1) GetBidID() string {
	if x != nil {
		return x.BidID
	}
	return ""
}

func (x *BidsRequestRollbackV1) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type BidsRequestStatusV1 struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	BidID         string                 `protobuf:"bytes,2,opt,name=BidID,proto3" json:"BidID,omitempty"`
	Status        string                 `protobuf:"bytes,3,opt,name=Status,proto3" json:"Status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BidsRequestStatusV1) Reset() {
	*x = BidsRequestStatusV1{}
	mi := &file_internal_protos_proto_bids_proto_bids_service_update_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BidsRequestStatusV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidsRequestStatusV1) ProtoMessage() {}

func (x *BidsRequestStatusV1) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protos_proto_bids_proto_bids_service_update_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidsRequestStatusV1.ProtoReflect.Descriptor instead.
func (*BidsRequestStatusV1) Descriptor() ([]byte, []int) {
	return file_internal_protos_proto_bids_proto_bids_service_update_proto_rawDescGZIP(), []int{3}
}

func (x *BidsRequestStatusV1) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *BidsRequestStatusV1) GetBidID() string {
	if x != nil {
		return x.BidID
	}
	return ""
}

func (x *BidsRequestStatusV1) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type BidsRequestSubmitDecisionV1 struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Username       string                 `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	BidID          string                 `protobuf:"bytes,2,opt,name=BidID,proto3" json:"BidID,omitempty"`
	Decision       string                 `protobuf:"bytes,3,opt,name=Decision,proto3" json:"Decision,omitempty"`
	OrganizationID string                 `protobuf:"bytes,4,opt,name=OrganizationID,proto3" json:"OrganizationID,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *BidsRequestSubmitDecisionV1) Reset() {
	*x = BidsRequestSubmitDecisionV1{}
	mi := &file_internal_protos_proto_bids_proto_bids_service_update_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BidsRequestSubmitDecisionV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidsRequestSubmitDecisionV1) ProtoMessage() {}

func (x *BidsRequestSubmitDecisionV1) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protos_proto_bids_proto_bids_service_update_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidsRequestSubmitDecisionV1.ProtoReflect.Descriptor instead.
func (*BidsRequestSubmitDecisionV1) Descriptor() ([]byte, []int) {
	return file_internal_protos_proto_bids_proto_bids_service_update_proto_rawDescGZIP(), []int{4}
}

func (x *BidsRequestSubmitDecisionV1) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *BidsRequestSubmitDecisionV1) GetBidID() string {
	if x != nil {
		return x.BidID
	}
	return ""
}

func (x *BidsRequestSubmitDecisionV1) GetDecision() string {
	if x != nil {
		return x.Decision
	}
	return ""
}

func (x *BidsRequestSubmitDecisionV1) GetOrganizationID() string {
	if x != nil {
		return x.OrganizationID
	}
	return ""
}

var File_internal_protos_proto_bids_proto_bids_service_update_proto protoreflect.FileDescriptor

const file_internal_protos_proto_bids_proto_bids_service_update_proto_rawDesc = "" +
	"\n" +
	":internal/protos/proto/bids/proto_bids_service_update.proto\x12\x06update\"\xc1\x01\n" +
	"\tBidEditV1\x12\x0e\n" +
	"\x02Id\x18\x01 \x01(\tR\x02Id\x12 \n" +
	"\vDescription\x18\x02 \x01(\tR\vDescription\x12\x12\n" +
	"\x04Name\x18\x03 \x01(\tR\x04Name\x12 \n" +
	"\vResponsible\x18\x04 \x01(\tR\vResponsible\x12\x16\n" +
	"\x06Status\x18\x05 \x01(\tR\x06Status\x12\x1a\n" +
	"\bTenderId\x18\x06 \x01(\tR\bTenderId\x12\x18\n" +
	"\aVersion\x18\a \x01(\x05R\aVersion\"{\n" +
	"\x11BidsRequestEditV1\x12\x1a\n" +
	"\bUsername\x18\x01 \x01(\tR\bUsername\x12\x14\n" +
	"\x05BidID\x18\x02 \x01(\tR\x05BidID\x12 \n" +
	"\vDescription\x18\x03 \x01(\tR\vDescription\x12\x12\n" +
	"\x04Name\x18\x04 \x01(\tR\x04Name\"c\n" +
	"\x15BidsRequestRollbackV1\x12\x1a\n" +
	"\bUsername\x18\x01 \x01(\tR\bUsername\x12\x14\n" +
	"\x05BidID\x18\x02 \x01(\tR\x05BidID\x12\x18\n" +
	"\aVersion\x18\x03 \x01(\tR\aVersion\"_\n" +
	"\x13BidsRequestStatusV1\x12\x1a\n" +
	"\bUsername\x18\x01 \x01(\tR\bUsername\x12\x14\n" +
	"\x05BidID\x18\x02 \x01(\tR\x05BidID\x12\x16\n" +
	"\x06Status\x18\x03 \x01(\tR\x06Status\"\x93\x01\n" +
	"\x1bBidsRequestSubmitDecisionV1\x12\x1a\n" +
	"\bUsername\x18\x01 \x01(\tR\bUsername\x12\x14\n" +
	"\x05BidID\x18\x02 \x01(\tR\x05BidID\x12\x1a\n" +
	"\bDecision\x18\x03 \x01(\tR\bDecision\x12&\n" +
	"\x0eOrganizationID\x18\x04 \x01(\tR\x0eOrganizationID2\x8a\x02\n" +
	"\x10BidsServiceFetch\x124\n" +
	"\x04Edit\x12\x19.update.BidsRequestEditV1\x1a\x11.update.BidEditV1\x12<\n" +
	"\bRollback\x12\x1d.update.BidsRequestRollbackV1\x1a\x11.update.BidEditV1\x128\n" +
	"\x06Status\x12\x1b.update.BidsRequestStatusV1\x1a\x11.update.BidEditV1\x12H\n" +
	"\x0eSubmitDecision\x12#.update.BidsRequestSubmitDecisionV1\x1a\x11.update.BidEditV1B\x12Z\x10/gen/bids/updateb\x06proto3"

var (
	file_internal_protos_proto_bids_proto_bids_service_update_proto_rawDescOnce sync.Once
	file_internal_protos_proto_bids_proto_bids_service_update_proto_rawDescData []byte
)

func file_internal_protos_proto_bids_proto_bids_service_update_proto_rawDescGZIP() []byte {
	file_internal_protos_proto_bids_proto_bids_service_update_proto_rawDescOnce.Do(func() {
		file_internal_protos_proto_bids_proto_bids_service_update_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_internal_protos_proto_bids_proto_bids_service_update_proto_rawDesc), len(file_internal_protos_proto_bids_proto_bids_service_update_proto_rawDesc)))
	})
	return file_internal_protos_proto_bids_proto_bids_service_update_proto_rawDescData
}

var file_internal_protos_proto_bids_proto_bids_service_update_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_internal_protos_proto_bids_proto_bids_service_update_proto_goTypes = []any{
	(*BidEditV1)(nil),                   // 0: update.BidEditV1
	(*BidsRequestEditV1)(nil),           // 1: update.BidsRequestEditV1
	(*BidsRequestRollbackV1)(nil),       // 2: update.BidsRequestRollbackV1
	(*BidsRequestStatusV1)(nil),         // 3: update.BidsRequestStatusV1
	(*BidsRequestSubmitDecisionV1)(nil), // 4: update.BidsRequestSubmitDecisionV1
}
var file_internal_protos_proto_bids_proto_bids_service_update_proto_depIdxs = []int32{
	1, // 0: update.BidsServiceFetch.Edit:input_type -> update.BidsRequestEditV1
	2, // 1: update.BidsServiceFetch.Rollback:input_type -> update.BidsRequestRollbackV1
	3, // 2: update.BidsServiceFetch.Status:input_type -> update.BidsRequestStatusV1
	4, // 3: update.BidsServiceFetch.SubmitDecision:input_type -> update.BidsRequestSubmitDecisionV1
	0, // 4: update.BidsServiceFetch.Edit:output_type -> update.BidEditV1
	0, // 5: update.BidsServiceFetch.Rollback:output_type -> update.BidEditV1
	0, // 6: update.BidsServiceFetch.Status:output_type -> update.BidEditV1
	0, // 7: update.BidsServiceFetch.SubmitDecision:output_type -> update.BidEditV1
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_protos_proto_bids_proto_bids_service_update_proto_init() }
func file_internal_protos_proto_bids_proto_bids_service_update_proto_init() {
	if File_internal_protos_proto_bids_proto_bids_service_update_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_internal_protos_proto_bids_proto_bids_service_update_proto_rawDesc), len(file_internal_protos_proto_bids_proto_bids_service_update_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_protos_proto_bids_proto_bids_service_update_proto_goTypes,
		DependencyIndexes: file_internal_protos_proto_bids_proto_bids_service_update_proto_depIdxs,
		MessageInfos:      file_internal_protos_proto_bids_proto_bids_service_update_proto_msgTypes,
	}.Build()
	File_internal_protos_proto_bids_proto_bids_service_update_proto = out.File
	file_internal_protos_proto_bids_proto_bids_service_update_proto_goTypes = nil
	file_internal_protos_proto_bids_proto_bids_service_update_proto_depIdxs = nil
}
