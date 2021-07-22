// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: rbac.proto

package rbac

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Rbac service

func NewRbacEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Rbac service

type RbacService interface {
	// 创建后台接口规则
	CreateRule(ctx context.Context, in *CreateRuleReq, opts ...client.CallOption) (*EmptyRsp, error)
	// 修改后台接口规则
	UpdateRule(ctx context.Context, in *UpdateRuleReq, opts ...client.CallOption) (*EmptyRsp, error)
	// 查询后台接口规则
	SelectRule(ctx context.Context, in *SelectRuleReq, opts ...client.CallOption) (*SelectRuleRsp, error)
	// 创建后台接口规则分组角色
	CreateAuthGroup(ctx context.Context, in *CreateAuthGroupReq, opts ...client.CallOption) (*EmptyRsp, error)
	// 修改后台接口规则分组角色
	UpdateAuthGroup(ctx context.Context, in *UpdateAuthGroupReq, opts ...client.CallOption) (*EmptyRsp, error)
	// 查询后台接口规则分组角色
	SelectAuthGroup(ctx context.Context, in *SelectAuthGroupReq, opts ...client.CallOption) (*SelectAuthGroupRsp, error)
	// 修改后台接口规则分组角色用户关系
	UpdateAuthGroupAccess(ctx context.Context, in *UpdateAuthGroupAccessReq, opts ...client.CallOption) (*EmptyRsp, error)
	// 校验后台访问权限
	VerifyUserAuth(ctx context.Context, in *VerifyUserAuthReq, opts ...client.CallOption) (*VerifyUserAuthRsp, error)
	// 获取用户前端菜单
	GetUserMenuReq(ctx context.Context, in *GetUserMenuReqReq, opts ...client.CallOption) (*GetUserMenuReqRsp, error)
}

type rbacService struct {
	c    client.Client
	name string
}

func NewRbacService(name string, c client.Client) RbacService {
	return &rbacService{
		c:    c,
		name: name,
	}
}

