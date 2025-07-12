package main

import (
	_ "github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.Info("test")
	// logrus.Info("Main")
	// m, err := migrate.New("path", "url")
	// migrate.NewWithDatabaseInstance("path", "postgres", &sql.DB{})
	// if err != nil {
	// 	logrus.Info("migrate error")
	// }
	// m.Up()
	// conn, err := sqlx.ConnectContext(context.Background(), "pgx", "postgres://postgres:postgres@127.0.0.1:5432/postgres")
	// if err != nil {
	// 	logrus.Info("connect error ")
	// 	logrus.Info(err.Error())
	// } else {
	// 	logrus.Info("connect not error")
	// }
	// if conn != nil {
	// 	conn.Close()
	// }
}
