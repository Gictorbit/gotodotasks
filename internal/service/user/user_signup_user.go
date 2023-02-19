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

// Signup signs up a new user and returns jwt token
func (us *UserService) Signup(ctx context.Context, req *userpb.SignupRequest) (*userpb.SignupResponse, error) {
	if err := us.ValidateSignupUser(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	userFound, err := us.dbConn.GetUserByEmail(ctx, req.Email)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		us.logger.Error("failed to get user by email",
			zap.Error(err),
			zap.String("email", req.Email),
		)
		return nil, ErrInternalError
	}
	if userFound != nil {
		return nil, status.Error(codes.AlreadyExists, "user already exists with this email")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		us.logger.Error("error hash password",
			zap.Error(err),
			zap.String("password", req.Password),
		)
		return nil, ErrInternalError
	}
	user := &userpb.User{
		Email:    req.Email,
		Name:     req.Name,
		Password: string(hashedPassword),
		Roles:    []userpb.UserRole{userpb.UserRole_USER_ROLE_NORMAL_USER},
	}
	if e := us.dbConn.CreateUser(ctx, user); e != nil {
		us.logger.Error("failed to create new user",
			zap.Error(e),
			zap.String("email", req.Email),
		)
		return nil, ErrInternalError
	}
	token, err := us.auth.NewToken(user)
	if err != nil {
		us.logger.Error("create token failed",
			zap.Error(err),
			zap.String("email", req.Email),
		)
		return nil, ErrInternalError
	}
	return &userpb.SignupResponse{
		Token: token,
	}, nil
}
