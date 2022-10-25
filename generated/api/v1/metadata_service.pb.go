// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: api/v1/metadata_service.proto

package v1

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

type Metadata_LicenseStatus int32

const (
	// Deprecated: Do not use.
	Metadata_NONE Metadata_LicenseStatus = 0
	// Deprecated: Do not use.
	Metadata_INVALID Metadata_LicenseStatus = 1
	// Deprecated: Do not use.
	Metadata_EXPIRED Metadata_LicenseStatus = 2
	// Deprecated: Do not use.
	Metadata_RESTARTING Metadata_LicenseStatus = 3
	Metadata_VALID      Metadata_LicenseStatus = 4
)

// Enum value maps for Metadata_LicenseStatus.
var (
	Metadata_LicenseStatus_name = map[int32]string{
		0: "NONE",
		1: "INVALID",
		2: "EXPIRED",
		3: "RESTARTING",
		4: "VALID",
	}
	Metadata_LicenseStatus_value = map[string]int32{
		"NONE":       0,
		"INVALID":    1,
		"EXPIRED":    2,
		"RESTARTING": 3,
		"VALID":      4,
	}
)

func (x Metadata_LicenseStatus) Enum() *Metadata_LicenseStatus {
	p := new(Metadata_LicenseStatus)
	*p = x
	return p
}

func (x Metadata_LicenseStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Metadata_LicenseStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_api_v1_metadata_service_proto_enumTypes[0].Descriptor()
}

func (Metadata_LicenseStatus) Type() protoreflect.EnumType {
	return &file_api_v1_metadata_service_proto_enumTypes[0]
}

