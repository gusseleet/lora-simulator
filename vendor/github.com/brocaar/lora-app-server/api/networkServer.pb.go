// Code generated by protoc-gen-go. DO NOT EDIT.
// source: networkServer.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CreateNetworkServerRequest struct {
	// Name of the network-server.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// hostname:ip of the network-server.
	Server string `protobuf:"bytes,2,opt,name=server" json:"server,omitempty"`
	// ca certificate for connecting to the network-server
	CaCert string `protobuf:"bytes,3,opt,name=caCert" json:"caCert,omitempty"`
	// tls (client) certificate for connecting to the network-server
	TlsCert string `protobuf:"bytes,4,opt,name=tlsCert" json:"tlsCert,omitempty"`
	// tls (client) key for connecting to the network-server
	TlsKey string `protobuf:"bytes,5,opt,name=tlsKey" json:"tlsKey,omitempty"`
	// routing-profile ca certificate (used by the network-server to connect
	// back to the application-server)
	RoutingProfileCACert string `protobuf:"bytes,6,opt,name=routingProfileCACert" json:"routingProfileCACert,omitempty"`
	// routing-profile tls certificate (used by the network-server to connect
	// back to the application-server)
	RoutingProfileTLSCert string `protobuf:"bytes,7,opt,name=routingProfileTLSCert" json:"routingProfileTLSCert,omitempty"`
	// routing-profile tls key (used by the network-server to connect
	// back to the application-server)
	RoutingProfileTLSKey string `protobuf:"bytes,8,opt,name=routingProfileTLSKey" json:"routingProfileTLSKey,omitempty"`
}

func (m *CreateNetworkServerRequest) Reset()                    { *m = CreateNetworkServerRequest{} }
func (m *CreateNetworkServerRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateNetworkServerRequest) ProtoMessage()               {}
func (*CreateNetworkServerRequest) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

func (m *CreateNetworkServerRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateNetworkServerRequest) GetServer() string {
	if m != nil {
		return m.Server
	}
	return ""
}

func (m *CreateNetworkServerRequest) GetCaCert() string {
	if m != nil {
		return m.CaCert
	}
	return ""
}

func (m *CreateNetworkServerRequest) GetTlsCert() string {
	if m != nil {
		return m.TlsCert
	}
	return ""
}

func (m *CreateNetworkServerRequest) GetTlsKey() string {
	if m != nil {
		return m.TlsKey
	}
	return ""
}

func (m *CreateNetworkServerRequest) GetRoutingProfileCACert() string {
	if m != nil {
		return m.RoutingProfileCACert
	}
	return ""
}

func (m *CreateNetworkServerRequest) GetRoutingProfileTLSCert() string {
	if m != nil {
		return m.RoutingProfileTLSCert
	}
	return ""
}

func (m *CreateNetworkServerRequest) GetRoutingProfileTLSKey() string {
	if m != nil {
		return m.RoutingProfileTLSKey
	}
	return ""
}

type CreateNetworkServerResponse struct {
	// ID of the network-server.
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *CreateNetworkServerResponse) Reset()                    { *m = CreateNetworkServerResponse{} }
func (m *CreateNetworkServerResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateNetworkServerResponse) ProtoMessage()               {}
func (*CreateNetworkServerResponse) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{1} }

func (m *CreateNetworkServerResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetNetworkServerRequest struct {
	// ID of the network-server.
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetNetworkServerRequest) Reset()                    { *m = GetNetworkServerRequest{} }
func (m *GetNetworkServerRequest) String() string            { return proto.CompactTextString(m) }
func (*GetNetworkServerRequest) ProtoMessage()               {}
func (*GetNetworkServerRequest) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{2} }

