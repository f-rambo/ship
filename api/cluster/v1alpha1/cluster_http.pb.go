// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.0
// - protoc             v5.27.1
// source: api/cluster/v1alpha1/cluster.proto

package v1alpha1

import (
	context "context"
	common "github.com/f-rambo/ship/api/common"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationClusterInterfaceGet = "/cluster.v1alpha1.ClusterInterface/Get"
const OperationClusterInterfaceGetRegions = "/cluster.v1alpha1.ClusterInterface/GetRegions"
const OperationClusterInterfaceList = "/cluster.v1alpha1.ClusterInterface/List"
const OperationClusterInterfacePing = "/cluster.v1alpha1.ClusterInterface/Ping"

type ClusterInterfaceHTTPServer interface {
	Get(context.Context, *ClusterArgs) (*Cluster, error)
	GetRegions(context.Context, *ClusterArgs) (*Regions, error)
	List(context.Context, *emptypb.Empty) (*ClusterList, error)
	Ping(context.Context, *emptypb.Empty) (*common.Msg, error)
}

func RegisterClusterInterfaceHTTPServer(s *http.Server, srv ClusterInterfaceHTTPServer) {
	r := s.Route("/")
	r.GET("/api/v1alpha1/cluster/ping", _ClusterInterface_Ping0_HTTP_Handler(srv))
	r.GET("/api/v1alpha1/cluster", _ClusterInterface_Get0_HTTP_Handler(srv))
	r.GET("/api/v1alpha1/cluster/list", _ClusterInterface_List0_HTTP_Handler(srv))
	r.GET("/api/v1alpha1/cluster/regions", _ClusterInterface_GetRegions0_HTTP_Handler(srv))
}

func _ClusterInterface_Ping0_HTTP_Handler(srv ClusterInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationClusterInterfacePing)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Ping(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*common.Msg)
		return ctx.Result(200, reply)
	}
}

func _ClusterInterface_Get0_HTTP_Handler(srv ClusterInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ClusterArgs
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationClusterInterfaceGet)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Get(ctx, req.(*ClusterArgs))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Cluster)
		return ctx.Result(200, reply)
	}
}

func _ClusterInterface_List0_HTTP_Handler(srv ClusterInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationClusterInterfaceList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.List(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ClusterList)
		return ctx.Result(200, reply)
	}
}

func _ClusterInterface_GetRegions0_HTTP_Handler(srv ClusterInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ClusterArgs
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationClusterInterfaceGetRegions)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetRegions(ctx, req.(*ClusterArgs))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Regions)
		return ctx.Result(200, reply)
	}
}

type ClusterInterfaceHTTPClient interface {
	Get(ctx context.Context, req *ClusterArgs, opts ...http.CallOption) (rsp *Cluster, err error)
	GetRegions(ctx context.Context, req *ClusterArgs, opts ...http.CallOption) (rsp *Regions, err error)
	List(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *ClusterList, err error)
	Ping(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *common.Msg, err error)
}

type ClusterInterfaceHTTPClientImpl struct {
	cc *http.Client
}

func NewClusterInterfaceHTTPClient(client *http.Client) ClusterInterfaceHTTPClient {
	return &ClusterInterfaceHTTPClientImpl{client}
}

func (c *ClusterInterfaceHTTPClientImpl) Get(ctx context.Context, in *ClusterArgs, opts ...http.CallOption) (*Cluster, error) {
	var out Cluster
	pattern := "/api/v1alpha1/cluster"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationClusterInterfaceGet))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *ClusterInterfaceHTTPClientImpl) GetRegions(ctx context.Context, in *ClusterArgs, opts ...http.CallOption) (*Regions, error) {
	var out Regions
	pattern := "/api/v1alpha1/cluster/regions"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationClusterInterfaceGetRegions))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *ClusterInterfaceHTTPClientImpl) List(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*ClusterList, error) {
	var out ClusterList
	pattern := "/api/v1alpha1/cluster/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationClusterInterfaceList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *ClusterInterfaceHTTPClientImpl) Ping(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*common.Msg, error) {
	var out common.Msg
	pattern := "/api/v1alpha1/cluster/ping"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationClusterInterfacePing))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