func (x Metadata_LicenseStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Metadata_LicenseStatus.Descriptor instead.
func (Metadata_LicenseStatus) EnumDescriptor() ([]byte, []int) {
	return file_api_v1_metadata_service_proto_rawDescGZIP(), []int{0, 0}
}

type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version      string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	BuildFlavor  string `protobuf:"bytes,2,opt,name=build_flavor,json=buildFlavor,proto3" json:"build_flavor,omitempty"`
	ReleaseBuild bool   `protobuf:"varint,3,opt,name=release_build,json=releaseBuild,proto3" json:"release_build,omitempty"`
	// Do not use this field. It will always contain "VALID"
	//
	// Deprecated: Do not use.
	LicenseStatus Metadata_LicenseStatus `protobuf:"varint,4,opt,name=license_status,json=licenseStatus,proto3,enum=v1.Metadata_LicenseStatus" json:"license_status,omitempty"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_metadata_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_metadata_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_api_v1_metadata_service_proto_rawDescGZIP(), []int{0}
}

func (x *Metadata) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Metadata) GetBuildFlavor() string {
	if x != nil {
		return x.BuildFlavor
	}
	return ""
}

func (x *Metadata) GetReleaseBuild() bool {
	if x != nil {
		return x.ReleaseBuild
	}
	return false
}

// Deprecated: Do not use.
func (x *Metadata) GetLicenseStatus() Metadata_LicenseStatus {
	if x != nil {
		return x.LicenseStatus
	}
	return Metadata_NONE
}

type TrustInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// holds the certificate chain hold by central
	CertChain [][]byte `protobuf:"bytes,1,rep,name=cert_chain,json=certChain,proto3" json:"cert_chain,omitempty"`
	// sensor challenge string
	SensorChallenge string `protobuf:"bytes,2,opt,name=sensor_challenge,json=sensorChallenge,proto3" json:"sensor_challenge,omitempty"`
	// central challenge string
	CentralChallenge string `protobuf:"bytes,3,opt,name=central_challenge,json=centralChallenge,proto3" json:"central_challenge,omitempty"`
	// additional CA certs configured in central in DER format
	AdditionalCas [][]byte `protobuf:"bytes,4,rep,name=additional_cas,json=additionalCas,proto3" json:"additional_cas,omitempty"`
}

func (x *TrustInfo) Reset() {
	*x = TrustInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_metadata_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrustInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrustInfo) ProtoMessage() {}

func (x *TrustInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_metadata_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrustInfo.ProtoReflect.Descriptor instead.
func (*TrustInfo) Descriptor() ([]byte, []int) {
	return file_api_v1_metadata_service_proto_rawDescGZIP(), []int{1}
}

func (x *TrustInfo) GetCertChain() [][]byte {
	if x != nil {
		return x.CertChain
	}
	return nil
}

func (x *TrustInfo) GetSensorChallenge() string {
	if x != nil {
		return x.SensorChallenge
	}
	return ""
}

func (x *TrustInfo) GetCentralChallenge() string {
	if x != nil {
		return x.CentralChallenge
	}
	return ""
}

func (x *TrustInfo) GetAdditionalCas() [][]byte {
	if x != nil {
		return x.AdditionalCas
	}
	return nil
}

type TLSChallengeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// signed data which is returned to the caller, is validated against the signature
	TrustInfoSerialized []byte `protobuf:"bytes,1,opt,name=trust_info_serialized,json=trustInfoSerialized,proto3" json:"trust_info_serialized,omitempty"`
	Signature           []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *TLSChallengeResponse) Reset() {
	*x = TLSChallengeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_metadata_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TLSChallengeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TLSChallengeResponse) ProtoMessage() {}

func (x *TLSChallengeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_metadata_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TLSChallengeResponse.ProtoReflect.Descriptor instead.
func (*TLSChallengeResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_metadata_service_proto_rawDescGZIP(), []int{2}
}

func (x *TLSChallengeResponse) GetTrustInfoSerialized() []byte {
	if x != nil {
		return x.TrustInfoSerialized
	}
	return nil
}

func (x *TLSChallengeResponse) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type TLSChallengeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// generated challenge token by the service asking for TLS certs
	ChallengeToken string `protobuf:"bytes,1,opt,name=challenge_token,json=challengeToken,proto3" json:"challenge_token,omitempty"`
}

func (x *TLSChallengeRequest) Reset() {
	*x = TLSChallengeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_metadata_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TLSChallengeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TLSChallengeRequest) ProtoMessage() {}

func (x *TLSChallengeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_metadata_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TLSChallengeRequest.ProtoReflect.Descriptor instead.
func (*TLSChallengeRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_metadata_service_proto_rawDescGZIP(), []int{3}
}

func (x *TLSChallengeRequest) GetChallengeToken() string {
	if x != nil {
		return x.ChallengeToken
	}
	return ""
}

var File_api_v1_metadata_service_proto protoreflect.FileDescriptor

var file_api_v1_metadata_service_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x02, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x66, 0x6c, 0x61, 0x76, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x46, 0x6c, 0x61, 0x76, 0x6f, 0x72, 0x12,
	0x23, 0x0a, 0x0d, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x42,
	0x75, 0x69, 0x6c, 0x64, 0x12, 0x45, 0x0a, 0x0e, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x76,
	0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4c, 0x69, 0x63, 0x65, 0x6e,
	0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x02, 0x18, 0x01, 0x52, 0x0d, 0x6c, 0x69,
	0x63, 0x65, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x5e, 0x0a, 0x0d, 0x4c,
	0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0c, 0x0a, 0x04,
	0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x1a, 0x02, 0x08, 0x01, 0x12, 0x0f, 0x0a, 0x07, 0x49, 0x4e,
	0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x01, 0x1a, 0x02, 0x08, 0x01, 0x12, 0x0f, 0x0a, 0x07, 0x45,
	0x58, 0x50, 0x49, 0x52, 0x45, 0x44, 0x10, 0x02, 0x1a, 0x02, 0x08, 0x01, 0x12, 0x12, 0x0a, 0x0a,
	0x52, 0x45, 0x53, 0x54, 0x41, 0x52, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x1a, 0x02, 0x08, 0x01,
	0x12, 0x09, 0x0a, 0x05, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x04, 0x22, 0xa9, 0x01, 0x0a, 0x09,
	0x54, 0x72, 0x75, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x65, 0x72,
	0x74, 0x5f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x09, 0x63,
	0x65, 0x72, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x73, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x5f, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65,
	0x6e, 0x67, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x5f, 0x63,
	0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10,
	0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65,
	0x12, 0x25, 0x0a, 0x0e, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x63,
	0x61, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0d, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x43, 0x61, 0x73, 0x22, 0x68, 0x0a, 0x14, 0x54, 0x4c, 0x53, 0x43, 0x68,
	0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x32, 0x0a, 0x15, 0x74, 0x72, 0x75, 0x73, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x73, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x13,
	0x74, 0x72, 0x75, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69,
	0x7a, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x22, 0x3e, 0x0a, 0x13, 0x54, 0x4c, 0x53, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x68, 0x61, 0x6c,
	0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x32, 0xad, 0x01, 0x0a, 0x0f, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x09, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x0c, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x14, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x5c, 0x0a, 0x0c, 0x54, 0x4c, 0x53, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65,
	0x6e, 0x67, 0x65, 0x12, 0x17, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x4c, 0x53, 0x43, 0x68, 0x61, 0x6c,
	0x6c, 0x65, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x4c, 0x53, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6c, 0x73, 0x2d, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67,
	0x65, 0x42, 0x1e, 0x0a, 0x18, 0x69, 0x6f, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x72, 0x6f, 0x78,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x5a, 0x02, 0x76,
	0x31, 0x58, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_metadata_service_proto_rawDescOnce sync.Once
	file_api_v1_metadata_service_proto_rawDescData = file_api_v1_metadata_service_proto_rawDesc
)

func file_api_v1_metadata_service_proto_rawDescGZIP() []byte {
	file_api_v1_metadata_service_proto_rawDescOnce.Do(func() {
		file_api_v1_metadata_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_metadata_service_proto_rawDescData)
	})
	return file_api_v1_metadata_service_proto_rawDescData
}

var file_api_v1_metadata_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_v1_metadata_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_v1_metadata_service_proto_goTypes = []interface{}{
	(Metadata_LicenseStatus)(0),  // 0: v1.Metadata.LicenseStatus
	(*Metadata)(nil),             // 1: v1.Metadata
	(*TrustInfo)(nil),            // 2: v1.TrustInfo
	(*TLSChallengeResponse)(nil), // 3: v1.TLSChallengeResponse
	(*TLSChallengeRequest)(nil),  // 4: v1.TLSChallengeRequest
	(*Empty)(nil),                // 5: v1.Empty
}
var file_api_v1_metadata_service_proto_depIdxs = []int32{
	0, // 0: v1.Metadata.license_status:type_name -> v1.Metadata.LicenseStatus
	5, // 1: v1.MetadataService.GetMetadata:input_type -> v1.Empty
	4, // 2: v1.MetadataService.TLSChallenge:input_type -> v1.TLSChallengeRequest
	1, // 3: v1.MetadataService.GetMetadata:output_type -> v1.Metadata
	3, // 4: v1.MetadataService.TLSChallenge:output_type -> v1.TLSChallengeResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_v1_metadata_service_proto_init() }
func file_api_v1_metadata_service_proto_init() {
	if File_api_v1_metadata_service_proto != nil {
		return
	}
	file_api_v1_empty_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_v1_metadata_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata); i {
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
		file_api_v1_metadata_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrustInfo); i {
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
		file_api_v1_metadata_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TLSChallengeResponse); i {
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
		file_api_v1_metadata_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TLSChallengeRequest); i {
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
			RawDescriptor: file_api_v1_metadata_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_metadata_service_proto_goTypes,
		DependencyIndexes: file_api_v1_metadata_service_proto_depIdxs,
		EnumInfos:         file_api_v1_metadata_service_proto_enumTypes,
		MessageInfos:      file_api_v1_metadata_service_proto_msgTypes,
	}.Build()
	File_api_v1_metadata_service_proto = out.File
	file_api_v1_metadata_service_proto_rawDesc = nil
	file_api_v1_metadata_service_proto_goTypes = nil
	file_api_v1_metadata_service_proto_depIdxs = nil
}
