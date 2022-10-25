// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: storage/resource_collection.proto

package storage

import (
	_ "github.com/gogo/protobuf/gogoproto"
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

type ResourceCollection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // primary key
	Name        string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	LastUpdated *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	CreatedBy   *SlimUser              `protobuf:"bytes,6,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	UpdatedBy   *SlimUser              `protobuf:"bytes,7,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by,omitempty"`
	// `resource_selectors` resolve as disjunction (OR) with each-other and with selectors from `embedded_collections`. For MVP, the size of resource_selectors will at most be 1 from UX standpoint.
	ResourceSelectors   []*ResourceSelector                              `protobuf:"bytes,8,rep,name=resource_selectors,json=resourceSelectors,proto3" json:"resource_selectors,omitempty"`
	EmbeddedCollections []*ResourceCollection_EmbeddedResourceCollection `protobuf:"bytes,9,rep,name=embedded_collections,json=embeddedCollections,proto3" json:"embedded_collections,omitempty"`
}

func (x *ResourceCollection) Reset() {
	*x = ResourceCollection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_resource_collection_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceCollection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceCollection) ProtoMessage() {}

func (x *ResourceCollection) ProtoReflect() protoreflect.Message {
	mi := &file_storage_resource_collection_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceCollection.ProtoReflect.Descriptor instead.
func (*ResourceCollection) Descriptor() ([]byte, []int) {
	return file_storage_resource_collection_proto_rawDescGZIP(), []int{0}
}

func (x *ResourceCollection) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ResourceCollection) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResourceCollection) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ResourceCollection) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ResourceCollection) GetLastUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUpdated
	}
	return nil
}

func (x *ResourceCollection) GetCreatedBy() *SlimUser {
	if x != nil {
		return x.CreatedBy
	}
	return nil
}

func (x *ResourceCollection) GetUpdatedBy() *SlimUser {
	if x != nil {
		return x.UpdatedBy
	}
	return nil
}

func (x *ResourceCollection) GetResourceSelectors() []*ResourceSelector {
	if x != nil {
		return x.ResourceSelectors
	}
	return nil
}

func (x *ResourceCollection) GetEmbeddedCollections() []*ResourceCollection_EmbeddedResourceCollection {
	if x != nil {
		return x.EmbeddedCollections
	}
	return nil
}

type ResourceSelector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// `rules` resolve as a conjunction (AND).
	Rules []*SelectorRule `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
}

func (x *ResourceSelector) Reset() {
	*x = ResourceSelector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_resource_collection_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceSelector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceSelector) ProtoMessage() {}

