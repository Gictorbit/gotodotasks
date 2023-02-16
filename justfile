composeFile := "docker-compose.yaml"
composeEnvFile := "compose.env"
sqlMigrationPath := "internal/migrations/sqls"
golangMigrationPath := "internal/migrations/golang"

# run docker compose up
dcompose-up:
    @echo "run docker compose up"
    docker-compose -f {{composeFile}} --env-file {{composeEnvFile}} up -d
    @echo "env variables are:"
    @cat compose.env

# stop docker compose containers
dcompose-stop:
    docker-compose -f {{composeFile}} --env-file {{composeEnvFile}} stop

# down and clean all compose file containers
dcompose-clean:
    docker-compose -f {{composeFile}} --env-file {{composeEnvFile}} down --volumes --remove-orphans --rmi local

# create new sql migration
create-sql MigrateName:
    @mkdir -p {{sqlMigrationPath}}
    goose -dir {{sqlMigrationPath}} create {{MigrateName}} sql

# create new golang migration
create-gomigrate MigrateName:
    @mkdir -p {{golangMigrationPath}}
    goose -dir {{golangMigrationPath}} create {{MigrateName}} go

# run all migrations
run-migrations:
    #!/usr/bin/env sh
    [ -d {{sqlMigrationPath}} ] && goose -v -dir {{sqlMigrationPath}} postgres "postgres://pgdbadmin:12345678@127.0.0.1:9475/postgresdb?sslmode=disable" up
    [ -d {{golangMigrationPath}} ] && goose -v -dir {{golangMigrationPath}} postgres "postgres://pgdbadmin:12345678@127.0.0.1:9475/postgresdb?sslmode=disable" up
    echo "all migrations applied"