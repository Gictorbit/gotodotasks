package user

import (
	"context"
	userpb "github.com/Gictorbit/gotodotasks/api/gen/proto/user/v1"
	sqlMaker "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type UserDataBase struct {
	rawConn    *pgxpool.Pool
	sqlBuilder sqlMaker.StatementBuilderType
}

type UserDBInterface interface {
	RawConn() *pgxpool.Pool
	GetSQLBuilder() sqlMaker.StatementBuilderType
	GetUserByEmail(ctx context.Context, email string) (*userpb.User, error)
	CreateUser(ctx context.Context, user *userpb.User) error
}

var _ UserDBInterface = &UserDataBase{}

func (udb *UserDataBase) RawConn() *pgxpool.Pool {
	return udb.rawConn
}

func (udb *UserDataBase) GetSQLBuilder() sqlMaker.StatementBuilderType {
	udb.sqlBuilder = sqlMaker.StatementBuilder.PlaceholderFormat(sqlMaker.Dollar)
	return udb.sqlBuilder
}

func NewUserDB(conn *pgxpool.Pool) *UserDataBase {
	return &UserDataBase{
		rawConn:    conn,
		sqlBuilder: sqlMaker.StatementBuilder.PlaceholderFormat(sqlMaker.Dollar),
	}
}

func NewUserDBFromDsn(databaseURL string) (*UserDataBase, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	rawConn, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		return nil, err
	}
	return &UserDataBase{
		rawConn:    rawConn,
		sqlBuilder: sqlMaker.StatementBuilder.PlaceholderFormat(sqlMaker.Dollar),
	}, nil
}
