// Code generated by Kitex v0.4.4. DO NOT EDIT.

package userservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	usermicro "usermicro/kitex_gen/usermicro"
)

func serviceInfo() *kitex.ServiceInfo {
	return userServiceServiceInfo
}

var userServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "UserService"
	handlerType := (*usermicro.UserService)(nil)
	methods := map[string]kitex.MethodInfo{
		"GetUserByNameMethod": kitex.NewMethodInfo(getUserByNameMethodHandler, newUserServiceGetUserByNameMethodArgs, newUserServiceGetUserByNameMethodResult, false),
		"LoginUserMethod":     kitex.NewMethodInfo(loginUserMethodHandler, newUserServiceLoginUserMethodArgs, newUserServiceLoginUserMethodResult, false),
		"CreateUserMethod":    kitex.NewMethodInfo(createUserMethodHandler, newUserServiceCreateUserMethodArgs, newUserServiceCreateUserMethodResult, false),
		"UpdateUserMethod":    kitex.NewMethodInfo(updateUserMethodHandler, newUserServiceUpdateUserMethodArgs, newUserServiceUpdateUserMethodResult, false),
		"FollowUserMethod":    kitex.NewMethodInfo(followUserMethodHandler, newUserServiceFollowUserMethodArgs, newUserServiceFollowUserMethodResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "usermicro",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func getUserByNameMethodHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*usermicro.UserServiceGetUserByNameMethodArgs)
	realResult := result.(*usermicro.UserServiceGetUserByNameMethodResult)
	success, err := handler.(usermicro.UserService).GetUserByNameMethod(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceGetUserByNameMethodArgs() interface{} {
	return usermicro.NewUserServiceGetUserByNameMethodArgs()
}

func newUserServiceGetUserByNameMethodResult() interface{} {
	return usermicro.NewUserServiceGetUserByNameMethodResult()
}

func loginUserMethodHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*usermicro.UserServiceLoginUserMethodArgs)
	realResult := result.(*usermicro.UserServiceLoginUserMethodResult)
	success, err := handler.(usermicro.UserService).LoginUserMethod(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceLoginUserMethodArgs() interface{} {
	return usermicro.NewUserServiceLoginUserMethodArgs()
}

func newUserServiceLoginUserMethodResult() interface{} {
	return usermicro.NewUserServiceLoginUserMethodResult()
}

func createUserMethodHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*usermicro.UserServiceCreateUserMethodArgs)
	realResult := result.(*usermicro.UserServiceCreateUserMethodResult)
	success, err := handler.(usermicro.UserService).CreateUserMethod(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceCreateUserMethodArgs() interface{} {
	return usermicro.NewUserServiceCreateUserMethodArgs()
}

func newUserServiceCreateUserMethodResult() interface{} {
	return usermicro.NewUserServiceCreateUserMethodResult()
}

func updateUserMethodHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*usermicro.UserServiceUpdateUserMethodArgs)
	realResult := result.(*usermicro.UserServiceUpdateUserMethodResult)
	success, err := handler.(usermicro.UserService).UpdateUserMethod(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceUpdateUserMethodArgs() interface{} {
	return usermicro.NewUserServiceUpdateUserMethodArgs()
}

func newUserServiceUpdateUserMethodResult() interface{} {
	return usermicro.NewUserServiceUpdateUserMethodResult()
}

func followUserMethodHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*usermicro.UserServiceFollowUserMethodArgs)
	realResult := result.(*usermicro.UserServiceFollowUserMethodResult)
	success, err := handler.(usermicro.UserService).FollowUserMethod(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceFollowUserMethodArgs() interface{} {
	return usermicro.NewUserServiceFollowUserMethodArgs()
}

func newUserServiceFollowUserMethodResult() interface{} {
	return usermicro.NewUserServiceFollowUserMethodResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) GetUserByNameMethod(ctx context.Context, request *usermicro.GetUserReq) (r *usermicro.GetUserResp, err error) {
	var _args usermicro.UserServiceGetUserByNameMethodArgs
	_args.Request = request
	var _result usermicro.UserServiceGetUserByNameMethodResult
	if err = p.c.Call(ctx, "GetUserByNameMethod", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) LoginUserMethod(ctx context.Context, request *usermicro.LoginUserReq) (r *usermicro.LoginUserResp, err error) {
	var _args usermicro.UserServiceLoginUserMethodArgs
	_args.Request = request
	var _result usermicro.UserServiceLoginUserMethodResult
	if err = p.c.Call(ctx, "LoginUserMethod", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CreateUserMethod(ctx context.Context, request *usermicro.CreateUserReq) (r *usermicro.CreateUserResp, err error) {
	var _args usermicro.UserServiceCreateUserMethodArgs
	_args.Request = request
	var _result usermicro.UserServiceCreateUserMethodResult
	if err = p.c.Call(ctx, "CreateUserMethod", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateUserMethod(ctx context.Context, request *usermicro.UpdateUserReq) (r *usermicro.UpdateUserResp, err error) {
	var _args usermicro.UserServiceUpdateUserMethodArgs
	_args.Request = request
	var _result usermicro.UserServiceUpdateUserMethodResult
	if err = p.c.Call(ctx, "UpdateUserMethod", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FollowUserMethod(ctx context.Context, request *usermicro.FollowUserReq) (r *usermicro.FollowUserResp, err error) {
	var _args usermicro.UserServiceFollowUserMethodArgs
	_args.Request = request
	var _result usermicro.UserServiceFollowUserMethodResult
	if err = p.c.Call(ctx, "FollowUserMethod", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
