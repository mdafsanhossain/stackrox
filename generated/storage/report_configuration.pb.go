// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: storage/report_configuration.proto

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

type ReportConfiguration_ReportType int32

const (
	ReportConfiguration_VULNERABILITY ReportConfiguration_ReportType = 0
)

// Enum value maps for ReportConfiguration_ReportType.
var (
	ReportConfiguration_ReportType_name = map[int32]string{
		0: "VULNERABILITY",
	}
	ReportConfiguration_ReportType_value = map[string]int32{
		"VULNERABILITY": 0,
	}
)

func (x ReportConfiguration_ReportType) Enum() *ReportConfiguration_ReportType {
	p := new(ReportConfiguration_ReportType)
	*p = x
	return p
}

func (x ReportConfiguration_ReportType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReportConfiguration_ReportType) Descriptor() protoreflect.EnumDescriptor {
	return file_storage_report_configuration_proto_enumTypes[0].Descriptor()
}

func (ReportConfiguration_ReportType) Type() protoreflect.EnumType {
	return &file_storage_report_configuration_proto_enumTypes[0]
}

func (x ReportConfiguration_ReportType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReportConfiguration_ReportType.Descriptor instead.
func (ReportConfiguration_ReportType) EnumDescriptor() ([]byte, []int) {
	return file_storage_report_configuration_proto_rawDescGZIP(), []int{0, 0}
}

type ReportLastRunStatus_RunStatus int32

const (
	ReportLastRunStatus_SUCCESS ReportLastRunStatus_RunStatus = 0
	ReportLastRunStatus_FAILURE ReportLastRunStatus_RunStatus = 1
)

// Enum value maps for ReportLastRunStatus_RunStatus.
var (
	ReportLastRunStatus_RunStatus_name = map[int32]string{
		0: "SUCCESS",
		1: "FAILURE",
	}
	ReportLastRunStatus_RunStatus_value = map[string]int32{
		"SUCCESS": 0,
		"FAILURE": 1,
	}
)

func (x ReportLastRunStatus_RunStatus) Enum() *ReportLastRunStatus_RunStatus {
	p := new(ReportLastRunStatus_RunStatus)
	*p = x
	return p
}

func (x ReportLastRunStatus_RunStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReportLastRunStatus_RunStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_storage_report_configuration_proto_enumTypes[1].Descriptor()
}

func (ReportLastRunStatus_RunStatus) Type() protoreflect.EnumType {
	return &file_storage_report_configuration_proto_enumTypes[1]
}

func (x ReportLastRunStatus_RunStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReportLastRunStatus_RunStatus.Descriptor instead.
func (ReportLastRunStatus_RunStatus) EnumDescriptor() ([]byte, []int) {
	return file_storage_report_configuration_proto_rawDescGZIP(), []int{1, 0}
}

type VulnerabilityReportFilters_Fixability int32

const (
	VulnerabilityReportFilters_BOTH        VulnerabilityReportFilters_Fixability = 0
	VulnerabilityReportFilters_FIXABLE     VulnerabilityReportFilters_Fixability = 1
	VulnerabilityReportFilters_NOT_FIXABLE VulnerabilityReportFilters_Fixability = 2
)

// Enum value maps for VulnerabilityReportFilters_Fixability.
var (
	VulnerabilityReportFilters_Fixability_name = map[int32]string{
		0: "BOTH",
		1: "FIXABLE",
		2: "NOT_FIXABLE",
	}
	VulnerabilityReportFilters_Fixability_value = map[string]int32{
		"BOTH":        0,
		"FIXABLE":     1,
		"NOT_FIXABLE": 2,
	}
)

func (x VulnerabilityReportFilters_Fixability) Enum() *VulnerabilityReportFilters_Fixability {
	p := new(VulnerabilityReportFilters_Fixability)
	*p = x
	return p
}

func (x VulnerabilityReportFilters_Fixability) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VulnerabilityReportFilters_Fixability) Descriptor() protoreflect.EnumDescriptor {
	return file_storage_report_configuration_proto_enumTypes[2].Descriptor()
}

func (VulnerabilityReportFilters_Fixability) Type() protoreflect.EnumType {
	return &file_storage_report_configuration_proto_enumTypes[2]
}

func (x VulnerabilityReportFilters_Fixability) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VulnerabilityReportFilters_Fixability.Descriptor instead.
func (VulnerabilityReportFilters_Fixability) EnumDescriptor() ([]byte, []int) {
	return file_storage_report_configuration_proto_rawDescGZIP(), []int{2, 0}
}

type ReportConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string                         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string                         `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Type        ReportConfiguration_ReportType `protobuf:"varint,4,opt,name=type,proto3,enum=storage.ReportConfiguration_ReportType" json:"type,omitempty"`
	// Types that are assignable to Filter:
	//	*ReportConfiguration_VulnReportFilters
	Filter  isReportConfiguration_Filter `protobuf_oneof:"filter"`
	ScopeId string                       `protobuf:"bytes,6,opt,name=scope_id,json=scopeId,proto3" json:"scope_id,omitempty"`
	// Types that are assignable to NotifierConfig:
	//	*ReportConfiguration_EmailConfig
	NotifierConfig        isReportConfiguration_NotifierConfig `protobuf_oneof:"notifier_config"`
	Schedule              *Schedule                            `protobuf:"bytes,8,opt,name=schedule,proto3" json:"schedule,omitempty"`
	LastRunStatus         *ReportLastRunStatus                 `protobuf:"bytes,9,opt,name=last_run_status,json=lastRunStatus,proto3" json:"last_run_status,omitempty"`
	LastSuccessfulRunTime *timestamppb.Timestamp               `protobuf:"bytes,10,opt,name=last_successful_run_time,json=lastSuccessfulRunTime,proto3" json:"last_successful_run_time,omitempty"`
}

func (x *ReportConfiguration) Reset() {
	*x = ReportConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_report_configuration_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportConfiguration) ProtoMessage() {}

func (x *ReportConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_storage_report_configuration_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportConfiguration.ProtoReflect.Descriptor instead.
func (*ReportConfiguration) Descriptor() ([]byte, []int) {
	return file_storage_report_configuration_proto_rawDescGZIP(), []int{0}
}

func (x *ReportConfiguration) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ReportConfiguration) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ReportConfiguration) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ReportConfiguration) GetType() ReportConfiguration_ReportType {
	if x != nil {
		return x.Type
	}
	return ReportConfiguration_VULNERABILITY
}

func (m *ReportConfiguration) GetFilter() isReportConfiguration_Filter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (x *ReportConfiguration) GetVulnReportFilters() *VulnerabilityReportFilters {
	if x, ok := x.GetFilter().(*ReportConfiguration_VulnReportFilters); ok {
		return x.VulnReportFilters
	}
	return nil
}

func (x *ReportConfiguration) GetScopeId() string {
	if x != nil {
		return x.ScopeId
	}
	return ""
}

func (m *ReportConfiguration) GetNotifierConfig() isReportConfiguration_NotifierConfig {
	if m != nil {
		return m.NotifierConfig
	}
	return nil
}

func (x *ReportConfiguration) GetEmailConfig() *EmailNotifierConfiguration {
	if x, ok := x.GetNotifierConfig().(*ReportConfiguration_EmailConfig); ok {
		return x.EmailConfig
	}
	return nil
}

func (x *ReportConfiguration) GetSchedule() *Schedule {
	if x != nil {
		return x.Schedule
	}
	return nil
}

func (x *ReportConfiguration) GetLastRunStatus() *ReportLastRunStatus {
	if x != nil {
		return x.LastRunStatus
	}
	return nil
}

func (x *ReportConfiguration) GetLastSuccessfulRunTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastSuccessfulRunTime
	}
	return nil
}

type isReportConfiguration_Filter interface {
	isReportConfiguration_Filter()
}

type ReportConfiguration_VulnReportFilters struct {
	VulnReportFilters *VulnerabilityReportFilters `protobuf:"bytes,5,opt,name=vuln_report_filters,json=vulnReportFilters,proto3,oneof"`
}

func (*ReportConfiguration_VulnReportFilters) isReportConfiguration_Filter() {}

type isReportConfiguration_NotifierConfig interface {
	isReportConfiguration_NotifierConfig()
}

type ReportConfiguration_EmailConfig struct {
	EmailConfig *EmailNotifierConfiguration `protobuf:"bytes,7,opt,name=email_config,json=emailConfig,proto3,oneof"`
}

func (*ReportConfiguration_EmailConfig) isReportConfiguration_NotifierConfig() {}

type ReportLastRunStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReportStatus ReportLastRunStatus_RunStatus `protobuf:"varint,1,opt,name=report_status,json=reportStatus,proto3,enum=storage.ReportLastRunStatus_RunStatus" json:"report_status,omitempty"`
	LastRunTime  *timestamppb.Timestamp        `protobuf:"bytes,2,opt,name=last_run_time,json=lastRunTime,proto3" json:"last_run_time,omitempty"`
	ErrorMsg     string                        `protobuf:"bytes,3,opt,name=error_msg,json=errorMsg,proto3" json:"error_msg,omitempty"`
}

func (x *ReportLastRunStatus) Reset() {
	*x = ReportLastRunStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_report_configuration_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportLastRunStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportLastRunStatus) ProtoMessage() {}

func (x *ReportLastRunStatus) ProtoReflect() protoreflect.Message {
	mi := &file_storage_report_configuration_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportLastRunStatus.ProtoReflect.Descriptor instead.
func (*ReportLastRunStatus) Descriptor() ([]byte, []int) {
	return file_storage_report_configuration_proto_rawDescGZIP(), []int{1}
}

func (x *ReportLastRunStatus) GetReportStatus() ReportLastRunStatus_RunStatus {
	if x != nil {
		return x.ReportStatus
	}
	return ReportLastRunStatus_SUCCESS
}

func (x *ReportLastRunStatus) GetLastRunTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastRunTime
	}
	return nil
}

func (x *ReportLastRunStatus) GetErrorMsg() string {
	if x != nil {
		return x.ErrorMsg
	}
	return ""
}

type VulnerabilityReportFilters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fixability      VulnerabilityReportFilters_Fixability `protobuf:"varint,1,opt,name=fixability,proto3,enum=storage.VulnerabilityReportFilters_Fixability" json:"fixability,omitempty"`
	SinceLastReport bool                                  `protobuf:"varint,2,opt,name=since_last_report,json=sinceLastReport,proto3" json:"since_last_report,omitempty"`
	Severities      []VulnerabilitySeverity               `protobuf:"varint,3,rep,packed,name=severities,proto3,enum=storage.VulnerabilitySeverity" json:"severities,omitempty"`
}

func (x *VulnerabilityReportFilters) Reset() {
	*x = VulnerabilityReportFilters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_report_configuration_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VulnerabilityReportFilters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VulnerabilityReportFilters) ProtoMessage() {}

func (x *VulnerabilityReportFilters) ProtoReflect() protoreflect.Message {
	mi := &file_storage_report_configuration_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VulnerabilityReportFilters.ProtoReflect.Descriptor instead.
func (*VulnerabilityReportFilters) Descriptor() ([]byte, []int) {
	return file_storage_report_configuration_proto_rawDescGZIP(), []int{2}
}

func (x *VulnerabilityReportFilters) GetFixability() VulnerabilityReportFilters_Fixability {
	if x != nil {
		return x.Fixability
	}
	return VulnerabilityReportFilters_BOTH
}

func (x *VulnerabilityReportFilters) GetSinceLastReport() bool {
	if x != nil {
		return x.SinceLastReport
	}
	return false
}

func (x *VulnerabilityReportFilters) GetSeverities() []VulnerabilitySeverity {
	if x != nil {
		return x.Severities
	}
	return nil
}

type EmailNotifierConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NotifierId   string   `protobuf:"bytes,1,opt,name=notifier_id,json=notifierId,proto3" json:"notifier_id,omitempty"`
	MailingLists []string `protobuf:"bytes,2,rep,name=mailing_lists,json=mailingLists,proto3" json:"mailing_lists,omitempty"`
}

func (x *EmailNotifierConfiguration) Reset() {
	*x = EmailNotifierConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_report_configuration_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailNotifierConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailNotifierConfiguration) ProtoMessage() {}

func (x *EmailNotifierConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_storage_report_configuration_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailNotifierConfiguration.ProtoReflect.Descriptor instead.
func (*EmailNotifierConfiguration) Descriptor() ([]byte, []int) {
	return file_storage_report_configuration_proto_rawDescGZIP(), []int{3}
}

func (x *EmailNotifierConfiguration) GetNotifierId() string {
	if x != nil {
		return x.NotifierId
	}
	return ""
}

func (x *EmailNotifierConfiguration) GetMailingLists() []string {
	if x != nil {
		return x.MailingLists
	}
	return nil
}

var File_storage_report_configuration_proto protoreflect.FileDescriptor

var file_storage_report_configuration_proto_rawDesc = []byte{
	0x0a, 0x22, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x1a, 0x11, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x63, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x16, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x9e, 0x05, 0x0a, 0x13, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x0c, 0xf2, 0xde, 0x1f, 0x08, 0x73, 0x71, 0x6c, 0x3a, 0x22, 0x70, 0x6b,
	0x22, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x18, 0xf2, 0xde, 0x1f, 0x14, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x3a,
	0x22, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x20, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x55, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x42, 0x18, 0xf2, 0xde,
	0x1f, 0x14, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x3a, 0x22, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x20, 0x54, 0x79, 0x70, 0x65, 0x22, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x55, 0x0a, 0x13,
	0x76, 0x75, 0x6c, 0x6e, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2e, 0x56, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x48, 0x00,
	0x52, 0x11, 0x76, 0x75, 0x6c, 0x6e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x49, 0x64, 0x12, 0x48,
	0x0a, 0x0c, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x01, 0x52, 0x0b, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2d, 0x0a, 0x08, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x08, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x44, 0x0a, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x72, 0x75, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x4c, 0x61, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0d,
	0x6c, 0x61, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x53, 0x0a,
	0x18, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c,
	0x5f, 0x72, 0x75, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x15, 0x6c, 0x61, 0x73,
	0x74, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x52, 0x75, 0x6e, 0x54, 0x69,
	0x6d, 0x65, 0x22, 0x1f, 0x0a, 0x0a, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x11, 0x0a, 0x0d, 0x56, 0x55, 0x4c, 0x4e, 0x45, 0x52, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54,
	0x59, 0x10, 0x00, 0x42, 0x08, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x42, 0x11, 0x0a,
	0x0f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x22, 0xe6, 0x01, 0x0a, 0x13, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x61, 0x73, 0x74, 0x52,
	0x75, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x4b, 0x0a, 0x0d, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x26, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x4c, 0x61, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x52, 0x75,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0c, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3e, 0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x72, 0x75,
	0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x75,
	0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d,
	0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d,
	0x73, 0x67, 0x22, 0x25, 0x0a, 0x09, 0x52, 0x75, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x10, 0x01, 0x22, 0x8e, 0x02, 0x0a, 0x1a, 0x56, 0x75,
	0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x4e, 0x0a, 0x0a, 0x66, 0x69, 0x78, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x56, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x73, 0x2e, 0x46, 0x69, 0x78, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x0a, 0x66, 0x69,
	0x78, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x69, 0x6e, 0x63,
	0x65, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0f, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x4c, 0x61, 0x73, 0x74, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x12, 0x3e, 0x0a, 0x0a, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2e, 0x56, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x52, 0x0a, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x22, 0x34, 0x0a, 0x0a, 0x46, 0x69, 0x78, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x4f, 0x54, 0x48, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x46, 0x49, 0x58, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x4e, 0x4f, 0x54,
	0x5f, 0x46, 0x49, 0x58, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x02, 0x22, 0x62, 0x0a, 0x1a, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x61, 0x69,
	0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0c, 0x6d, 0x61, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x42, 0x24,
	0x0a, 0x19, 0x69, 0x6f, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x72, 0x6f, 0x78, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5a, 0x07, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_storage_report_configuration_proto_rawDescOnce sync.Once
	file_storage_report_configuration_proto_rawDescData = file_storage_report_configuration_proto_rawDesc
)

func file_storage_report_configuration_proto_rawDescGZIP() []byte {
	file_storage_report_configuration_proto_rawDescOnce.Do(func() {
		file_storage_report_configuration_proto_rawDescData = protoimpl.X.CompressGZIP(file_storage_report_configuration_proto_rawDescData)
	})
	return file_storage_report_configuration_proto_rawDescData
}

var file_storage_report_configuration_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_storage_report_configuration_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_storage_report_configuration_proto_goTypes = []interface{}{
	(ReportConfiguration_ReportType)(0),        // 0: storage.ReportConfiguration.ReportType
	(ReportLastRunStatus_RunStatus)(0),         // 1: storage.ReportLastRunStatus.RunStatus
	(VulnerabilityReportFilters_Fixability)(0), // 2: storage.VulnerabilityReportFilters.Fixability
	(*ReportConfiguration)(nil),                // 3: storage.ReportConfiguration
	(*ReportLastRunStatus)(nil),                // 4: storage.ReportLastRunStatus
	(*VulnerabilityReportFilters)(nil),         // 5: storage.VulnerabilityReportFilters
	(*EmailNotifierConfiguration)(nil),         // 6: storage.EmailNotifierConfiguration
	(*Schedule)(nil),                           // 7: storage.Schedule
	(*timestamppb.Timestamp)(nil),              // 8: google.protobuf.Timestamp
	(VulnerabilitySeverity)(0),                 // 9: storage.VulnerabilitySeverity
}
var file_storage_report_configuration_proto_depIdxs = []int32{
	0,  // 0: storage.ReportConfiguration.type:type_name -> storage.ReportConfiguration.ReportType
	5,  // 1: storage.ReportConfiguration.vuln_report_filters:type_name -> storage.VulnerabilityReportFilters
	6,  // 2: storage.ReportConfiguration.email_config:type_name -> storage.EmailNotifierConfiguration
	7,  // 3: storage.ReportConfiguration.schedule:type_name -> storage.Schedule
	4,  // 4: storage.ReportConfiguration.last_run_status:type_name -> storage.ReportLastRunStatus
	8,  // 5: storage.ReportConfiguration.last_successful_run_time:type_name -> google.protobuf.Timestamp
	1,  // 6: storage.ReportLastRunStatus.report_status:type_name -> storage.ReportLastRunStatus.RunStatus
	8,  // 7: storage.ReportLastRunStatus.last_run_time:type_name -> google.protobuf.Timestamp
	2,  // 8: storage.VulnerabilityReportFilters.fixability:type_name -> storage.VulnerabilityReportFilters.Fixability
	9,  // 9: storage.VulnerabilityReportFilters.severities:type_name -> storage.VulnerabilitySeverity
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_storage_report_configuration_proto_init() }
func file_storage_report_configuration_proto_init() {
	if File_storage_report_configuration_proto != nil {
		return
	}
	file_storage_cve_proto_init()
	file_storage_schedule_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_storage_report_configuration_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportConfiguration); i {
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
		file_storage_report_configuration_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportLastRunStatus); i {
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
		file_storage_report_configuration_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VulnerabilityReportFilters); i {
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
		file_storage_report_configuration_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailNotifierConfiguration); i {
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
	file_storage_report_configuration_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ReportConfiguration_VulnReportFilters)(nil),
		(*ReportConfiguration_EmailConfig)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_storage_report_configuration_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_storage_report_configuration_proto_goTypes,
		DependencyIndexes: file_storage_report_configuration_proto_depIdxs,
		EnumInfos:         file_storage_report_configuration_proto_enumTypes,
		MessageInfos:      file_storage_report_configuration_proto_msgTypes,
	}.Build()
	File_storage_report_configuration_proto = out.File
	file_storage_report_configuration_proto_rawDesc = nil
	file_storage_report_configuration_proto_goTypes = nil
	file_storage_report_configuration_proto_depIdxs = nil
}
