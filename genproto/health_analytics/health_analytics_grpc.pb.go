// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: health_analytics.proto

package health_analytics

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// HealthAnalyticsServiceClient is the client API for HealthAnalyticsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HealthAnalyticsServiceClient interface {
	AddMedicalRecord(ctx context.Context, in *AddMedicalRecordRequest, opts ...grpc.CallOption) (*AddMedicalRecordResponse, error)
	GetMedicalRecord(ctx context.Context, in *GetMedicalRecordRequest, opts ...grpc.CallOption) (*GetMedicalRecordResponse, error)
	UpdateMedicalRecord(ctx context.Context, in *UpdateMedicalRecordRequest, opts ...grpc.CallOption) (*UpdateMedicalRecordResponse, error)
	DeleteMedicalRecord(ctx context.Context, in *DeleteMedicalRecordRequest, opts ...grpc.CallOption) (*DeleteMedicalRecordResponse, error)
	ListMedicalRecords(ctx context.Context, in *ListMedicalRecordsRequest, opts ...grpc.CallOption) (*ListMedicalRecordsResponse, error)
	AddLifestyleData(ctx context.Context, in *AddLifestyleDataRequest, opts ...grpc.CallOption) (*AddLifestyleDataResponse, error)
	GetLifestyleData(ctx context.Context, in *GetLifestyleDataRequest, opts ...grpc.CallOption) (*GetLifestyleDataResponse, error)
	UpdateLifestyleData(ctx context.Context, in *UpdateLifestyleDataRequest, opts ...grpc.CallOption) (*UpdateLifestyleDataResponse, error)
	DeleteLifestyleData(ctx context.Context, in *DeleteLifestyleDataRequest, opts ...grpc.CallOption) (*DeleteLifestyleDataResponse, error)
	AddWearableData(ctx context.Context, in *AddWearableDataRequest, opts ...grpc.CallOption) (*AddWearableDataResponse, error)
	GetWearableData(ctx context.Context, in *GetWearableDataRequest, opts ...grpc.CallOption) (*GetWearableDataResponse, error)
	UpdateWearableData(ctx context.Context, in *UpdateWearableDataRequest, opts ...grpc.CallOption) (*UpdateWearableDataResponse, error)
	DeleteWearableData(ctx context.Context, in *DeleteWearableDataRequest, opts ...grpc.CallOption) (*DeleteWearableDataResponse, error)
	GenerateHealthRecommendations(ctx context.Context, in *GenerateHealthRecommendationsRequest, opts ...grpc.CallOption) (*GenerateHealthRecommendationsResponse, error)
	GetRealtimeHealthMonitoring(ctx context.Context, in *GetRealtimeHealthMonitoringRequest, opts ...grpc.CallOption) (HealthAnalyticsService_GetRealtimeHealthMonitoringClient, error)
	GetDailyHealthSummary(ctx context.Context, in *GetDailyHealthSummaryRequest, opts ...grpc.CallOption) (*GetDailyHealthSummaryResponse, error)
	GetWeeklyHealthSummary(ctx context.Context, in *GetWeeklyHealthSummaryRequest, opts ...grpc.CallOption) (*GetWeeklyHealthSummaryResponse, error)
}

type healthAnalyticsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHealthAnalyticsServiceClient(cc grpc.ClientConnInterface) HealthAnalyticsServiceClient {
	return &healthAnalyticsServiceClient{cc}
}

