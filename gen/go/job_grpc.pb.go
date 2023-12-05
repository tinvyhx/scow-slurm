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

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: job.proto

package gen

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

const (
	JobService_GetJobs_FullMethodName            = "/scow.scheduler_adapter.JobService/GetJobs"
	JobService_GetJobById_FullMethodName         = "/scow.scheduler_adapter.JobService/GetJobById"
	JobService_ChangeJobTimeLimit_FullMethodName = "/scow.scheduler_adapter.JobService/ChangeJobTimeLimit"
	JobService_QueryJobTimeLimit_FullMethodName  = "/scow.scheduler_adapter.JobService/QueryJobTimeLimit"
	JobService_SubmitJob_FullMethodName          = "/scow.scheduler_adapter.JobService/SubmitJob"
	JobService_CancelJob_FullMethodName          = "/scow.scheduler_adapter.JobService/CancelJob"
	JobService_SubmitScriptAsJob_FullMethodName  = "/scow.scheduler_adapter.JobService/SubmitScriptAsJob"
)

// JobServiceClient is the client API for JobService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JobServiceClient interface {
	// description: get jobs with filter options
	// special case:
	// - some of fields not exist, discard them
	GetJobs(ctx context.Context, in *GetJobsRequest, opts ...grpc.CallOption) (*GetJobsResponse, error)
	// description: get job info by id
	// special case:
	// - job id not exist, don't throw
	// - some of fields not exist, discard them
	GetJobById(ctx context.Context, in *GetJobByIdRequest, opts ...grpc.CallOption) (*GetJobByIdResponse, error)
	// description: change a job's time limit
	// errors:
	//   - job not found
	//     NOT_FOUND, JOB_NOT_FOUND, {}
	ChangeJobTimeLimit(ctx context.Context, in *ChangeJobTimeLimitRequest, opts ...grpc.CallOption) (*ChangeJobTimeLimitResponse, error)
	// description: query time limit
	// errors:
	//   - job not found
	//     NOT_FOUND, JOB_NOT_FOUND, {}
	QueryJobTimeLimit(ctx context.Context, in *QueryJobTimeLimitRequest, opts ...grpc.CallOption) (*QueryJobTimeLimitResponse, error)
	// description: submit job
	// errors:
	//   - sbatch failed
	//     UNKNOWN, SBATCH_FAILED, {
	//     reason: string
	//     }
	//   - user not exist
	//     NOT_FOUND, USER_NOT_FOUND, {}
	SubmitJob(ctx context.Context, in *SubmitJobRequest, opts ...grpc.CallOption) (*SubmitJobResponse, error)
	// description: cancel a job
	// errors:
	//   - user not exist
	//     NOT_FOUND, USER_NOT_FOUND, {}
	//   - job not found
	//     NOT_FOUND, JOB_NOT_FOUND, {}
	CancelJob(ctx context.Context, in *CancelJobRequest, opts ...grpc.CallOption) (*CancelJobResponse, error)
	// description: submit a script  as a job
	// errors:
	//   - sbatch failed
	//     UNKNOWN, SBATCH_FAILED, {
	//     reason: string
	//     }
	//   - user not exist
	//     NOT_FOUND, USER_NOT_FOUND, {}
	SubmitScriptAsJob(ctx context.Context, in *SubmitScriptAsJobRequest, opts ...grpc.CallOption) (*SubmitScriptAsJobResponse, error)
}

type jobServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewJobServiceClient(cc grpc.ClientConnInterface) JobServiceClient {
	return &jobServiceClient{cc}
}

