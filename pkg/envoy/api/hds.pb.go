// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/hds.proto

package envoy_api_v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf2 "github.com/golang/protobuf/ptypes/duration"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Different Envoy instances may have different capabilities (e.g. Redis)
// and/or have ports enabled for different protocols.
type Capability_Protocol int32

const (
	Capability_HTTP  Capability_Protocol = 0
	Capability_TCP   Capability_Protocol = 1
	Capability_REDIS Capability_Protocol = 2
)

var Capability_Protocol_name = map[int32]string{
	0: "HTTP",
	1: "TCP",
	2: "REDIS",
}
var Capability_Protocol_value = map[string]int32{
	"HTTP":  0,
	"TCP":   1,
	"REDIS": 2,
}

func (x Capability_Protocol) String() string {
	return proto.EnumName(Capability_Protocol_name, int32(x))
}
func (Capability_Protocol) EnumDescriptor() ([]byte, []int) { return fileDescriptor7, []int{0, 0} }

// Defines supported protocols etc, so the management server can assign proper
// endpoints to healthcheck.
type Capability struct {
	HealthCheckProtocol []Capability_Protocol `protobuf:"varint,1,rep,packed,name=health_check_protocol,json=healthCheckProtocol,enum=envoy.api.v2.Capability_Protocol" json:"health_check_protocol,omitempty"`
}

func (m *Capability) Reset()                    { *m = Capability{} }
func (m *Capability) String() string            { return proto.CompactTextString(m) }
func (*Capability) ProtoMessage()               {}
func (*Capability) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *Capability) GetHealthCheckProtocol() []Capability_Protocol {
	if m != nil {
		return m.HealthCheckProtocol
	}
	return nil
}

type HealthCheckRequest struct {
	Node       *Node       `protobuf:"bytes,1,opt,name=node" json:"node,omitempty"`
	Capability *Capability `protobuf:"bytes,2,opt,name=capability" json:"capability,omitempty"`
}

func (m *HealthCheckRequest) Reset()                    { *m = HealthCheckRequest{} }
func (m *HealthCheckRequest) String() string            { return proto.CompactTextString(m) }
func (*HealthCheckRequest) ProtoMessage()               {}
func (*HealthCheckRequest) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

func (m *HealthCheckRequest) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *HealthCheckRequest) GetCapability() *Capability {
	if m != nil {
		return m.Capability
	}
	return nil
}

type EndpointHealth struct {
	Endpoint     *Endpoint    `protobuf:"bytes,1,opt,name=endpoint" json:"endpoint,omitempty"`
	HealthStatus HealthStatus `protobuf:"varint,2,opt,name=health_status,json=healthStatus,enum=envoy.api.v2.HealthStatus" json:"health_status,omitempty"`
}

func (m *EndpointHealth) Reset()                    { *m = EndpointHealth{} }
func (m *EndpointHealth) String() string            { return proto.CompactTextString(m) }
func (*EndpointHealth) ProtoMessage()               {}
func (*EndpointHealth) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{2} }

func (m *EndpointHealth) GetEndpoint() *Endpoint {
	if m != nil {
		return m.Endpoint
	}
	return nil
}

func (m *EndpointHealth) GetHealthStatus() HealthStatus {
	if m != nil {
		return m.HealthStatus
	}
	return HealthStatus_UNKNOWN
}

type EndpointHealthResponse struct {
	EndpointsHealth []*EndpointHealth `protobuf:"bytes,1,rep,name=endpoints_health,json=endpointsHealth" json:"endpoints_health,omitempty"`
}

func (m *EndpointHealthResponse) Reset()                    { *m = EndpointHealthResponse{} }
func (m *EndpointHealthResponse) String() string            { return proto.CompactTextString(m) }
func (*EndpointHealthResponse) ProtoMessage()               {}
func (*EndpointHealthResponse) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{3} }

