package service

import (
	"context"
	"sync"

	"github.com/ffauzann/authentication/internal/constant"
	"github.com/ffauzann/authentication/internal/model"
	"github.com/ffauzann/authentication/internal/util"
)

func (s *service) Register(ctx context.Context, req *model.RegisterRequest) (res *model.RegisterResponse, err error) {
	// Validate user existence
	isUserExist, err := s.IsUserExist(ctx, &model.IsUserExistRequest{
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Username:    req.Username,
	})
	if err != nil {
		util.Log().Error(err.Error())
		return
	}

	if isUserExist.IsExist {
		res = &model.RegisterResponse{
			StatusCode: constant.RSCFailed,
			Reasons:    isUserExist.Reasons,
		}
		return
	}

	// Begin to register new user
	var user *model.User
	if user, err = s.createUser(ctx, req); err != nil {
		util.Log().Error(err.Error())
		return
	}

	// Currently only support for user
	if err = s.repository.db.BatchAssignRoles(ctx, user.Id, []uint8{constant.RoleIdUser}); err != nil {
		util.Log().Error(err.Error())
		return
	}

	res = &model.RegisterResponse{
		StatusCode: constant.RSCSucceed,
	}

	return
}

func (s *service) createUser(ctx context.Context, req *model.RegisterRequest) (user *model.User, err error) {
	// Prepare concurrent for hashing due it could take quite some times
	var wg sync.WaitGroup
	chErr := make(chan error, 2)
	fnHash := func(pwd string, pwdType uint8) {
		defer wg.Done()
		hashedPassword, err := util.HashPassword(pwd)
		if err != nil {
			chErr <- err
			return
		}

		switch pwdType {
		case constant.MasterPasswordType:
			req.MasterPassword = hashedPassword
		case constant.UserPasswordType:
			req.UserPassword = hashedPassword
		}
	}

	// Begin concurrent
	wg.Add(2)
	go fnHash(s.config.Encryption.MasterPassword, constant.MasterPasswordType)
	go fnHash(req.PlainPassword, constant.UserPasswordType)
	wg.Wait()

	// Begin non-blocking read channel
	select {
	case err = <-chErr: // Error occured
		util.Log().Error(err.Error())
		return
	default: // No error, moving on
	}

	user = req.ToUser()
	if err = s.repository.db.CreateUser(ctx, user); err != nil {
		util.Log().Error(err.Error())
		return
	}

	return
}
