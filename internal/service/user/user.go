package user

import (
	userpb "github.com/Gictorbit/gotodotasks/api/gen/proto/user/v1"
	"github.com/Gictorbit/gotodotasks/internal/authutil"
	userdb "github.com/Gictorbit/gotodotasks/internal/db/postgres/user"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	logger *zap.Logger
	dbConn userdb.UserDBInterface
	auth   *authutil.AuthManager
	userpb.UnimplementedUserServiceServer
}

var (
	ErrInternalError = status.Error(codes.Internal, "internal error")
)

// NewUserService returns new NewUserService
func NewUserService(logger *zap.Logger, dbConn userdb.UserDBInterface, auth *authutil.AuthManager) *UserService {
	return &UserService{
		logger: logger,
		dbConn: dbConn,
		auth:   auth,
	}
}