func (m *EndpointHealthResponse) GetEndpointsHealth() []*EndpointHealth {
	if m != nil {
		return m.EndpointsHealth
	}
	return nil
}

type HealthCheckRequestOrEndpointHealthResponse struct {
	// Types that are valid to be assigned to RequestType:
	//	*HealthCheckRequestOrEndpointHealthResponse_HealthCheckRequest
	//	*HealthCheckRequestOrEndpointHealthResponse_EndpointHealthResponse
	RequestType isHealthCheckRequestOrEndpointHealthResponse_RequestType `protobuf_oneof:"request_type"`
}

func (m *HealthCheckRequestOrEndpointHealthResponse) Reset() {
	*m = HealthCheckRequestOrEndpointHealthResponse{}
}
func (m *HealthCheckRequestOrEndpointHealthResponse) String() string {
	return proto.CompactTextString(m)
}
func (*HealthCheckRequestOrEndpointHealthResponse) ProtoMessage() {}
func (*HealthCheckRequestOrEndpointHealthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor7, []int{4}
}

type isHealthCheckRequestOrEndpointHealthResponse_RequestType interface {
	isHealthCheckRequestOrEndpointHealthResponse_RequestType()
}

type HealthCheckRequestOrEndpointHealthResponse_HealthCheckRequest struct {
	HealthCheckRequest *HealthCheckRequest `protobuf:"bytes,1,opt,name=health_check_request,json=healthCheckRequest,oneof"`
}
type HealthCheckRequestOrEndpointHealthResponse_EndpointHealthResponse struct {
	EndpointHealthResponse *EndpointHealthResponse `protobuf:"bytes,2,opt,name=endpoint_health_response,json=endpointHealthResponse,oneof"`
}

func (*HealthCheckRequestOrEndpointHealthResponse_HealthCheckRequest) isHealthCheckRequestOrEndpointHealthResponse_RequestType() {
}
func (*HealthCheckRequestOrEndpointHealthResponse_EndpointHealthResponse) isHealthCheckRequestOrEndpointHealthResponse_RequestType() {
}

func (m *HealthCheckRequestOrEndpointHealthResponse) GetRequestType() isHealthCheckRequestOrEndpointHealthResponse_RequestType {
	if m != nil {
		return m.RequestType
	}
	return nil
}

func (m *HealthCheckRequestOrEndpointHealthResponse) GetHealthCheckRequest() *HealthCheckRequest {
	if x, ok := m.GetRequestType().(*HealthCheckRequestOrEndpointHealthResponse_HealthCheckRequest); ok {
		return x.HealthCheckRequest
	}
	return nil
}

func (m *HealthCheckRequestOrEndpointHealthResponse) GetEndpointHealthResponse() *EndpointHealthResponse {
	if x, ok := m.GetRequestType().(*HealthCheckRequestOrEndpointHealthResponse_EndpointHealthResponse); ok {
		return x.EndpointHealthResponse
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*HealthCheckRequestOrEndpointHealthResponse) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _HealthCheckRequestOrEndpointHealthResponse_OneofMarshaler, _HealthCheckRequestOrEndpointHealthResponse_OneofUnmarshaler, _HealthCheckRequestOrEndpointHealthResponse_OneofSizer, []interface{}{
		(*HealthCheckRequestOrEndpointHealthResponse_HealthCheckRequest)(nil),
		(*HealthCheckRequestOrEndpointHealthResponse_EndpointHealthResponse)(nil),
	}
}

func _HealthCheckRequestOrEndpointHealthResponse_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*HealthCheckRequestOrEndpointHealthResponse)
	// request_type
	switch x := m.RequestType.(type) {
	case *HealthCheckRequestOrEndpointHealthResponse_HealthCheckRequest:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.HealthCheckRequest); err != nil {
			return err
		}
	case *HealthCheckRequestOrEndpointHealthResponse_EndpointHealthResponse:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.EndpointHealthResponse); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("HealthCheckRequestOrEndpointHealthResponse.RequestType has unexpected type %T", x)
	}
	return nil
}