func (x *ResourceSelector) ProtoReflect() protoreflect.Message {
	mi := &file_storage_resource_collection_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceSelector.ProtoReflect.Descriptor instead.
func (*ResourceSelector) Descriptor() ([]byte, []int) {
	return file_storage_resource_collection_proto_rawDescGZIP(), []int{1}
}

func (x *ResourceSelector) GetRules() []*SelectorRule {
	if x != nil {
		return x.Rules
	}
	return nil
}

type SelectorRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// `field_name` can be one of the following:
	// - Cluster
	// - Cluster Label
	// - Namespace
	// - Namespace Label
	// - Namespace Annotation
	// - Deployment
	// - Deployment Label
	// - Deployment Annotation
	FieldName string `protobuf:"bytes,1,opt,name=field_name,json=fieldName,proto3" json:"field_name,omitempty"`
	// 'operator' only supports disjunction (OR) currently
	Operator BooleanOperator `protobuf:"varint,2,opt,name=operator,proto3,enum=storage.BooleanOperator" json:"operator,omitempty"`
	// `values` resolve as a conjunction (AND) or disjunction (OR) depending on operator. For MVP, only OR is supported from UX standpoint.
	Values []*RuleValue `protobuf:"bytes,3,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *SelectorRule) Reset() {
	*x = SelectorRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_resource_collection_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelectorRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectorRule) ProtoMessage() {}

func (x *SelectorRule) ProtoReflect() protoreflect.Message {
	mi := &file_storage_resource_collection_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectorRule.ProtoReflect.Descriptor instead.
func (*SelectorRule) Descriptor() ([]byte, []int) {
	return file_storage_resource_collection_proto_rawDescGZIP(), []int{2}
}

func (x *SelectorRule) GetFieldName() string {
	if x != nil {
		return x.FieldName
	}
	return ""
}

func (x *SelectorRule) GetOperator() BooleanOperator {
	if x != nil {
		return x.Operator
	}
	return BooleanOperator_OR
}

func (x *SelectorRule) GetValues() []*RuleValue {
	if x != nil {
		return x.Values
	}
	return nil
}

type RuleValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *RuleValue) Reset() {
	*x = RuleValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_resource_collection_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RuleValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuleValue) ProtoMessage() {}

func (x *RuleValue) ProtoReflect() protoreflect.Message {
	mi := &file_storage_resource_collection_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RuleValue.ProtoReflect.Descriptor instead.
func (*RuleValue) Descriptor() ([]byte, []int) {
	return file_storage_resource_collection_proto_rawDescGZIP(), []int{3}
}

func (x *RuleValue) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type ResourceCollection_EmbeddedResourceCollection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 'id' is searchable to force a separate table
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ResourceCollection_EmbeddedResourceCollection) Reset() {
	*x = ResourceCollection_EmbeddedResourceCollection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_resource_collection_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceCollection_EmbeddedResourceCollection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceCollection_EmbeddedResourceCollection) ProtoMessage() {}

func (x *ResourceCollection_EmbeddedResourceCollection) ProtoReflect() protoreflect.Message {
	mi := &file_storage_resource_collection_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceCollection_EmbeddedResourceCollection.ProtoReflect.Descriptor instead.
func (*ResourceCollection_EmbeddedResourceCollection) Descriptor() ([]byte, []int) {
	return file_storage_resource_collection_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ResourceCollection_EmbeddedResourceCollection) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_storage_resource_collection_proto protoreflect.FileDescriptor

var file_storage_resource_collection_proto_rawDesc = []byte{
	0x0a, 0x21, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x67,
	0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb1, 0x05,
	0x0a, 0x12, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0c, 0xf2, 0xde, 0x1f, 0x08, 0x73, 0x71, 0x6c, 0x3a, 0x22, 0x70, 0x6b, 0x22, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x3d, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x29, 0xf2, 0xde, 0x1f, 0x25, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x3a, 0x22, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x20, 0x73,
	0x71, 0x6c, 0x3a, 0x22, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x22, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3d,
	0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x30, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x6c, 0x69, 0x6d,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12,
	0x30, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x6c,
	0x69, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42,
	0x79, 0x12, 0x48, 0x0a, 0x12, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x73, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x11, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x69, 0x0a, 0x14, 0x65,
	0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x13, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x43, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x88, 0x01, 0x0a, 0x1a, 0x45, 0x6d, 0x62, 0x65, 0x64,
	0x64, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x6a, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x5a, 0xf2, 0xde, 0x1f, 0x56, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x3a, 0x22, 0x45,
	0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x20, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x20, 0x49, 0x44, 0x2c, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x22, 0x20, 0x73, 0x71,
	0x6c, 0x3a, 0x22, 0x66, 0x6b, 0x28, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x69, 0x64, 0x29, 0x2c, 0x72, 0x65, 0x73,
	0x74, 0x72, 0x69, 0x63, 0x74, 0x2d, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x22, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x3f, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x2b, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x53,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x75, 0x6c,
	0x65, 0x73, 0x22, 0x8f, 0x01, 0x0a, 0x0c, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52,
	0x75, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x34, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x42,
	0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x08,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x2a, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x22, 0x21, 0x0a, 0x09, 0x52, 0x75, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x24, 0x0a, 0x19, 0x69, 0x6f, 0x2e, 0x73, 0x74,
	0x61, 0x63, 0x6b, 0x72, 0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x5a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_storage_resource_collection_proto_rawDescOnce sync.Once
	file_storage_resource_collection_proto_rawDescData = file_storage_resource_collection_proto_rawDesc
)

func file_storage_resource_collection_proto_rawDescGZIP() []byte {
	file_storage_resource_collection_proto_rawDescOnce.Do(func() {
		file_storage_resource_collection_proto_rawDescData = protoimpl.X.CompressGZIP(file_storage_resource_collection_proto_rawDescData)
	})
	return file_storage_resource_collection_proto_rawDescData
}

var file_storage_resource_collection_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_storage_resource_collection_proto_goTypes = []interface{}{
	(*ResourceCollection)(nil),                            // 0: storage.ResourceCollection
	(*ResourceSelector)(nil),                              // 1: storage.ResourceSelector
	(*SelectorRule)(nil),                                  // 2: storage.SelectorRule
	(*RuleValue)(nil),                                     // 3: storage.RuleValue
	(*ResourceCollection_EmbeddedResourceCollection)(nil), // 4: storage.ResourceCollection.EmbeddedResourceCollection
	(*timestamppb.Timestamp)(nil),                         // 5: google.protobuf.Timestamp
	(*SlimUser)(nil),                                      // 6: storage.SlimUser
	(BooleanOperator)(0),                                  // 7: storage.BooleanOperator
}
var file_storage_resource_collection_proto_depIdxs = []int32{
	5, // 0: storage.ResourceCollection.created_at:type_name -> google.protobuf.Timestamp
	5, // 1: storage.ResourceCollection.last_updated:type_name -> google.protobuf.Timestamp
	6, // 2: storage.ResourceCollection.created_by:type_name -> storage.SlimUser
	6, // 3: storage.ResourceCollection.updated_by:type_name -> storage.SlimUser
	1, // 4: storage.ResourceCollection.resource_selectors:type_name -> storage.ResourceSelector
	4, // 5: storage.ResourceCollection.embedded_collections:type_name -> storage.ResourceCollection.EmbeddedResourceCollection
	2, // 6: storage.ResourceSelector.rules:type_name -> storage.SelectorRule
	7, // 7: storage.SelectorRule.operator:type_name -> storage.BooleanOperator
	3, // 8: storage.SelectorRule.values:type_name -> storage.RuleValue
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_storage_resource_collection_proto_init() }
func file_storage_resource_collection_proto_init() {
	if File_storage_resource_collection_proto != nil {
		return
	}
	file_storage_user_proto_init()
	file_storage_policy_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_storage_resource_collection_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceCollection); i {
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
		file_storage_resource_collection_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceSelector); i {
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
		file_storage_resource_collection_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SelectorRule); i {
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
		file_storage_resource_collection_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RuleValue); i {
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
		file_storage_resource_collection_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceCollection_EmbeddedResourceCollection); i {
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
			RawDescriptor: file_storage_resource_collection_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_storage_resource_collection_proto_goTypes,
		DependencyIndexes: file_storage_resource_collection_proto_depIdxs,
		MessageInfos:      file_storage_resource_collection_proto_msgTypes,
	}.Build()
	File_storage_resource_collection_proto = out.File
	file_storage_resource_collection_proto_rawDesc = nil
	file_storage_resource_collection_proto_goTypes = nil
	file_storage_resource_collection_proto_depIdxs = nil
}
