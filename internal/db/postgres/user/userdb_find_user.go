package user

import (
	"context"
	userpb "github.com/Gictorbit/gotodotasks/api/gen/proto/user/v1"
)

const getUserByEmailQuery = `
	SELECT 
	    id,
	    name,
	    email,
	    password,
	    roles
	FROM 
	    users 
	WHERE 
	    email=$1;
`

// GetUserByEmail returns a user by email
func (udb *UserDataBase) GetUserByEmail(ctx context.Context, email string) (*userpb.User, error) {
	user := &userpb.User{}
	err := udb.RawConn().QueryRow(ctx, getUserByEmailQuery, email).Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Roles,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}
