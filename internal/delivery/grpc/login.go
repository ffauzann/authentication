package grpc

import (
	"context"

	"github.com/ffauzann/authentication/internal/constant"
	"github.com/ffauzann/authentication/internal/model"
	"github.com/ffauzann/authentication/internal/util"
	"github.com/ffauzann/authentication/proto"
)

func validateLogin(req *proto.LoginRequest) error {
	if len(req.UserId) < 6 || len(req.Password) < 8 {
		return constant.ErrInvalidUsernamePassword
	}
	return nil
}

func (s *srv) Login(ctx context.Context, req *proto.LoginRequest) (res *proto.LoginResponse, err error) {
	if err = validateLogin(req); err != nil {
		return
	}

	var method constant.LoginMethod
	switch req.Method {
	case proto.LoginMethod_METHOD_LOGIN:
		method = constant.LoginMethodLogin
	case proto.LoginMethod_METHOD_RECOVERY:
		method = constant.LoginMethodRecovery
	default:
		err = constant.ErrInvalidMethod
		return
	}

	result, err := s.service.Login(ctx, &model.LoginRequest{
		UserId:      req.UserId,
		Password:    req.Password,
		LoginMethod: method,
		DeviceInfo: model.DeviceInfo{
			Id:   req.DeviceInfo.Id,
			Name: req.DeviceInfo.GetName(),
			OSInfo: model.OSInfo{
				Name:    req.DeviceInfo.GetOsInfo().GetName(),
				Version: req.DeviceInfo.GetOsInfo().GetVersion(),
			},
		},
	})
	if err != nil {
		util.Log().Error(err.Error())
		return
	}

	res = &proto.LoginResponse{
		AccessToken:  result.Token.AccessToken,
		RefreshToken: result.Token.RefreshToken,
	}

	return
}
