package service

import (
	"context"
	"fmt"
	"strings"
	"sync"

	commonUtil "github.com/ffauzann/common/util"

	"github.com/ffauzann/authentication/internal/constant"
	"github.com/ffauzann/authentication/internal/model"
	"github.com/ffauzann/authentication/internal/util"
)

func (s *service) IsUserExist(ctx context.Context, req *model.IsUserExistRequest) (res *model.IsUserExistResponse, err error) {
	// Prepare concurrent
	var wg sync.WaitGroup
	chErr := make(chan error, 3)
	chReason := make(chan string, 3)
	fnIsExist := func(userIdType constant.UserIdType, userIdVal string) {
		defer wg.Done()
		isExist, err := s.repository.db.IsUserExist(ctx, userIdType, userIdVal)
		if err != nil {
			chErr <- err
			return
		}

		if isExist {
			chReason <- fmt.Sprintf("user with %s %s already exist", strings.Replace(string(userIdType), "_", " ", 1), userIdVal)
		}
	}

	// Begin concurrent
	if req.Email != "" {
		wg.Add(1)
		go fnIsExist(constant.UserIdTypeEmail, req.Email)
	}
	if req.PhoneNumber != "" {
		wg.Add(1)
		go fnIsExist(constant.UserIdTypePhoneNumber, req.PhoneNumber)
	}
	if req.Username != "" {
		wg.Add(1)
		go fnIsExist(constant.UserIdTypeUsername, req.Username)
	}
	wg.Wait()

	// Begin non-blocking read channel
	if err = commonUtil.ErrorFromChannel(chErr); err != nil {
		util.Log().Error(err.Error())
		return
	}

	// Read all channel values if any
	close(chReason)
	reasons := []string{}
	for s := range chReason {
		reasons = append(reasons, s)
	}

	// Format response
	res = &model.IsUserExistResponse{
		IsExist: len(reasons) > 0,
		Reasons: reasons,
	}

	return
}