func (c *jobServiceClient) GetJobs(ctx context.Context, in *GetJobsRequest, opts ...grpc.CallOption) (*GetJobsResponse, error) {
	out := new(GetJobsResponse)
	err := c.cc.Invoke(ctx, JobService_GetJobs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) GetJobById(ctx context.Context, in *GetJobByIdRequest, opts ...grpc.CallOption) (*GetJobByIdResponse, error) {
	out := new(GetJobByIdResponse)
	err := c.cc.Invoke(ctx, JobService_GetJobById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) ChangeJobTimeLimit(ctx context.Context, in *ChangeJobTimeLimitRequest, opts ...grpc.CallOption) (*ChangeJobTimeLimitResponse, error) {
	out := new(ChangeJobTimeLimitResponse)
	err := c.cc.Invoke(ctx, JobService_ChangeJobTimeLimit_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) QueryJobTimeLimit(ctx context.Context, in *QueryJobTimeLimitRequest, opts ...grpc.CallOption) (*QueryJobTimeLimitResponse, error) {
	out := new(QueryJobTimeLimitResponse)
	err := c.cc.Invoke(ctx, JobService_QueryJobTimeLimit_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) SubmitJob(ctx context.Context, in *SubmitJobRequest, opts ...grpc.CallOption) (*SubmitJobResponse, error) {
	out := new(SubmitJobResponse)
	err := c.cc.Invoke(ctx, JobService_SubmitJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) CancelJob(ctx context.Context, in *CancelJobRequest, opts ...grpc.CallOption) (*CancelJobResponse, error) {
	out := new(CancelJobResponse)
	err := c.cc.Invoke(ctx, JobService_CancelJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) SubmitScriptAsJob(ctx context.Context, in *SubmitScriptAsJobRequest, opts ...grpc.CallOption) (*SubmitScriptAsJobResponse, error) {
	out := new(SubmitScriptAsJobResponse)
	err := c.cc.Invoke(ctx, JobService_SubmitScriptAsJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JobServiceServer is the server API for JobService service.
// All implementations should embed UnimplementedJobServiceServer
// for forward compatibility
type JobServiceServer interface {
	// description: get jobs with filter options
	// special case:
	// - some of fields not exist, discard them
	GetJobs(context.Context, *GetJobsRequest) (*GetJobsResponse, error)
	// description: get job info by id
	// special case:
	// - job id not exist, don't throw
	// - some of fields not exist, discard them
	GetJobById(context.Context, *GetJobByIdRequest) (*GetJobByIdResponse, error)
	// description: change a job's time limit
	// errors:
	//   - job not found
	//     NOT_FOUND, JOB_NOT_FOUND, {}
	ChangeJobTimeLimit(context.Context, *ChangeJobTimeLimitRequest) (*ChangeJobTimeLimitResponse, error)
	// description: query time limit
	// errors:
	//   - job not found
	//     NOT_FOUND, JOB_NOT_FOUND, {}
	QueryJobTimeLimit(context.Context, *QueryJobTimeLimitRequest) (*QueryJobTimeLimitResponse, error)
	// description: submit job
	// errors:
	//   - sbatch failed
	//     UNKNOWN, SBATCH_FAILED, {
	//     reason: string
	//     }
	//   - user not exist
	//     NOT_FOUND, USER_NOT_FOUND, {}
	SubmitJob(context.Context, *SubmitJobRequest) (*SubmitJobResponse, error)
	// description: cancel a job
	// errors:
	//   - user not exist
	//     NOT_FOUND, USER_NOT_FOUND, {}
	//   - job not found
	//     NOT_FOUND, JOB_NOT_FOUND, {}
	CancelJob(context.Context, *CancelJobRequest) (*CancelJobResponse, error)
	// description: submit a script  as a job
	// errors:
	//   - sbatch failed
	//     UNKNOWN, SBATCH_FAILED, {
	//     reason: string
	//     }
	//   - user not exist
	//     NOT_FOUND, USER_NOT_FOUND, {}
	SubmitScriptAsJob(context.Context, *SubmitScriptAsJobRequest) (*SubmitScriptAsJobResponse, error)
}

// UnimplementedJobServiceServer should be embedded to have forward compatible implementations.
type UnimplementedJobServiceServer struct {
}

func (UnimplementedJobServiceServer) GetJobs(context.Context, *GetJobsRequest) (*GetJobsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJobs not implemented")
}
func (UnimplementedJobServiceServer) GetJobById(context.Context, *GetJobByIdRequest) (*GetJobByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJobById not implemented")
}
func (UnimplementedJobServiceServer) ChangeJobTimeLimit(context.Context, *ChangeJobTimeLimitRequest) (*ChangeJobTimeLimitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeJobTimeLimit not implemented")
}
func (UnimplementedJobServiceServer) QueryJobTimeLimit(context.Context, *QueryJobTimeLimitRequest) (*QueryJobTimeLimitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryJobTimeLimit not implemented")
}
func (UnimplementedJobServiceServer) SubmitJob(context.Context, *SubmitJobRequest) (*SubmitJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitJob not implemented")
}
func (UnimplementedJobServiceServer) CancelJob(context.Context, *CancelJobRequest) (*CancelJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelJob not implemented")
}
func (UnimplementedJobServiceServer) SubmitScriptAsJob(context.Context, *SubmitScriptAsJobRequest) (*SubmitScriptAsJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitScriptAsJob not implemented")
}

// UnsafeJobServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JobServiceServer will
// result in compilation errors.
type UnsafeJobServiceServer interface {
	mustEmbedUnimplementedJobServiceServer()
}

func RegisterJobServiceServer(s grpc.ServiceRegistrar, srv JobServiceServer) {
	s.RegisterService(&JobService_ServiceDesc, srv)
}

func _JobService_GetJobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetJobsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).GetJobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobService_GetJobs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).GetJobs(ctx, req.(*GetJobsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_GetJobById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetJobByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).GetJobById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobService_GetJobById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).GetJobById(ctx, req.(*GetJobByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_ChangeJobTimeLimit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeJobTimeLimitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).ChangeJobTimeLimit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobService_ChangeJobTimeLimit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).ChangeJobTimeLimit(ctx, req.(*ChangeJobTimeLimitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_QueryJobTimeLimit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryJobTimeLimitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).QueryJobTimeLimit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobService_QueryJobTimeLimit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).QueryJobTimeLimit(ctx, req.(*QueryJobTimeLimitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_SubmitJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).SubmitJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobService_SubmitJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).SubmitJob(ctx, req.(*SubmitJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_CancelJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).CancelJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobService_CancelJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).CancelJob(ctx, req.(*CancelJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_SubmitScriptAsJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitScriptAsJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).SubmitScriptAsJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobService_SubmitScriptAsJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).SubmitScriptAsJob(ctx, req.(*SubmitScriptAsJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// JobService_ServiceDesc is the grpc.ServiceDesc for JobService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var JobService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "scow.scheduler_adapter.JobService",
	HandlerType: (*JobServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetJobs",
			Handler:    _JobService_GetJobs_Handler,
		},
		{
			MethodName: "GetJobById",
			Handler:    _JobService_GetJobById_Handler,
		},
		{
			MethodName: "ChangeJobTimeLimit",
			Handler:    _JobService_ChangeJobTimeLimit_Handler,
		},
		{
			MethodName: "QueryJobTimeLimit",
			Handler:    _JobService_QueryJobTimeLimit_Handler,
		},
		{
			MethodName: "SubmitJob",
			Handler:    _JobService_SubmitJob_Handler,
		},
		{
			MethodName: "CancelJob",
			Handler:    _JobService_CancelJob_Handler,
		},
		{
			MethodName: "SubmitScriptAsJob",
			Handler:    _JobService_SubmitScriptAsJob_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "job.proto",
}