// Code generated by protoc-gen-go.
// source: db.proto
// DO NOT EDIT!

/*
Package db is a generated protocol buffer package.

It is generated from these files:
	db.proto

It has these top-level messages:
	ProjectData
	Expert
	DatasetData
	Job
	TenantStats
	AssetCache
	ChunkData
	SourceLink
	VersionData
	FieldSchema
	DatasetDescriptor
*/
package db

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import events "mlp/catalog/events"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Range int32

const (
	Range_Events           Range = 0
	Range_PROJECTS         Range = 1
	Range_DATASETS         Range = 2
	Range_PROJECT_DATASETS Range = 3
	Range_JOBS             Range = 4
	Range_STATS            Range = 5
	Range_ASSET_CACHE      Range = 6
	Range_EXPERTS          Range = 7
)

var Range_name = map[int32]string{
	0: "Events",
	1: "PROJECTS",
	2: "DATASETS",
	3: "PROJECT_DATASETS",
	4: "JOBS",
	5: "STATS",
	6: "ASSET_CACHE",
	7: "EXPERTS",
}
var Range_value = map[string]int32{
	"Events":           0,
	"PROJECTS":         1,
	"DATASETS":         2,
	"PROJECT_DATASETS": 3,
	"JOBS":             4,
	"STATS":            5,
	"ASSET_CACHE":      6,
	"EXPERTS":          7,
}

func (x Range) String() string {
	return proto.EnumName(Range_name, int32(x))
}
func (Range) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type LinkType int32

const (
	LinkType_UNIVERSAL LinkType = 0
	LinkType_UPLOAD    LinkType = 1
)

var LinkType_name = map[int32]string{
	0: "UNIVERSAL",
	1: "UPLOAD",
}
var LinkType_value = map[string]int32{
	"UNIVERSAL": 0,
	"UPLOAD":    1,
}

