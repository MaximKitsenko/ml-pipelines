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
	DatasetId          string             `protobuf:"bytes,1,opt,name=dataset_id,json=datasetId" json:"dataset_id,omitempty"`
	Name               string             `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	ProjectId          string             `protobuf:"bytes,3,opt,name=project_id,json=projectId" json:"project_id,omitempty"`
	RecordCount        int64              `protobuf:"varint,4,opt,name=record_count,json=recordCount" json:"record_count,omitempty"`
	RecordCountSet     bool               `protobuf:"varint,5,opt,name=record_count_set,json=recordCountSet" json:"record_count_set,omitempty"`
	FileCount          int64              `protobuf:"varint,6,opt,name=file_count,json=fileCount" json:"file_count,omitempty"`
	FileCountSet       bool               `protobuf:"varint,7,opt,name=file_count_set,json=fileCountSet" json:"file_count_set,omitempty"`
	RawBytes           int64              `protobuf:"varint,8,opt,name=raw_bytes,json=rawBytes" json:"raw_bytes,omitempty"`
	RawBytesSet        bool               `protobuf:"varint,9,opt,name=raw_bytes_set,json=rawBytesSet" json:"raw_bytes_set,omitempty"`
	ZipBytes           int64              `protobuf:"varint,10,opt,name=zip_bytes,json=zipBytes" json:"zip_bytes,omitempty"`
	ZipBytesSet        bool               `protobuf:"varint,11,opt,name=zip_bytes_set,json=zipBytesSet" json:"zip_bytes_set,omitempty"`
	SampleBody         []byte             `protobuf:"bytes,12,opt,name=sample_body,json=sampleBody,proto3" json:"sample_body,omitempty"`
	SampleKind         events.SAMPLE_KIND `protobuf:"varint,13,opt,name=sample_kind,json=sampleKind,enum=SAMPLE_KIND" json:"sample_kind,omitempty"`
	SampleSet          bool               `protobuf:"varint,14,opt,name=sample_set,json=sampleSet" json:"sample_set,omitempty"`
	UpdateTimestamp    int64              `protobuf:"varint,15,opt,name=update_timestamp,json=updateTimestamp" json:"update_timestamp,omitempty"`
	UpdateTimestampSet bool               `protobuf:"varint,16,opt,name=update_timestamp_set,json=updateTimestampSet" json:"update_timestamp_set,omitempty"`
	DataFormat         string             `protobuf:"bytes,17,opt,name=data_format,json=dataFormat" json:"data_format,omitempty"`
	DataFormatSet      bool               `protobuf:"varint,18,opt,name=data_format_set,json=dataFormatSet" json:"data_format_set,omitempty"`
	Description        string             `protobuf:"bytes,19,opt,name=description" json:"description,omitempty"`
	DescriptionSet     bool               `protobuf:"varint,20,opt,name=description_set,json=descriptionSet" json:"description_set,omitempty"`
	UpstreamJobs       []string           `protobuf:"bytes,21,rep,name=upstream_jobs,json=upstreamJobs" json:"upstream_jobs,omitempty"`
	DownstreamJobs     []string           `protobuf:"bytes,22,rep,name=downstream_jobs,json=downstreamJobs" json:"downstream_jobs,omitempty"`
	Experts            []*Expert          `protobuf:"bytes,23,rep,name=experts" json:"experts,omitempty"`
	LocationId         string             `protobuf:"bytes,24,opt,name=location_id,json=locationId" json:"location_id,omitempty"`
	LocationIdSet      bool               `protobuf:"varint,25,opt,name=location_id_set,json=locationIdSet" json:"location_id_set,omitempty"`
	LocationUri        string             `protobuf:"bytes,26,opt,name=location_uri,json=locationUri" json:"location_uri,omitempty"`
	LocationUriSet     bool               `protobuf:"varint,27,opt,name=location_uri_set,json=locationUriSet" json:"location_uri_set,omitempty"`
}

func (m *DatasetData) Reset()                    { *m = DatasetData{} }
func (m *DatasetData) String() string            { return proto.CompactTextString(m) }
func (*DatasetData) ProtoMessage()               {}
func (*DatasetData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DatasetData) GetExperts() []*Expert {
	if m != nil {
		return m.Experts
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
	Digest string `protobuf:"bytes,1,opt,name=digest" json:"digest,omitempty"`
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
	// 1167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x56, 0xdd, 0x6e, 0xdb, 0x36,
	0x18, 0xad, 0xfc, 0x27, 0xeb, 0x93, 0xec, 0xa8, 0x5c, 0xdb, 0xa9, 0x2d, 0x82, 0x38, 0x6a, 0xbb,
	0x79, 0x05, 0x66, 0x0c, 0x19, 0x06, 0xec, 0xd6, 0x49, 0x1c, 0xcc, 0x69, 0x96, 0x04, 0x92, 0x53,
	0xec, 0x4e, 0x90, 0x2d, 0x36, 0xa1, 0x13, 0x8b, 0x82, 0x44, 0x37, 0x4d, 0xde, 0x63, 0x4f, 0x31,
	0xec, 0x85, 0xb6, 0x47, 0xd9, 0xcd, 0xc0, 0x8f, 0x92, 0x2c, 0x3b, 0x19, 0xb0, 0x3b, 0xf1, 0x7c,
	0xc7, 0x87, 0xe4, 0x21, 0x79, 0x3e, 0x43, 0x3b, 0x9a, 0x0e, 0x92, 0x94, 0x0b, 0xfe, 0xca, 0xa2,
	0x9f, 0x69, 0x2c, 0x32, 0x35, 0x72, 0xff, 0xd4, 0xc0, 0x3c, 0x4f, 0xf9, 0x9c, 0xce, 0xc4, 0x61,
	0x28, 0x42, 0x42, 0xa0, 0x11, 0x87, 0x0b, 0xea, 0x68, 0x3d, 0xad, 0x6f, 0x78, 0xf8, 0x4d, 0xba,
	0x50, 0x63, 0x91, 0x53, 0x43, 0xa4, 0xc6, 0x22, 0xf2, 0x06, 0x3a, 0x51, 0x28, 0xc2, 0x8c, 0x8a,
	0x60, 0xc6, 0x97, 0xb1, 0x70, 0xea, 0x3d, 0xad, 0xdf, 0xf4, 0xac, 0x1c, 0x3c, 0x90, 0x18, 0x79,
	0x0d, 0x46, 0x1a, 0xde, 0x06, 0xd3, 0x3b, 0x41, 0x33, 0xa7, 0xd1, 0xd3, 0xfa, 0x75, 0xaf, 0x9d,
	0x86, 0xb7, 0xfb, 0x72, 0x2c, 0x8b, 0xf7, 0x2c, 0xc9, 0x8b, 0x4d, 0x55, 0xbc, 0x67, 0x49, 0x59,
	0x9c, 0xf3, 0x69, 0x2e, 0xdd, 0x42, 0xe9, 0xf6, 0x9c, 0x4f, 0x51, 0xd6, 0x3d, 0x82, 0xd6, 0xe8,
	0x4b, 0x42, 0x53, 0x9c, 0x80, 0xe2, 0x57, 0xc0, 0xa2, 0x7c, 0xb9, 0x6d, 0x05, 0x8c, 0x23, 0xb2,
	0x03, 0x66, 0x5e, 0xc4, 0xdd, 0xa8, 0xb5, 0x83, 0x82, 0x4e, 0xc3, 0x05, 0x75, 0xff, 0xd2, 0xc1,
	0x3c, 0x54, 0xeb, 0xc5, 0x7d, 0x6f, 0x03, 0x14, 0x7b, 0x2a, 0xe5, 0x8c, 0x1c, 0x19, 0x47, 0xa5,
	0x2d, 0xb5, 0x8a, 0x2d, 0xdb, 0x00, 0x89, 0x72, 0x4e, 0xfe, 0xa4, 0xae, 0x7e, 0x92, 0x23, 0xe3,
	0x88, 0xec, 0x82, 0x95, 0xd2, 0x19, 0x4f, 0xa3, 0x7c, 0x27, 0xca, 0x03, 0x53, 0x61, 0xca, 0xa3,
	0x3e, 0xd8, 0x55, 0x4a, 0x90, 0x51, 0x81, 0x6e, 0xb4, 0xbd, 0x6e, 0x85, 0xe6, 0x53, 0x21, 0xe7,
	0xfa, 0xc4, 0x6e, 0x68, 0xc5, 0x94, 0xba, 0x67, 0x48, 0x44, 0x09, 0xbd, 0x85, 0xee, 0xaa, 0x8c,
	0x32, 0x3a, 0xca, 0x58, 0x25, 0x45, 0x8a, 0xac, 0x1d, 0x49, 0x7b, 0xe3, 0x48, 0x5c, 0xe8, 0x94,
	0x45, 0x54, 0x30, 0x50, 0xc1, 0x2c, 0x08, 0xb9, 0xc0, 0xea, 0xd8, 0x60, 0xe3, 0xd8, 0x5c, 0xe8,
	0x94, 0x45, 0x14, 0x30, 0x95, 0x40, 0x41, 0x90, 0x02, 0x3b, 0x60, 0x66, 0xe1, 0x22, 0xb9, 0xa1,
	0xc1, 0x94, 0x47, 0x77, 0x8e, 0xd5, 0xd3, 0xfa, 0x96, 0x07, 0x0a, 0xda, 0xe7, 0xd1, 0x1d, 0xf9,
	0xbe, 0x24, 0x5c, 0xb3, 0x38, 0x72, 0x3a, 0x3d, 0xad, 0xdf, 0xdd, 0xb3, 0x06, 0xfe, 0xf0, 0xd7,
	0xf3, 0x93, 0x51, 0xf0, 0x61, 0x7c, 0x7a, 0x58, 0xd0, 0x3f, 0xb0, 0x38, 0x92, 0xb6, 0xe4, 0x74,
	0x39, 0x61, 0x17, 0x27, 0x34, 0x14, 0x22, 0xa7, 0xfb, 0x0e, 0xec, 0x65, 0x12, 0x85, 0x82, 0x06,
	0x82, 0x2d, 0x68, 0x26, 0xc2, 0x45, 0xe2, 0x6c, 0xe1, 0xb2, 0xb7, 0x14, 0x3e, 0x29, 0x60, 0xf2,
	0x03, 0x3c, 0xdb, 0xa4, 0xa2, 0xa6, 0x8d, 0x9a, 0x64, 0x83, 0x9e, 0xef, 0x45, 0xde, 0x8f, 0xe0,
	0x13, 0x4f, 0x17, 0xa1, 0x70, 0x9e, 0xaa, 0x2b, 0x26, 0xa1, 0x23, 0x44, 0xc8, 0x37, 0xb0, 0x55,
	0x21, 0xa0, 0x1a, 0x41, 0xb5, 0xce, 0x8a, 0x24, 0x85, 0x7a, 0x60, 0x46, 0x34, 0x9b, 0xa5, 0x2c,
	0x11, 0x8c, 0xc7, 0xce, 0x57, 0x28, 0x54, 0x85, 0xc8, 0xb7, 0xb0, 0x55, 0x19, 0xa2, 0xd2, 0x33,
	0x75, 0x4d, 0x2a, 0xb0, 0x94, 0x7a, 0x03, 0x9d, 0x65, 0x92, 0x89, 0x94, 0x86, 0x8b, 0x60, 0xce,
	0xa7, 0x99, 0xf3, 0xbc, 0x57, 0xef, 0x1b, 0x9e, 0x55, 0x80, 0xc7, 0x7c, 0x9a, 0xa1, 0x1a, 0xbf,
	0x8d, 0xab, 0xb4, 0x17, 0x48, 0xeb, 0xae, 0x60, 0x24, 0xee, 0x82, 0xae, 0x5e, 0x4c, 0xe6, 0x7c,
	0xdd, 0xab, 0xf7, 0xcd, 0x3d, 0x7d, 0xa0, 0xde, 0x9e, 0x57, 0xe0, 0xd2, 0x84, 0x1b, 0x3e, 0x0b,
	0x71, 0x59, 0x2c, 0x72, 0x1c, 0x65, 0x42, 0x01, 0x8d, 0x23, 0x69, 0x42, 0x85, 0x80, 0x4b, 0x7f,
	0xa9, 0x4c, 0x58, 0x91, 0xe4, 0xca, 0x77, 0xc1, 0x2a, 0x79, 0xcb, 0x94, 0x39, 0xaf, 0x94, 0x0b,
	0x05, 0x76, 0x91, 0x32, 0xf9, 0x5a, 0xaa, 0x14, 0xd4, 0x7a, 0xad, 0x6c, 0xa8, 0xd0, 0x7c, 0x2a,
	0xdc, 0x3f, 0x34, 0xa8, 0x1f, 0xf3, 0x29, 0x79, 0x0e, 0x2d, 0x99, 0x24, 0xe5, 0x83, 0x6e, 0xce,
	0xf9, 0x74, 0x1c, 0x91, 0x97, 0x20, 0xf3, 0xa4, 0x9a, 0x0c, 0xfa, 0x9c, 0x4f, 0x4f, 0xff, 0xc7,
	0x9b, 0x7e, 0x01, 0x2d, 0x16, 0x27, 0x4b, 0x21, 0x13, 0x4d, 0x3a, 0x96, 0x8f, 0x88, 0x03, 0x3a,
	0x5f, 0x0a, 0x2c, 0x34, 0xb1, 0x50, 0x0c, 0xab, 0x1e, 0xb6, 0x1e, 0xf7, 0xd0, 0xfd, 0x5b, 0x03,
	0x73, 0x42, 0xe3, 0x30, 0x16, 0xbe, 0x08, 0x45, 0x26, 0x0f, 0xb1, 0x58, 0x83, 0x7a, 0xee, 0x9a,
	0x8a, 0xd7, 0x1c, 0x54, 0x2f, 0xfe, 0x41, 0x06, 0xd7, 0x1e, 0xcf, 0xe0, 0x55, 0x92, 0xd6, 0xd7,
	0x93, 0x54, 0xe5, 0x53, 0xc2, 0x53, 0x51, 0xc9, 0xa7, 0xa6, 0xcc, 0x27, 0x89, 0x95, 0x94, 0x3c,
	0x45, 0x15, 0xa5, 0xa9, 0x28, 0x0a, 0x53, 0x94, 0x1d, 0x30, 0x17, 0x3c, 0xa2, 0x37, 0x6b, 0x71,
	0x0d, 0x08, 0xa9, 0xc0, 0xfe, 0x19, 0x60, 0x98, 0xc9, 0x15, 0x85, 0xb3, 0x2b, 0x2a, 0x0d, 0x8c,
	0xd8, 0x25, 0xcd, 0x44, 0x7e, 0x22, 0xf9, 0x48, 0xe6, 0x2b, 0x26, 0x42, 0x0d, 0x13, 0x01, 0xbf,
	0x5d, 0x06, 0xc6, 0xc1, 0xd5, 0x32, 0xbe, 0x2e, 0xfa, 0xd2, 0x55, 0x98, 0x5d, 0x15, 0x7d, 0x49,
	0x7e, 0x3f, 0x1a, 0xca, 0x2f, 0x41, 0x46, 0x5a, 0x90, 0xb1, 0x7b, 0x8a, 0x3b, 0xae, 0x7b, 0x7a,
	0x1a, 0xde, 0xfa, 0xec, 0x1e, 0x4b, 0x32, 0xa0, 0xb0, 0xa4, 0xc2, 0x58, 0xbf, 0x67, 0x89, 0x2c,
	0xb9, 0x0c, 0xc0, 0xe7, 0xcb, 0x74, 0x46, 0x4f, 0x58, 0x7c, 0x4d, 0xb6, 0xa1, 0x81, 0xe9, 0xa3,
	0x61, 0xfa, 0x18, 0x03, 0x09, 0x4e, 0xee, 0x12, 0xea, 0x21, 0x8c, 0x21, 0x86, 0xe4, 0xb5, 0xde,
	0xa2, 0x20, 0xbc, 0x44, 0xaf, 0xc1, 0x58, 0x66, 0x34, 0x55, 0x65, 0x75, 0x87, 0xda, 0x12, 0xc0,
	0xc6, 0xf3, 0x8f, 0x06, 0xe6, 0x47, 0x9a, 0x66, 0x8c, 0xc7, 0xb8, 0x31, 0x07, 0xf4, 0xcf, 0x6a,
	0x98, 0x9f, 0x73, 0x31, 0x24, 0x2e, 0xb4, 0x66, 0x72, 0xff, 0x99, 0x53, 0xc3, 0x9b, 0x03, 0x83,
	0xd2, 0x0e, 0x2f, 0xaf, 0x90, 0x77, 0xa0, 0xab, 0x89, 0x33, 0xa7, 0x8e, 0x24, 0x73, 0xb0, 0xda,
	0x88, 0x57, 0xd4, 0xd6, 0x5c, 0x69, 0xfc, 0xb7, 0x2b, 0xcd, 0x35, 0x57, 0x36, 0x7a, 0x62, 0x6b,
	0xb3, 0x27, 0xee, 0x01, 0x14, 0xf1, 0xc3, 0x53, 0x6c, 0x38, 0xe6, 0x1e, 0x19, 0x14, 0x4d, 0xb5,
	0xac, 0x78, 0x15, 0x96, 0xfb, 0x13, 0x98, 0x47, 0x8c, 0xde, 0x44, 0xfe, 0xec, 0x8a, 0x2e, 0x1e,
	0xff, 0xb7, 0x41, 0xa0, 0x21, 0xee, 0x92, 0xf2, 0x54, 0xe5, 0xb7, 0xfb, 0xbb, 0x06, 0x4f, 0x1f,
	0x08, 0x97, 0x4c, 0x6d, 0xc5, 0x7c, 0xd0, 0x75, 0x6b, 0x0f, 0xbb, 0xee, 0x5b, 0x68, 0x65, 0x38,
	0x7d, 0x6e, 0x99, 0x35, 0xa8, 0x2c, 0xc9, 0xcb, 0x6b, 0xf2, 0x81, 0x15, 0xad, 0x45, 0xa4, 0x2c,
	0xbe, 0x44, 0xdf, 0x0c, 0xcf, 0xca, 0xbb, 0x0b, 0x62, 0xef, 0xbf, 0x40, 0xd3, 0x0b, 0xe3, 0x4b,
	0x4a, 0x00, 0x5a, 0x23, 0xfc, 0x5b, 0x65, 0x3f, 0x21, 0x16, 0xb4, 0xcf, 0xbd, 0xb3, 0xe3, 0xd1,
	0xc1, 0xc4, 0xb7, 0x35, 0x39, 0x3a, 0x1c, 0x4e, 0x86, 0xfe, 0x68, 0xe2, 0xdb, 0x35, 0xf2, 0x0c,
	0xec, 0xbc, 0x16, 0x94, 0x68, 0x9d, 0xb4, 0xa1, 0x71, 0x7c, 0xb6, 0xef, 0xdb, 0x0d, 0x62, 0x40,
	0xd3, 0x9f, 0x0c, 0x27, 0xbe, 0xdd, 0x24, 0x5b, 0x60, 0x0e, 0x7d, 0x7f, 0x34, 0x09, 0x0e, 0x86,
	0x07, 0xbf, 0x8c, 0xec, 0x16, 0x31, 0x41, 0x1f, 0xfd, 0x76, 0x3e, 0xf2, 0x26, 0xbe, 0xad, 0xbf,
	0x7f, 0x07, 0xed, 0xe2, 0x5a, 0x92, 0x0e, 0x18, 0x17, 0xa7, 0xe3, 0x8f, 0x23, 0xcf, 0x1f, 0x9e,
	0xd8, 0x4f, 0xe4, 0x5a, 0x2e, 0xce, 0x4f, 0xce, 0x86, 0x87, 0xb6, 0x36, 0x6d, 0xe1, 0xbf, 0xbc,
	0x1f, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x89, 0x20, 0x82, 0x88, 0xff, 0x09, 0x00, 0x00,
}
