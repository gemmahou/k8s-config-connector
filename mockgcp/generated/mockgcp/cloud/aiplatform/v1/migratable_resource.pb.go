// Copyright 2024 Google LLC
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
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: mockgcp/cloud/aiplatform/v1/migratable_resource.proto

package aiplatformpb

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Represents one resource that exists in automl.googleapis.com,
// datalabeling.googleapis.com or ml.googleapis.com.
type MigratableResource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Resource:
	//
	//	*MigratableResource_MlEngineModelVersion_
	//	*MigratableResource_AutomlModel_
	//	*MigratableResource_AutomlDataset_
	//	*MigratableResource_DataLabelingDataset_
	Resource isMigratableResource_Resource `protobuf_oneof:"resource"`
	// Output only. Timestamp when the last migration attempt on this
	// MigratableResource started. Will not be set if there's no migration attempt
	// on this MigratableResource.
	LastMigrateTime *timestamp.Timestamp `protobuf:"bytes,5,opt,name=last_migrate_time,json=lastMigrateTime,proto3" json:"last_migrate_time,omitempty"`
	// Output only. Timestamp when this MigratableResource was last updated.
	LastUpdateTime *timestamp.Timestamp `protobuf:"bytes,6,opt,name=last_update_time,json=lastUpdateTime,proto3" json:"last_update_time,omitempty"`
}

func (x *MigratableResource) Reset() {
	*x = MigratableResource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgcp_cloud_aiplatform_v1_migratable_resource_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MigratableResource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MigratableResource) ProtoMessage() {}

