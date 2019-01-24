// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vault/request_forwarding_service.proto

package vault

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	forwarding "github.com/hashicorp/vault/helper/forwarding"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	math "math"
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

type EchoRequest struct {
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	// ClusterAddr is used to send up a standby node's address to the active
	// node upon heartbeat
	ClusterAddr string `protobuf:"bytes,2,opt,name=cluster_addr,json=clusterAddr,proto3" json:"cluster_addr,omitempty"`
	// ClusterAddrs is used to send up a list of cluster addresses to a dr
	// primary from a dr secondary
	ClusterAddrs         []string `protobuf:"bytes,3,rep,name=cluster_addrs,json=clusterAddrs,proto3" json:"cluster_addrs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoRequest) Reset()         { *m = EchoRequest{} }
func (m *EchoRequest) String() string { return proto.CompactTextString(m) }
func (*EchoRequest) ProtoMessage()    {}
func (*EchoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5f7512e4ab7b58a, []int{0}
}

func (m *EchoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoRequest.Unmarshal(m, b)
}
func (m *EchoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoRequest.Marshal(b, m, deterministic)
}
func (m *EchoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoRequest.Merge(m, src)
}
func (m *EchoRequest) XXX_Size() int {
	return xxx_messageInfo_EchoRequest.Size(m)
}
func (m *EchoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EchoRequest proto.InternalMessageInfo

func (m *EchoRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *EchoRequest) GetClusterAddr() string {
	if m != nil {
		return m.ClusterAddr
	}
	return ""
}

func (m *EchoRequest) GetClusterAddrs() []string {
	if m != nil {
		return m.ClusterAddrs
	}
	return nil
}

type EchoReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	ClusterAddrs         []string `protobuf:"bytes,2,rep,name=cluster_addrs,json=clusterAddrs,proto3" json:"cluster_addrs,omitempty"`
	ReplicationState     uint32   `protobuf:"varint,3,opt,name=replication_state,json=replicationState,proto3" json:"replication_state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoReply) Reset()         { *m = EchoReply{} }
func (m *EchoReply) String() string { return proto.CompactTextString(m) }
func (*EchoReply) ProtoMessage()    {}
func (*EchoReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5f7512e4ab7b58a, []int{1}
}

func (m *EchoReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoReply.Unmarshal(m, b)
}
func (m *EchoReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoReply.Marshal(b, m, deterministic)
}
func (m *EchoReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoReply.Merge(m, src)
}
func (m *EchoReply) XXX_Size() int {
	return xxx_messageInfo_EchoReply.Size(m)
}
func (m *EchoReply) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoReply.DiscardUnknown(m)
}

var xxx_messageInfo_EchoReply proto.InternalMessageInfo

func (m *EchoReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *EchoReply) GetClusterAddrs() []string {
	if m != nil {
		return m.ClusterAddrs
	}
	return nil
}

func (m *EchoReply) GetReplicationState() uint32 {
	if m != nil {
		return m.ReplicationState
	}
	return 0
}