func (c *rbacService) CreateRule(ctx context.Context, in *CreateRuleReq, opts ...client.CallOption) (*EmptyRsp, error) {
	req := c.c.NewRequest(c.name, "Rbac.CreateRule", in)
	out := new(EmptyRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacService) UpdateRule(ctx context.Context, in *UpdateRuleReq, opts ...client.CallOption) (*EmptyRsp, error) {
	req := c.c.NewRequest(c.name, "Rbac.UpdateRule", in)
	out := new(EmptyRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacService) SelectRule(ctx context.Context, in *SelectRuleReq, opts ...client.CallOption) (*SelectRuleRsp, error) {
	req := c.c.NewRequest(c.name, "Rbac.SelectRule", in)
	out := new(SelectRuleRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacService) CreateAuthGroup(ctx context.Context, in *CreateAuthGroupReq, opts ...client.CallOption) (*EmptyRsp, error) {
	req := c.c.NewRequest(c.name, "Rbac.CreateAuthGroup", in)
	out := new(EmptyRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacService) UpdateAuthGroup(ctx context.Context, in *UpdateAuthGroupReq, opts ...client.CallOption) (*EmptyRsp, error) {
	req := c.c.NewRequest(c.name, "Rbac.UpdateAuthGroup", in)
	out := new(EmptyRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacService) SelectAuthGroup(ctx context.Context, in *SelectAuthGroupReq, opts ...client.CallOption) (*SelectAuthGroupRsp, error) {
	req := c.c.NewRequest(c.name, "Rbac.SelectAuthGroup", in)
	out := new(SelectAuthGroupRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacService) UpdateAuthGroupAccess(ctx context.Context, in *UpdateAuthGroupAccessReq, opts ...client.CallOption) (*EmptyRsp, error) {
	req := c.c.NewRequest(c.name, "Rbac.UpdateAuthGroupAccess", in)
	out := new(EmptyRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacService) VerifyUserAuth(ctx context.Context, in *VerifyUserAuthReq, opts ...client.CallOption) (*VerifyUserAuthRsp, error) {
	req := c.c.NewRequest(c.name, "Rbac.VerifyUserAuth", in)
	out := new(VerifyUserAuthRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacService) GetUserMenuReq(ctx context.Context, in *GetUserMenuReqReq, opts ...client.CallOption) (*GetUserMenuReqRsp, error) {
	req := c.c.NewRequest(c.name, "Rbac.GetUserMenuReq", in)
	out := new(GetUserMenuReqRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Rbac service

type RbacHandler interface {
	// 创建后台接口规则
	CreateRule(context.Context, *CreateRuleReq, *EmptyRsp) error
	// 修改后台接口规则
	UpdateRule(context.Context, *UpdateRuleReq, *EmptyRsp) error
	// 查询后台接口规则
	SelectRule(context.Context, *SelectRuleReq, *SelectRuleRsp) error
	// 创建后台接口规则分组角色
	CreateAuthGroup(context.Context, *CreateAuthGroupReq, *EmptyRsp) error
	// 修改后台接口规则分组角色
	UpdateAuthGroup(context.Context, *UpdateAuthGroupReq, *EmptyRsp) error
	// 查询后台接口规则分组角色
	SelectAuthGroup(context.Context, *SelectAuthGroupReq, *SelectAuthGroupRsp) error
	// 修改后台接口规则分组角色用户关系
	UpdateAuthGroupAccess(context.Context, *UpdateAuthGroupAccessReq, *EmptyRsp) error
	// 校验后台访问权限
	VerifyUserAuth(context.Context, *VerifyUserAuthReq, *VerifyUserAuthRsp) error
	// 获取用户前端菜单
	GetUserMenuReq(context.Context, *GetUserMenuReqReq, *GetUserMenuReqRsp) error
}

func RegisterRbacHandler(s server.Server, hdlr RbacHandler, opts ...server.HandlerOption) error {
	type rbac interface {
		CreateRule(ctx context.Context, in *CreateRuleReq, out *EmptyRsp) error
		UpdateRule(ctx context.Context, in *UpdateRuleReq, out *EmptyRsp) error
		SelectRule(ctx context.Context, in *SelectRuleReq, out *SelectRuleRsp) error
		CreateAuthGroup(ctx context.Context, in *CreateAuthGroupReq, out *EmptyRsp) error
		UpdateAuthGroup(ctx context.Context, in *UpdateAuthGroupReq, out *EmptyRsp) error
		SelectAuthGroup(ctx context.Context, in *SelectAuthGroupReq, out *SelectAuthGroupRsp) error
		UpdateAuthGroupAccess(ctx context.Context, in *UpdateAuthGroupAccessReq, out *EmptyRsp) error
		VerifyUserAuth(ctx context.Context, in *VerifyUserAuthReq, out *VerifyUserAuthRsp) error
		GetUserMenuReq(ctx context.Context, in *GetUserMenuReqReq, out *GetUserMenuReqRsp) error
	}
	type Rbac struct {
		rbac
	}
	h := &rbacHandler{hdlr}
	return s.Handle(s.NewHandler(&Rbac{h}, opts...))
}

type rbacHandler struct {
	RbacHandler
}

func (h *rbacHandler) CreateRule(ctx context.Context, in *CreateRuleReq, out *EmptyRsp) error {
	return h.RbacHandler.CreateRule(ctx, in, out)
}

func (h *rbacHandler) UpdateRule(ctx context.Context, in *UpdateRuleReq, out *EmptyRsp) error {
	return h.RbacHandler.UpdateRule(ctx, in, out)
}

func (h *rbacHandler) SelectRule(ctx context.Context, in *SelectRuleReq, out *SelectRuleRsp) error {
	return h.RbacHandler.SelectRule(ctx, in, out)
}

func (h *rbacHandler) CreateAuthGroup(ctx context.Context, in *CreateAuthGroupReq, out *EmptyRsp) error {
	return h.RbacHandler.CreateAuthGroup(ctx, in, out)
}

func (h *rbacHandler) UpdateAuthGroup(ctx context.Context, in *UpdateAuthGroupReq, out *EmptyRsp) error {
	return h.RbacHandler.UpdateAuthGroup(ctx, in, out)
}

func (h *rbacHandler) SelectAuthGroup(ctx context.Context, in *SelectAuthGroupReq, out *SelectAuthGroupRsp) error {
	return h.RbacHandler.SelectAuthGroup(ctx, in, out)
}

func (h *rbacHandler) UpdateAuthGroupAccess(ctx context.Context, in *UpdateAuthGroupAccessReq, out *EmptyRsp) error {
	return h.RbacHandler.UpdateAuthGroupAccess(ctx, in, out)
}

func (h *rbacHandler) VerifyUserAuth(ctx context.Context, in *VerifyUserAuthReq, out *VerifyUserAuthRsp) error {
	return h.RbacHandler.VerifyUserAuth(ctx, in, out)
}

func (h *rbacHandler) GetUserMenuReq(ctx context.Context, in *GetUserMenuReqReq, out *GetUserMenuReqRsp) error {
	return h.RbacHandler.GetUserMenuReq(ctx, in, out)
}
