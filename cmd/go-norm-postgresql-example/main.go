package main

import (
	"context"

	"github.com/perfectgentlemande/go-norm-postgresql-example/internal/database/dbuser"
	"github.com/perfectgentlemande/go-norm-postgresql-example/internal/logger"
	dbUserMigrations "github.com/perfectgentlemande/go-norm-postgresql-example/migrations/dbuser"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	ctx := context.Background()
	log := logger.DefaultLogger()

	dbUser, err := dbuser.NewDatabase(ctx, "postgres://postgres:postgres@postgres:5432/postgres?sslmode=disable", dbUserMigrations.MigrationAssets)
	if err != nil {
		log.WithError(err).Fatal("cannot create db")
	}
	defer dbUser.Close(ctx)

	err = dbUser.Ping(ctx)
	if err != nil {
		log.WithField("conn_string", *configPath).WithError(err).Error("cannot ping db")
		return
	}
}
