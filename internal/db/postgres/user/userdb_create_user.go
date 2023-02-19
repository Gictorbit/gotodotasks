package user

import (
	"context"
	userpb "github.com/Gictorbit/gotodotasks/api/gen/proto/user/v1"
)

const createNewUserQuery = `
	INSERT INTO users(id,name,email,password,roles)
	VALUES(DEFAULT,$1,$2,$3,$4)
	RETURNING id;
`

// CreateUser creates a new user
func (udb *UserDataBase) CreateUser(ctx context.Context, user *userpb.User) error {
	return udb.RawConn().QueryRow(ctx, createNewUserQuery,
		user.Name,
		user.Email,
		user.Password,
		user.Roles,
	).Scan(&user.Id)
}
