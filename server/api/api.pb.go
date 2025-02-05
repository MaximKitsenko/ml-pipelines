// Code generated by protoc-gen-go.
// source: api.proto
// DO NOT EDIT!

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	CreateProjectRequest
	CreateProjectResponse
	ApplyRequest
	ApplyResponse
	CreateDataset
	DatasetRequestMeta
	OkResponse
	ScenarioRequest
	Error
	ScenarioResponse
	KillRequest
	PingRequest
	WipeDatabase
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import events "mlp/catalog/events"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateProjectRequest struct {
	ProjectId   string `protobuf:"bytes,1,opt,name=ProjectId" json:"ProjectId,omitempty"`
	ProjectName string `protobuf:"bytes,2,opt,name=ProjectName" json:"ProjectName,omitempty"`
}

func (m *CreateProjectRequest) Reset()                    { *m = CreateProjectRequest{} }
func (m *CreateProjectRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateProjectRequest) ProtoMessage()               {}
func (*CreateProjectRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type CreateProjectResponse struct {
}

func (m *CreateProjectResponse) Reset()                    { *m = CreateProjectResponse{} }
func (m *CreateProjectResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateProjectResponse) ProtoMessage()               {}
func (*CreateProjectResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ApplyRequest struct {
	Events []*events.Event `protobuf:"bytes,1,rep,name=Events" json:"Events,omitempty"`
}

func (m *ApplyRequest) Reset()                    { *m = ApplyRequest{} }
func (m *ApplyRequest) String() string            { return proto.CompactTextString(m) }
func (*ApplyRequest) ProtoMessage()               {}
func (*ApplyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ApplyRequest) GetEvents() []*events.Event {
	if m != nil {
		return m.Events
	}
	return nil
}

type ApplyResponse struct {
	Version uint64 `protobuf:"varint,1,opt,name=Version" json:"Version,omitempty"`
}

func (m *ApplyResponse) Reset()                    { *m = ApplyResponse{} }
func (m *ApplyResponse) String() string            { return proto.CompactTextString(m) }
func (*ApplyResponse) ProtoMessage()               {}
func (*ApplyResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type CreateDataset struct {
	DatasetId string `protobuf:"bytes,1,opt,name=dataset_id,json=datasetId" json:"dataset_id,omitempty"`
	ProjectId string `protobuf:"bytes,2,opt,name=project_id,json=projectId" json:"project_id,omitempty"`
}

func (m *CreateDataset) Reset()                    { *m = CreateDataset{} }
func (m *CreateDataset) String() string            { return proto.CompactTextString(m) }
func (*CreateDataset) ProtoMessage()               {}
func (*CreateDataset) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type DatasetRequestMeta struct {
}

func (m *DatasetRequestMeta) Reset()                    { *m = DatasetRequestMeta{} }
func (m *DatasetRequestMeta) String() string            { return proto.CompactTextString(m) }
func (*DatasetRequestMeta) ProtoMessage()               {}
func (*DatasetRequestMeta) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type OkResponse struct {
}

func (m *OkResponse) Reset()                    { *m = OkResponse{} }
func (m *OkResponse) String() string            { return proto.CompactTextString(m) }
func (*OkResponse) ProtoMessage()               {}
func (*OkResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type ScenarioRequest struct {
	Name      string          `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	Events    []*events.Event `protobuf:"bytes,2,rep,name=Events" json:"Events,omitempty"`
	Timestamp int64           `protobuf:"varint,3,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *ScenarioRequest) Reset()                    { *m = ScenarioRequest{} }
func (m *ScenarioRequest) String() string            { return proto.CompactTextString(m) }
func (*ScenarioRequest) ProtoMessage()               {}
func (*ScenarioRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ScenarioRequest) GetEvents() []*events.Event {
	if m != nil {
		return m.Events
	}
	return nil
}

type Error struct {
	Message string `protobuf:"bytes,1,opt,name=Message" json:"Message,omitempty"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type ScenarioResponse struct {
	Error *Error `protobuf:"bytes,1,opt,name=Error" json:"Error,omitempty"`
}

func (m *ScenarioResponse) Reset()                    { *m = ScenarioResponse{} }
func (m *ScenarioResponse) String() string            { return proto.CompactTextString(m) }
func (*ScenarioResponse) ProtoMessage()               {}
func (*ScenarioResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ScenarioResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type KillRequest struct {
}

func (m *KillRequest) Reset()                    { *m = KillRequest{} }
func (m *KillRequest) String() string            { return proto.CompactTextString(m) }
func (*KillRequest) ProtoMessage()               {}
func (*KillRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type PingRequest struct {
}

func (m *PingRequest) Reset()                    { *m = PingRequest{} }
func (m *PingRequest) String() string            { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()               {}
func (*PingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

type WipeDatabase struct {
}

func (m *WipeDatabase) Reset()                    { *m = WipeDatabase{} }
func (m *WipeDatabase) String() string            { return proto.CompactTextString(m) }
func (*WipeDatabase) ProtoMessage()               {}
func (*WipeDatabase) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func init() {
	proto.RegisterType((*CreateProjectRequest)(nil), "CreateProjectRequest")
	proto.RegisterType((*CreateProjectResponse)(nil), "CreateProjectResponse")
	proto.RegisterType((*ApplyRequest)(nil), "ApplyRequest")
	proto.RegisterType((*ApplyResponse)(nil), "ApplyResponse")
	proto.RegisterType((*CreateDataset)(nil), "CreateDataset")
	proto.RegisterType((*DatasetRequestMeta)(nil), "DatasetRequestMeta")
	proto.RegisterType((*OkResponse)(nil), "OkResponse")
	proto.RegisterType((*ScenarioRequest)(nil), "ScenarioRequest")
	proto.RegisterType((*Error)(nil), "Error")
	proto.RegisterType((*ScenarioResponse)(nil), "ScenarioResponse")
	proto.RegisterType((*KillRequest)(nil), "KillRequest")
	proto.RegisterType((*PingRequest)(nil), "PingRequest")
	proto.RegisterType((*WipeDatabase)(nil), "WipeDatabase")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Catalog service

type CatalogClient interface {
	CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...grpc.CallOption) (*CreateProjectResponse, error)
	Apply(ctx context.Context, in *ApplyRequest, opts ...grpc.CallOption) (*ApplyResponse, error)
}

type catalogClient struct {
	cc *grpc.ClientConn
}

func NewCatalogClient(cc *grpc.ClientConn) CatalogClient {
	return &catalogClient{cc}
}

func (c *catalogClient) CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...grpc.CallOption) (*CreateProjectResponse, error) {
	out := new(CreateProjectResponse)
	err := grpc.Invoke(ctx, "/Catalog/CreateProject", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogClient) Apply(ctx context.Context, in *ApplyRequest, opts ...grpc.CallOption) (*ApplyResponse, error) {
	out := new(ApplyResponse)
	err := grpc.Invoke(ctx, "/Catalog/Apply", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Catalog service

type CatalogServer interface {
	CreateProject(context.Context, *CreateProjectRequest) (*CreateProjectResponse, error)
	Apply(context.Context, *ApplyRequest) (*ApplyResponse, error)
}

func RegisterCatalogServer(s *grpc.Server, srv CatalogServer) {
	s.RegisterService(&_Catalog_serviceDesc, srv)
}

func _Catalog_CreateProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServer).CreateProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Catalog/CreateProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServer).CreateProject(ctx, req.(*CreateProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catalog_Apply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServer).Apply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Catalog/Apply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServer).Apply(ctx, req.(*ApplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Catalog_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Catalog",
	HandlerType: (*CatalogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProject",
			Handler:    _Catalog_CreateProject_Handler,
		},
		{
			MethodName: "Apply",
			Handler:    _Catalog_Apply_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

// Client API for Test service

type TestClient interface {
	// Setup a given state in the database
	Setup(ctx context.Context, in *ScenarioRequest, opts ...grpc.CallOption) (*OkResponse, error)
	Kill(ctx context.Context, in *KillRequest, opts ...grpc.CallOption) (*OkResponse, error)
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*OkResponse, error)
	Wipe(ctx context.Context, in *WipeDatabase, opts ...grpc.CallOption) (*OkResponse, error)
}

type testClient struct {
	cc *grpc.ClientConn
}

func NewTestClient(cc *grpc.ClientConn) TestClient {
	return &testClient{cc}
}

func (c *testClient) Setup(ctx context.Context, in *ScenarioRequest, opts ...grpc.CallOption) (*OkResponse, error) {
	out := new(OkResponse)
	err := grpc.Invoke(ctx, "/Test/Setup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testClient) Kill(ctx context.Context, in *KillRequest, opts ...grpc.CallOption) (*OkResponse, error) {
	out := new(OkResponse)
	err := grpc.Invoke(ctx, "/Test/Kill", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*OkResponse, error) {
	out := new(OkResponse)
	err := grpc.Invoke(ctx, "/Test/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testClient) Wipe(ctx context.Context, in *WipeDatabase, opts ...grpc.CallOption) (*OkResponse, error) {
	out := new(OkResponse)
	err := grpc.Invoke(ctx, "/Test/Wipe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Test service

type TestServer interface {
	// Setup a given state in the database
	Setup(context.Context, *ScenarioRequest) (*OkResponse, error)
	Kill(context.Context, *KillRequest) (*OkResponse, error)
	Ping(context.Context, *PingRequest) (*OkResponse, error)
	Wipe(context.Context, *WipeDatabase) (*OkResponse, error)
}

func RegisterTestServer(s *grpc.Server, srv TestServer) {
	s.RegisterService(&_Test_serviceDesc, srv)
}

func _Test_Setup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScenarioRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServer).Setup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Test/Setup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServer).Setup(ctx, req.(*ScenarioRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Test_Kill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServer).Kill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Test/Kill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServer).Kill(ctx, req.(*KillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Test_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Test/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Test_Wipe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WipeDatabase)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServer).Wipe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Test/Wipe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServer).Wipe(ctx, req.(*WipeDatabase))
	}
	return interceptor(ctx, in, info, handler)
}

var _Test_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Test",
	HandlerType: (*TestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Setup",
			Handler:    _Test_Setup_Handler,
		},
		{
			MethodName: "Kill",
			Handler:    _Test_Kill_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Test_Ping_Handler,
		},
		{
			MethodName: "Wipe",
			Handler:    _Test_Wipe_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 433 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x53, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0x6d, 0xd6, 0xb4, 0x53, 0x6f, 0x93, 0x31, 0x59, 0x1b, 0x44, 0xd5, 0x40, 0xc5, 0xf0, 0x10,
	0x5e, 0x2c, 0x54, 0xfe, 0x00, 0x68, 0xec, 0x61, 0x42, 0x85, 0x29, 0x43, 0xe3, 0x11, 0x79, 0xed,
	0x55, 0x65, 0x68, 0x63, 0x63, 0xbb, 0x48, 0xfc, 0x1b, 0x7e, 0x2a, 0xf2, 0x57, 0x1b, 0xaa, 0xbe,
	0xf9, 0x1c, 0xdf, 0xdc, 0x7b, 0xce, 0xb9, 0x0e, 0x8c, 0xb8, 0x12, 0x4c, 0x69, 0x69, 0xe5, 0xa4,
	0xc0, 0xdf, 0xd8, 0x5a, 0x13, 0x10, 0x7d, 0x80, 0x8b, 0x6b, 0x8d, 0xdc, 0xe2, 0x9d, 0x96, 0x3f,
	0x70, 0x61, 0x1b, 0xfc, 0xb5, 0x45, 0x63, 0xc9, 0x15, 0x8c, 0x22, 0x73, 0xbb, 0xac, 0xb2, 0x69,
	0x56, 0x8f, 0x9a, 0x3d, 0x41, 0xa6, 0x30, 0x8e, 0xe0, 0x33, 0xdf, 0x60, 0x75, 0xe2, 0xef, 0xbb,
	0x14, 0x7d, 0x06, 0x97, 0x07, 0x7d, 0x8d, 0x92, 0xad, 0x41, 0xca, 0xa0, 0xf8, 0xa0, 0xd4, 0xfa,
	0x4f, 0x1a, 0xf4, 0x02, 0x86, 0x37, 0x5e, 0x50, 0x95, 0x4d, 0xfb, 0xf5, 0x78, 0x36, 0x64, 0x1e,
	0x36, 0x91, 0xa5, 0x6f, 0xa0, 0x8c, 0xf5, 0xa1, 0x01, 0xa9, 0xe0, 0xf4, 0x01, 0xb5, 0x11, 0xb2,
	0xf5, 0xba, 0xf2, 0x26, 0x41, 0x3a, 0x87, 0x32, 0xcc, 0xfc, 0xc8, 0x2d, 0x37, 0x68, 0xc9, 0x73,
	0x80, 0x65, 0x38, 0x7e, 0x17, 0x3b, 0x17, 0x91, 0xb9, 0x5d, 0xba, 0x6b, 0x15, 0xd4, 0xb9, 0xeb,
	0x60, 0x62, 0xa4, 0x92, 0x49, 0x7a, 0x01, 0x24, 0x36, 0x8a, 0x5a, 0xe7, 0x68, 0x39, 0x2d, 0x00,
	0xbe, 0xfc, 0xdc, 0xb9, 0x59, 0xc0, 0x93, 0xfb, 0x05, 0xb6, 0x5c, 0x0b, 0x99, 0x0c, 0x11, 0xc8,
	0x7d, 0x28, 0x61, 0x9c, 0x3f, 0x77, 0x4c, 0x9e, 0x1c, 0x33, 0xe9, 0xd2, 0xb6, 0x62, 0x83, 0xc6,
	0xf2, 0x8d, 0xaa, 0xfa, 0xd3, 0xac, 0xee, 0x37, 0x7b, 0x82, 0xbe, 0x84, 0xc1, 0x8d, 0xd6, 0x52,
	0x3b, 0xeb, 0x73, 0x34, 0x86, 0xaf, 0x52, 0xf7, 0x04, 0xe9, 0x5b, 0x38, 0xdf, 0xeb, 0x88, 0x41,
	0x5d, 0xc5, 0xcf, 0x7c, 0xad, 0x9f, 0xe9, 0x50, 0x13, 0x48, 0x5a, 0xc2, 0xf8, 0x93, 0x58, 0xaf,
	0xa3, 0x6a, 0x07, 0xef, 0x44, 0xbb, 0x4a, 0xf0, 0x0c, 0x8a, 0x6f, 0x42, 0xf9, 0x20, 0x1f, 0xb9,
	0xc1, 0xd9, 0x16, 0x4e, 0xaf, 0xb9, 0xe5, 0x6b, 0xb9, 0x22, 0xef, 0x53, 0xca, 0x71, 0xb3, 0xe4,
	0x92, 0x1d, 0x7b, 0x41, 0x93, 0xa7, 0xec, 0xf8, 0x03, 0xe8, 0x91, 0x1a, 0x06, 0x7e, 0xa5, 0xa4,
	0x64, 0xdd, 0xa7, 0x30, 0x39, 0x63, 0xff, 0x6d, 0x9a, 0xf6, 0x66, 0x7f, 0x33, 0xc8, 0xbf, 0xba,
	0x50, 0x6b, 0x18, 0xdc, 0xa3, 0xdd, 0x2a, 0x72, 0xce, 0x0e, 0xf2, 0x9e, 0x8c, 0x59, 0x67, 0x1f,
	0x3d, 0xf2, 0x0a, 0x72, 0xe7, 0x8b, 0x14, 0xac, 0x63, 0xef, 0x48, 0x91, 0x73, 0x4b, 0x0a, 0xd6,
	0x31, 0x7d, 0x58, 0xf4, 0x1a, 0x72, 0x97, 0x01, 0x29, 0x59, 0x37, 0x8a, 0x83, 0xaa, 0xc7, 0xa1,
	0xff, 0x8f, 0xde, 0xfd, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x8f, 0xbc, 0x26, 0x62, 0x03, 0x00,
	0x00,
}
