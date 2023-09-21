package service

import (
	"context"
	"time"

	"hangoutin/authentication/internal/constant"
	"hangoutin/authentication/internal/model"
	"hangoutin/authentication/internal/util"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Login(ctx context.Context, req *model.LoginRequest) (res *model.LoginResponse, err error) {
	user, err := s.repository.db.GetUserByOneOfIdentifier(ctx, req.UserId)
	if err != nil {
		util.Log().Error(err.Error())
		return
	}

	if user.IsBlocked {
		return nil, constant.ErrUserIsBlocked
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return nil, constant.ErrInvalidUsernamePassword
		}
		util.Log().Error(err.Error())
		return
	}

	var token model.Token
	// Generate access_token
	token.AccessToken, err = s.generateToken(ctx, user, constant.TokenTypeAccess)
	if err != nil {
		util.Log().Error(err.Error())
		return
	}

	// Generate refresh_token
	token.RefreshToken, err = s.generateToken(ctx, user, constant.TokenTypeRefresh)
	if err != nil {
		util.Log().Error(err.Error())
		return
	}

	err = s.registerUserDevice(ctx, req.ToUserDevice(user.Id), &token)
	if err != nil {
		util.Log().Error(err.Error())
		return
	}

	res = &model.LoginResponse{
		Token: token,
	}

	return
}

func (s *service) registerUserDevice(ctx context.Context, userDevice *model.UserDevice, token *model.Token) (err error) {
	err = s.repository.redis.RegisterUserDevice(ctx, userDevice.DeviceId, token)
	if err != nil {
		util.Log().Error(err.Error())
		return
	}

	err = s.repository.db.RegisterUserDevice(ctx, userDevice)
	if err != nil {
		util.Log().Error(err.Error())
		return
	}

	return
}

func (s *service) generateToken(ctx context.Context, user *model.User, tokenType constant.TokenType) (token string, err error) {
	// Construct base claims.
	claims := model.Claims{
		UserId:      user.Id,
		Name:        user.Name,
		Username:    user.Username.String,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber.String,
		Roles:       user.Roles,
		TokenType:   tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: user.Email,
			Issuer:  s.config.Jwt.Iss,
		},
	}

	// Determine whether to set expiry time.
	// Currently refresh_token doesn't have an exp.
	// Instead, validate device_id.
	switch tokenType {
	case constant.TokenTypeAccess:
		d, err := time.ParseDuration(s.config.Jwt.Exp)
		if err != nil {
			util.Log().Error(err.Error())
			return "", err
		}
		claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(d))
	}

	// Create token & sign.
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString([]byte(s.config.Jwt.SigningKey))
}
