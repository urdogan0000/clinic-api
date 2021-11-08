package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"main/pkg/env"
	"os"
)

var (
	DB  *pgxpool.Pool
	Ctx context.Context
	err error
)

func init() {
	println("Driver PostgreSQL Initialized")
	ConnStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s pool_max_conns=%d", os.Getenv("pg_db_host"), env.EnvInt("pg_db_port"), os.Getenv("pg_db_user"), os.Getenv("pg_db_pass"), os.Getenv("pg_db_default"), env.EnvInt("pg_db_max_conn"))

	DB, err = pgxpool.Connect(context.Background(), ConnStr)
	if err != nil {
		println(err.Error())
	} else {
		print("conn ok")
	}

	Ctx = context.Background()
}
