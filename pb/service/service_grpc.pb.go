// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: service/service.proto

package service

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	model "github.com/lancer2672/DandelionServer_Go/pb/model"
	request "github.com/lancer2672/DandelionServer_Go/pb/request"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Dandelion_CreateMovieHistory_FullMethodName = "/pb.Dandelion/CreateMovieHistory"
	Dandelion_GetMovieHistory_FullMethodName    = "/pb.Dandelion/GetMovieHistory"
	Dandelion_GetListGenres_FullMethodName      = "/pb.Dandelion/GetListGenres"
	Dandelion_CreateMovie_FullMethodName        = "/pb.Dandelion/CreateMovie"
	Dandelion_GetListMovies_FullMethodName      = "/pb.Dandelion/GetListMovies"
	Dandelion_GetMovie_FullMethodName           = "/pb.Dandelion/GetMovie"
	Dandelion_GetMoviesByGenre_FullMethodName   = "/pb.Dandelion/GetMoviesByGenre"
	Dandelion_GetMoviesBySerie_FullMethodName   = "/pb.Dandelion/GetMoviesBySerie"
	Dandelion_GetRecentMovies_FullMethodName    = "/pb.Dandelion/GetRecentMovies"
	Dandelion_SearchMovies_FullMethodName       = "/pb.Dandelion/SearchMovies"
	Dandelion_GetVotesByUser_FullMethodName     = "/pb.Dandelion/GetVotesByUser"
)

// DandelionClient is the client API for Dandelion service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DandelionClient interface {
	// movie-history
	CreateMovieHistory(ctx context.Context, in *request.CreateMovieHistoryRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetMovieHistory(ctx context.Context, in *request.GetMovieHistoryRequest, opts ...grpc.CallOption) (*request.GetMovieHistoryResponse, error)
	// genre
	GetListGenres(ctx context.Context, in *request.GetListGenresRequest, opts ...grpc.CallOption) (*request.GetListGenresResponse, error)
	// movie
	CreateMovie(ctx context.Context, in *request.CreateMovieRequest, opts ...grpc.CallOption) (*model.Movie, error)
	GetListMovies(ctx context.Context, in *request.GetListMoviesRequest, opts ...grpc.CallOption) (*request.GetListMoviesResponse, error)
	GetMovie(ctx context.Context, in *request.GetMovieRequest, opts ...grpc.CallOption) (*model.Movie, error)
	GetMoviesByGenre(ctx context.Context, in *request.GetMoviesByGenreRequest, opts ...grpc.CallOption) (*request.GetMoviesByGenreResponse, error)
	GetMoviesBySerie(ctx context.Context, in *request.GetMoviesBySerieRequest, opts ...grpc.CallOption) (*request.GetMoviesBySerieResponse, error)
	GetRecentMovies(ctx context.Context, in *request.GetRecentMoviesRequest, opts ...grpc.CallOption) (*request.GetRecentMoviesResponse, error)
	SearchMovies(ctx context.Context, in *request.SearchMoviesRequest, opts ...grpc.CallOption) (*request.SearchMoviesResponse, error)
	// Vote
	GetVotesByUser(ctx context.Context, in *request.GetVotesByUserRequest, opts ...grpc.CallOption) (*request.GetVotesByUserResponse, error)
}

type dandelionClient struct {
	cc grpc.ClientConnInterface
}

func NewDandelionClient(cc grpc.ClientConnInterface) DandelionClient {
	return &dandelionClient{cc}
}

