package model

import (
	"hangoutin/authentication/internal/constant"
	"time"
)

type LoginRequest struct {
	UserId      string
	Password    string
	LoginMethod constant.LoginMethod
	DeviceInfo  DeviceInfo
}

func (l *LoginRequest) ToUserDevice(userId uint64) *UserDevice {
	return &UserDevice{
		UserId:      userId,
		DeviceId:    l.DeviceInfo.Id,
		DeviceName:  l.DeviceInfo.Name,
		DeviceModel: l.DeviceInfo.Model,
		OSName:      l.DeviceInfo.OSInfo.Name,
		OSVersion:   l.DeviceInfo.OSInfo.Version,
		LastLogin:   time.Now(),
	}
}

type LoginResponse struct {
	Token    Token
	DeviceId string
}
