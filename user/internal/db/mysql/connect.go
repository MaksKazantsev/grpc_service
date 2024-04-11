package mysql

import (
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

func MustConnect(dsn string) *sqlx.DB {
	db, err := sqlx.Open("mysql", dsn+"?multiStatements=true")
	if err != nil {
		panic("failed to connect to db" + err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		panic("failed to ping db" + err.Error())
	}
	// TODO: migrations

	initAdmin(db)

	return db
}

func initAdmin(db *sqlx.DB) {
	if err := godotenv.Load(".env"); err != nil {
		panic("failed to find .env file: " + err.Error())
	}

	adminEmail := os.Getenv("ADMIN_EMAIL")
	panicIfEmpty(adminEmail, "admin email")
	adminPass := os.Getenv("ADMIN_PASS")
	panicIfEmpty(adminPass, "admin password")

	q := `INSERT INTO users (uuid, username, password, permlvl, email, phone_number) VALUES(?,?,?,?,?,?)`

	hashed, err := bcrypt.GenerateFromPassword([]byte(adminPass), bcrypt.DefaultCost)
	if err != nil {
		panic("failed to hash password")
	}

	if err = db.QueryRowx(q, uuid.New, "ADMIN", hashed, "admin", adminEmail, "+79099991389").Err(); err != nil {
		panic("failed to query admin" + err.Error())
	}
}

func panicIfEmpty(value, msg string) {
	if value == "" {
		panic("can not be empty:" + msg)
	}
}
