//*
// Copyright (c) 2022 Peking University and Peking University Institute for Computing and Digital Economy
// SCOW is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//          http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
// EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
// MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: config.proto

package gen

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

type PartitionInfo_PartitionStatus int32

const (
	PartitionInfo_NOT_AVAILABLE PartitionInfo_PartitionStatus = 0
	PartitionInfo_AVAILABLE     PartitionInfo_PartitionStatus = 1
)

// Enum value maps for PartitionInfo_PartitionStatus.
var (
	PartitionInfo_PartitionStatus_name = map[int32]string{
		0: "NOT_AVAILABLE",
		1: "AVAILABLE",
	}
	PartitionInfo_PartitionStatus_value = map[string]int32{
		"NOT_AVAILABLE": 0,
		"AVAILABLE":     1,
	}
)

func (x PartitionInfo_PartitionStatus) Enum() *PartitionInfo_PartitionStatus {
	p := new(PartitionInfo_PartitionStatus)
	*p = x
	return p
}

func (x PartitionInfo_PartitionStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PartitionInfo_PartitionStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_config_proto_enumTypes[0].Descriptor()
}

func (PartitionInfo_PartitionStatus) Type() protoreflect.EnumType {
	return &file_config_proto_enumTypes[0]
}

func (x PartitionInfo_PartitionStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PartitionInfo_PartitionStatus.Descriptor instead.
func (PartitionInfo_PartitionStatus) EnumDescriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{5, 0}
}

type GetClusterConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountName string `protobuf:"bytes,1,opt,name=account_name,json=accountName,proto3" json:"account_name,omitempty"`
}

func (x *GetClusterConfigRequest) Reset() {
	*x = GetClusterConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClusterConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClusterConfigRequest) ProtoMessage() {}

func (x *GetClusterConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClusterConfigRequest.ProtoReflect.Descriptor instead.
func (*GetClusterConfigRequest) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{0}
}

func (x *GetClusterConfigRequest) GetAccountName() string {
	if x != nil {
		return x.AccountName
	}
	return ""
}

// static configuration of partition
type Partition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// mem: memory size in M
	MemMb uint64 `protobuf:"varint,2,opt,name=mem_mb,json=memMb,proto3" json:"mem_mb,omitempty"`
	// cores: number of cores
	Cores uint32 `protobuf:"varint,3,opt,name=cores,proto3" json:"cores,omitempty"`
	// gpus: number of gpu
	Gpus uint32 `protobuf:"varint,4,opt,name=gpus,proto3" json:"gpus,omitempty"`
	// nodes: number of nodes
	Nodes uint32 `protobuf:"varint,5,opt,name=nodes,proto3" json:"nodes,omitempty"`
	// list that stores qos. the list can be empty.
	Qos []string `protobuf:"bytes,6,rep,name=qos,proto3" json:"qos,omitempty"`
	// price item description
	Comment *string `protobuf:"bytes,7,opt,name=comment,proto3,oneof" json:"comment,omitempty"`
}

func (x *Partition) Reset() {
	*x = Partition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Partition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Partition) ProtoMessage() {}

func (x *Partition) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Partition.ProtoReflect.Descriptor instead.
func (*Partition) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{1}
}

func (x *Partition) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Partition) GetMemMb() uint64 {
	if x != nil {
		return x.MemMb
	}
	return 0
}

func (x *Partition) GetCores() uint32 {
	if x != nil {
		return x.Cores
	}
	return 0
}

func (x *Partition) GetGpus() uint32 {
	if x != nil {
		return x.Gpus
	}
	return 0
}

func (x *Partition) GetNodes() uint32 {
	if x != nil {
		return x.Nodes
	}
	return 0
}

func (x *Partition) GetQos() []string {
	if x != nil {
		return x.Qos
	}
	return nil
}

func (x *Partition) GetComment() string {
	if x != nil && x.Comment != nil {
		return *x.Comment
	}
	return ""
}

type GetClusterConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Partitions    []*Partition `protobuf:"bytes,1,rep,name=partitions,proto3" json:"partitions,omitempty"`
	SchedulerName string       `protobuf:"bytes,2,opt,name=scheduler_name,json=schedulerName,proto3" json:"scheduler_name,omitempty"`
}

func (x *GetClusterConfigResponse) Reset() {
	*x = GetClusterConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClusterConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClusterConfigResponse) ProtoMessage() {}

func (x *GetClusterConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClusterConfigResponse.ProtoReflect.Descriptor instead.
func (*GetClusterConfigResponse) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{2}
}