func (x *MigratableResource) ProtoReflect() protoreflect.Message {
	mi := &file_mockgcp_cloud_aiplatform_v1_migratable_resource_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MigratableResource.ProtoReflect.Descriptor instead.
func (*MigratableResource) Descriptor() ([]byte, []int) {
	return file_mockgcp_cloud_aiplatform_v1_migratable_resource_proto_rawDescGZIP(), []int{0}
}

func (m *MigratableResource) GetResource() isMigratableResource_Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (x *MigratableResource) GetMlEngineModelVersion() *MigratableResource_MlEngineModelVersion {
	if x, ok := x.GetResource().(*MigratableResource_MlEngineModelVersion_); ok {
		return x.MlEngineModelVersion
	}
	return nil
}

func (x *MigratableResource) GetAutomlModel() *MigratableResource_AutomlModel {
	if x, ok := x.GetResource().(*MigratableResource_AutomlModel_); ok {
		return x.AutomlModel
	}
	return nil
}

func (x *MigratableResource) GetAutomlDataset() *MigratableResource_AutomlDataset {
	if x, ok := x.GetResource().(*MigratableResource_AutomlDataset_); ok {
		return x.AutomlDataset
	}
	return nil
}

func (x *MigratableResource) GetDataLabelingDataset() *MigratableResource_DataLabelingDataset {
	if x, ok := x.GetResource().(*MigratableResource_DataLabelingDataset_); ok {
		return x.DataLabelingDataset
	}
	return nil
}

func (x *MigratableResource) GetLastMigrateTime() *timestamp.Timestamp {
	if x != nil {
		return x.LastMigrateTime
	}
	return nil
}

func (x *MigratableResource) GetLastUpdateTime() *timestamp.Timestamp {
	if x != nil {
		return x.LastUpdateTime
	}
	return nil
}

type isMigratableResource_Resource interface {
	isMigratableResource_Resource()
}

type MigratableResource_MlEngineModelVersion_ struct {
	// Output only. Represents one Version in ml.googleapis.com.
	MlEngineModelVersion *MigratableResource_MlEngineModelVersion `protobuf:"bytes,1,opt,name=ml_engine_model_version,json=mlEngineModelVersion,proto3,oneof"`
}

type MigratableResource_AutomlModel_ struct {
	// Output only. Represents one Model in automl.googleapis.com.
	AutomlModel *MigratableResource_AutomlModel `protobuf:"bytes,2,opt,name=automl_model,json=automlModel,proto3,oneof"`
}

type MigratableResource_AutomlDataset_ struct {
	// Output only. Represents one Dataset in automl.googleapis.com.
	AutomlDataset *MigratableResource_AutomlDataset `protobuf:"bytes,3,opt,name=automl_dataset,json=automlDataset,proto3,oneof"`
}

type MigratableResource_DataLabelingDataset_ struct {
	// Output only. Represents one Dataset in datalabeling.googleapis.com.
	DataLabelingDataset *MigratableResource_DataLabelingDataset `protobuf:"bytes,4,opt,name=data_labeling_dataset,json=dataLabelingDataset,proto3,oneof"`
}

func (*MigratableResource_MlEngineModelVersion_) isMigratableResource_Resource() {}

func (*MigratableResource_AutomlModel_) isMigratableResource_Resource() {}

func (*MigratableResource_AutomlDataset_) isMigratableResource_Resource() {}

func (*MigratableResource_DataLabelingDataset_) isMigratableResource_Resource() {}

// Represents one model Version in ml.googleapis.com.
type MigratableResource_MlEngineModelVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ml.googleapis.com endpoint that this model Version currently lives
	// in.
	// Example values:
	//
	// * ml.googleapis.com
	// * us-centrall-ml.googleapis.com
	// * europe-west4-ml.googleapis.com
	// * asia-east1-ml.googleapis.com
	Endpoint string `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	// Full resource name of ml engine model Version.
	// Format: `projects/{project}/models/{model}/versions/{version}`.
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *MigratableResource_MlEngineModelVersion) Reset() {
	*x = MigratableResource_MlEngineModelVersion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgcp_cloud_aiplatform_v1_migratable_resource_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MigratableResource_MlEngineModelVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MigratableResource_MlEngineModelVersion) ProtoMessage() {}

func (x *MigratableResource_MlEngineModelVersion) ProtoReflect() protoreflect.Message {
	mi := &file_mockgcp_cloud_aiplatform_v1_migratable_resource_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MigratableResource_MlEngineModelVersion.ProtoReflect.Descriptor instead.
func (*MigratableResource_MlEngineModelVersion) Descriptor() ([]byte, []int) {
	return file_mockgcp_cloud_aiplatform_v1_migratable_resource_proto_rawDescGZIP(), []int{0, 0}
}

func (x *MigratableResource_MlEngineModelVersion) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *MigratableResource_MlEngineModelVersion) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

// Represents one Model in automl.googleapis.com.
type MigratableResource_AutomlModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Full resource name of automl Model.
	// Format:
	// `projects/{project}/locations/{location}/models/{model}`.
	Model string `protobuf:"bytes,1,opt,name=model,proto3" json:"model,omitempty"`
	// The Model's display name in automl.googleapis.com.
	ModelDisplayName string `protobuf:"bytes,3,opt,name=model_display_name,json=modelDisplayName,proto3" json:"model_display_name,omitempty"`
}

func (x *MigratableResource_AutomlModel) Reset() {
	*x = MigratableResource_AutomlModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgcp_cloud_aiplatform_v1_migratable_resource_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MigratableResource_AutomlModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MigratableResource_AutomlModel) ProtoMessage() {}

func (x *MigratableResource_AutomlModel) ProtoReflect() protoreflect.Message {
	mi := &file_mockgcp_cloud_aiplatform_v1_migratable_resource_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MigratableResource_AutomlModel.ProtoReflect.Descriptor instead.
func (*MigratableResource_AutomlModel) Descriptor() ([]byte, []int) {
	return file_mockgcp_cloud_aiplatform_v1_migratable_resource_proto_rawDescGZIP(), []int{0, 1}
}

func (x *MigratableResource_AutomlModel) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *MigratableResource_AutomlModel) GetModelDisplayName() string {
	if x != nil {
		return x.ModelDisplayName
	}
	return ""
}

// Represents one Dataset in automl.googleapis.com.
type MigratableResource_AutomlDataset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Full resource name of automl Dataset.
	// Format:
	// `projects/{project}/locations/{location}/datasets/{dataset}`.
	Dataset string `protobuf:"bytes,1,opt,name=dataset,proto3" json:"dataset,omitempty"`
	// The Dataset's display name in automl.googleapis.com.
	DatasetDisplayName string `protobuf:"bytes,4,opt,name=dataset_display_name,json=datasetDisplayName,proto3" json:"dataset_display_name,omitempty"`
}

func (x *MigratableResource_AutomlDataset) Reset() {
	*x = MigratableResource_AutomlDataset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgcp_cloud_aiplatform_v1_migratable_resource_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MigratableResource_AutomlDataset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MigratableResource_AutomlDataset) ProtoMessage() {}

func (x *MigratableResource_AutomlDataset) ProtoReflect() protoreflect.Message {
	mi := &file_mockgcp_cloud_aiplatform_v1_migratable_resource_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MigratableResource_AutomlDataset.ProtoReflect.Descriptor instead.
func (*MigratableResource_AutomlDataset) Descriptor() ([]byte, []int) {
	return file_mockgcp_cloud_aiplatform_v1_migratable_resource_proto_rawDescGZIP(), []int{0, 2}
}

func (x *MigratableResource_AutomlDataset) GetDataset() string {
	if x != nil {
		return x.Dataset
	}
	return ""
}

func (x *MigratableResource_AutomlDataset) GetDatasetDisplayName() string {
	if x != nil {
		return x.DatasetDisplayName
	}
	return ""
}

// Represents one Dataset in datalabeling.googleapis.com.
type MigratableResource_DataLabelingDataset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Full resource name of data labeling Dataset.
	// Format:
	// `projects/{project}/datasets/{dataset}`.
	Dataset string `protobuf:"bytes,1,opt,name=dataset,proto3" json:"dataset,omitempty"`
	// The Dataset's display name in datalabeling.googleapis.com.
	DatasetDisplayName string `protobuf:"bytes,4,opt,name=dataset_display_name,json=datasetDisplayName,proto3" json:"dataset_display_name,omitempty"`
	// The migratable AnnotatedDataset in datalabeling.googleapis.com belongs to
	// the data labeling Dataset.
	DataLabelingAnnotatedDatasets []*MigratableResource_DataLabelingDataset_DataLabelingAnnotatedDataset `protobuf:"bytes,3,rep,name=data_labeling_annotated_datasets,json=dataLabelingAnnotatedDatasets,proto3" json:"data_labeling_annotated_datasets,omitempty"`
}

func (x *MigratableResource_DataLabelingDataset) Reset() {
	*x = MigratableResource_DataLabelingDataset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgcp_cloud_aiplatform_v1_migratable_resource_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MigratableResource_DataLabelingDataset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MigratableResource_DataLabelingDataset) ProtoMessage() {}

func (x *MigratableResource_DataLabelingDataset) ProtoReflect() protoreflect.Message {
	mi := &file_mockgcp_cloud_aiplatform_v1_migratable_resource_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MigratableResource_DataLabelingDataset.ProtoReflect.Descriptor instead.
func (*MigratableResource_DataLabelingDataset) Descriptor() ([]byte, []int) {
	return file_mockgcp_cloud_aiplatform_v1_migratable_resource_proto_rawDescGZIP(), []int{0, 3}
}

func (x *MigratableResource_DataLabelingDataset) GetDataset() string {
	if x != nil {
		return x.Dataset
	}
	return ""
}

func (x *MigratableResource_DataLabelingDataset) GetDatasetDisplayName() string {
	if x != nil {
		return x.DatasetDisplayName
	}
	return ""
}

func (x *MigratableResource_DataLabelingDataset) GetDataLabelingAnnotatedDatasets() []*MigratableResource_DataLabelingDataset_DataLabelingAnnotatedDataset {
	if x != nil {
		return x.DataLabelingAnnotatedDatasets
	}
	return nil
}

// Represents one AnnotatedDataset in datalabeling.googleapis.com.
type MigratableResource_DataLabelingDataset_DataLabelingAnnotatedDataset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Full resource name of data labeling AnnotatedDataset.
	// Format:
	// `projects/{project}/datasets/{dataset}/annotatedDatasets/{annotated_dataset}`.
	AnnotatedDataset string `protobuf:"bytes,1,opt,name=annotated_dataset,json=annotatedDataset,proto3" json:"annotated_dataset,omitempty"`
	// The AnnotatedDataset's display name in datalabeling.googleapis.com.
	AnnotatedDatasetDisplayName string `protobuf:"bytes,3,opt,name=annotated_dataset_display_name,json=annotatedDatasetDisplayName,proto3" json:"annotated_dataset_display_name,omitempty"`
}

func (x *MigratableResource_DataLabelingDataset_DataLabelingAnnotatedDataset) Reset() {
	*x = MigratableResource_DataLabelingDataset_DataLabelingAnnotatedDataset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgcp_cloud_aiplatform_v1_migratable_resource_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MigratableResource_DataLabelingDataset_DataLabelingAnnotatedDataset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MigratableResource_DataLabelingDataset_DataLabelingAnnotatedDataset) ProtoMessage() {}

func (x *MigratableResource_DataLabelingDataset_DataLabelingAnnotatedDataset) ProtoReflect() protoreflect.Message {
	mi := &file_mockgcp_cloud_aiplatform_v1_migratable_resource_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MigratableResource_DataLabelingDataset_DataLabelingAnnotatedDataset.ProtoReflect.Descriptor instead.
func (*MigratableResource_DataLabelingDataset_DataLabelingAnnotatedDataset) Descriptor() ([]byte, []int) {
	return file_mockgcp_cloud_aiplatform_v1_migratable_resource_proto_rawDescGZIP(), []int{0, 3, 0}
}

func (x *MigratableResource_DataLabelingDataset_DataLabelingAnnotatedDataset) GetAnnotatedDataset() string {
	if x != nil {
		return x.AnnotatedDataset
	}
	return ""
}

func (x *MigratableResource_DataLabelingDataset_DataLabelingAnnotatedDataset) GetAnnotatedDatasetDisplayName() string {
	if x != nil {
		return x.AnnotatedDatasetDisplayName
	}
	return ""
}

var File_mockgcp_cloud_aiplatform_v1_migratable_resource_proto protoreflect.FileDescriptor

var file_mockgcp_cloud_aiplatform_v1_migratable_resource_proto_rawDesc = []byte{
	0x0a, 0x35, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x69,
	0x67, 0x72, 0x61, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xf5, 0x0b, 0x0a, 0x12, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x82, 0x01, 0x0a, 0x17, 0x6d, 0x6c, 0x5f,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x44, 0x2e, 0x6d, 0x6f, 0x63,
	0x6b, 0x67, 0x63, 0x70, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x4d, 0x6c, 0x45, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x00, 0x52, 0x14, 0x6d, 0x6c, 0x45, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x65, 0x0a,
	0x0c, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76,
	0x31, 0x2e, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x00, 0x52, 0x0b, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x12, 0x6b, 0x0a, 0x0e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x6d,
	0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x69, 0x67, 0x72, 0x61,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x41, 0x75,
	0x74, 0x6f, 0x6d, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x48, 0x00, 0x52, 0x0d, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65,
	0x74, 0x12, 0x7e, 0x0a, 0x15, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x69,
	0x6e, 0x67, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x43, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4d,
	0x69, 0x67, 0x72, 0x61, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x44, 0x61,
	0x74, 0x61, 0x73, 0x65, 0x74, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x00, 0x52, 0x13, 0x64, 0x61,
	0x74, 0x61, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65,
	0x74, 0x12, 0x4b, 0x0a, 0x11, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0f, 0x6c,
	0x61, 0x73, 0x74, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x49,
	0x0a, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x1a, 0x6c, 0x0a, 0x14, 0x4d, 0x6c, 0x45,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x38, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1e,
	0xfa, 0x41, 0x1b, 0x0a, 0x19, 0x6d, 0x6c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x73, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x6f, 0x6d,
	0x6c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x36, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x20, 0xfa, 0x41, 0x1d, 0x0a, 0x1b, 0x61, 0x75, 0x74, 0x6f,
	0x6d, 0x6c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x2c,
	0x0a, 0x12, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x7f, 0x0a, 0x0d,
	0x41, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x12, 0x3c, 0x0a,
	0x07, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x22,
	0xfa, 0x41, 0x1f, 0x0a, 0x1d, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x52, 0x07, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x12, 0x30, 0x0a, 0x14, 0x64,
	0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x64, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0xfd, 0x03,
	0x0a, 0x13, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x44, 0x61,
	0x74, 0x61, 0x73, 0x65, 0x74, 0x12, 0x42, 0x0a, 0x07, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x28, 0xfa, 0x41, 0x25, 0x0a, 0x23, 0x64, 0x61, 0x74,
	0x61, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74,
	0x52, 0x07, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x12, 0x30, 0x0a, 0x14, 0x64, 0x61, 0x74,
	0x61, 0x73, 0x65, 0x74, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74,
	0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0xa9, 0x01, 0x0a, 0x20,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x60, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x65,
	0x64, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x52, 0x1d, 0x64, 0x61, 0x74, 0x61, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x44,
	0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x1a, 0xc3, 0x01, 0x0a, 0x1c, 0x44, 0x61, 0x74, 0x61,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x65,
	0x64, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x12, 0x5e, 0x0a, 0x11, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x31, 0xfa, 0x41, 0x2e, 0x0a, 0x2c, 0x64, 0x61, 0x74, 0x61, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x44,
	0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x52, 0x10, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x65,
	0x64, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x12, 0x43, 0x0a, 0x1e, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x5f, 0x64, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x1b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x0a, 0x0a,
	0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0xad, 0x05, 0x0a, 0x1f, 0x63, 0x6f,
	0x6d, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x17, 0x4d,
	0x69, 0x67, 0x72, 0x61, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x69,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x70, 0x62, 0x3b, 0x61, 0x69, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x70, 0x62, 0xaa, 0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5c,
	0x56, 0x31, 0xea, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x3a, 0x3a, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x3a,
	0x56, 0x31, 0xea, 0x41, 0x51, 0x0a, 0x19, 0x6d, 0x6c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x34, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x7b, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x7d, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x7d, 0xea, 0x41, 0x55, 0x0a, 0x1b, 0x61, 0x75, 0x74, 0x6f, 0x6d,
	0x6c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x36, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x7b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x7d, 0xea, 0x41,
	0x5b, 0x0a, 0x1d, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74,
	0x12, 0x3a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65,
	0x74, 0x73, 0x2f, 0x7b, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x7d, 0xea, 0x41, 0x4c, 0x0a,
	0x23, 0x64, 0x61, 0x74, 0x61, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x61, 0x74,
	0x61, 0x73, 0x65, 0x74, 0x12, 0x25, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74,
	0x73, 0x2f, 0x7b, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x7d, 0xea, 0x41, 0x7b, 0x0a, 0x2c,
	0x64, 0x61, 0x74, 0x61, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x12, 0x4b, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x7b, 0x64, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x7d, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x44, 0x61, 0x74,
	0x61, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x7b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x7d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_mockgcp_cloud_aiplatform_v1_migratable_resource_proto_rawDescOnce sync.Once
	file_mockgcp_cloud_aiplatform_v1_migratable_resource_proto_rawDescData = file_mockgcp_cloud_aiplatform_v1_migratable_resource_proto_rawDesc
)

func file_mockgcp_cloud_aiplatform_v1_migratable_resource_proto_rawDescGZIP() []byte {
	file_mockgcp_cloud_aiplatform_v1_migratable_resource_proto_rawDescOnce.Do(func() {
		file_mockgcp_cloud_aiplatform_v1_migratable_resource_proto_rawDescData = protoimpl.X.CompressGZIP(file_mockgcp_cloud_aiplatform_v1_migratable_resource_proto_rawDescData)
	})
	return file_mockgcp_cloud_aiplatform_v1_migratable_resource_proto_rawDescData
}

var file_mockgcp_cloud_aiplatform_v1_migratable_resource_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_mockgcp_cloud_aiplatform_v1_migratable_resource_proto_goTypes = []interface{}{
	(*MigratableResource)(nil),                                                  // 0: mockgcp.cloud.aiplatform.v1.MigratableResource
	(*MigratableResource_MlEngineModelVersion)(nil),                             // 1: mockgcp.cloud.aiplatform.v1.MigratableResource.MlEngineModelVersion
	(*MigratableResource_AutomlModel)(nil),                                      // 2: mockgcp.cloud.aiplatform.v1.MigratableResource.AutomlModel
	(*MigratableResource_AutomlDataset)(nil),                                    // 3: mockgcp.cloud.aiplatform.v1.MigratableResource.AutomlDataset
	(*MigratableResource_DataLabelingDataset)(nil),                              // 4: mockgcp.cloud.aiplatform.v1.MigratableResource.DataLabelingDataset
	(*MigratableResource_DataLabelingDataset_DataLabelingAnnotatedDataset)(nil), // 5: mockgcp.cloud.aiplatform.v1.MigratableResource.DataLabelingDataset.DataLabelingAnnotatedDataset
	(*timestamp.Timestamp)(nil),                                                 // 6: google.protobuf.Timestamp
}
var file_mockgcp_cloud_aiplatform_v1_migratable_resource_proto_depIdxs = []int32{
	1, // 0: mockgcp.cloud.aiplatform.v1.MigratableResource.ml_engine_model_version:type_name -> mockgcp.cloud.aiplatform.v1.MigratableResource.MlEngineModelVersion
	2, // 1: mockgcp.cloud.aiplatform.v1.MigratableResource.automl_model:type_name -> mockgcp.cloud.aiplatform.v1.MigratableResource.AutomlModel
	3, // 2: mockgcp.cloud.aiplatform.v1.MigratableResource.automl_dataset:type_name -> mockgcp.cloud.aiplatform.v1.MigratableResource.AutomlDataset
	4, // 3: mockgcp.cloud.aiplatform.v1.MigratableResource.data_labeling_dataset:type_name -> mockgcp.cloud.aiplatform.v1.MigratableResource.DataLabelingDataset
	6, // 4: mockgcp.cloud.aiplatform.v1.MigratableResource.last_migrate_time:type_name -> google.protobuf.Timestamp
	6, // 5: mockgcp.cloud.aiplatform.v1.MigratableResource.last_update_time:type_name -> google.protobuf.Timestamp
	5, // 6: mockgcp.cloud.aiplatform.v1.MigratableResource.DataLabelingDataset.data_labeling_annotated_datasets:type_name -> mockgcp.cloud.aiplatform.v1.MigratableResource.DataLabelingDataset.DataLabelingAnnotatedDataset
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_mockgcp_cloud_aiplatform_v1_migratable_resource_proto_init() }
func file_mockgcp_cloud_aiplatform_v1_migratable_resource_proto_init() {
	if File_mockgcp_cloud_aiplatform_v1_migratable_resource_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mockgcp_cloud_aiplatform_v1_migratable_resource_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MigratableResource); i {
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
		file_mockgcp_cloud_aiplatform_v1_migratable_resource_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MigratableResource_MlEngineModelVersion); i {
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
		file_mockgcp_cloud_aiplatform_v1_migratable_resource_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MigratableResource_AutomlModel); i {
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
		file_mockgcp_cloud_aiplatform_v1_migratable_resource_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MigratableResource_AutomlDataset); i {
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
		file_mockgcp_cloud_aiplatform_v1_migratable_resource_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MigratableResource_DataLabelingDataset); i {
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
		file_mockgcp_cloud_aiplatform_v1_migratable_resource_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MigratableResource_DataLabelingDataset_DataLabelingAnnotatedDataset); i {
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
	file_mockgcp_cloud_aiplatform_v1_migratable_resource_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*MigratableResource_MlEngineModelVersion_)(nil),
		(*MigratableResource_AutomlModel_)(nil),
		(*MigratableResource_AutomlDataset_)(nil),
		(*MigratableResource_DataLabelingDataset_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mockgcp_cloud_aiplatform_v1_migratable_resource_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mockgcp_cloud_aiplatform_v1_migratable_resource_proto_goTypes,
		DependencyIndexes: file_mockgcp_cloud_aiplatform_v1_migratable_resource_proto_depIdxs,
		MessageInfos:      file_mockgcp_cloud_aiplatform_v1_migratable_resource_proto_msgTypes,
	}.Build()
	File_mockgcp_cloud_aiplatform_v1_migratable_resource_proto = out.File
	file_mockgcp_cloud_aiplatform_v1_migratable_resource_proto_rawDesc = nil
	file_mockgcp_cloud_aiplatform_v1_migratable_resource_proto_goTypes = nil
	file_mockgcp_cloud_aiplatform_v1_migratable_resource_proto_depIdxs = nil
}
