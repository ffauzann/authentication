package grpc

import (
	"context"
	"regexp"

	"github.com/ffauzann/authentication/internal/constant"
	"github.com/ffauzann/authentication/internal/model"
	"github.com/ffauzann/authentication/internal/util"
	"github.com/ffauzann/authentication/proto"
)

func (s *srv) IsUserExist(ctx context.Context, req *proto.IsUserExistRequest) (res *proto.IsUserExistResponse, err error) {
	if err = validateIsUserExist(req); err != nil {
		return
	}

	result, err := s.service.IsUserExist(ctx, &model.IsUserExistRequest{
		Username:    req.GetUsername(),
		Email:       req.GetEmail(),
		PhoneNumber: req.GetPhoneNumber(),
	})
	if err != nil {
		util.Log().Error(err.Error())
		return
	}

	res = &proto.IsUserExistResponse{
		IsExist: result.IsExist,
		Reasons: result.Reasons,
	}

	return
}

func validateIsUserExist(req *proto.IsUserExistRequest) error {
	if req.Email != nil {
		regexEmail := regexp.MustCompile(constant.RegexEmail)
		if !regexEmail.MatchString(req.GetEmail()) {
			return constant.ErrMalformedEmail
		}
	}

	if len(req.GetPhoneNumber()) < 4 {
		req.PhoneNumber = nil
	}

	if len(req.GetUsername()) < 4 {
		req.Username = nil
	}

	if req.Email == nil && req.PhoneNumber == nil && req.Username == nil {
		return constant.ErrNoArg
	}

	return nil
}