func (c *dandelionClient) CreateMovieHistory(ctx context.Context, in *request.CreateMovieHistoryRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Dandelion_CreateMovieHistory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dandelionClient) GetMovieHistory(ctx context.Context, in *request.GetMovieHistoryRequest, opts ...grpc.CallOption) (*request.GetMovieHistoryResponse, error) {
	out := new(request.GetMovieHistoryResponse)
	err := c.cc.Invoke(ctx, Dandelion_GetMovieHistory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dandelionClient) GetListGenres(ctx context.Context, in *request.GetListGenresRequest, opts ...grpc.CallOption) (*request.GetListGenresResponse, error) {
	out := new(request.GetListGenresResponse)
	err := c.cc.Invoke(ctx, Dandelion_GetListGenres_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dandelionClient) CreateMovie(ctx context.Context, in *request.CreateMovieRequest, opts ...grpc.CallOption) (*model.Movie, error) {
	out := new(model.Movie)
	err := c.cc.Invoke(ctx, Dandelion_CreateMovie_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dandelionClient) GetListMovies(ctx context.Context, in *request.GetListMoviesRequest, opts ...grpc.CallOption) (*request.GetListMoviesResponse, error) {
	out := new(request.GetListMoviesResponse)
	err := c.cc.Invoke(ctx, Dandelion_GetListMovies_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dandelionClient) GetMovie(ctx context.Context, in *request.GetMovieRequest, opts ...grpc.CallOption) (*model.Movie, error) {
	out := new(model.Movie)
	err := c.cc.Invoke(ctx, Dandelion_GetMovie_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dandelionClient) GetMoviesByGenre(ctx context.Context, in *request.GetMoviesByGenreRequest, opts ...grpc.CallOption) (*request.GetMoviesByGenreResponse, error) {
	out := new(request.GetMoviesByGenreResponse)
	err := c.cc.Invoke(ctx, Dandelion_GetMoviesByGenre_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dandelionClient) GetMoviesBySerie(ctx context.Context, in *request.GetMoviesBySerieRequest, opts ...grpc.CallOption) (*request.GetMoviesBySerieResponse, error) {
	out := new(request.GetMoviesBySerieResponse)
	err := c.cc.Invoke(ctx, Dandelion_GetMoviesBySerie_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dandelionClient) GetRecentMovies(ctx context.Context, in *request.GetRecentMoviesRequest, opts ...grpc.CallOption) (*request.GetRecentMoviesResponse, error) {
	out := new(request.GetRecentMoviesResponse)
	err := c.cc.Invoke(ctx, Dandelion_GetRecentMovies_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dandelionClient) SearchMovies(ctx context.Context, in *request.SearchMoviesRequest, opts ...grpc.CallOption) (*request.SearchMoviesResponse, error) {
	out := new(request.SearchMoviesResponse)
	err := c.cc.Invoke(ctx, Dandelion_SearchMovies_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dandelionClient) GetVotesByUser(ctx context.Context, in *request.GetVotesByUserRequest, opts ...grpc.CallOption) (*request.GetVotesByUserResponse, error) {
	out := new(request.GetVotesByUserResponse)
	err := c.cc.Invoke(ctx, Dandelion_GetVotesByUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DandelionServer is the server API for Dandelion service.
// All implementations must embed UnimplementedDandelionServer
// for forward compatibility
type DandelionServer interface {
	// movie-history
	CreateMovieHistory(context.Context, *request.CreateMovieHistoryRequest) (*empty.Empty, error)
	GetMovieHistory(context.Context, *request.GetMovieHistoryRequest) (*request.GetMovieHistoryResponse, error)
	// genre
	GetListGenres(context.Context, *request.GetListGenresRequest) (*request.GetListGenresResponse, error)
	// movie
	CreateMovie(context.Context, *request.CreateMovieRequest) (*model.Movie, error)
	GetListMovies(context.Context, *request.GetListMoviesRequest) (*request.GetListMoviesResponse, error)
	GetMovie(context.Context, *request.GetMovieRequest) (*model.Movie, error)
	GetMoviesByGenre(context.Context, *request.GetMoviesByGenreRequest) (*request.GetMoviesByGenreResponse, error)
	GetMoviesBySerie(context.Context, *request.GetMoviesBySerieRequest) (*request.GetMoviesBySerieResponse, error)
	GetRecentMovies(context.Context, *request.GetRecentMoviesRequest) (*request.GetRecentMoviesResponse, error)
	SearchMovies(context.Context, *request.SearchMoviesRequest) (*request.SearchMoviesResponse, error)
	// Vote
	GetVotesByUser(context.Context, *request.GetVotesByUserRequest) (*request.GetVotesByUserResponse, error)
	mustEmbedUnimplementedDandelionServer()
}

// UnimplementedDandelionServer must be embedded to have forward compatible implementations.
type UnimplementedDandelionServer struct {
}

func (UnimplementedDandelionServer) CreateMovieHistory(context.Context, *request.CreateMovieHistoryRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMovieHistory not implemented")
}
func (UnimplementedDandelionServer) GetMovieHistory(context.Context, *request.GetMovieHistoryRequest) (*request.GetMovieHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMovieHistory not implemented")
}
func (UnimplementedDandelionServer) GetListGenres(context.Context, *request.GetListGenresRequest) (*request.GetListGenresResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListGenres not implemented")
}
func (UnimplementedDandelionServer) CreateMovie(context.Context, *request.CreateMovieRequest) (*model.Movie, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMovie not implemented")
}
func (UnimplementedDandelionServer) GetListMovies(context.Context, *request.GetListMoviesRequest) (*request.GetListMoviesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListMovies not implemented")
}
func (UnimplementedDandelionServer) GetMovie(context.Context, *request.GetMovieRequest) (*model.Movie, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMovie not implemented")
}
func (UnimplementedDandelionServer) GetMoviesByGenre(context.Context, *request.GetMoviesByGenreRequest) (*request.GetMoviesByGenreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMoviesByGenre not implemented")
}
func (UnimplementedDandelionServer) GetMoviesBySerie(context.Context, *request.GetMoviesBySerieRequest) (*request.GetMoviesBySerieResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMoviesBySerie not implemented")
}
func (UnimplementedDandelionServer) GetRecentMovies(context.Context, *request.GetRecentMoviesRequest) (*request.GetRecentMoviesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecentMovies not implemented")
}
func (UnimplementedDandelionServer) SearchMovies(context.Context, *request.SearchMoviesRequest) (*request.SearchMoviesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchMovies not implemented")
}
func (UnimplementedDandelionServer) GetVotesByUser(context.Context, *request.GetVotesByUserRequest) (*request.GetVotesByUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVotesByUser not implemented")
}
func (UnimplementedDandelionServer) mustEmbedUnimplementedDandelionServer() {}

// UnsafeDandelionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DandelionServer will
// result in compilation errors.
type UnsafeDandelionServer interface {
	mustEmbedUnimplementedDandelionServer()
}

func RegisterDandelionServer(s grpc.ServiceRegistrar, srv DandelionServer) {
	s.RegisterService(&Dandelion_ServiceDesc, srv)
}

func _Dandelion_CreateMovieHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.CreateMovieHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DandelionServer).CreateMovieHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Dandelion_CreateMovieHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DandelionServer).CreateMovieHistory(ctx, req.(*request.CreateMovieHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dandelion_GetMovieHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.GetMovieHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DandelionServer).GetMovieHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Dandelion_GetMovieHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DandelionServer).GetMovieHistory(ctx, req.(*request.GetMovieHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dandelion_GetListGenres_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.GetListGenresRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DandelionServer).GetListGenres(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Dandelion_GetListGenres_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DandelionServer).GetListGenres(ctx, req.(*request.GetListGenresRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dandelion_CreateMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.CreateMovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DandelionServer).CreateMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Dandelion_CreateMovie_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DandelionServer).CreateMovie(ctx, req.(*request.CreateMovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dandelion_GetListMovies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.GetListMoviesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DandelionServer).GetListMovies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Dandelion_GetListMovies_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DandelionServer).GetListMovies(ctx, req.(*request.GetListMoviesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dandelion_GetMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.GetMovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DandelionServer).GetMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Dandelion_GetMovie_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DandelionServer).GetMovie(ctx, req.(*request.GetMovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dandelion_GetMoviesByGenre_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.GetMoviesByGenreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DandelionServer).GetMoviesByGenre(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Dandelion_GetMoviesByGenre_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DandelionServer).GetMoviesByGenre(ctx, req.(*request.GetMoviesByGenreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dandelion_GetMoviesBySerie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.GetMoviesBySerieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DandelionServer).GetMoviesBySerie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Dandelion_GetMoviesBySerie_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DandelionServer).GetMoviesBySerie(ctx, req.(*request.GetMoviesBySerieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dandelion_GetRecentMovies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.GetRecentMoviesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DandelionServer).GetRecentMovies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Dandelion_GetRecentMovies_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DandelionServer).GetRecentMovies(ctx, req.(*request.GetRecentMoviesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dandelion_SearchMovies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.SearchMoviesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DandelionServer).SearchMovies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Dandelion_SearchMovies_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DandelionServer).SearchMovies(ctx, req.(*request.SearchMoviesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dandelion_GetVotesByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.GetVotesByUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DandelionServer).GetVotesByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Dandelion_GetVotesByUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DandelionServer).GetVotesByUser(ctx, req.(*request.GetVotesByUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Dandelion_ServiceDesc is the grpc.ServiceDesc for Dandelion service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Dandelion_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Dandelion",
	HandlerType: (*DandelionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMovieHistory",
			Handler:    _Dandelion_CreateMovieHistory_Handler,
		},
		{
			MethodName: "GetMovieHistory",
			Handler:    _Dandelion_GetMovieHistory_Handler,
		},
		{
			MethodName: "GetListGenres",
			Handler:    _Dandelion_GetListGenres_Handler,
		},
		{
			MethodName: "CreateMovie",
			Handler:    _Dandelion_CreateMovie_Handler,
		},
		{
			MethodName: "GetListMovies",
			Handler:    _Dandelion_GetListMovies_Handler,
		},
		{
			MethodName: "GetMovie",
			Handler:    _Dandelion_GetMovie_Handler,
		},
		{
			MethodName: "GetMoviesByGenre",
			Handler:    _Dandelion_GetMoviesByGenre_Handler,
		},
		{
			MethodName: "GetMoviesBySerie",
			Handler:    _Dandelion_GetMoviesBySerie_Handler,
		},
		{
			MethodName: "GetRecentMovies",
			Handler:    _Dandelion_GetRecentMovies_Handler,
		},
		{
			MethodName: "SearchMovies",
			Handler:    _Dandelion_SearchMovies_Handler,
		},
		{
			MethodName: "GetVotesByUser",
			Handler:    _Dandelion_GetVotesByUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/service.proto",
}