func _HealthCheckRequestOrEndpointHealthResponse_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*HealthCheckRequestOrEndpointHealthResponse)
	switch tag {
	case 1: // request_type.health_check_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(HealthCheckRequest)
		err := b.DecodeMessage(msg)
		m.RequestType = &HealthCheckRequestOrEndpointHealthResponse_HealthCheckRequest{msg}
		return true, err
	case 2: // request_type.endpoint_health_response
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EndpointHealthResponse)
		err := b.DecodeMessage(msg)
		m.RequestType = &HealthCheckRequestOrEndpointHealthResponse_EndpointHealthResponse{msg}
		return true, err
	default:
		return false, nil
	}
}

func _HealthCheckRequestOrEndpointHealthResponse_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*HealthCheckRequestOrEndpointHealthResponse)
	// request_type
	switch x := m.RequestType.(type) {
	case *HealthCheckRequestOrEndpointHealthResponse_HealthCheckRequest:
		s := proto.Size(x.HealthCheckRequest)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *HealthCheckRequestOrEndpointHealthResponse_EndpointHealthResponse:
		s := proto.Size(x.EndpointHealthResponse)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type LocalityEndpoints struct {
	Locality  *Locality   `protobuf:"bytes,1,opt,name=locality" json:"locality,omitempty"`
	Endpoints []*Endpoint `protobuf:"bytes,2,rep,name=endpoints" json:"endpoints,omitempty"`
}

func (m *LocalityEndpoints) Reset()                    { *m = LocalityEndpoints{} }
func (m *LocalityEndpoints) String() string            { return proto.CompactTextString(m) }
func (*LocalityEndpoints) ProtoMessage()               {}
func (*LocalityEndpoints) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{5} }

func (m *LocalityEndpoints) GetLocality() *Locality {
	if m != nil {
		return m.Locality
	}
	return nil
}

func (m *LocalityEndpoints) GetEndpoints() []*Endpoint {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

// The cluster name and locality is provided to Envoy for the endpoints that it
// health checks to support statistics reporting, logging and debugging by the
// Envoy instance (outside of HDS). For maximum usefulness, it should match the
// same cluster structure as that provided by EDS.
type ClusterHealthCheck struct {
	ClusterName  string               `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName" json:"cluster_name,omitempty"`
	HealthChecks []*HealthCheck       `protobuf:"bytes,2,rep,name=health_checks,json=healthChecks" json:"health_checks,omitempty"`
	Endpoints    []*LocalityEndpoints `protobuf:"bytes,3,rep,name=endpoints" json:"endpoints,omitempty"`
}

func (m *ClusterHealthCheck) Reset()                    { *m = ClusterHealthCheck{} }
func (m *ClusterHealthCheck) String() string            { return proto.CompactTextString(m) }
func (*ClusterHealthCheck) ProtoMessage()               {}
func (*ClusterHealthCheck) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{6} }

func (m *ClusterHealthCheck) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *ClusterHealthCheck) GetHealthChecks() []*HealthCheck {
	if m != nil {
		return m.HealthChecks
	}
	return nil
}

func (m *ClusterHealthCheck) GetEndpoints() []*LocalityEndpoints {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

type HealthCheckSpecifier struct {
	HealthCheck []*ClusterHealthCheck `protobuf:"bytes,1,rep,name=health_check,json=healthCheck" json:"health_check,omitempty"`
	// The default is 1 second.
	Interval *google_protobuf2.Duration `protobuf:"bytes,2,opt,name=interval" json:"interval,omitempty"`
}

func (m *HealthCheckSpecifier) Reset()                    { *m = HealthCheckSpecifier{} }
func (m *HealthCheckSpecifier) String() string            { return proto.CompactTextString(m) }
func (*HealthCheckSpecifier) ProtoMessage()               {}
func (*HealthCheckSpecifier) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{7} }