type ClientKey struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	X                    []byte   `protobuf:"bytes,2,opt,name=x,proto3" json:"x,omitempty"`
	Y                    []byte   `protobuf:"bytes,3,opt,name=y,proto3" json:"y,omitempty"`
	D                    []byte   `protobuf:"bytes,4,opt,name=d,proto3" json:"d,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientKey) Reset()         { *m = ClientKey{} }
func (m *ClientKey) String() string { return proto.CompactTextString(m) }
func (*ClientKey) ProtoMessage()    {}
func (*ClientKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5f7512e4ab7b58a, []int{2}
}

func (m *ClientKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientKey.Unmarshal(m, b)
}
func (m *ClientKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientKey.Marshal(b, m, deterministic)
}
func (m *ClientKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientKey.Merge(m, src)
}
func (m *ClientKey) XXX_Size() int {
	return xxx_messageInfo_ClientKey.Size(m)
}
func (m *ClientKey) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientKey.DiscardUnknown(m)
}

var xxx_messageInfo_ClientKey proto.InternalMessageInfo

func (m *ClientKey) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ClientKey) GetX() []byte {
	if m != nil {
		return m.X
	}
	return nil
}

func (m *ClientKey) GetY() []byte {
	if m != nil {
		return m.Y
	}
	return nil
}

func (m *ClientKey) GetD() []byte {
	if m != nil {
		return m.D
	}
	return nil
}

type PerfStandbyElectionInput struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PerfStandbyElectionInput) Reset()         { *m = PerfStandbyElectionInput{} }
func (m *PerfStandbyElectionInput) String() string { return proto.CompactTextString(m) }
func (*PerfStandbyElectionInput) ProtoMessage()    {}
func (*PerfStandbyElectionInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5f7512e4ab7b58a, []int{3}
}

func (m *PerfStandbyElectionInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PerfStandbyElectionInput.Unmarshal(m, b)
}
func (m *PerfStandbyElectionInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PerfStandbyElectionInput.Marshal(b, m, deterministic)
}
func (m *PerfStandbyElectionInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PerfStandbyElectionInput.Merge(m, src)
}
func (m *PerfStandbyElectionInput) XXX_Size() int {
	return xxx_messageInfo_PerfStandbyElectionInput.Size(m)
}
func (m *PerfStandbyElectionInput) XXX_DiscardUnknown() {
	xxx_messageInfo_PerfStandbyElectionInput.DiscardUnknown(m)
}

var xxx_messageInfo_PerfStandbyElectionInput proto.InternalMessageInfo

type PerfStandbyElectionResponse struct {
	Id                   string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ClusterId            string     `protobuf:"bytes,2,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	PrimaryClusterAddr   string     `protobuf:"bytes,3,opt,name=primary_cluster_addr,json=primaryClusterAddr,proto3" json:"primary_cluster_addr,omitempty"`
	CaCert               []byte     `protobuf:"bytes,4,opt,name=ca_cert,json=caCert,proto3" json:"ca_cert,omitempty"`
	ClientCert           []byte     `protobuf:"bytes,5,opt,name=client_cert,json=clientCert,proto3" json:"client_cert,omitempty"`
	ClientKey            *ClientKey `protobuf:"bytes,6,opt,name=client_key,json=clientKey,proto3" json:"client_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *PerfStandbyElectionResponse) Reset()         { *m = PerfStandbyElectionResponse{} }
func (m *PerfStandbyElectionResponse) String() string { return proto.CompactTextString(m) }
func (*PerfStandbyElectionResponse) ProtoMessage()    {}
func (*PerfStandbyElectionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5f7512e4ab7b58a, []int{4}
}

func (m *PerfStandbyElectionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PerfStandbyElectionResponse.Unmarshal(m, b)
}
func (m *PerfStandbyElectionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PerfStandbyElectionResponse.Marshal(b, m, deterministic)
}
func (m *PerfStandbyElectionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PerfStandbyElectionResponse.Merge(m, src)
}
func (m *PerfStandbyElectionResponse) XXX_Size() int {
	return xxx_messageInfo_PerfStandbyElectionResponse.Size(m)
}
func (m *PerfStandbyElectionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PerfStandbyElectionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PerfStandbyElectionResponse proto.InternalMessageInfo

func (m *PerfStandbyElectionResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PerfStandbyElectionResponse) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *PerfStandbyElectionResponse) GetPrimaryClusterAddr() string {
	if m != nil {
		return m.PrimaryClusterAddr
	}
	return ""
}

func (m *PerfStandbyElectionResponse) GetCaCert() []byte {
	if m != nil {
		return m.CaCert
	}
	return nil
}

func (m *PerfStandbyElectionResponse) GetClientCert() []byte {
	if m != nil {
		return m.ClientCert
	}
	return nil
}

func (m *PerfStandbyElectionResponse) GetClientKey() *ClientKey {
	if m != nil {
		return m.ClientKey
	}
	return nil
}

func init() {
	proto.RegisterType((*EchoRequest)(nil), "vault.EchoRequest")
	proto.RegisterType((*EchoReply)(nil), "vault.EchoReply")
	proto.RegisterType((*ClientKey)(nil), "vault.ClientKey")
	proto.RegisterType((*PerfStandbyElectionInput)(nil), "vault.PerfStandbyElectionInput")
	proto.RegisterType((*PerfStandbyElectionResponse)(nil), "vault.PerfStandbyElectionResponse")
}

func init() {
	proto.RegisterFile("vault/request_forwarding_service.proto", fileDescriptor_f5f7512e4ab7b58a)
}

var fileDescriptor_f5f7512e4ab7b58a = []byte{
	// 493 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0x41, 0x6f, 0x1a, 0x3d,
	0x10, 0x8d, 0x81, 0x10, 0x31, 0x90, 0x88, 0xf8, 0x8b, 0xf4, 0xad, 0xa8, 0xa2, 0x90, 0xad, 0x54,
	0x21, 0x55, 0xda, 0x8d, 0xd2, 0x73, 0x0f, 0x2d, 0x4a, 0x25, 0xd4, 0x4b, 0xb5, 0xb9, 0xf5, 0xb2,
	0x32, 0xf6, 0x04, 0xac, 0x2e, 0x6b, 0xd7, 0x36, 0x49, 0xf6, 0x27, 0xf7, 0xd6, 0x9f, 0x50, 0xad,
	0xd7, 0x04, 0x10, 0x4d, 0x2f, 0x68, 0xe7, 0xcd, 0x63, 0xde, 0xf8, 0xf9, 0x19, 0xde, 0x3d, 0xb2,
	0x75, 0xe1, 0x52, 0x83, 0x3f, 0xd7, 0x68, 0x5d, 0xfe, 0xa0, 0xcc, 0x13, 0x33, 0x42, 0x96, 0x8b,
	0xdc, 0xa2, 0x79, 0x94, 0x1c, 0x13, 0x6d, 0x94, 0x53, 0xf4, 0xd8, 0xf3, 0x46, 0x97, 0x4b, 0x2c,
	0x34, 0x9a, 0x74, 0xcb, 0x4b, 0x5d, 0xa5, 0xd1, 0x36, 0xac, 0x58, 0x41, 0xff, 0x8e, 0x2f, 0x55,
	0xd6, 0x4c, 0xa3, 0x11, 0x9c, 0xac, 0xd0, 0x5a, 0xb6, 0xc0, 0x88, 0x8c, 0xc9, 0xa4, 0x97, 0x6d,
	0x4a, 0x7a, 0x0d, 0x03, 0x5e, 0xac, 0xad, 0x43, 0x93, 0x33, 0x21, 0x4c, 0xd4, 0xf2, 0xed, 0x7e,
	0xc0, 0x3e, 0x09, 0x61, 0xe8, 0x5b, 0x38, 0xdd, 0xa5, 0xd8, 0xa8, 0x3d, 0x6e, 0x4f, 0x7a, 0xd9,
	0x60, 0x87, 0x63, 0xe3, 0x27, 0xe8, 0x35, 0x82, 0xba, 0xa8, 0xfe, 0x21, 0x77, 0x30, 0xab, 0x75,
	0x38, 0x8b, 0xbe, 0x87, 0x73, 0x83, 0xba, 0x90, 0x9c, 0x39, 0xa9, 0xca, 0xdc, 0x3a, 0xe6, 0x30,
	0x6a, 0x8f, 0xc9, 0xe4, 0x34, 0x1b, 0xee, 0x34, 0xee, 0x6b, 0x3c, 0x9e, 0x41, 0x6f, 0x5a, 0x48,
	0x2c, 0xdd, 0x57, 0xac, 0x28, 0x85, 0x4e, 0xed, 0x42, 0x50, 0xf5, 0xdf, 0x74, 0x00, 0xe4, 0xd9,
	0x1f, 0x6b, 0x90, 0x91, 0xe7, 0xba, 0xaa, 0xfc, 0xac, 0x41, 0x46, 0xaa, 0xba, 0x12, 0x51, 0xa7,
	0xa9, 0x44, 0x3c, 0x82, 0xe8, 0x1b, 0x9a, 0x87, 0x7b, 0xc7, 0x4a, 0x31, 0xaf, 0xee, 0x0a, 0xe4,
	0xb5, 0xcc, 0xac, 0xd4, 0x6b, 0x17, 0xff, 0x22, 0xf0, 0xe6, 0x2f, 0xcd, 0x0c, 0xad, 0x56, 0xa5,
	0x45, 0x7a, 0x06, 0x2d, 0x29, 0x82, 0x6e, 0x4b, 0x0a, 0x7a, 0x09, 0xb0, 0x39, 0xa8, 0x14, 0xc1,
	0xd5, 0x5e, 0x40, 0x66, 0x82, 0xde, 0xc0, 0x85, 0x36, 0x72, 0xc5, 0x4c, 0x95, 0xef, 0xd9, 0xdf,
	0xf6, 0x44, 0x1a, 0x7a, 0xd3, 0x9d, 0x5b, 0xf8, 0x1f, 0x4e, 0x38, 0xcb, 0x39, 0x1a, 0x17, 0x16,
	0xee, 0x72, 0x36, 0x45, 0xe3, 0xe8, 0x15, 0xf4, 0xb9, 0x37, 0xa0, 0x69, 0x1e, 0xfb, 0x26, 0x34,
	0x90, 0x27, 0xa4, 0x10, 0xaa, 0xfc, 0x07, 0x56, 0x51, 0x77, 0x4c, 0x26, 0xfd, 0xdb, 0x61, 0xe2,
	0x63, 0x94, 0xbc, 0x58, 0x57, 0x2f, 0x17, 0x3e, 0x6f, 0x7f, 0x13, 0x38, 0x0f, 0xc9, 0xf9, 0xf2,
	0x12, 0x2f, 0xfa, 0x11, 0xce, 0x42, 0xb5, 0x49, 0xd5, 0x7f, 0xc9, 0x36, 0x7d, 0x49, 0x00, 0x47,
	0x17, 0xfb, 0x60, 0x63, 0x4f, 0x7c, 0x44, 0x13, 0xe8, 0xd4, 0x01, 0xa1, 0x34, 0x28, 0xef, 0xc4,
	0x73, 0x34, 0xdc, 0xc3, 0x74, 0x51, 0xc5, 0x47, 0xb4, 0x80, 0xeb, 0xda, 0x6f, 0x65, 0x56, 0xac,
	0xe4, 0x78, 0x60, 0x7b, 0xb3, 0xc1, 0x55, 0xf8, 0xe3, 0x6b, 0xd7, 0x36, 0x8a, 0x5f, 0x27, 0x6c,
	0x77, 0xbb, 0x21, 0x9f, 0xe3, 0xef, 0xe3, 0x85, 0x74, 0xcb, 0xf5, 0x3c, 0xe1, 0x6a, 0x95, 0x2e,
	0x99, 0x5d, 0x4a, 0xae, 0x8c, 0x4e, 0x9b, 0x47, 0xe9, 0x7f, 0xe7, 0x5d, 0xff, 0xb4, 0x3e, 0xfc,
	0x09, 0x00, 0x00, 0xff, 0xff, 0x03, 0x94, 0x0a, 0x17, 0xaa, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RequestForwardingClient is the client API for RequestForwarding service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RequestForwardingClient interface {
	ForwardRequest(ctx context.Context, in *forwarding.Request, opts ...grpc.CallOption) (*forwarding.Response, error)
	Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoReply, error)
	PerformanceStandbyElectionRequest(ctx context.Context, in *PerfStandbyElectionInput, opts ...grpc.CallOption) (RequestForwarding_PerformanceStandbyElectionRequestClient, error)
}

type requestForwardingClient struct {
	cc *grpc.ClientConn
}

func NewRequestForwardingClient(cc *grpc.ClientConn) RequestForwardingClient {
	return &requestForwardingClient{cc}
}

func (c *requestForwardingClient) ForwardRequest(ctx context.Context, in *forwarding.Request, opts ...grpc.CallOption) (*forwarding.Response, error) {
	out := new(forwarding.Response)
	err := c.cc.Invoke(ctx, "/vault.RequestForwarding/ForwardRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requestForwardingClient) Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoReply, error) {
	out := new(EchoReply)
	err := c.cc.Invoke(ctx, "/vault.RequestForwarding/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requestForwardingClient) PerformanceStandbyElectionRequest(ctx context.Context, in *PerfStandbyElectionInput, opts ...grpc.CallOption) (RequestForwarding_PerformanceStandbyElectionRequestClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RequestForwarding_serviceDesc.Streams[0], "/vault.RequestForwarding/PerformanceStandbyElectionRequest", opts...)
	if err != nil {
		return nil, err
	}
	x := &requestForwardingPerformanceStandbyElectionRequestClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RequestForwarding_PerformanceStandbyElectionRequestClient interface {
	Recv() (*PerfStandbyElectionResponse, error)
	grpc.ClientStream
}

type requestForwardingPerformanceStandbyElectionRequestClient struct {
	grpc.ClientStream
}

func (x *requestForwardingPerformanceStandbyElectionRequestClient) Recv() (*PerfStandbyElectionResponse, error) {
	m := new(PerfStandbyElectionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RequestForwardingServer is the server API for RequestForwarding service.
type RequestForwardingServer interface {
	ForwardRequest(context.Context, *forwarding.Request) (*forwarding.Response, error)
	Echo(context.Context, *EchoRequest) (*EchoReply, error)
	PerformanceStandbyElectionRequest(*PerfStandbyElectionInput, RequestForwarding_PerformanceStandbyElectionRequestServer) error
}

func RegisterRequestForwardingServer(s *grpc.Server, srv RequestForwardingServer) {
	s.RegisterService(&_RequestForwarding_serviceDesc, srv)
}

func _RequestForwarding_ForwardRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(forwarding.Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestForwardingServer).ForwardRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vault.RequestForwarding/ForwardRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestForwardingServer).ForwardRequest(ctx, req.(*forwarding.Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _RequestForwarding_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestForwardingServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vault.RequestForwarding/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestForwardingServer).Echo(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RequestForwarding_PerformanceStandbyElectionRequest_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PerfStandbyElectionInput)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RequestForwardingServer).PerformanceStandbyElectionRequest(m, &requestForwardingPerformanceStandbyElectionRequestServer{stream})
}

type RequestForwarding_PerformanceStandbyElectionRequestServer interface {
	Send(*PerfStandbyElectionResponse) error
	grpc.ServerStream
}

type requestForwardingPerformanceStandbyElectionRequestServer struct {
	grpc.ServerStream
}

func (x *requestForwardingPerformanceStandbyElectionRequestServer) Send(m *PerfStandbyElectionResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _RequestForwarding_serviceDesc = grpc.ServiceDesc{
	ServiceName: "vault.RequestForwarding",
	HandlerType: (*RequestForwardingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ForwardRequest",
			Handler:    _RequestForwarding_ForwardRequest_Handler,
		},
		{
			MethodName: "Echo",
			Handler:    _RequestForwarding_Echo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PerformanceStandbyElectionRequest",
			Handler:       _RequestForwarding_PerformanceStandbyElectionRequest_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "vault/request_forwarding_service.proto",
}