func (x *GetClusterConfigResponse) GetPartitions() []*Partition {
	if x != nil {
		return x.Partitions
	}
	return nil
}

func (x *GetClusterConfigResponse) GetSchedulerName() string {
	if x != nil {
		return x.SchedulerName
	}
	return ""
}

type GetAvailablePartitionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountName string `protobuf:"bytes,1,opt,name=account_name,json=accountName,proto3" json:"account_name,omitempty"`
	UserId      string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetAvailablePartitionsRequest) Reset() {
	*x = GetAvailablePartitionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAvailablePartitionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAvailablePartitionsRequest) ProtoMessage() {}

func (x *GetAvailablePartitionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAvailablePartitionsRequest.ProtoReflect.Descriptor instead.
func (*GetAvailablePartitionsRequest) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{3}
}

func (x *GetAvailablePartitionsRequest) GetAccountName() string {
	if x != nil {
		return x.AccountName
	}
	return ""
}

func (x *GetAvailablePartitionsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetAvailablePartitionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Partitions []*Partition `protobuf:"bytes,1,rep,name=partitions,proto3" json:"partitions,omitempty"`
}

func (x *GetAvailablePartitionsResponse) Reset() {
	*x = GetAvailablePartitionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAvailablePartitionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAvailablePartitionsResponse) ProtoMessage() {}

func (x *GetAvailablePartitionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAvailablePartitionsResponse.ProtoReflect.Descriptor instead.
func (*GetAvailablePartitionsResponse) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{4}
}

func (x *GetAvailablePartitionsResponse) GetPartitions() []*Partition {
	if x != nil {
		return x.Partitions
	}
	return nil
}

// the runtime state of the partition
type PartitionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PartitionName         string `protobuf:"bytes,1,opt,name=partition_name,json=partitionName,proto3" json:"partition_name,omitempty"`
	NodeCount             uint32 `protobuf:"varint,2,opt,name=node_count,json=nodeCount,proto3" json:"node_count,omitempty"`
	RunningNodeCount      uint32 `protobuf:"varint,3,opt,name=running_node_count,json=runningNodeCount,proto3" json:"running_node_count,omitempty"`
	IdleNodeCount         uint32 `protobuf:"varint,4,opt,name=idle_node_count,json=idleNodeCount,proto3" json:"idle_node_count,omitempty"`
	NotAvailableNodeCount uint32 `protobuf:"varint,5,opt,name=not_available_node_count,json=notAvailableNodeCount,proto3" json:"not_available_node_count,omitempty"`
	CpuCoreCount          uint32 `protobuf:"varint,6,opt,name=cpu_core_count,json=cpuCoreCount,proto3" json:"cpu_core_count,omitempty"`
	RunningCpuCount       uint32 `protobuf:"varint,7,opt,name=running_cpu_count,json=runningCpuCount,proto3" json:"running_cpu_count,omitempty"`
	IdleCpuCount          uint32 `protobuf:"varint,8,opt,name=idle_cpu_count,json=idleCpuCount,proto3" json:"idle_cpu_count,omitempty"`
	NotAvailableCpuCount  uint32 `protobuf:"varint,9,opt,name=not_available_cpu_count,json=notAvailableCpuCount,proto3" json:"not_available_cpu_count,omitempty"`
	GpuCoreCount          uint32 `protobuf:"varint,10,opt,name=gpu_core_count,json=gpuCoreCount,proto3" json:"gpu_core_count,omitempty"`
	RunningGpuCount       uint32 `protobuf:"varint,11,opt,name=running_gpu_count,json=runningGpuCount,proto3" json:"running_gpu_count,omitempty"`
	IdleGpuCount          uint32 `protobuf:"varint,12,opt,name=idle_gpu_count,json=idleGpuCount,proto3" json:"idle_gpu_count,omitempty"`
	NotAvailableGpuCount  uint32 `protobuf:"varint,13,opt,name=not_available_gpu_count,json=notAvailableGpuCount,proto3" json:"not_available_gpu_count,omitempty"`
	JobCount              uint32 `protobuf:"varint,14,opt,name=job_count,json=jobCount,proto3" json:"job_count,omitempty"`
	RunningJobCount       uint32 `protobuf:"varint,15,opt,name=running_job_count,json=runningJobCount,proto3" json:"running_job_count,omitempty"`
	PendingJobCount       uint32 `protobuf:"varint,16,opt,name=pending_job_count,json=pendingJobCount,proto3" json:"pending_job_count,omitempty"`
	// node utilization rate
	UsageRatePercentage uint32                        `protobuf:"varint,17,opt,name=usage_rate_percentage,json=usageRatePercentage,proto3" json:"usage_rate_percentage,omitempty"`
	PartitionStatus     PartitionInfo_PartitionStatus `protobuf:"varint,18,opt,name=partition_status,json=partitionStatus,proto3,enum=scow.scheduler_adapter.PartitionInfo_PartitionStatus" json:"partition_status,omitempty"`
}