func (m *HealthCheckSpecifier) GetHealthCheck() []*ClusterHealthCheck {
	if m != nil {
		return m.HealthCheck
	}
	return nil
}

func (m *HealthCheckSpecifier) GetInterval() *google_protobuf2.Duration {
	if m != nil {
		return m.Interval
	}
	return nil
}

func init() {
	proto.RegisterType((*Capability)(nil), "envoy.api.v2.Capability")
	proto.RegisterType((*HealthCheckRequest)(nil), "envoy.api.v2.HealthCheckRequest")
	proto.RegisterType((*EndpointHealth)(nil), "envoy.api.v2.EndpointHealth")
	proto.RegisterType((*EndpointHealthResponse)(nil), "envoy.api.v2.EndpointHealthResponse")
	proto.RegisterType((*HealthCheckRequestOrEndpointHealthResponse)(nil), "envoy.api.v2.HealthCheckRequestOrEndpointHealthResponse")
	proto.RegisterType((*LocalityEndpoints)(nil), "envoy.api.v2.LocalityEndpoints")
	proto.RegisterType((*ClusterHealthCheck)(nil), "envoy.api.v2.ClusterHealthCheck")
	proto.RegisterType((*HealthCheckSpecifier)(nil), "envoy.api.v2.HealthCheckSpecifier")
	proto.RegisterEnum("envoy.api.v2.Capability_Protocol", Capability_Protocol_name, Capability_Protocol_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for HealthDiscoveryService service

type HealthDiscoveryServiceClient interface {
	// 1. Envoy starts up and if its can_healthcheck option in the static
	//    bootstrap config is enabled, sends HealthCheckRequest to the management
	//    server. It supplies its capabilities (which protocol it can health check
	//    with, what zone it resides in, etc.).
	// 2. In response to (1), the management server designates this Envoy as a
	//    healthchecker to health check a subset of all upstream hosts for a given
	//    cluster (for example upstream Host 1 and Host 2). It streams
	//    HealthCheckSpecifier messages with cluster related configuration for all
	//    clusters this Envoy is designated to health check. Subsequent
	//    HealthCheckSpecifier message will be sent on changes to:
	//    a. Endpoints to health checks
	//    b. Per cluster configuration change
	// 3. Envoy creates a health probe based on the HealthCheck config and sends
	//    it to endpoint(ip:port) of Host 1 and 2. Based on the HealthCheck
	//    configuration Envoy waits upon the arrival of the probe response and
	//    looks at the content of the response to decide whether the endpoint is
	//    healthy or not. If a response hasn’t been received within the timeout
	//    interval, the endpoint health status is considered TIMEOUT.
	// 4. Envoy reports results back in an EndpointHealthResponse message.
	//    Envoy streams responses as often as the interval configured by the
	//    management server in HealthCheckSpecifier.
	// 5. The management Server collects health statuses for all endpoints in the
	//    cluster (for all clusters) and uses this information to construct
	//    EndpointDiscoveryResponse messages.
	// 6. Once Envoy has a list of upstream endpoints to send traffic to, it load
	//    balances traffic to them without additional health checking. It may
	//    use inline healthcheck (i.e. consider endpoint UNHEALTHY if connection
	//    failed to a particular endpoint to account for health status propagation
	//    delay between HDS and EDS).
	// By default, can_healthcheck is true. If can_healthcheck is false, Cluster
	// configuration may not contain HealthCheck message.
	// TODO(htuch): How is can_healthcheck communicated to CDS to ensure the above
	// invariant?
	// TODO(htuch): Add @amb67's diagram.
	StreamHealthCheck(ctx context.Context, opts ...grpc.CallOption) (HealthDiscoveryService_StreamHealthCheckClient, error)
	// TODO(htuch): Unlike the gRPC version, there is no stream-based binding of
	// request/response. Should we add an identifier to the HealthCheckSpecifier
	// to bind with the response?
	FetchHealthCheck(ctx context.Context, in *HealthCheckRequestOrEndpointHealthResponse, opts ...grpc.CallOption) (*HealthCheckSpecifier, error)
}

type healthDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewHealthDiscoveryServiceClient(cc *grpc.ClientConn) HealthDiscoveryServiceClient {
	return &healthDiscoveryServiceClient{cc}
}

func (c *healthDiscoveryServiceClient) StreamHealthCheck(ctx context.Context, opts ...grpc.CallOption) (HealthDiscoveryService_StreamHealthCheckClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_HealthDiscoveryService_serviceDesc.Streams[0], c.cc, "/envoy.api.v2.HealthDiscoveryService/StreamHealthCheck", opts...)
	if err != nil {
		return nil, err
	}
	x := &healthDiscoveryServiceStreamHealthCheckClient{stream}
	return x, nil
}

type HealthDiscoveryService_StreamHealthCheckClient interface {
	Send(*HealthCheckRequestOrEndpointHealthResponse) error
	Recv() (*HealthCheckSpecifier, error)
	grpc.ClientStream
}

type healthDiscoveryServiceStreamHealthCheckClient struct {
	grpc.ClientStream
}

func (x *healthDiscoveryServiceStreamHealthCheckClient) Send(m *HealthCheckRequestOrEndpointHealthResponse) error {
	return x.ClientStream.SendMsg(m)
}

func (x *healthDiscoveryServiceStreamHealthCheckClient) Recv() (*HealthCheckSpecifier, error) {
	m := new(HealthCheckSpecifier)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *healthDiscoveryServiceClient) FetchHealthCheck(ctx context.Context, in *HealthCheckRequestOrEndpointHealthResponse, opts ...grpc.CallOption) (*HealthCheckSpecifier, error) {
	out := new(HealthCheckSpecifier)
	err := grpc.Invoke(ctx, "/envoy.api.v2.HealthDiscoveryService/FetchHealthCheck", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for HealthDiscoveryService service

type HealthDiscoveryServiceServer interface {
	// 1. Envoy starts up and if its can_healthcheck option in the static
	//    bootstrap config is enabled, sends HealthCheckRequest to the management
	//    server. It supplies its capabilities (which protocol it can health check
	//    with, what zone it resides in, etc.).
	// 2. In response to (1), the management server designates this Envoy as a
	//    healthchecker to health check a subset of all upstream hosts for a given
	//    cluster (for example upstream Host 1 and Host 2). It streams
	//    HealthCheckSpecifier messages with cluster related configuration for all
	//    clusters this Envoy is designated to health check. Subsequent
	//    HealthCheckSpecifier message will be sent on changes to:
	//    a. Endpoints to health checks
	//    b. Per cluster configuration change
	// 3. Envoy creates a health probe based on the HealthCheck config and sends
	//    it to endpoint(ip:port) of Host 1 and 2. Based on the HealthCheck
	//    configuration Envoy waits upon the arrival of the probe response and
	//    looks at the content of the response to decide whether the endpoint is
	//    healthy or not. If a response hasn’t been received within the timeout
	//    interval, the endpoint health status is considered TIMEOUT.
	// 4. Envoy reports results back in an EndpointHealthResponse message.
	//    Envoy streams responses as often as the interval configured by the
	//    management server in HealthCheckSpecifier.
	// 5. The management Server collects health statuses for all endpoints in the
	//    cluster (for all clusters) and uses this information to construct
	//    EndpointDiscoveryResponse messages.
	// 6. Once Envoy has a list of upstream endpoints to send traffic to, it load
	//    balances traffic to them without additional health checking. It may
	//    use inline healthcheck (i.e. consider endpoint UNHEALTHY if connection
	//    failed to a particular endpoint to account for health status propagation
	//    delay between HDS and EDS).
	// By default, can_healthcheck is true. If can_healthcheck is false, Cluster
	// configuration may not contain HealthCheck message.
	// TODO(htuch): How is can_healthcheck communicated to CDS to ensure the above
	// invariant?
	// TODO(htuch): Add @amb67's diagram.
	StreamHealthCheck(HealthDiscoveryService_StreamHealthCheckServer) error
	// TODO(htuch): Unlike the gRPC version, there is no stream-based binding of
	// request/response. Should we add an identifier to the HealthCheckSpecifier
	// to bind with the response?
	FetchHealthCheck(context.Context, *HealthCheckRequestOrEndpointHealthResponse) (*HealthCheckSpecifier, error)
}

func RegisterHealthDiscoveryServiceServer(s *grpc.Server, srv HealthDiscoveryServiceServer) {
	s.RegisterService(&_HealthDiscoveryService_serviceDesc, srv)
}

func _HealthDiscoveryService_StreamHealthCheck_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HealthDiscoveryServiceServer).StreamHealthCheck(&healthDiscoveryServiceStreamHealthCheckServer{stream})
}

type HealthDiscoveryService_StreamHealthCheckServer interface {
	Send(*HealthCheckSpecifier) error
	Recv() (*HealthCheckRequestOrEndpointHealthResponse, error)
	grpc.ServerStream
}

type healthDiscoveryServiceStreamHealthCheckServer struct {
	grpc.ServerStream
}

func (x *healthDiscoveryServiceStreamHealthCheckServer) Send(m *HealthCheckSpecifier) error {
	return x.ServerStream.SendMsg(m)
}

func (x *healthDiscoveryServiceStreamHealthCheckServer) Recv() (*HealthCheckRequestOrEndpointHealthResponse, error) {
	m := new(HealthCheckRequestOrEndpointHealthResponse)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _HealthDiscoveryService_FetchHealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequestOrEndpointHealthResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthDiscoveryServiceServer).FetchHealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.api.v2.HealthDiscoveryService/FetchHealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthDiscoveryServiceServer).FetchHealthCheck(ctx, req.(*HealthCheckRequestOrEndpointHealthResponse))
	}
	return interceptor(ctx, in, info, handler)
}

