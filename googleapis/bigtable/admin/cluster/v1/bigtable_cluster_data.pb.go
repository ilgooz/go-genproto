// Copyright 2017 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.12.2
// source: google/bigtable/admin/cluster/v1/bigtable_cluster_data.proto

package cluster

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	longrunning "google.golang.org/genproto/googleapis/longrunning"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type StorageType int32

const (
	// The storage type used is unspecified.
	StorageType_STORAGE_UNSPECIFIED StorageType = 0
	// Data will be stored in SSD, providing low and consistent latencies.
	StorageType_STORAGE_SSD StorageType = 1
	// Data will be stored in HDD, providing high and less predictable
	// latencies.
	StorageType_STORAGE_HDD StorageType = 2
)

// Enum value maps for StorageType.
var (
	StorageType_name = map[int32]string{
		0: "STORAGE_UNSPECIFIED",
		1: "STORAGE_SSD",
		2: "STORAGE_HDD",
	}
	StorageType_value = map[string]int32{
		"STORAGE_UNSPECIFIED": 0,
		"STORAGE_SSD":         1,
		"STORAGE_HDD":         2,
	}
)

func (x StorageType) Enum() *StorageType {
	p := new(StorageType)
	*p = x
	return p
}

func (x StorageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StorageType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_bigtable_admin_cluster_v1_bigtable_cluster_data_proto_enumTypes[0].Descriptor()
}

func (StorageType) Type() protoreflect.EnumType {
	return &file_google_bigtable_admin_cluster_v1_bigtable_cluster_data_proto_enumTypes[0]
}

func (x StorageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StorageType.Descriptor instead.
func (StorageType) EnumDescriptor() ([]byte, []int) {
	return file_google_bigtable_admin_cluster_v1_bigtable_cluster_data_proto_rawDescGZIP(), []int{0}
}

// Possible states of a zone.
type Zone_Status int32

const (
	// The state of the zone is unknown or unspecified.
	Zone_UNKNOWN Zone_Status = 0
	// The zone is in a good state.
	Zone_OK Zone_Status = 1
	// The zone is down for planned maintenance.
	Zone_PLANNED_MAINTENANCE Zone_Status = 2
	// The zone is down for emergency or unplanned maintenance.
	Zone_EMERGENCY_MAINENANCE Zone_Status = 3
)

// Enum value maps for Zone_Status.
var (
	Zone_Status_name = map[int32]string{
		0: "UNKNOWN",
		1: "OK",
		2: "PLANNED_MAINTENANCE",
		3: "EMERGENCY_MAINENANCE",
	}
	Zone_Status_value = map[string]int32{
		"UNKNOWN":              0,
		"OK":                   1,
		"PLANNED_MAINTENANCE":  2,
		"EMERGENCY_MAINENANCE": 3,
	}
)

func (x Zone_Status) Enum() *Zone_Status {
	p := new(Zone_Status)
	*p = x
	return p
}

func (x Zone_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Zone_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_google_bigtable_admin_cluster_v1_bigtable_cluster_data_proto_enumTypes[1].Descriptor()
}

func (Zone_Status) Type() protoreflect.EnumType {
	return &file_google_bigtable_admin_cluster_v1_bigtable_cluster_data_proto_enumTypes[1]
}

func (x Zone_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Zone_Status.Descriptor instead.
func (Zone_Status) EnumDescriptor() ([]byte, []int) {
	return file_google_bigtable_admin_cluster_v1_bigtable_cluster_data_proto_rawDescGZIP(), []int{0, 0}
}

// A physical location in which a particular project can allocate Cloud BigTable
// resources.
type Zone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A permanent unique identifier for the zone.
	// Values are of the form projects/<project>/zones/[a-z][-a-z0-9]*
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The name of this zone as it appears in UIs.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// The current state of this zone.
	Status Zone_Status `protobuf:"varint,3,opt,name=status,proto3,enum=google.bigtable.admin.cluster.v1.Zone_Status" json:"status,omitempty"`
}

func (x *Zone) Reset() {
	*x = Zone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_bigtable_admin_cluster_v1_bigtable_cluster_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Zone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Zone) ProtoMessage() {}

func (x *Zone) ProtoReflect() protoreflect.Message {
	mi := &file_google_bigtable_admin_cluster_v1_bigtable_cluster_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Zone.ProtoReflect.Descriptor instead.
func (*Zone) Descriptor() ([]byte, []int) {
	return file_google_bigtable_admin_cluster_v1_bigtable_cluster_data_proto_rawDescGZIP(), []int{0}
}

func (x *Zone) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Zone) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Zone) GetStatus() Zone_Status {
	if x != nil {
		return x.Status
	}
	return Zone_UNKNOWN
}