func (x LinkType) String() string {
	return proto.EnumName(LinkType_name, int32(x))
}
func (LinkType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ProjectData struct {
	Name         string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Id           string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	DatasetCount int32  `protobuf:"varint,3,opt,name=dataset_count,json=datasetCount" json:"dataset_count,omitempty"`
	RawBytes     int64  `protobuf:"varint,4,opt,name=raw_bytes,json=rawBytes" json:"raw_bytes,omitempty"`
	ZipBytes     int64  `protobuf:"varint,5,opt,name=zip_bytes,json=zipBytes" json:"zip_bytes,omitempty"`
	JobCount     int32  `protobuf:"varint,6,opt,name=job_count,json=jobCount" json:"job_count,omitempty"`
}

func (m *ProjectData) Reset()                    { *m = ProjectData{} }
func (m *ProjectData) String() string            { return proto.CompactTextString(m) }
func (*ProjectData) ProtoMessage()               {}
func (*ProjectData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Expert struct {
	ExpertId   string `protobuf:"bytes,1,opt,name=expert_id,json=expertId" json:"expert_id,omitempty"`
	ExpertName string `protobuf:"bytes,2,opt,name=expert_name,json=expertName" json:"expert_name,omitempty"`
}

func (m *Expert) Reset()                    { *m = Expert{} }
func (m *Expert) String() string            { return proto.CompactTextString(m) }
func (*Expert) ProtoMessage()               {}
func (*Expert) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type DatasetData struct {
	DatasetId       string                `protobuf:"bytes,1,opt,name=dataset_id,json=datasetId" json:"dataset_id,omitempty"`
	Name            string                `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	ProjectId       string                `protobuf:"bytes,3,opt,name=project_id,json=projectId" json:"project_id,omitempty"`
	RecordCount     int64                 `protobuf:"varint,4,opt,name=record_count,json=recordCount" json:"record_count,omitempty"`
	FileCount       int64                 `protobuf:"varint,5,opt,name=file_count,json=fileCount" json:"file_count,omitempty"`
	StorageBytes    int64                 `protobuf:"varint,6,opt,name=storage_bytes,json=storageBytes" json:"storage_bytes,omitempty"`
	Sample          *events.DatasetSample `protobuf:"bytes,7,opt,name=sample" json:"sample,omitempty"`
	UpdateTimestamp int64                 `protobuf:"varint,8,opt,name=update_timestamp,json=updateTimestamp" json:"update_timestamp,omitempty"`
	DataFormat      string                `protobuf:"bytes,9,opt,name=data_format,json=dataFormat" json:"data_format,omitempty"`
	Description     string                `protobuf:"bytes,10,opt,name=description" json:"description,omitempty"`
	LocationId      string                `protobuf:"bytes,11,opt,name=location_id,json=locationId" json:"location_id,omitempty"`
	LocationUri     string                `protobuf:"bytes,12,opt,name=location_uri,json=locationUri" json:"location_uri,omitempty"`
	Experts         []string              `protobuf:"bytes,13,rep,name=experts" json:"experts,omitempty"`
	UpstreamJobs    []string              `protobuf:"bytes,14,rep,name=upstream_jobs,json=upstreamJobs" json:"upstream_jobs,omitempty"`
	DownstreamJobs  []string              `protobuf:"bytes,15,rep,name=downstream_jobs,json=downstreamJobs" json:"downstream_jobs,omitempty"`
}

func (m *DatasetData) Reset()                    { *m = DatasetData{} }
func (m *DatasetData) String() string            { return proto.CompactTextString(m) }
func (*DatasetData) ProtoMessage()               {}
func (*DatasetData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DatasetData) GetSample() *events.DatasetSample {
	if m != nil {
		return m.Sample
	}
	return nil
}

type Job struct {
	JobId     string    `protobuf:"bytes,1,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	JobName   string    `protobuf:"bytes,2,opt,name=job_name,json=jobName" json:"job_name,omitempty"`
	ProjectId string    `protobuf:"bytes,3,opt,name=project_id,json=projectId" json:"project_id,omitempty"`
	Inputs    []string  `protobuf:"bytes,4,rep,name=inputs" json:"inputs,omitempty"`
	Outputs   []string  `protobuf:"bytes,5,rep,name=outputs" json:"outputs,omitempty"`
	Experts   []*Expert `protobuf:"bytes,6,rep,name=experts" json:"experts,omitempty"`
}

func (m *Job) Reset()                    { *m = Job{} }
func (m *Job) String() string            { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()               {}
func (*Job) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Job) GetExperts() []*Expert {
	if m != nil {
		return m.Experts
	}
	return nil
}

type TenantStats struct {
	ProjectCount int32 `protobuf:"varint,1,opt,name=project_count,json=projectCount" json:"project_count,omitempty"`
	DatasetCount int32 `protobuf:"varint,2,opt,name=dataset_count,json=datasetCount" json:"dataset_count,omitempty"`
	JobCount     int32 `protobuf:"varint,3,opt,name=job_count,json=jobCount" json:"job_count,omitempty"`
	ReportCount  int32 `protobuf:"varint,4,opt,name=report_count,json=reportCount" json:"report_count,omitempty"`
	ExpertCount  int32 `protobuf:"varint,5,opt,name=expert_count,json=expertCount" json:"expert_count,omitempty"`
	ModelCount   int32 `protobuf:"varint,6,opt,name=model_count,json=modelCount" json:"model_count,omitempty"`
}

func (m *TenantStats) Reset()                    { *m = TenantStats{} }
func (m *TenantStats) String() string            { return proto.CompactTextString(m) }
func (*TenantStats) ProtoMessage()               {}
func (*TenantStats) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type AssetCache struct {
	Digest []byte `protobuf:"bytes,1,opt,name=digest,proto3" json:"digest,omitempty"`
	Body   []byte `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *AssetCache) Reset()                    { *m = AssetCache{} }
func (m *AssetCache) String() string            { return proto.CompactTextString(m) }
func (*AssetCache) ProtoMessage()               {}
func (*AssetCache) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type ChunkData struct {
	Hash    string `protobuf:"bytes,1,opt,name=hash" json:"hash,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	RawSize int64  `protobuf:"varint,3,opt,name=raw_size,json=rawSize" json:"raw_size,omitempty"`
	ZipSize int64  `protobuf:"varint,4,opt,name=zip_size,json=zipSize" json:"zip_size,omitempty"`
}

func (m *ChunkData) Reset()                    { *m = ChunkData{} }
func (m *ChunkData) String() string            { return proto.CompactTextString(m) }
func (*ChunkData) ProtoMessage()               {}
func (*ChunkData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type SourceLink struct {
	Kind       LinkType `protobuf:"varint,1,opt,name=kind,enum=LinkType" json:"kind,omitempty"`
	SourceName string   `protobuf:"bytes,2,opt,name=source_name,json=sourceName" json:"source_name,omitempty"`
	UserName   string   `protobuf:"bytes,3,opt,name=user_name,json=userName" json:"user_name,omitempty"`
}

func (m *SourceLink) Reset()                    { *m = SourceLink{} }
func (m *SourceLink) String() string            { return proto.CompactTextString(m) }
func (*SourceLink) ProtoMessage()               {}
func (*SourceLink) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type VersionData struct {
	Version     int32              `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	Chunks      []*ChunkData       `protobuf:"bytes,2,rep,name=chunks" json:"chunks,omitempty"`
	Sources     []*SourceLink      `protobuf:"bytes,3,rep,name=sources" json:"sources,omitempty"`
	RawSize     int64              `protobuf:"varint,4,opt,name=raw_size,json=rawSize" json:"raw_size,omitempty"`
	ZipSize     int64              `protobuf:"varint,5,opt,name=zip_size,json=zipSize" json:"zip_size,omitempty"`
	DatasetId   string             `protobuf:"bytes,6,opt,name=dataset_id,json=datasetId" json:"dataset_id,omitempty"`
	Descriptor_ *DatasetDescriptor `protobuf:"bytes,7,opt,name=descriptor" json:"descriptor,omitempty"`
}

func (m *VersionData) Reset()                    { *m = VersionData{} }
func (m *VersionData) String() string            { return proto.CompactTextString(m) }
func (*VersionData) ProtoMessage()               {}
func (*VersionData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *VersionData) GetChunks() []*ChunkData {
	if m != nil {
		return m.Chunks
	}
	return nil
}

func (m *VersionData) GetSources() []*SourceLink {
	if m != nil {
		return m.Sources
	}
	return nil
}

func (m *VersionData) GetDescriptor_() *DatasetDescriptor {
	if m != nil {
		return m.Descriptor_
	}
	return nil
}

type FieldSchema struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
}

func (m *FieldSchema) Reset()                    { *m = FieldSchema{} }
func (m *FieldSchema) String() string            { return proto.CompactTextString(m) }
func (*FieldSchema) ProtoMessage()               {}
func (*FieldSchema) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type DatasetDescriptor struct {
	Type         string         `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	RecordCount  int64          `protobuf:"varint,2,opt,name=record_count,json=recordCount" json:"record_count,omitempty"`
	Schema       []*FieldSchema `protobuf:"bytes,3,rep,name=schema" json:"schema,omitempty"`
	SampleString string         `protobuf:"bytes,4,opt,name=sample_string,json=sampleString" json:"sample_string,omitempty"`
}

func (m *DatasetDescriptor) Reset()                    { *m = DatasetDescriptor{} }
func (m *DatasetDescriptor) String() string            { return proto.CompactTextString(m) }
func (*DatasetDescriptor) ProtoMessage()               {}
func (*DatasetDescriptor) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *DatasetDescriptor) GetSchema() []*FieldSchema {
	if m != nil {
		return m.Schema
	}
	return nil
}

func init() {
	proto.RegisterType((*ProjectData)(nil), "ProjectData")
	proto.RegisterType((*Expert)(nil), "Expert")
	proto.RegisterType((*DatasetData)(nil), "DatasetData")
	proto.RegisterType((*Job)(nil), "Job")
	proto.RegisterType((*TenantStats)(nil), "TenantStats")
	proto.RegisterType((*AssetCache)(nil), "AssetCache")
	proto.RegisterType((*ChunkData)(nil), "ChunkData")
	proto.RegisterType((*SourceLink)(nil), "SourceLink")
	proto.RegisterType((*VersionData)(nil), "VersionData")
	proto.RegisterType((*FieldSchema)(nil), "FieldSchema")
	proto.RegisterType((*DatasetDescriptor)(nil), "DatasetDescriptor")
	proto.RegisterEnum("Range", Range_name, Range_value)
	proto.RegisterEnum("LinkType", LinkType_name, LinkType_value)
}

func init() { proto.RegisterFile("db.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1012 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x56, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x0e, 0xf5, 0x43, 0x89, 0x43, 0xda, 0x66, 0x17, 0x6d, 0xc1, 0x34, 0x30, 0xa2, 0x28, 0x4d,
	0xab, 0xe6, 0xa0, 0x83, 0x8b, 0x02, 0xbd, 0x2a, 0xb6, 0x8c, 0xda, 0x30, 0x6c, 0x83, 0xa4, 0x83,
	0xde, 0x88, 0xa5, 0xb8, 0xb1, 0x57, 0xb6, 0xb8, 0xc4, 0xee, 0x2a, 0x8e, 0xf5, 0x04, 0x7d, 0x81,
	0x3e, 0x45, 0xd1, 0x27, 0xea, 0xa3, 0xf4, 0x52, 0xec, 0x0f, 0x29, 0x29, 0x76, 0x81, 0xdc, 0x76,
	0xbe, 0x6f, 0x34, 0xdc, 0xf9, 0xf8, 0xcd, 0x50, 0xd0, 0x2f, 0xf2, 0x71, 0xc5, 0x99, 0x64, 0xdf,
	0x05, 0xe4, 0x23, 0x29, 0xa5, 0x30, 0xd1, 0xf0, 0x6f, 0x07, 0xfc, 0x4b, 0xce, 0xe6, 0x64, 0x26,
	0x8f, 0xb0, 0xc4, 0x08, 0x41, 0xa7, 0xc4, 0x0b, 0x12, 0x39, 0x03, 0x67, 0xe4, 0xc5, 0xfa, 0x8c,
	0x76, 0xa1, 0x45, 0x8b, 0xa8, 0xa5, 0x91, 0x16, 0x2d, 0xd0, 0x6b, 0xd8, 0x29, 0xb0, 0xc4, 0x82,
	0xc8, 0x6c, 0xc6, 0x96, 0xa5, 0x8c, 0xda, 0x03, 0x67, 0xd4, 0x8d, 0x03, 0x0b, 0x1e, 0x2a, 0x0c,
	0xbd, 0x00, 0x8f, 0xe3, 0xfb, 0x2c, 0x7f, 0x90, 0x44, 0x44, 0x9d, 0x81, 0x33, 0x6a, 0xc7, 0x7d,
	0x8e, 0xef, 0xdf, 0xa9, 0x58, 0x91, 0x2b, 0x5a, 0x59, 0xb2, 0x6b, 0xc8, 0x15, 0xad, 0x1a, 0x72,
	0xce, 0x72, 0x5b, 0xda, 0xd5, 0xa5, 0xfb, 0x73, 0x96, 0xeb, 0xb2, 0xc3, 0x63, 0x70, 0xa7, 0x9f,
	0x2a, 0xc2, 0xf5, 0x03, 0x88, 0x3e, 0x65, 0xb4, 0xb0, 0xd7, 0xed, 0x1b, 0xe0, 0xa4, 0x40, 0x2f,
	0xc1, 0xb7, 0xa4, 0xee, 0xc6, 0xdc, 0x1d, 0x0c, 0x74, 0x8e, 0x17, 0x64, 0xf8, 0x47, 0x07, 0xfc,
	0x23, 0x73, 0x5f, 0xdd, 0xf7, 0x3e, 0x40, 0xdd, 0x53, 0x53, 0xce, 0xb3, 0xc8, 0x49, 0xd1, 0xc8,
	0xd2, 0xda, 0x90, 0x65, 0x1f, 0xa0, 0x32, 0xca, 0xa9, 0x9f, 0xb4, 0xcd, 0x4f, 0x2c, 0x72, 0x52,
	0xa0, 0x57, 0x10, 0x70, 0x32, 0x63, 0xbc, 0xb0, 0x9d, 0x18, 0x0d, 0x7c, 0x83, 0x19, 0x8d, 0xf6,
	0x01, 0x3e, 0xd0, 0x3b, 0x62, 0x13, 0x8c, 0x0e, 0x9e, 0x42, 0x0c, 0xfd, 0x1a, 0x76, 0x84, 0x64,
	0x1c, 0x5f, 0x13, 0xab, 0x94, 0xab, 0x33, 0x02, 0x0b, 0x1a, 0xb5, 0x7e, 0x00, 0x57, 0xe0, 0x45,
	0x75, 0x47, 0xa2, 0xde, 0xc0, 0x19, 0xf9, 0x07, 0xbb, 0x63, 0xdb, 0x56, 0xa2, 0xd1, 0xd8, 0xb2,
	0xe8, 0x27, 0x08, 0x97, 0x55, 0x81, 0x25, 0xc9, 0x24, 0x5d, 0x10, 0x21, 0xf1, 0xa2, 0x8a, 0xfa,
	0xba, 0xde, 0x9e, 0xc1, 0xd3, 0x1a, 0x56, 0xe2, 0xa9, 0xce, 0xb3, 0x0f, 0x8c, 0x2f, 0xb0, 0x8c,
	0x3c, 0x23, 0x9e, 0x82, 0x8e, 0x35, 0x82, 0x06, 0xe0, 0x17, 0x44, 0xcc, 0x38, 0xad, 0x24, 0x65,
	0x65, 0x04, 0x3a, 0x61, 0x13, 0x52, 0x25, 0xee, 0xd8, 0x0c, 0xab, 0xb3, 0x12, 0xc7, 0x37, 0x25,
	0x6a, 0xc8, 0xa8, 0xd3, 0x24, 0x2c, 0x39, 0x8d, 0x02, 0x53, 0xa3, 0xc6, 0xae, 0x38, 0x45, 0x11,
	0xf4, 0xcc, 0x0b, 0x13, 0xd1, 0xce, 0xa0, 0x3d, 0xf2, 0xe2, 0x3a, 0x54, 0xc2, 0x2c, 0x2b, 0x21,
	0x39, 0xc1, 0x8b, 0x6c, 0xce, 0x72, 0x11, 0xed, 0x6a, 0x3e, 0xa8, 0xc1, 0x53, 0x96, 0x0b, 0xf4,
	0x23, 0xec, 0x15, 0xec, 0xbe, 0xdc, 0x4c, 0xdb, 0xd3, 0x69, 0xbb, 0x6b, 0x58, 0x25, 0x0e, 0xff,
	0x72, 0xa0, 0x7d, 0xca, 0x72, 0xf4, 0x0d, 0xb8, 0xca, 0x77, 0xcd, 0xeb, 0xef, 0xce, 0x59, 0x7e,
	0x52, 0xa0, 0xe7, 0xa0, 0xdc, 0xb7, 0xe9, 0xa3, 0xde, 0x9c, 0xe5, 0xe7, 0x5f, 0xe0, 0x80, 0x6f,
	0xc1, 0xa5, 0x65, 0xb5, 0x94, 0xca, 0xff, 0xea, 0xc1, 0x36, 0x52, 0x8d, 0xb1, 0xa5, 0xd4, 0x44,
	0xd7, 0x34, 0x66, 0x43, 0xf4, 0x6a, 0xdd, 0xb2, 0x3b, 0x68, 0x8f, 0xfc, 0x83, 0xde, 0xd8, 0xb8,
	0xbd, 0xe9, 0x7d, 0xf8, 0x8f, 0x03, 0x7e, 0x4a, 0x4a, 0x5c, 0xca, 0x44, 0x62, 0xa3, 0x45, 0x7d,
	0x07, 0x63, 0x23, 0xc7, 0x0c, 0xa3, 0x05, 0x1b, 0x27, 0x6d, 0x4f, 0x6c, 0xeb, 0xe9, 0x89, 0x5d,
	0xcf, 0x5d, 0x7b, 0x7b, 0xee, 0x8c, 0x9b, 0x2b, 0xc6, 0xe5, 0x86, 0x9b, 0xbb, 0xca, 0xcd, 0x0a,
	0x6b, 0x52, 0xec, 0xcc, 0xad, 0xfd, 0xdc, 0x8d, 0xed, 0x1c, 0x9a, 0x94, 0x97, 0xe0, 0x2f, 0x58,
	0x41, 0xee, 0xb6, 0x86, 0x1b, 0x34, 0x64, 0xc6, 0xfb, 0x57, 0x80, 0x89, 0x50, 0x37, 0xc2, 0xb3,
	0x1b, 0xa2, 0x04, 0x2c, 0xe8, 0x35, 0x11, 0xa6, 0xa9, 0x20, 0xb6, 0x91, 0x9a, 0xc6, 0x9c, 0x15,
	0x0f, 0xba, 0x8b, 0x20, 0xd6, 0xe7, 0x21, 0x05, 0xef, 0xf0, 0x66, 0x59, 0xde, 0xd6, 0x5b, 0xec,
	0x06, 0x8b, 0x9b, 0x7a, 0x8b, 0xa9, 0xf3, 0x93, 0x23, 0xfc, 0x1c, 0xd4, 0x4e, 0xca, 0x04, 0x5d,
	0x11, 0xdd, 0x71, 0x3b, 0xee, 0x71, 0x7c, 0x9f, 0xd0, 0x95, 0xa6, 0xd4, 0x8a, 0xd2, 0x94, 0x19,
	0xdd, 0xde, 0x8a, 0x56, 0x8a, 0x1a, 0x52, 0x80, 0x84, 0x2d, 0xf9, 0x8c, 0x9c, 0xd1, 0xf2, 0x16,
	0xed, 0x43, 0xe7, 0x96, 0x96, 0xc6, 0x34, 0xbb, 0x07, 0xde, 0x58, 0x81, 0xe9, 0x43, 0x45, 0x62,
	0x0d, 0xab, 0x96, 0x85, 0x4e, 0xde, 0xda, 0x44, 0x06, 0xd2, 0x26, 0x7a, 0x01, 0xde, 0x52, 0x10,
	0x6e, 0x68, 0xe3, 0xa1, 0xbe, 0x02, 0xf4, 0x9a, 0xfa, 0xd7, 0x01, 0xff, 0x3d, 0xe1, 0x82, 0xb2,
	0x52, 0x37, 0x16, 0x41, 0xef, 0xa3, 0x09, 0xed, 0x7b, 0xae, 0x43, 0x34, 0x04, 0x77, 0xa6, 0xfa,
	0x17, 0x51, 0x4b, 0x3b, 0x07, 0xc6, 0x8d, 0x1c, 0xb1, 0x65, 0xd0, 0x1b, 0xe8, 0x99, 0x07, 0x8b,
	0xa8, 0xad, 0x93, 0xfc, 0xf1, 0xba, 0x91, 0xb8, 0xe6, 0xb6, 0x54, 0xe9, 0xfc, 0xbf, 0x2a, 0xdd,
	0x2d, 0x55, 0x3e, 0xdb, 0xa0, 0xee, 0xe7, 0x1b, 0xf4, 0x00, 0xa0, 0x5e, 0x10, 0x8c, 0xdb, 0x5d,
	0x85, 0xea, 0x5d, 0x75, 0xd4, 0x30, 0xf1, 0x46, 0xd6, 0xf0, 0x17, 0xf0, 0x8f, 0x29, 0xb9, 0x2b,
	0x92, 0xd9, 0x0d, 0x59, 0x3c, 0xfd, 0x6d, 0x42, 0xd0, 0x91, 0x0f, 0x55, 0xf3, 0x56, 0xd5, 0x79,
	0xf8, 0xa7, 0x03, 0x5f, 0x3d, 0x2a, 0xdc, 0x64, 0x3a, 0xeb, 0xcc, 0x47, 0x3b, 0xba, 0xf5, 0x78,
	0x47, 0x7f, 0x0f, 0xae, 0xd0, 0x8f, 0xb7, 0x92, 0x05, 0xe3, 0x8d, 0x2b, 0xc5, 0x96, 0xd3, 0xab,
	0x5a, 0xef, 0xd9, 0x4c, 0x48, 0x4e, 0xcb, 0x6b, 0xad, 0x9b, 0x17, 0x07, 0x06, 0x4c, 0x34, 0xf6,
	0xf6, 0x13, 0x74, 0x63, 0x5c, 0x5e, 0x13, 0x04, 0xe0, 0x4e, 0xf5, 0x47, 0x38, 0x7c, 0x86, 0x02,
	0xe8, 0x5f, 0xc6, 0x17, 0xa7, 0xd3, 0xc3, 0x34, 0x09, 0x1d, 0x15, 0x1d, 0x4d, 0xd2, 0x49, 0x32,
	0x4d, 0x93, 0xb0, 0x85, 0xbe, 0x86, 0xd0, 0x72, 0x59, 0x83, 0xb6, 0x51, 0x1f, 0x3a, 0xa7, 0x17,
	0xef, 0x92, 0xb0, 0x83, 0x3c, 0xe8, 0x26, 0xe9, 0x24, 0x4d, 0xc2, 0x2e, 0xda, 0x03, 0x7f, 0x92,
	0x24, 0xd3, 0x34, 0x3b, 0x9c, 0x1c, 0xfe, 0x36, 0x0d, 0x5d, 0xe4, 0x43, 0x6f, 0xfa, 0xfb, 0xe5,
	0x34, 0x4e, 0x93, 0xb0, 0xf7, 0xf6, 0x0d, 0xf4, 0x6b, 0x5b, 0xa2, 0x1d, 0xf0, 0xae, 0xce, 0x4f,
	0xde, 0x4f, 0xe3, 0x64, 0x72, 0x16, 0x3e, 0x53, 0x77, 0xb9, 0xba, 0x3c, 0xbb, 0x98, 0x1c, 0x85,
	0x4e, 0xee, 0xea, 0xff, 0x04, 0x3f, 0xff, 0x17, 0x00, 0x00, 0xff, 0xff, 0xe9, 0x68, 0x72, 0xda,
	0x2d, 0x08, 0x00, 0x00,
}
