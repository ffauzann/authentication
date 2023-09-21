package repository

import (
	"context"
	"fmt"
	"hangoutin/authentication/internal/util"
	"strings"
)

func (r *dbRepository) BatchAssignRoles(ctx context.Context, userId uint64, roleIds []uint8) (err error) {
	if len(roleIds) == 0 {
		return
	}

	// Prepare values to be inserted
	values := []string{}
	for i := range roleIds {
		values = append(values, fmt.Sprintf("(%d, %d)", userId, roleIds[i]))
	}

	// Exec query
	query := fmt.Sprintf("INSERT INTO user_roles(user_id, role_id) VALUES %s", strings.Join(values, ","))
	_, err = r.db.ExecContext(ctx, query)
	if err != nil {
		util.Log().Error(err.Error())
		return
	}

	return
}