// An isolated set of Cloud BigTable resources on which tables can be hosted.
type Cluster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A permanent unique identifier for the cluster. For technical reasons, the
	// zone in which the cluster resides is included here.
	// Values are of the form
	// projects/<project>/zones/<zone>/clusters/[a-z][-a-z0-9]*
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The operation currently running on the cluster, if any.
	// This cannot be set directly, only through CreateCluster, UpdateCluster,
	// or UndeleteCluster. Calls to these methods will be rejected if
	// "current_operation" is already set.
	CurrentOperation *longrunning.Operation `protobuf:"bytes,3,opt,name=current_operation,json=currentOperation,proto3" json:"current_operation,omitempty"`
	// The descriptive name for this cluster as it appears in UIs.
	// Must be unique per zone.
	DisplayName string `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// The number of serve nodes allocated to this cluster.
	ServeNodes int32 `protobuf:"varint,5,opt,name=serve_nodes,json=serveNodes,proto3" json:"serve_nodes,omitempty"`
	// What storage type to use for tables in this cluster. Only configurable at
	// cluster creation time. If unspecified, STORAGE_SSD will be used.
	DefaultStorageType StorageType `protobuf:"varint,8,opt,name=default_storage_type,json=defaultStorageType,proto3,enum=google.bigtable.admin.cluster.v1.StorageType" json:"default_storage_type,omitempty"`
}

func (x *Cluster) Reset() {
	*x = Cluster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_bigtable_admin_cluster_v1_bigtable_cluster_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cluster) ProtoMessage() {}

func (x *Cluster) ProtoReflect() protoreflect.Message {
	mi := &file_google_bigtable_admin_cluster_v1_bigtable_cluster_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cluster.ProtoReflect.Descriptor instead.
func (*Cluster) Descriptor() ([]byte, []int) {
	return file_google_bigtable_admin_cluster_v1_bigtable_cluster_data_proto_rawDescGZIP(), []int{1}
}

func (x *Cluster) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Cluster) GetCurrentOperation() *longrunning.Operation {
	if x != nil {
		return x.CurrentOperation
	}
	return nil
}

func (x *Cluster) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Cluster) GetServeNodes() int32 {
	if x != nil {
		return x.ServeNodes
	}
	return 0
}

func (x *Cluster) GetDefaultStorageType() StorageType {
	if x != nil {
		return x.DefaultStorageType
	}
	return StorageType_STORAGE_UNSPECIFIED
}

var File_google_bigtable_admin_cluster_v1_bigtable_cluster_data_proto protoreflect.FileDescriptor

var file_google_bigtable_admin_cluster_v1_bigtable_cluster_data_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x62, 0x69, 0x67, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x2f, 0x62, 0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x62, 0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6c, 0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e, 0x6e, 0x69,
	0x6e, 0x67, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd6, 0x01, 0x0a, 0x04, 0x5a, 0x6f, 0x6e, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x45, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x62, 0x69,
	0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x5a, 0x6f, 0x6e, 0x65, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x50, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x4c,
	0x41, 0x4e, 0x4e, 0x45, 0x44, 0x5f, 0x4d, 0x41, 0x49, 0x4e, 0x54, 0x45, 0x4e, 0x41, 0x4e, 0x43,
	0x45, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x45, 0x4d, 0x45, 0x52, 0x47, 0x45, 0x4e, 0x43, 0x59,
	0x5f, 0x4d, 0x41, 0x49, 0x4e, 0x45, 0x4e, 0x41, 0x4e, 0x43, 0x45, 0x10, 0x03, 0x22, 0x8e, 0x02,
	0x0a, 0x07, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x4a, 0x0a,
	0x11, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x6c, 0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x5f, 0x0a,
	0x14, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x62, 0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x12, 0x64, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x2a, 0x48,
	0x0a, 0x0b, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a,
	0x13, 0x53, 0x54, 0x4f, 0x52, 0x41, 0x47, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x54, 0x4f, 0x52, 0x41, 0x47,
	0x45, 0x5f, 0x53, 0x53, 0x44, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x54, 0x4f, 0x52, 0x41,
	0x47, 0x45, 0x5f, 0x48, 0x44, 0x44, 0x10, 0x02, 0x42, 0x8b, 0x01, 0x0a, 0x24, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x62, 0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x42, 0x18, 0x42, 0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x47, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x62, 0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2f, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_bigtable_admin_cluster_v1_bigtable_cluster_data_proto_rawDescOnce sync.Once
	file_google_bigtable_admin_cluster_v1_bigtable_cluster_data_proto_rawDescData = file_google_bigtable_admin_cluster_v1_bigtable_cluster_data_proto_rawDesc
)

func file_google_bigtable_admin_cluster_v1_bigtable_cluster_data_proto_rawDescGZIP() []byte {
	file_google_bigtable_admin_cluster_v1_bigtable_cluster_data_proto_rawDescOnce.Do(func() {
		file_google_bigtable_admin_cluster_v1_bigtable_cluster_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_bigtable_admin_cluster_v1_bigtable_cluster_data_proto_rawDescData)
	})
	return file_google_bigtable_admin_cluster_v1_bigtable_cluster_data_proto_rawDescData
}

var file_google_bigtable_admin_cluster_v1_bigtable_cluster_data_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_google_bigtable_admin_cluster_v1_bigtable_cluster_data_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_bigtable_admin_cluster_v1_bigtable_cluster_data_proto_goTypes = []interface{}{
	(StorageType)(0),              // 0: google.bigtable.admin.cluster.v1.StorageType
	(Zone_Status)(0),              // 1: google.bigtable.admin.cluster.v1.Zone.Status
	(*Zone)(nil),                  // 2: google.bigtable.admin.cluster.v1.Zone
	(*Cluster)(nil),               // 3: google.bigtable.admin.cluster.v1.Cluster
	(*longrunning.Operation)(nil), // 4: google.longrunning.Operation
}
var file_google_bigtable_admin_cluster_v1_bigtable_cluster_data_proto_depIdxs = []int32{
	1, // 0: google.bigtable.admin.cluster.v1.Zone.status:type_name -> google.bigtable.admin.cluster.v1.Zone.Status
	4, // 1: google.bigtable.admin.cluster.v1.Cluster.current_operation:type_name -> google.longrunning.Operation
	0, // 2: google.bigtable.admin.cluster.v1.Cluster.default_storage_type:type_name -> google.bigtable.admin.cluster.v1.StorageType
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_bigtable_admin_cluster_v1_bigtable_cluster_data_proto_init() }
func file_google_bigtable_admin_cluster_v1_bigtable_cluster_data_proto_init() {
	if File_google_bigtable_admin_cluster_v1_bigtable_cluster_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_bigtable_admin_cluster_v1_bigtable_cluster_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Zone); i {
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
		file_google_bigtable_admin_cluster_v1_bigtable_cluster_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cluster); i {
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
			RawDescriptor: file_google_bigtable_admin_cluster_v1_bigtable_cluster_data_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_bigtable_admin_cluster_v1_bigtable_cluster_data_proto_goTypes,
		DependencyIndexes: file_google_bigtable_admin_cluster_v1_bigtable_cluster_data_proto_depIdxs,
		EnumInfos:         file_google_bigtable_admin_cluster_v1_bigtable_cluster_data_proto_enumTypes,
		MessageInfos:      file_google_bigtable_admin_cluster_v1_bigtable_cluster_data_proto_msgTypes,
	}.Build()
	File_google_bigtable_admin_cluster_v1_bigtable_cluster_data_proto = out.File
	file_google_bigtable_admin_cluster_v1_bigtable_cluster_data_proto_rawDesc = nil
	file_google_bigtable_admin_cluster_v1_bigtable_cluster_data_proto_goTypes = nil
	file_google_bigtable_admin_cluster_v1_bigtable_cluster_data_proto_depIdxs = nil
}