var _HealthDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.api.v2.HealthDiscoveryService",
	HandlerType: (*HealthDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchHealthCheck",
			Handler:    _HealthDiscoveryService_FetchHealthCheck_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamHealthCheck",
			Handler:       _HealthDiscoveryService_StreamHealthCheck_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api/hds.proto",
}

func init() { proto.RegisterFile("api/hds.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 666 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0xc1, 0x6e, 0xd3, 0x4c,
	0x10, 0xce, 0xa6, 0xfd, 0x7f, 0xd2, 0x49, 0x1a, 0xd2, 0xa5, 0x44, 0x69, 0x54, 0xd1, 0xd6, 0x02,
	0x14, 0xf5, 0xe0, 0x20, 0x03, 0x52, 0x55, 0x09, 0x90, 0x48, 0x0b, 0x41, 0x42, 0xa5, 0xda, 0x84,
	0x73, 0xd8, 0x38, 0xdb, 0xc6, 0xc2, 0xf5, 0x1a, 0x7b, 0x63, 0x94, 0x03, 0x17, 0x24, 0x8e, 0x5c,
	0xe0, 0xc0, 0x43, 0x70, 0xe7, 0x45, 0x78, 0x05, 0x1e, 0x81, 0x07, 0x40, 0xd9, 0x5d, 0xbb, 0x1b,
	0x12, 0xc3, 0x8d, 0x9b, 0x3d, 0xf3, 0xcd, 0xcc, 0x37, 0xdf, 0x37, 0x5a, 0x58, 0xa7, 0xa1, 0xd7,
	0x1e, 0x8f, 0x62, 0x3b, 0x8c, 0xb8, 0xe0, 0xb8, 0xc2, 0x82, 0x84, 0x4f, 0x6d, 0x1a, 0x7a, 0x76,
	0xe2, 0x34, 0xab, 0xb3, 0xe4, 0x90, 0xc6, 0x4c, 0x65, 0x9b, 0x75, 0x09, 0x66, 0xd4, 0x17, 0xe3,
	0x81, 0x3b, 0x66, 0xee, 0x6b, 0x1d, 0xdf, 0x3e, 0xe7, 0xfc, 0xdc, 0x67, 0xed, 0x59, 0x9a, 0x06,
	0x01, 0x17, 0x54, 0x78, 0x3c, 0xd0, 0x3d, 0x9b, 0x37, 0x74, 0x56, 0xfe, 0x0d, 0x27, 0x67, 0xed,
	0xd1, 0x24, 0x92, 0x00, 0x95, 0xb7, 0x3e, 0x22, 0x80, 0x0e, 0x0d, 0xe9, 0xd0, 0xf3, 0x3d, 0x31,
	0xc5, 0x2f, 0xe1, 0xba, 0x39, 0x62, 0x20, 0x41, 0x2e, 0xf7, 0x1b, 0x68, 0x77, 0xa5, 0x55, 0x75,
	0xf6, 0x6c, 0x93, 0xa2, 0x7d, 0x59, 0x68, 0x9f, 0x6a, 0x20, 0xb9, 0xa6, 0xea, 0x3b, 0xb3, 0xf2,
	0x34, 0x68, 0xb5, 0xa0, 0x94, 0x7e, 0xe3, 0x12, 0xac, 0x76, 0xfb, 0xfd, 0xd3, 0x5a, 0x01, 0x5f,
	0x81, 0x95, 0x7e, 0xe7, 0xb4, 0x86, 0xf0, 0x1a, 0xfc, 0x47, 0x8e, 0x8f, 0x9e, 0xf5, 0x6a, 0x45,
	0x2b, 0x01, 0xdc, 0xbd, 0x6c, 0x40, 0xd8, 0x9b, 0x09, 0x8b, 0x05, 0xbe, 0x0d, 0xab, 0x01, 0x1f,
	0xb1, 0x06, 0xda, 0x45, 0xad, 0xb2, 0x83, 0xe7, 0x59, 0x9c, 0xf0, 0x11, 0x23, 0x32, 0x8f, 0x0f,
	0x00, 0xdc, 0x8c, 0x53, 0xa3, 0x28, 0xd1, 0x8d, 0x3c, 0xce, 0xc4, 0xc0, 0x5a, 0x1f, 0x10, 0x54,
	0x8f, 0x83, 0x51, 0xc8, 0xbd, 0x40, 0x28, 0x02, 0xd8, 0x81, 0x12, 0xd3, 0x11, 0x3d, 0xb8, 0x3e,
	0xdf, 0x2a, 0xc5, 0x93, 0x0c, 0x87, 0x1f, 0xc1, 0xba, 0xd6, 0x2f, 0x16, 0x54, 0x4c, 0x62, 0xc9,
	0xa1, 0xea, 0x34, 0xe7, 0x0b, 0xd5, 0x80, 0x9e, 0x44, 0x90, 0xca, 0xd8, 0xf8, 0xb3, 0x28, 0xd4,
	0xe7, 0x69, 0x10, 0x16, 0x87, 0x3c, 0x88, 0x19, 0x7e, 0x0a, 0xb5, 0x74, 0x4c, 0x3c, 0x50, 0x35,
	0xd2, 0x95, 0xb2, 0xb3, 0xbd, 0x9c, 0x96, 0xae, 0xbf, 0x9a, 0x55, 0xa9, 0x80, 0xf5, 0x13, 0xc1,
	0xfe, 0xa2, 0xc6, 0x2f, 0xa2, 0x9c, 0xb9, 0x7d, 0xd8, 0x9c, 0x3b, 0x89, 0x48, 0xe1, 0xb5, 0x24,
	0xbb, 0xcb, 0x36, 0x33, 0xfb, 0x76, 0x0b, 0x04, 0x8f, 0x17, 0x1d, 0x7d, 0x05, 0x8d, 0x94, 0x97,
	0x5e, 0x66, 0x10, 0xe9, 0x89, 0xda, 0xb7, 0x9b, 0x7f, 0xdc, 0x4a, 0x63, 0xbb, 0x05, 0x52, 0x67,
	0x4b, 0x33, 0x8f, 0xab, 0x50, 0xd1, 0x54, 0x07, 0x62, 0x1a, 0x32, 0xeb, 0x1d, 0x6c, 0x3c, 0xe7,
	0x2e, 0x9d, 0xb9, 0x9d, 0xf6, 0x8a, 0x67, 0x1e, 0xfb, 0x3a, 0xb8, 0xdc, 0xe3, 0xb4, 0x84, 0x64,
	0x38, 0x7c, 0x0f, 0xd6, 0x32, 0x49, 0x1b, 0x45, 0xe9, 0x40, 0xde, 0x61, 0x5c, 0x02, 0xad, 0x6f,
	0x08, 0x70, 0xc7, 0x9f, 0xc4, 0x82, 0x45, 0x86, 0x48, 0x78, 0x0f, 0x2a, 0xae, 0x8a, 0x0e, 0x02,
	0x7a, 0xa1, 0x2e, 0x7c, 0x8d, 0x94, 0x75, 0xec, 0x84, 0x5e, 0x30, 0xfc, 0x30, 0xbb, 0x29, 0x69,
	0x40, 0x3a, 0x73, 0x2b, 0x5f, 0xf9, 0x8a, 0x21, 0x78, 0x8c, 0x1f, 0x98, 0x7c, 0x57, 0x64, 0xed,
	0xce, 0xf2, 0x25, 0x33, 0x5d, 0x4c, 0xe2, 0x9f, 0x10, 0x6c, 0x1a, 0xcd, 0x7b, 0x21, 0x73, 0xbd,
	0x33, 0x8f, 0x45, 0xb8, 0x03, 0x15, 0x93, 0x97, 0x3e, 0xc6, 0xdf, 0x0e, 0x62, 0x71, 0x65, 0x52,
	0x36, 0xd8, 0xe1, 0xfb, 0x50, 0xf2, 0x02, 0xc1, 0xa2, 0x84, 0xfa, 0xda, 0xf7, 0x2d, 0x5b, 0x3d,
	0x59, 0x76, 0xfa, 0x64, 0xd9, 0x47, 0xfa, 0xc9, 0x22, 0x19, 0xd4, 0xf9, 0x5a, 0x84, 0xba, 0xea,
	0x79, 0xe4, 0xc5, 0x2e, 0x4f, 0x58, 0x34, 0xed, 0xb1, 0x28, 0xf1, 0x5c, 0x86, 0xdf, 0xc2, 0x46,
	0x4f, 0x44, 0x8c, 0x5e, 0x98, 0x32, 0x1f, 0xfc, 0xed, 0x4c, 0xf3, 0xce, 0xbf, 0x69, 0xe5, 0x56,
	0x66, 0x4a, 0x58, 0x85, 0x16, 0xba, 0x83, 0xf0, 0x17, 0x04, 0xb5, 0x27, 0x4c, 0xb8, 0xe3, 0x7f,
	0x37, 0xf8, 0xd6, 0xfb, 0xef, 0x3f, 0x3e, 0x17, 0x77, 0xac, 0x66, 0x3b, 0x71, 0xda, 0xa3, 0x54,
	0x89, 0x43, 0xd3, 0x96, 0x43, 0xb4, 0x3f, 0xfc, 0x5f, 0x4a, 0x79, 0xf7, 0x57, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x15, 0xb8, 0x1d, 0x3d, 0x70, 0x06, 0x00, 0x00,
}
