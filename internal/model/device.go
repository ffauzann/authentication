package model

import "time"

type DeviceInfo struct {
	Id     string
	Name   string
	Model  string
	OSInfo OSInfo
}

type OSInfo struct {
	Name    string
	Version string
}

type UserDevice struct {
	Id          uint64    `db:"id"`
	UserId      uint64    `db:"user_id"`
	DeviceId    string    `db:"device_id"`
	DeviceName  string    `db:"device_name"`
	DeviceModel string    `db:"device_model"`
	OSName      string    `db:"os_name"`
	OSVersion   string    `db:"os_version"`
	LastLogin   time.Time `db:"last_login"`
	IsRevoked   bool      `db:"is_revoked"`
}
