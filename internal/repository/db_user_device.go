package repository

import (
	"context"

	"github.com/ffauzann/authentication/internal/model"
	"github.com/ffauzann/authentication/internal/util"
)

func (r *dbRepository) RegisterUserDevice(ctx context.Context, ud *model.UserDevice) (err error) {
	query := `
	INSERT INTO 
		user_devices (user_id, device_id, device_name, device_model, os_name, os_version, last_login)
		VALUES (:user_id, :device_id, :device_name, :device_model, :os_name, :os_version, :last_login)
	`

	arg := map[string]interface{}{
		"user_id":      ud.UserId,
		"device_id":    ud.DeviceId,
		"device_name":  ud.DeviceName,
		"device_model": ud.DeviceModel,
		"os_name":      ud.OSName,
		"os_version":   ud.OSVersion,
		"last_login":   ud.LastLogin,
	}

	result, err := r.db.NamedExecContext(ctx, query, arg)
	if err != nil {
		util.Log().Error(err.Error())
		return
	}

	id, _ := result.LastInsertId()
	ud.Id = uint64(id)
	return
}