func (c *healthAnalyticsServiceClient) AddMedicalRecord(ctx context.Context, in *AddMedicalRecordRequest, opts ...grpc.CallOption) (*AddMedicalRecordResponse, error) {
	out := new(AddMedicalRecordResponse)
	err := c.cc.Invoke(ctx, "/health_analytics.HealthAnalyticsService/AddMedicalRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthAnalyticsServiceClient) GetMedicalRecord(ctx context.Context, in *GetMedicalRecordRequest, opts ...grpc.CallOption) (*GetMedicalRecordResponse, error) {
	out := new(GetMedicalRecordResponse)
	err := c.cc.Invoke(ctx, "/health_analytics.HealthAnalyticsService/GetMedicalRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthAnalyticsServiceClient) UpdateMedicalRecord(ctx context.Context, in *UpdateMedicalRecordRequest, opts ...grpc.CallOption) (*UpdateMedicalRecordResponse, error) {
	out := new(UpdateMedicalRecordResponse)
	err := c.cc.Invoke(ctx, "/health_analytics.HealthAnalyticsService/UpdateMedicalRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthAnalyticsServiceClient) DeleteMedicalRecord(ctx context.Context, in *DeleteMedicalRecordRequest, opts ...grpc.CallOption) (*DeleteMedicalRecordResponse, error) {
	out := new(DeleteMedicalRecordResponse)
	err := c.cc.Invoke(ctx, "/health_analytics.HealthAnalyticsService/DeleteMedicalRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthAnalyticsServiceClient) ListMedicalRecords(ctx context.Context, in *ListMedicalRecordsRequest, opts ...grpc.CallOption) (*ListMedicalRecordsResponse, error) {
	out := new(ListMedicalRecordsResponse)
	err := c.cc.Invoke(ctx, "/health_analytics.HealthAnalyticsService/ListMedicalRecords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthAnalyticsServiceClient) AddLifestyleData(ctx context.Context, in *AddLifestyleDataRequest, opts ...grpc.CallOption) (*AddLifestyleDataResponse, error) {
	out := new(AddLifestyleDataResponse)
	err := c.cc.Invoke(ctx, "/health_analytics.HealthAnalyticsService/AddLifestyleData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthAnalyticsServiceClient) GetLifestyleData(ctx context.Context, in *GetLifestyleDataRequest, opts ...grpc.CallOption) (*GetLifestyleDataResponse, error) {
	out := new(GetLifestyleDataResponse)
	err := c.cc.Invoke(ctx, "/health_analytics.HealthAnalyticsService/GetLifestyleData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthAnalyticsServiceClient) UpdateLifestyleData(ctx context.Context, in *UpdateLifestyleDataRequest, opts ...grpc.CallOption) (*UpdateLifestyleDataResponse, error) {
	out := new(UpdateLifestyleDataResponse)
	err := c.cc.Invoke(ctx, "/health_analytics.HealthAnalyticsService/UpdateLifestyleData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthAnalyticsServiceClient) DeleteLifestyleData(ctx context.Context, in *DeleteLifestyleDataRequest, opts ...grpc.CallOption) (*DeleteLifestyleDataResponse, error) {
	out := new(DeleteLifestyleDataResponse)
	err := c.cc.Invoke(ctx, "/health_analytics.HealthAnalyticsService/DeleteLifestyleData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthAnalyticsServiceClient) AddWearableData(ctx context.Context, in *AddWearableDataRequest, opts ...grpc.CallOption) (*AddWearableDataResponse, error) {
	out := new(AddWearableDataResponse)
	err := c.cc.Invoke(ctx, "/health_analytics.HealthAnalyticsService/AddWearableData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthAnalyticsServiceClient) GetWearableData(ctx context.Context, in *GetWearableDataRequest, opts ...grpc.CallOption) (*GetWearableDataResponse, error) {
	out := new(GetWearableDataResponse)
	err := c.cc.Invoke(ctx, "/health_analytics.HealthAnalyticsService/GetWearableData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthAnalyticsServiceClient) UpdateWearableData(ctx context.Context, in *UpdateWearableDataRequest, opts ...grpc.CallOption) (*UpdateWearableDataResponse, error) {
	out := new(UpdateWearableDataResponse)
	err := c.cc.Invoke(ctx, "/health_analytics.HealthAnalyticsService/UpdateWearableData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthAnalyticsServiceClient) DeleteWearableData(ctx context.Context, in *DeleteWearableDataRequest, opts ...grpc.CallOption) (*DeleteWearableDataResponse, error) {
	out := new(DeleteWearableDataResponse)
	err := c.cc.Invoke(ctx, "/health_analytics.HealthAnalyticsService/DeleteWearableData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthAnalyticsServiceClient) GenerateHealthRecommendations(ctx context.Context, in *GenerateHealthRecommendationsRequest, opts ...grpc.CallOption) (*GenerateHealthRecommendationsResponse, error) {
	out := new(GenerateHealthRecommendationsResponse)
	err := c.cc.Invoke(ctx, "/health_analytics.HealthAnalyticsService/GenerateHealthRecommendations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthAnalyticsServiceClient) GetRealtimeHealthMonitoring(ctx context.Context, in *GetRealtimeHealthMonitoringRequest, opts ...grpc.CallOption) (HealthAnalyticsService_GetRealtimeHealthMonitoringClient, error) {
	stream, err := c.cc.NewStream(ctx, &HealthAnalyticsService_ServiceDesc.Streams[0], "/health_analytics.HealthAnalyticsService/GetRealtimeHealthMonitoring", opts...)
	if err != nil {
		return nil, err
	}
	x := &healthAnalyticsServiceGetRealtimeHealthMonitoringClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type HealthAnalyticsService_GetRealtimeHealthMonitoringClient interface {
	Recv() (*GetRealtimeHealthMonitoringResponse, error)
	grpc.ClientStream
}

type healthAnalyticsServiceGetRealtimeHealthMonitoringClient struct {
	grpc.ClientStream
}

func (x *healthAnalyticsServiceGetRealtimeHealthMonitoringClient) Recv() (*GetRealtimeHealthMonitoringResponse, error) {
	m := new(GetRealtimeHealthMonitoringResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *healthAnalyticsServiceClient) GetDailyHealthSummary(ctx context.Context, in *GetDailyHealthSummaryRequest, opts ...grpc.CallOption) (*GetDailyHealthSummaryResponse, error) {
	out := new(GetDailyHealthSummaryResponse)
	err := c.cc.Invoke(ctx, "/health_analytics.HealthAnalyticsService/GetDailyHealthSummary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthAnalyticsServiceClient) GetWeeklyHealthSummary(ctx context.Context, in *GetWeeklyHealthSummaryRequest, opts ...grpc.CallOption) (*GetWeeklyHealthSummaryResponse, error) {
	out := new(GetWeeklyHealthSummaryResponse)
	err := c.cc.Invoke(ctx, "/health_analytics.HealthAnalyticsService/GetWeeklyHealthSummary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthAnalyticsServiceServer is the server API for HealthAnalyticsService service.
// All implementations must embed UnimplementedHealthAnalyticsServiceServer
// for forward compatibility
type HealthAnalyticsServiceServer interface {
	AddMedicalRecord(context.Context, *AddMedicalRecordRequest) (*AddMedicalRecordResponse, error)
	GetMedicalRecord(context.Context, *GetMedicalRecordRequest) (*GetMedicalRecordResponse, error)
	UpdateMedicalRecord(context.Context, *UpdateMedicalRecordRequest) (*UpdateMedicalRecordResponse, error)
	DeleteMedicalRecord(context.Context, *DeleteMedicalRecordRequest) (*DeleteMedicalRecordResponse, error)
	ListMedicalRecords(context.Context, *ListMedicalRecordsRequest) (*ListMedicalRecordsResponse, error)
	AddLifestyleData(context.Context, *AddLifestyleDataRequest) (*AddLifestyleDataResponse, error)
	GetLifestyleData(context.Context, *GetLifestyleDataRequest) (*GetLifestyleDataResponse, error)
	UpdateLifestyleData(context.Context, *UpdateLifestyleDataRequest) (*UpdateLifestyleDataResponse, error)
	DeleteLifestyleData(context.Context, *DeleteLifestyleDataRequest) (*DeleteLifestyleDataResponse, error)
	AddWearableData(context.Context, *AddWearableDataRequest) (*AddWearableDataResponse, error)
	GetWearableData(context.Context, *GetWearableDataRequest) (*GetWearableDataResponse, error)
	UpdateWearableData(context.Context, *UpdateWearableDataRequest) (*UpdateWearableDataResponse, error)
	DeleteWearableData(context.Context, *DeleteWearableDataRequest) (*DeleteWearableDataResponse, error)
	GenerateHealthRecommendations(context.Context, *GenerateHealthRecommendationsRequest) (*GenerateHealthRecommendationsResponse, error)
	GetRealtimeHealthMonitoring(*GetRealtimeHealthMonitoringRequest, HealthAnalyticsService_GetRealtimeHealthMonitoringServer) error
	GetDailyHealthSummary(context.Context, *GetDailyHealthSummaryRequest) (*GetDailyHealthSummaryResponse, error)
	GetWeeklyHealthSummary(context.Context, *GetWeeklyHealthSummaryRequest) (*GetWeeklyHealthSummaryResponse, error)
	mustEmbedUnimplementedHealthAnalyticsServiceServer()
}

// UnimplementedHealthAnalyticsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHealthAnalyticsServiceServer struct {
}

func (UnimplementedHealthAnalyticsServiceServer) AddMedicalRecord(context.Context, *AddMedicalRecordRequest) (*AddMedicalRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMedicalRecord not implemented")
}
func (UnimplementedHealthAnalyticsServiceServer) GetMedicalRecord(context.Context, *GetMedicalRecordRequest) (*GetMedicalRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMedicalRecord not implemented")
}
func (UnimplementedHealthAnalyticsServiceServer) UpdateMedicalRecord(context.Context, *UpdateMedicalRecordRequest) (*UpdateMedicalRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMedicalRecord not implemented")
}
func (UnimplementedHealthAnalyticsServiceServer) DeleteMedicalRecord(context.Context, *DeleteMedicalRecordRequest) (*DeleteMedicalRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMedicalRecord not implemented")
}
func (UnimplementedHealthAnalyticsServiceServer) ListMedicalRecords(context.Context, *ListMedicalRecordsRequest) (*ListMedicalRecordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMedicalRecords not implemented")
}
func (UnimplementedHealthAnalyticsServiceServer) AddLifestyleData(context.Context, *AddLifestyleDataRequest) (*AddLifestyleDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddLifestyleData not implemented")
}
func (UnimplementedHealthAnalyticsServiceServer) GetLifestyleData(context.Context, *GetLifestyleDataRequest) (*GetLifestyleDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLifestyleData not implemented")
}
func (UnimplementedHealthAnalyticsServiceServer) UpdateLifestyleData(context.Context, *UpdateLifestyleDataRequest) (*UpdateLifestyleDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLifestyleData not implemented")
}
func (UnimplementedHealthAnalyticsServiceServer) DeleteLifestyleData(context.Context, *DeleteLifestyleDataRequest) (*DeleteLifestyleDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLifestyleData not implemented")
}
func (UnimplementedHealthAnalyticsServiceServer) AddWearableData(context.Context, *AddWearableDataRequest) (*AddWearableDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddWearableData not implemented")
}
func (UnimplementedHealthAnalyticsServiceServer) GetWearableData(context.Context, *GetWearableDataRequest) (*GetWearableDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWearableData not implemented")
}
func (UnimplementedHealthAnalyticsServiceServer) UpdateWearableData(context.Context, *UpdateWearableDataRequest) (*UpdateWearableDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWearableData not implemented")
}
func (UnimplementedHealthAnalyticsServiceServer) DeleteWearableData(context.Context, *DeleteWearableDataRequest) (*DeleteWearableDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWearableData not implemented")
}
func (UnimplementedHealthAnalyticsServiceServer) GenerateHealthRecommendations(context.Context, *GenerateHealthRecommendationsRequest) (*GenerateHealthRecommendationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateHealthRecommendations not implemented")
}
func (UnimplementedHealthAnalyticsServiceServer) GetRealtimeHealthMonitoring(*GetRealtimeHealthMonitoringRequest, HealthAnalyticsService_GetRealtimeHealthMonitoringServer) error {
	return status.Errorf(codes.Unimplemented, "method GetRealtimeHealthMonitoring not implemented")
}
func (UnimplementedHealthAnalyticsServiceServer) GetDailyHealthSummary(context.Context, *GetDailyHealthSummaryRequest) (*GetDailyHealthSummaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDailyHealthSummary not implemented")
}
func (UnimplementedHealthAnalyticsServiceServer) GetWeeklyHealthSummary(context.Context, *GetWeeklyHealthSummaryRequest) (*GetWeeklyHealthSummaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWeeklyHealthSummary not implemented")
}
func (UnimplementedHealthAnalyticsServiceServer) mustEmbedUnimplementedHealthAnalyticsServiceServer() {
}

// UnsafeHealthAnalyticsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HealthAnalyticsServiceServer will
// result in compilation errors.
type UnsafeHealthAnalyticsServiceServer interface {
	mustEmbedUnimplementedHealthAnalyticsServiceServer()
}

func RegisterHealthAnalyticsServiceServer(s grpc.ServiceRegistrar, srv HealthAnalyticsServiceServer) {
	s.RegisterService(&HealthAnalyticsService_ServiceDesc, srv)
}

func _HealthAnalyticsService_AddMedicalRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMedicalRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthAnalyticsServiceServer).AddMedicalRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/health_analytics.HealthAnalyticsService/AddMedicalRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthAnalyticsServiceServer).AddMedicalRecord(ctx, req.(*AddMedicalRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthAnalyticsService_GetMedicalRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMedicalRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthAnalyticsServiceServer).GetMedicalRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/health_analytics.HealthAnalyticsService/GetMedicalRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthAnalyticsServiceServer).GetMedicalRecord(ctx, req.(*GetMedicalRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthAnalyticsService_UpdateMedicalRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMedicalRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthAnalyticsServiceServer).UpdateMedicalRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/health_analytics.HealthAnalyticsService/UpdateMedicalRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthAnalyticsServiceServer).UpdateMedicalRecord(ctx, req.(*UpdateMedicalRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthAnalyticsService_DeleteMedicalRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMedicalRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthAnalyticsServiceServer).DeleteMedicalRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/health_analytics.HealthAnalyticsService/DeleteMedicalRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthAnalyticsServiceServer).DeleteMedicalRecord(ctx, req.(*DeleteMedicalRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthAnalyticsService_ListMedicalRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMedicalRecordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthAnalyticsServiceServer).ListMedicalRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/health_analytics.HealthAnalyticsService/ListMedicalRecords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthAnalyticsServiceServer).ListMedicalRecords(ctx, req.(*ListMedicalRecordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthAnalyticsService_AddLifestyleData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddLifestyleDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthAnalyticsServiceServer).AddLifestyleData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/health_analytics.HealthAnalyticsService/AddLifestyleData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthAnalyticsServiceServer).AddLifestyleData(ctx, req.(*AddLifestyleDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthAnalyticsService_GetLifestyleData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLifestyleDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthAnalyticsServiceServer).GetLifestyleData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/health_analytics.HealthAnalyticsService/GetLifestyleData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthAnalyticsServiceServer).GetLifestyleData(ctx, req.(*GetLifestyleDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthAnalyticsService_UpdateLifestyleData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateLifestyleDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthAnalyticsServiceServer).UpdateLifestyleData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/health_analytics.HealthAnalyticsService/UpdateLifestyleData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthAnalyticsServiceServer).UpdateLifestyleData(ctx, req.(*UpdateLifestyleDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthAnalyticsService_DeleteLifestyleData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLifestyleDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthAnalyticsServiceServer).DeleteLifestyleData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/health_analytics.HealthAnalyticsService/DeleteLifestyleData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthAnalyticsServiceServer).DeleteLifestyleData(ctx, req.(*DeleteLifestyleDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthAnalyticsService_AddWearableData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddWearableDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthAnalyticsServiceServer).AddWearableData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/health_analytics.HealthAnalyticsService/AddWearableData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthAnalyticsServiceServer).AddWearableData(ctx, req.(*AddWearableDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthAnalyticsService_GetWearableData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWearableDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthAnalyticsServiceServer).GetWearableData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/health_analytics.HealthAnalyticsService/GetWearableData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthAnalyticsServiceServer).GetWearableData(ctx, req.(*GetWearableDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthAnalyticsService_UpdateWearableData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWearableDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthAnalyticsServiceServer).UpdateWearableData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/health_analytics.HealthAnalyticsService/UpdateWearableData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthAnalyticsServiceServer).UpdateWearableData(ctx, req.(*UpdateWearableDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthAnalyticsService_DeleteWearableData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteWearableDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthAnalyticsServiceServer).DeleteWearableData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/health_analytics.HealthAnalyticsService/DeleteWearableData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthAnalyticsServiceServer).DeleteWearableData(ctx, req.(*DeleteWearableDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthAnalyticsService_GenerateHealthRecommendations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateHealthRecommendationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthAnalyticsServiceServer).GenerateHealthRecommendations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/health_analytics.HealthAnalyticsService/GenerateHealthRecommendations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthAnalyticsServiceServer).GenerateHealthRecommendations(ctx, req.(*GenerateHealthRecommendationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthAnalyticsService_GetRealtimeHealthMonitoring_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetRealtimeHealthMonitoringRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HealthAnalyticsServiceServer).GetRealtimeHealthMonitoring(m, &healthAnalyticsServiceGetRealtimeHealthMonitoringServer{stream})
}

type HealthAnalyticsService_GetRealtimeHealthMonitoringServer interface {
	Send(*GetRealtimeHealthMonitoringResponse) error
	grpc.ServerStream
}

type healthAnalyticsServiceGetRealtimeHealthMonitoringServer struct {
	grpc.ServerStream
}

func (x *healthAnalyticsServiceGetRealtimeHealthMonitoringServer) Send(m *GetRealtimeHealthMonitoringResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _HealthAnalyticsService_GetDailyHealthSummary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDailyHealthSummaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthAnalyticsServiceServer).GetDailyHealthSummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/health_analytics.HealthAnalyticsService/GetDailyHealthSummary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthAnalyticsServiceServer).GetDailyHealthSummary(ctx, req.(*GetDailyHealthSummaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthAnalyticsService_GetWeeklyHealthSummary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWeeklyHealthSummaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthAnalyticsServiceServer).GetWeeklyHealthSummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/health_analytics.HealthAnalyticsService/GetWeeklyHealthSummary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthAnalyticsServiceServer).GetWeeklyHealthSummary(ctx, req.(*GetWeeklyHealthSummaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HealthAnalyticsService_ServiceDesc is the grpc.ServiceDesc for HealthAnalyticsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HealthAnalyticsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "health_analytics.HealthAnalyticsService",
	HandlerType: (*HealthAnalyticsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddMedicalRecord",
			Handler:    _HealthAnalyticsService_AddMedicalRecord_Handler,
		},
		{
			MethodName: "GetMedicalRecord",
			Handler:    _HealthAnalyticsService_GetMedicalRecord_Handler,
		},
		{
			MethodName: "UpdateMedicalRecord",
			Handler:    _HealthAnalyticsService_UpdateMedicalRecord_Handler,
		},
		{
			MethodName: "DeleteMedicalRecord",
			Handler:    _HealthAnalyticsService_DeleteMedicalRecord_Handler,
		},
		{
			MethodName: "ListMedicalRecords",
			Handler:    _HealthAnalyticsService_ListMedicalRecords_Handler,
		},
		{
			MethodName: "AddLifestyleData",
			Handler:    _HealthAnalyticsService_AddLifestyleData_Handler,
		},
		{
			MethodName: "GetLifestyleData",
			Handler:    _HealthAnalyticsService_GetLifestyleData_Handler,
		},
		{
			MethodName: "UpdateLifestyleData",
			Handler:    _HealthAnalyticsService_UpdateLifestyleData_Handler,
		},
		{
			MethodName: "DeleteLifestyleData",
			Handler:    _HealthAnalyticsService_DeleteLifestyleData_Handler,
		},
		{
			MethodName: "AddWearableData",
			Handler:    _HealthAnalyticsService_AddWearableData_Handler,
		},
		{
			MethodName: "GetWearableData",
			Handler:    _HealthAnalyticsService_GetWearableData_Handler,
		},
		{
			MethodName: "UpdateWearableData",
			Handler:    _HealthAnalyticsService_UpdateWearableData_Handler,
		},
		{
			MethodName: "DeleteWearableData",
			Handler:    _HealthAnalyticsService_DeleteWearableData_Handler,
		},
		{
			MethodName: "GenerateHealthRecommendations",
			Handler:    _HealthAnalyticsService_GenerateHealthRecommendations_Handler,
		},
		{
			MethodName: "GetDailyHealthSummary",
			Handler:    _HealthAnalyticsService_GetDailyHealthSummary_Handler,
		},
		{
			MethodName: "GetWeeklyHealthSummary",
			Handler:    _HealthAnalyticsService_GetWeeklyHealthSummary_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetRealtimeHealthMonitoring",
			Handler:       _HealthAnalyticsService_GetRealtimeHealthMonitoring_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "health_analytics.proto",
}