func (m *GetNetworkServerRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetNetworkServerResponse struct {
	// ID of the network-server.
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// Timestamp when the record was created.
	CreatedAt string `protobuf:"bytes,2,opt,name=createdAt" json:"createdAt,omitempty"`
	// Timestamp when the record was last updated.
	UpdatedAt string `protobuf:"bytes,3,opt,name=updatedAt" json:"updatedAt,omitempty"`
	// Name of the network-server.
	Name string `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	// hostname:ip of the network-server.
	Server string `protobuf:"bytes,5,opt,name=server" json:"server,omitempty"`
	// ca certificate for connecting to the network-server
	CaCert string `protobuf:"bytes,6,opt,name=caCert" json:"caCert,omitempty"`
	// tls (client) certificate for connecting to the network-server
	TlsCert string `protobuf:"bytes,7,opt,name=tlsCert" json:"tlsCert,omitempty"`
	// routing-profile ca certificate (used by the network-server to connect
	// back to the application-server)
	RoutingProfileCACert string `protobuf:"bytes,8,opt,name=routingProfileCACert" json:"routingProfileCACert,omitempty"`
	// routing-profile tls certificate (used by the network-server to connect
	// back to the application-server)
	RoutingProfileTLSCert string `protobuf:"bytes,9,opt,name=routingProfileTLSCert" json:"routingProfileTLSCert,omitempty"`
}

func (m *GetNetworkServerResponse) Reset()                    { *m = GetNetworkServerResponse{} }
func (m *GetNetworkServerResponse) String() string            { return proto.CompactTextString(m) }
func (*GetNetworkServerResponse) ProtoMessage()               {}
func (*GetNetworkServerResponse) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{3} }

func (m *GetNetworkServerResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GetNetworkServerResponse) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *GetNetworkServerResponse) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *GetNetworkServerResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetNetworkServerResponse) GetServer() string {
	if m != nil {
		return m.Server
	}
	return ""
}

func (m *GetNetworkServerResponse) GetCaCert() string {
	if m != nil {
		return m.CaCert
	}
	return ""
}

func (m *GetNetworkServerResponse) GetTlsCert() string {
	if m != nil {
		return m.TlsCert
	}
	return ""
}

func (m *GetNetworkServerResponse) GetRoutingProfileCACert() string {
	if m != nil {
		return m.RoutingProfileCACert
	}
	return ""
}

func (m *GetNetworkServerResponse) GetRoutingProfileTLSCert() string {
	if m != nil {
		return m.RoutingProfileTLSCert
	}
	return ""
}

type UpdateNetworkServerRequest struct {
	// ID of the network-server.
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// Name of the network-server.
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// hostname:ip of the network-server.
	Server string `protobuf:"bytes,3,opt,name=server" json:"server,omitempty"`
	// ca certificate for connecting to the network-server
	CaCert string `protobuf:"bytes,4,opt,name=caCert" json:"caCert,omitempty"`
	// tls (client) certificate for connecting to the network-server
	TlsCert string `protobuf:"bytes,5,opt,name=tlsCert" json:"tlsCert,omitempty"`
	// tls (client) key for connecting to the network-server
	TlsKey string `protobuf:"bytes,6,opt,name=tlsKey" json:"tlsKey,omitempty"`
	// routing-profile ca certificate (used by the network-server to connect
	// back to the application-server)
	RoutingProfileCACert string `protobuf:"bytes,7,opt,name=routingProfileCACert" json:"routingProfileCACert,omitempty"`
	// routing-profile tls certificate (used by the network-server to connect
	// back to the application-server)
	RoutingProfileTLSCert string `protobuf:"bytes,8,opt,name=routingProfileTLSCert" json:"routingProfileTLSCert,omitempty"`
	// routing-profile tls key (used by the network-server to connect
	// back to the application-server)
	RoutingProfileTLSKey string `protobuf:"bytes,9,opt,name=routingProfileTLSKey" json:"routingProfileTLSKey,omitempty"`
}

func (m *UpdateNetworkServerRequest) Reset()                    { *m = UpdateNetworkServerRequest{} }
func (m *UpdateNetworkServerRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateNetworkServerRequest) ProtoMessage()               {}
func (*UpdateNetworkServerRequest) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{4} }

func (m *UpdateNetworkServerRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateNetworkServerRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateNetworkServerRequest) GetServer() string {
	if m != nil {
		return m.Server
	}
	return ""
}

func (m *UpdateNetworkServerRequest) GetCaCert() string {
	if m != nil {
		return m.CaCert
	}
	return ""
}

func (m *UpdateNetworkServerRequest) GetTlsCert() string {
	if m != nil {
		return m.TlsCert
	}
	return ""
}

func (m *UpdateNetworkServerRequest) GetTlsKey() string {
	if m != nil {
		return m.TlsKey
	}
	return ""
}

func (m *UpdateNetworkServerRequest) GetRoutingProfileCACert() string {
	if m != nil {
		return m.RoutingProfileCACert
	}
	return ""
}

func (m *UpdateNetworkServerRequest) GetRoutingProfileTLSCert() string {
	if m != nil {
		return m.RoutingProfileTLSCert
	}
	return ""
}

func (m *UpdateNetworkServerRequest) GetRoutingProfileTLSKey() string {
	if m != nil {
		return m.RoutingProfileTLSKey
	}
	return ""
}

type UpdateNetworkServerResponse struct {
}

func (m *UpdateNetworkServerResponse) Reset()                    { *m = UpdateNetworkServerResponse{} }
func (m *UpdateNetworkServerResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateNetworkServerResponse) ProtoMessage()               {}
func (*UpdateNetworkServerResponse) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{5} }

type DeleteNetworkServerRequest struct {
	// ID of the network-server.
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteNetworkServerRequest) Reset()                    { *m = DeleteNetworkServerRequest{} }
func (m *DeleteNetworkServerRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteNetworkServerRequest) ProtoMessage()               {}
func (*DeleteNetworkServerRequest) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{6} }

func (m *DeleteNetworkServerRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeleteNetworkServerResponse struct {
}

func (m *DeleteNetworkServerResponse) Reset()                    { *m = DeleteNetworkServerResponse{} }
func (m *DeleteNetworkServerResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteNetworkServerResponse) ProtoMessage()               {}
func (*DeleteNetworkServerResponse) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{7} }

type ListNetworkServerRequest struct {
	// Max number of items to return.
	Limit int64 `protobuf:"varint,1,opt,name=limit" json:"limit,omitempty"`
	// Offset in the result-set (for pagination).
	Offset int64 `protobuf:"varint,2,opt,name=offset" json:"offset,omitempty"`
	// Organization id to filter on.
	OrganizationID int64 `protobuf:"varint,3,opt,name=organizationID" json:"organizationID,omitempty"`
}

func (m *ListNetworkServerRequest) Reset()                    { *m = ListNetworkServerRequest{} }
func (m *ListNetworkServerRequest) String() string            { return proto.CompactTextString(m) }
func (*ListNetworkServerRequest) ProtoMessage()               {}
func (*ListNetworkServerRequest) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{8} }

func (m *ListNetworkServerRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListNetworkServerRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ListNetworkServerRequest) GetOrganizationID() int64 {
	if m != nil {
		return m.OrganizationID
	}
	return 0
}

type ListNetworkServerResponse struct {
	// Total number of network-servers.
	TotalCount int64 `protobuf:"varint,1,opt,name=totalCount" json:"totalCount,omitempty"`
	// Network-servers within the result-set.
	Result []*GetNetworkServerResponse `protobuf:"bytes,2,rep,name=result" json:"result,omitempty"`
}

func (m *ListNetworkServerResponse) Reset()                    { *m = ListNetworkServerResponse{} }
func (m *ListNetworkServerResponse) String() string            { return proto.CompactTextString(m) }
func (*ListNetworkServerResponse) ProtoMessage()               {}
func (*ListNetworkServerResponse) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{9} }

func (m *ListNetworkServerResponse) GetTotalCount() int64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *ListNetworkServerResponse) GetResult() []*GetNetworkServerResponse {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateNetworkServerRequest)(nil), "api.CreateNetworkServerRequest")
	proto.RegisterType((*CreateNetworkServerResponse)(nil), "api.CreateNetworkServerResponse")
	proto.RegisterType((*GetNetworkServerRequest)(nil), "api.GetNetworkServerRequest")
	proto.RegisterType((*GetNetworkServerResponse)(nil), "api.GetNetworkServerResponse")
	proto.RegisterType((*UpdateNetworkServerRequest)(nil), "api.UpdateNetworkServerRequest")
	proto.RegisterType((*UpdateNetworkServerResponse)(nil), "api.UpdateNetworkServerResponse")
	proto.RegisterType((*DeleteNetworkServerRequest)(nil), "api.DeleteNetworkServerRequest")
	proto.RegisterType((*DeleteNetworkServerResponse)(nil), "api.DeleteNetworkServerResponse")
	proto.RegisterType((*ListNetworkServerRequest)(nil), "api.ListNetworkServerRequest")
	proto.RegisterType((*ListNetworkServerResponse)(nil), "api.ListNetworkServerResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for NetworkServer service

type NetworkServerClient interface {
	// Create creates the given network-server.
	Create(ctx context.Context, in *CreateNetworkServerRequest, opts ...grpc.CallOption) (*CreateNetworkServerResponse, error)
	// Get returns the network-server matching the given id.
	Get(ctx context.Context, in *GetNetworkServerRequest, opts ...grpc.CallOption) (*GetNetworkServerResponse, error)
	// Update updates the given network-server.
	Update(ctx context.Context, in *UpdateNetworkServerRequest, opts ...grpc.CallOption) (*UpdateNetworkServerResponse, error)
	// Delete deletes the network-server matching the given id.
	Delete(ctx context.Context, in *DeleteNetworkServerRequest, opts ...grpc.CallOption) (*DeleteNetworkServerResponse, error)
	// List lists the available network-servers.
	List(ctx context.Context, in *ListNetworkServerRequest, opts ...grpc.CallOption) (*ListNetworkServerResponse, error)
}

type networkServerClient struct {
	cc *grpc.ClientConn
}

func NewNetworkServerClient(cc *grpc.ClientConn) NetworkServerClient {
	return &networkServerClient{cc}
}

func (c *networkServerClient) Create(ctx context.Context, in *CreateNetworkServerRequest, opts ...grpc.CallOption) (*CreateNetworkServerResponse, error) {
	out := new(CreateNetworkServerResponse)
	err := grpc.Invoke(ctx, "/api.NetworkServer/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServerClient) Get(ctx context.Context, in *GetNetworkServerRequest, opts ...grpc.CallOption) (*GetNetworkServerResponse, error) {
	out := new(GetNetworkServerResponse)
	err := grpc.Invoke(ctx, "/api.NetworkServer/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServerClient) Update(ctx context.Context, in *UpdateNetworkServerRequest, opts ...grpc.CallOption) (*UpdateNetworkServerResponse, error) {
	out := new(UpdateNetworkServerResponse)
	err := grpc.Invoke(ctx, "/api.NetworkServer/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServerClient) Delete(ctx context.Context, in *DeleteNetworkServerRequest, opts ...grpc.CallOption) (*DeleteNetworkServerResponse, error) {
	out := new(DeleteNetworkServerResponse)
	err := grpc.Invoke(ctx, "/api.NetworkServer/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServerClient) List(ctx context.Context, in *ListNetworkServerRequest, opts ...grpc.CallOption) (*ListNetworkServerResponse, error) {
	out := new(ListNetworkServerResponse)
	err := grpc.Invoke(ctx, "/api.NetworkServer/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NetworkServer service

type NetworkServerServer interface {
	// Create creates the given network-server.
	Create(context.Context, *CreateNetworkServerRequest) (*CreateNetworkServerResponse, error)
	// Get returns the network-server matching the given id.
	Get(context.Context, *GetNetworkServerRequest) (*GetNetworkServerResponse, error)
	// Update updates the given network-server.
	Update(context.Context, *UpdateNetworkServerRequest) (*UpdateNetworkServerResponse, error)
	// Delete deletes the network-server matching the given id.
	Delete(context.Context, *DeleteNetworkServerRequest) (*DeleteNetworkServerResponse, error)
	// List lists the available network-servers.
	List(context.Context, *ListNetworkServerRequest) (*ListNetworkServerResponse, error)
}

func RegisterNetworkServerServer(s *grpc.Server, srv NetworkServerServer) {
	s.RegisterService(&_NetworkServer_serviceDesc, srv)
}

func _NetworkServer_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNetworkServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.NetworkServer/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServerServer).Create(ctx, req.(*CreateNetworkServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkServer_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNetworkServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.NetworkServer/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServerServer).Get(ctx, req.(*GetNetworkServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkServer_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNetworkServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.NetworkServer/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServerServer).Update(ctx, req.(*UpdateNetworkServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkServer_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNetworkServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.NetworkServer/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServerServer).Delete(ctx, req.(*DeleteNetworkServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkServer_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNetworkServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServerServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.NetworkServer/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServerServer).List(ctx, req.(*ListNetworkServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NetworkServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.NetworkServer",
	HandlerType: (*NetworkServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _NetworkServer_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _NetworkServer_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _NetworkServer_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _NetworkServer_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _NetworkServer_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "networkServer.proto",
}

func init() { proto.RegisterFile("networkServer.proto", fileDescriptor8) }

var fileDescriptor8 = []byte{
	// 612 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x56, 0x93, 0x34, 0x6d, 0x0f, 0x62, 0x17, 0xa6, 0x8c, 0x34, 0xed, 0xba, 0x11, 0x21, 0x34,
	0x26, 0xd6, 0x4a, 0x05, 0x6e, 0xb8, 0x9b, 0x3a, 0x69, 0x42, 0x54, 0x08, 0xa5, 0xe3, 0x01, 0x42,
	0xeb, 0x56, 0x16, 0x59, 0x9c, 0xd9, 0x0e, 0x08, 0x10, 0x37, 0xbc, 0x02, 0xaf, 0x81, 0x10, 0x8f,
	0xc0, 0x3b, 0xf0, 0x0a, 0x3c, 0x08, 0x8a, 0xed, 0xfd, 0x74, 0xb3, 0xa3, 0xaa, 0x77, 0x3d, 0x3f,
	0x3e, 0x5f, 0xbe, 0xf3, 0x7d, 0x76, 0xe1, 0x5e, 0x86, 0xc5, 0x27, 0xca, 0x3e, 0x4c, 0x31, 0xfb,
	0x88, 0xd9, 0x20, 0x67, 0x54, 0x50, 0xe4, 0x26, 0x39, 0x09, 0x7b, 0x4b, 0x4a, 0x97, 0x29, 0x1e,
	0x26, 0x39, 0x19, 0x26, 0x59, 0x46, 0x45, 0x22, 0x08, 0xcd, 0xb8, 0x6a, 0x89, 0x7e, 0x3b, 0x10,
	0x8e, 0x19, 0x4e, 0x04, 0x7e, 0x73, 0x7d, 0x40, 0x8c, 0xcf, 0x0b, 0xcc, 0x05, 0x42, 0xe0, 0x65,
	0xc9, 0x19, 0x0e, 0x6a, 0x7b, 0xb5, 0xfd, 0x56, 0x2c, 0x7f, 0xa3, 0x6d, 0xf0, 0xb9, 0x6c, 0x0a,
	0x1c, 0x99, 0xd5, 0x51, 0x99, 0x9f, 0x25, 0x63, 0xcc, 0x44, 0xe0, 0xaa, 0xbc, 0x8a, 0x50, 0x00,
	0x0d, 0x91, 0x72, 0x59, 0xf0, 0x64, 0xe1, 0x22, 0x2c, 0x4f, 0x88, 0x94, 0xbf, 0xc6, 0x9f, 0x83,
	0xba, 0x3a, 0xa1, 0x22, 0x34, 0x82, 0x36, 0xa3, 0x85, 0x20, 0xd9, 0xf2, 0x2d, 0xa3, 0x0b, 0x92,
	0xe2, 0xf1, 0x91, 0x3c, 0xee, 0xcb, 0x2e, 0x63, 0x0d, 0x3d, 0x87, 0xfb, 0xab, 0xf9, 0xd3, 0xc9,
	0x54, 0x1e, 0x6a, 0xc8, 0x43, 0xe6, 0xe2, 0x6d, 0xa4, 0xd3, 0xc9, 0xb4, 0xfc, 0x9e, 0xa6, 0x09,
	0x49, 0xd5, 0xa2, 0x43, 0xe8, 0x1a, 0x37, 0xc6, 0x73, 0x9a, 0x71, 0x8c, 0xb6, 0xc0, 0x21, 0x73,
	0xb9, 0x30, 0x37, 0x76, 0xc8, 0x3c, 0x7a, 0x02, 0x0f, 0x4e, 0xb0, 0x30, 0x6e, 0xf7, 0x66, 0xeb,
	0x2f, 0x07, 0x82, 0xdb, 0xbd, 0xe6, 0xb9, 0xa8, 0x07, 0xad, 0x99, 0xfc, 0x8c, 0xf9, 0x91, 0xd0,
	0x4a, 0x5c, 0x25, 0xca, 0x6a, 0x91, 0xcf, 0x75, 0x55, 0xe9, 0x71, 0x95, 0xb8, 0x94, 0xd5, 0x33,
	0xca, 0x5a, 0xb7, 0xc8, 0xea, 0xdb, 0x64, 0x6d, 0xac, 0xca, 0x6a, 0x93, 0xaf, 0xb9, 0x89, 0x7c,
	0xad, 0x0a, 0xf9, 0xa2, 0x3f, 0x0e, 0x84, 0xef, 0x24, 0xab, 0x75, 0xf6, 0x7b, 0x49, 0xdb, 0x31,
	0xd2, 0x76, 0x2d, 0xb4, 0x3d, 0x1b, 0xed, 0xba, 0xcd, 0xcd, 0xfe, 0x5a, 0x6e, 0x6e, 0x6c, 0xb2,
	0x8e, 0xe6, 0x26, 0x6e, 0x6e, 0x55, 0xb8, 0x79, 0x07, 0xba, 0xc6, 0x0d, 0x2a, 0xd7, 0x45, 0x4f,
	0x21, 0x3c, 0xc6, 0x29, 0x5e, 0x6f, 0xc1, 0xe5, 0x30, 0x63, 0xb7, 0x1e, 0x96, 0x43, 0x30, 0x21,
	0xdc, 0x7c, 0x17, 0xda, 0x50, 0x4f, 0xc9, 0x19, 0x11, 0x7a, 0x9a, 0x0a, 0xca, 0x9d, 0xd2, 0xc5,
	0x82, 0x63, 0xe5, 0x70, 0x37, 0xd6, 0x11, 0x7a, 0x0c, 0x5b, 0x94, 0x2d, 0x93, 0x8c, 0x7c, 0x91,
	0xaf, 0xd9, 0xab, 0x63, 0xa9, 0x9e, 0x1b, 0xdf, 0xc8, 0x46, 0x0c, 0x3a, 0x06, 0x44, 0x7d, 0xa3,
	0xfa, 0x00, 0x82, 0x8a, 0x24, 0x1d, 0xd3, 0x22, 0xbb, 0xc0, 0xbd, 0x96, 0x41, 0x2f, 0xc0, 0x67,
	0x98, 0x17, 0x69, 0x09, 0xee, 0xee, 0xdf, 0x19, 0xed, 0x0c, 0x92, 0x9c, 0x0c, 0x6c, 0x17, 0x34,
	0xd6, 0xcd, 0xa3, 0x9f, 0x1e, 0xdc, 0x5d, 0xe9, 0x40, 0x29, 0xf8, 0xea, 0xc5, 0x40, 0xbb, 0x72,
	0x84, 0xfd, 0xc1, 0x0d, 0xf7, 0xec, 0x0d, 0x7a, 0x89, 0xbb, 0xdf, 0xff, 0xfe, 0xfb, 0xe1, 0x74,
	0xa2, 0xb6, 0x7c, 0xd1, 0xf5, 0xb3, 0x7f, 0xa8, 0x5c, 0xcb, 0x5f, 0xd6, 0x0e, 0x10, 0x06, 0xf7,
	0x04, 0x0b, 0xd4, 0xb3, 0x7c, 0xad, 0xc2, 0xa9, 0xe6, 0x12, 0x3d, 0x94, 0x20, 0x5d, 0xd4, 0x31,
	0x81, 0x0c, 0xbf, 0x92, 0xf9, 0x37, 0x74, 0x0e, 0xbe, 0x32, 0x8e, 0x26, 0x65, 0xbf, 0x87, 0x9a,
	0x54, 0x95, 0xcd, 0x1e, 0x49, 0xbc, 0x7e, 0x68, 0xc7, 0x2b, 0x99, 0x65, 0xe0, 0x2b, 0x7b, 0x69,
	0x48, 0xbb, 0x33, 0x35, 0x64, 0x95, 0x19, 0x35, 0xc5, 0x83, 0x0a, 0x8a, 0x33, 0xf0, 0x4a, 0xf7,
	0x20, 0xb5, 0x2c, 0x9b, 0x75, 0xc3, 0xbe, 0xad, 0xac, 0x91, 0x7a, 0x12, 0x69, 0x1b, 0x19, 0x15,
	0x7b, 0xef, 0xcb, 0x3f, 0xe2, 0x67, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0x96, 0x0d, 0x92, 0x0f,
	0xc2, 0x07, 0x00, 0x00,
}