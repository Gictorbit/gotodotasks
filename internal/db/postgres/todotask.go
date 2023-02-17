package postgres

import (
	"context"
	sqlMaker "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type TodoTaskDataBase struct {
	rawConn    *pgxpool.Pool
	sqlBuilder sqlMaker.StatementBuilderType
}

type TodoTaskDBInterface interface {
	RawConn() *pgxpool.Pool
	GetSQLBuilder() sqlMaker.StatementBuilderType
}

func (ttd *TodoTaskDataBase) RawConn() *pgxpool.Pool {
	return ttd.rawConn
}

func (ttd *TodoTaskDataBase) GetSQLBuilder() sqlMaker.StatementBuilderType {
	ttd.sqlBuilder = sqlMaker.StatementBuilder.PlaceholderFormat(sqlMaker.Dollar)
	return ttd.sqlBuilder
}

func NewTodoTaskDB(conn *pgxpool.Pool) *TodoTaskDataBase {
	return &TodoTaskDataBase{
		rawConn:    conn,
		sqlBuilder: sqlMaker.StatementBuilder.PlaceholderFormat(sqlMaker.Dollar),
	}
}

func NewTodoTaskFromDsn(databaseURL string) (*TodoTaskDataBase, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	rawConn, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		return nil, err
	}
	return &TodoTaskDataBase{
		rawConn:    rawConn,
		sqlBuilder: sqlMaker.StatementBuilder.PlaceholderFormat(sqlMaker.Dollar),
	}, nil
}
