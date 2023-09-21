package repository

import (
	"context"
	"database/sql"
	"fmt"
	"hangoutin/authentication/internal/constant"
	"hangoutin/authentication/internal/model"
	"hangoutin/authentication/internal/util"
	"strings"
)

func (r *dbRepository) CreateUser(ctx context.Context, user *model.User) (err error) {
	query := `
	INSERT INTO 
		users(name, username, email, phone_number, password, master_password)
		VALUES(:name, :username, :email, :phone_number, :password, :master_password)
	`

	arg := map[string]interface{}{
		"name":            user.Name,
		"username":        user.Username,
		"email":           user.Email,
		"phone_number":    user.PhoneNumber,
		"password":        user.Password,
		"master_password": user.MasterPassword,
	}

	result, err := r.db.NamedExecContext(ctx, query, arg)
	if err != nil {
		util.Log().Error(err.Error())
		return
	}

	id, _ := result.LastInsertId()
	user.Id = uint64(id)

	return
}

func (r *dbRepository) IsUserExist(ctx context.Context, userIdType constant.UserIdType, userIdVal string) (isExist bool, err error) {
	if err = userIdType.Validate(); err != nil {
		return
	}

	var count int
	query := fmt.Sprintf("SELECT COUNT(1) FROM users WHERE %s = ?", userIdType)
	if err = r.db.QueryRowContext(ctx, query, userIdVal).Scan(&count); err != nil {
		util.Log().Error(err.Error())
		return
	}

	return count > 0, nil
}

// GetUserByOneOfIdentifier returns user data if there's any match with one the username/email/phone_number values.
func (r *dbRepository) GetUserByOneOfIdentifier(ctx context.Context, val string) (user *model.User, err error) {
	user = new(model.User)
	query := `SELECT 
			u.id,
			u.name,
			u.username,
			u.email,
			u.phone_number,
			u.password,
			u.master_password,
			u.is_blocked,
			GROUP_CONCAT(r.name) roles
		FROM users u
		LEFT JOIN user_roles ur ON u.id = ur.user_id
		LEFT JOIN roles r ON r.id = ur.role_id 
		WHERE 
			(email = ? OR username = ? OR phone_number = ?)
			AND deleted_at IS NULL
		GROUP BY u.id
		LIMIT 1`

	var roles string
	if err = r.db.QueryRowContext(ctx, query, val, val, val).Scan(
		&user.Id,
		&user.Name,
		&user.Username,
		&user.Email,
		&user.PhoneNumber,
		&user.Password,
		&user.MasterPassword,
		&user.IsBlocked,
		&roles,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, constant.ErrUserNotFound
		}
		util.Log().Error(err.Error())
		return
	}

	user.Roles = strings.Split(roles, ",")

	return
}
