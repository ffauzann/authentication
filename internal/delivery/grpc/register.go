package grpc

import (
	"context"
	"regexp"
	"unicode"

	"hangoutin/authentication/internal/constant"
	"hangoutin/authentication/internal/model"
	"hangoutin/authentication/internal/util"
	"hangoutin/authentication/proto"
)

func (s *srv) Register(ctx context.Context, req *proto.RegisterRequest) (res *proto.RegisterResponse, err error) {
	if err = validateRegister(req); err != nil {
		return
	}

	var result *model.RegisterResponse
	if result, err = s.service.Register(ctx, &model.RegisterRequest{
		Name: req.Name,

		Email:       req.Email,
		PhoneNumber: req.GetPhoneNumber(),
		Username:    req.GetUsername(),

		PlainPassword: req.Password,
	}); err != nil {
		util.Log().Error(err.Error())
		return
	}

	res = &proto.RegisterResponse{
		Code:    proto.RegisterStatusCode(result.StatusCode),
		Reasons: result.Reasons,
	}

	return
}

func validateRegister(req *proto.RegisterRequest) error {
	regexEmail := regexp.MustCompile(constant.RegexEmail)
	if !regexEmail.MatchString(req.Email) {
		return constant.ErrMalformedEmail
	}

	if !validatePassword(req.Password) {
		return constant.ErrPasswordIsTooWeak
	}

	return nil
}

func validatePassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	var number, upper, special bool
	for _, r := range password {
		switch {
		case unicode.IsNumber(r):
			number = true
		case unicode.IsUpper(r):
			upper = true
		case unicode.IsPunct(r) || unicode.IsSymbol(r):
			special = true
		case unicode.IsLetter(r) || r == ' ':
		}
	}

	return number && upper && special
}
