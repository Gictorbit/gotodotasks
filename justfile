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

# format and lint and generate proto file using buf
proto:
    rm -rf api/gen
    @echo "run proto linter..."
    @cd api && buf lint && cd -

    @echo "format proto..."
    @cd api && buf format -w && cd -

    @echo "generate proto..."
    @cd api && buf generate && cd -

build: clean
     go build -o ./bin/gotodotasks -ldflags="-s -w" github.com/Gictorbit/gotodotasks/cmd
run: build
    ./bin/gotodotasks --lgr --database "postgres://pgdbadmin:12345678@127.0.0.1:9475/postgresdb?sslmode=disable" taskmanager
clean:
     @[ -d "./bin" ] && rm -r ./bin && echo "bin directory cleaned" || true