func (x *PartitionInfo) Reset() {
	*x = PartitionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartitionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartitionInfo) ProtoMessage() {}

func (x *PartitionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartitionInfo.ProtoReflect.Descriptor instead.
func (*PartitionInfo) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{5}
}

func (x *PartitionInfo) GetPartitionName() string {
	if x != nil {
		return x.PartitionName
	}
	return ""
}

func (x *PartitionInfo) GetNodeCount() uint32 {
	if x != nil {
		return x.NodeCount
	}
	return 0
}

func (x *PartitionInfo) GetRunningNodeCount() uint32 {
	if x != nil {
		return x.RunningNodeCount
	}
	return 0
}

func (x *PartitionInfo) GetIdleNodeCount() uint32 {
	if x != nil {
		return x.IdleNodeCount
	}
	return 0
}

func (x *PartitionInfo) GetNotAvailableNodeCount() uint32 {
	if x != nil {
		return x.NotAvailableNodeCount
	}
	return 0
}

func (x *PartitionInfo) GetCpuCoreCount() uint32 {
	if x != nil {
		return x.CpuCoreCount
	}
	return 0
}

func (x *PartitionInfo) GetRunningCpuCount() uint32 {
	if x != nil {
		return x.RunningCpuCount
	}
	return 0
}

func (x *PartitionInfo) GetIdleCpuCount() uint32 {
	if x != nil {
		return x.IdleCpuCount
	}
	return 0
}

func (x *PartitionInfo) GetNotAvailableCpuCount() uint32 {
	if x != nil {
		return x.NotAvailableCpuCount
	}
	return 0
}

func (x *PartitionInfo) GetGpuCoreCount() uint32 {
	if x != nil {
		return x.GpuCoreCount
	}
	return 0
}

func (x *PartitionInfo) GetRunningGpuCount() uint32 {
	if x != nil {
		return x.RunningGpuCount
	}
	return 0
}

func (x *PartitionInfo) GetIdleGpuCount() uint32 {
	if x != nil {
		return x.IdleGpuCount
	}
	return 0
}

func (x *PartitionInfo) GetNotAvailableGpuCount() uint32 {
	if x != nil {
		return x.NotAvailableGpuCount
	}
	return 0
}

func (x *PartitionInfo) GetJobCount() uint32 {
	if x != nil {
		return x.JobCount
	}
	return 0
}

func (x *PartitionInfo) GetRunningJobCount() uint32 {
	if x != nil {
		return x.RunningJobCount
	}
	return 0
}

func (x *PartitionInfo) GetPendingJobCount() uint32 {
	if x != nil {
		return x.PendingJobCount
	}
	return 0
}

func (x *PartitionInfo) GetUsageRatePercentage() uint32 {
	if x != nil {
		return x.UsageRatePercentage
	}
	return 0
}

func (x *PartitionInfo) GetPartitionStatus() PartitionInfo_PartitionStatus {
	if x != nil {
		return x.PartitionStatus
	}
	return PartitionInfo_NOT_AVAILABLE
}

type GetClusterInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetClusterInfoRequest) Reset() {
	*x = GetClusterInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClusterInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClusterInfoRequest) ProtoMessage() {}

func (x *GetClusterInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClusterInfoRequest.ProtoReflect.Descriptor instead.
func (*GetClusterInfoRequest) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{6}
}

type GetClusterInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterName string           `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	Partitions  []*PartitionInfo `protobuf:"bytes,2,rep,name=partitions,proto3" json:"partitions,omitempty"`
}

func (x *GetClusterInfoResponse) Reset() {
	*x = GetClusterInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClusterInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClusterInfoResponse) ProtoMessage() {}

func (x *GetClusterInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClusterInfoResponse.ProtoReflect.Descriptor instead.
func (*GetClusterInfoResponse) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{7}
}

func (x *GetClusterInfoResponse) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *GetClusterInfoResponse) GetPartitions() []*PartitionInfo {
	if x != nil {
		return x.Partitions
	}
	return nil
}

var File_config_proto protoreflect.FileDescriptor

var file_config_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16,
	0x73, 0x63, 0x6f, 0x77, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x61,
	0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x22, 0x3c, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0xb3, 0x01, 0x0a, 0x09, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x5f, 0x6d, 0x62,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6d, 0x65, 0x6d, 0x4d, 0x62, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f,
	0x72, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x67, 0x70, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x04, 0x67, 0x70, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x10, 0x0a,
	0x03, 0x71, 0x6f, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x71, 0x6f, 0x73, 0x12,
	0x1d, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x42, 0x0a,
	0x0a, 0x08, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x84, 0x01, 0x0a, 0x18, 0x47,
	0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x74, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x73, 0x63,
	0x6f, 0x77, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x61,
	0x70, 0x74, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a,
	0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x22, 0x5b, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x63,
	0x0a, 0x1e, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x61,
	0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x41, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x73, 0x63, 0x6f, 0x77, 0x2e, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2e, 0x50, 0x61,
	0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x22, 0x82, 0x07, 0x0a, 0x0d, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70,
	0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x72,
	0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67,
	0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x69, 0x64, 0x6c,
	0x65, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0d, 0x69, 0x64, 0x6c, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x37, 0x0a, 0x18, 0x6e, 0x6f, 0x74, 0x5f, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62,
	0x6c, 0x65, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x15, 0x6e, 0x6f, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x63, 0x70,
	0x75, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0c, 0x63, 0x70, 0x75, 0x43, 0x6f, 0x72, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x2a, 0x0a, 0x11, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x70, 0x75, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x72, 0x75, 0x6e,
	0x6e, 0x69, 0x6e, 0x67, 0x43, 0x70, 0x75, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0e,
	0x69, 0x64, 0x6c, 0x65, 0x5f, 0x63, 0x70, 0x75, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x69, 0x64, 0x6c, 0x65, 0x43, 0x70, 0x75, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x35, 0x0a, 0x17, 0x6e, 0x6f, 0x74, 0x5f, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x6c, 0x65, 0x5f, 0x63, 0x70, 0x75, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x14, 0x6e, 0x6f, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x43, 0x70, 0x75, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x67, 0x70, 0x75,
	0x5f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0c, 0x67, 0x70, 0x75, 0x43, 0x6f, 0x72, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x2a, 0x0a, 0x11, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x67, 0x70, 0x75, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x72, 0x75, 0x6e, 0x6e,
	0x69, 0x6e, 0x67, 0x47, 0x70, 0x75, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x69,
	0x64, 0x6c, 0x65, 0x5f, 0x67, 0x70, 0x75, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0c, 0x69, 0x64, 0x6c, 0x65, 0x47, 0x70, 0x75, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x35, 0x0a, 0x17, 0x6e, 0x6f, 0x74, 0x5f, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62,
	0x6c, 0x65, 0x5f, 0x67, 0x70, 0x75, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x14, 0x6e, 0x6f, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65,
	0x47, 0x70, 0x75, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6a, 0x6f, 0x62, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6a, 0x6f, 0x62,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67,
	0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0f, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x4a, 0x6f, 0x62, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x2a, 0x0a, 0x11, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x6a, 0x6f, 0x62,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x70, 0x65,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4a, 0x6f, 0x62, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x32, 0x0a,
	0x15, 0x75, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x75, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67,
	0x65, 0x12, 0x60, 0x0a, 0x10, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x35, 0x2e, 0x73, 0x63,
	0x6f, 0x77, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x61,
	0x70, 0x74, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x0f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x33, 0x0a, 0x0f, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x11, 0x0a, 0x0d, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x56,
	0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x56, 0x41,
	0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x22, 0x17, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x82, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x45, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x63, 0x6f, 0x77, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x72,
	0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x74,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0x81, 0x03, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x75, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2f, 0x2e, 0x73,
	0x63, 0x6f, 0x77, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x61, 0x64,
	0x61, 0x70, 0x74, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e,
	0x73, 0x63, 0x6f, 0x77, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x61,
	0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x87, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65,
	0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x35, 0x2e, 0x73, 0x63, 0x6f,
	0x77, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x61, 0x70,
	0x74, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65,
	0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x36, 0x2e, 0x73, 0x63, 0x6f, 0x77, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x72, 0x5f, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x76,
	0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6f, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2d, 0x2e, 0x73, 0x63,
	0x6f, 0x77, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x61,
	0x70, 0x74, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x73, 0x63, 0x6f,
	0x77, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x61, 0x70,
	0x74, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xb6, 0x01, 0x0a, 0x1a, 0x63,
	0x6f, 0x6d, 0x2e, 0x73, 0x63, 0x6f, 0x77, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x72, 0x5f, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x42, 0x0b, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x16, 0x73, 0x63, 0x6f, 0x77, 0x2d, 0x73,
	0x6c, 0x75, 0x72, 0x6d, 0x2d, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e,
	0xa2, 0x02, 0x03, 0x53, 0x53, 0x58, 0xaa, 0x02, 0x15, 0x53, 0x63, 0x6f, 0x77, 0x2e, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x41, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0xca, 0x02,
	0x15, 0x53, 0x63, 0x6f, 0x77, 0x5c, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x41,
	0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0xe2, 0x02, 0x21, 0x53, 0x63, 0x6f, 0x77, 0x5c, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x41, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x16, 0x53, 0x63, 0x6f,
	0x77, 0x3a, 0x3a, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x41, 0x64, 0x61, 0x70,
	0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_proto_rawDescOnce sync.Once
	file_config_proto_rawDescData = file_config_proto_rawDesc
)

func file_config_proto_rawDescGZIP() []byte {
	file_config_proto_rawDescOnce.Do(func() {
		file_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_proto_rawDescData)
	})
	return file_config_proto_rawDescData
}

var file_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_config_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_config_proto_goTypes = []interface{}{
	(PartitionInfo_PartitionStatus)(0),     // 0: scow.scheduler_adapter.PartitionInfo.PartitionStatus
	(*GetClusterConfigRequest)(nil),        // 1: scow.scheduler_adapter.GetClusterConfigRequest
	(*Partition)(nil),                      // 2: scow.scheduler_adapter.Partition
	(*GetClusterConfigResponse)(nil),       // 3: scow.scheduler_adapter.GetClusterConfigResponse
	(*GetAvailablePartitionsRequest)(nil),  // 4: scow.scheduler_adapter.GetAvailablePartitionsRequest
	(*GetAvailablePartitionsResponse)(nil), // 5: scow.scheduler_adapter.GetAvailablePartitionsResponse
	(*PartitionInfo)(nil),                  // 6: scow.scheduler_adapter.PartitionInfo
	(*GetClusterInfoRequest)(nil),          // 7: scow.scheduler_adapter.GetClusterInfoRequest
	(*GetClusterInfoResponse)(nil),         // 8: scow.scheduler_adapter.GetClusterInfoResponse
}
var file_config_proto_depIdxs = []int32{
	2, // 0: scow.scheduler_adapter.GetClusterConfigResponse.partitions:type_name -> scow.scheduler_adapter.Partition
	2, // 1: scow.scheduler_adapter.GetAvailablePartitionsResponse.partitions:type_name -> scow.scheduler_adapter.Partition
	0, // 2: scow.scheduler_adapter.PartitionInfo.partition_status:type_name -> scow.scheduler_adapter.PartitionInfo.PartitionStatus
	6, // 3: scow.scheduler_adapter.GetClusterInfoResponse.partitions:type_name -> scow.scheduler_adapter.PartitionInfo
	1, // 4: scow.scheduler_adapter.ConfigService.GetClusterConfig:input_type -> scow.scheduler_adapter.GetClusterConfigRequest
	4, // 5: scow.scheduler_adapter.ConfigService.GetAvailablePartitions:input_type -> scow.scheduler_adapter.GetAvailablePartitionsRequest
	7, // 6: scow.scheduler_adapter.ConfigService.GetClusterInfo:input_type -> scow.scheduler_adapter.GetClusterInfoRequest
	3, // 7: scow.scheduler_adapter.ConfigService.GetClusterConfig:output_type -> scow.scheduler_adapter.GetClusterConfigResponse
	5, // 8: scow.scheduler_adapter.ConfigService.GetAvailablePartitions:output_type -> scow.scheduler_adapter.GetAvailablePartitionsResponse
	8, // 9: scow.scheduler_adapter.ConfigService.GetClusterInfo:output_type -> scow.scheduler_adapter.GetClusterInfoResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_config_proto_init() }
func file_config_proto_init() {
	if File_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClusterConfigRequest); i {
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
		file_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Partition); i {
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
		file_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClusterConfigResponse); i {
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
		file_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAvailablePartitionsRequest); i {
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
		file_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAvailablePartitionsResponse); i {
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
		file_config_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartitionInfo); i {
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
		file_config_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClusterInfoRequest); i {
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
		file_config_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClusterInfoResponse); i {
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
	file_config_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_config_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_config_proto_goTypes,
		DependencyIndexes: file_config_proto_depIdxs,
		EnumInfos:         file_config_proto_enumTypes,
		MessageInfos:      file_config_proto_msgTypes,
	}.Build()
	File_config_proto = out.File
	file_config_proto_rawDesc = nil
	file_config_proto_goTypes = nil
	file_config_proto_depIdxs = nil
}
