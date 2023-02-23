// Code generated by Kitex v0.4.4. DO NOT EDIT.

package userserver

import (
	"context"
	douyinuser "douyin-user/idl/douyin_user/kitex_gen/douyinuser"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return userServerServiceInfo
}

var userServerServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "UserServer"
	handlerType := (*douyinuser.UserServer)(nil)
	methods := map[string]kitex.MethodInfo{
		"CreateUser":  kitex.NewMethodInfo(createUserHandler, newUserServerCreateUserArgs, newUserServerCreateUserResult, false),
		"CheckUser":   kitex.NewMethodInfo(checkUserHandler, newUserServerCheckUserArgs, newUserServerCheckUserResult, false),
		"GetUserInfo": kitex.NewMethodInfo(getUserInfoHandler, newUserServerGetUserInfoArgs, newUserServerGetUserInfoResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "douyinfavorite",
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

func createUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*douyinuser.UserServerCreateUserArgs)
	realResult := result.(*douyinuser.UserServerCreateUserResult)
	success, err := handler.(douyinuser.UserServer).CreateUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServerCreateUserArgs() interface{} {
	return douyinuser.NewUserServerCreateUserArgs()
}

func newUserServerCreateUserResult() interface{} {
	return douyinuser.NewUserServerCreateUserResult()
}

func checkUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*douyinuser.UserServerCheckUserArgs)
	realResult := result.(*douyinuser.UserServerCheckUserResult)
	success, err := handler.(douyinuser.UserServer).CheckUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServerCheckUserArgs() interface{} {
	return douyinuser.NewUserServerCheckUserArgs()
}

func newUserServerCheckUserResult() interface{} {
	return douyinuser.NewUserServerCheckUserResult()
}

func getUserInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*douyinuser.UserServerGetUserInfoArgs)
	realResult := result.(*douyinuser.UserServerGetUserInfoResult)
	success, err := handler.(douyinuser.UserServer).GetUserInfo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServerGetUserInfoArgs() interface{} {
	return douyinuser.NewUserServerGetUserInfoArgs()
}

func newUserServerGetUserInfoResult() interface{} {
	return douyinuser.NewUserServerGetUserInfoResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CreateUser(ctx context.Context, req *douyinuser.CreateUserRequest) (r *douyinuser.CreateUserResponse, err error) {
	var _args douyinuser.UserServerCreateUserArgs
	_args.Req = req
	var _result douyinuser.UserServerCreateUserResult
	if err = p.c.Call(ctx, "CreateUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CheckUser(ctx context.Context, req *douyinuser.CheckUserRequest) (r *douyinuser.CheckUserResponse, err error) {
	var _args douyinuser.UserServerCheckUserArgs
	_args.Req = req
	var _result douyinuser.UserServerCheckUserResult
	if err = p.c.Call(ctx, "CheckUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUserInfo(ctx context.Context, req *douyinuser.GetUserInfoRequest) (r *douyinuser.GetUserInfoResponse, err error) {
	var _args douyinuser.UserServerGetUserInfoArgs
	_args.Req = req
	var _result douyinuser.UserServerGetUserInfoResult
	if err = p.c.Call(ctx, "GetUserInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}