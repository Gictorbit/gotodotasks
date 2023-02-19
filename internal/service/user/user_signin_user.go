package user

import (
	"context"
	"errors"
	userpb "github.com/Gictorbit/gotodotasks/api/gen/proto/user/v1"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrInvalidEmailORPassword = status.Error(codes.Unauthenticated, "invalid email or password")
)

// SignIn signs in an exists user by email and password
func (us *UserService) SignIn(ctx context.Context, req *userpb.SignInRequest) (*userpb.SignInResponse, error) {
	if err := us.ValidateSignInUser(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	user, err := us.dbConn.GetUserByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrInvalidEmailORPassword
		}
		us.logger.Error("failed to get user by email",
			zap.Error(err),
			zap.String("email", req.Email),
		)
		return nil, ErrInternalError
	}
	if e := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); e != nil {
		return nil, ErrInvalidEmailORPassword
	}
	token, err := us.auth.GenerateNewToken(user)
	if err != nil {
		us.logger.Error("create token failed",
			zap.Error(err),
			zap.String("email", req.Email),
		)
		return nil, ErrInternalError
	}
	return &userpb.SignInResponse{
		Token: token,
	}, nil
}
