package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"

	"github.com/zhutiancheng1997/wechat_motion_cheat/config"
)

var (
	client *DB
)

func GetMysqlClient() *DB {
	return client
}

type DB struct {
	*sql.DB
	// Add more DAO instances here for other tables
}

func InitDB(cfg *config.MySQLConfig) error {
	// => root:123456@tcp(127.0.0.1:3306)/test
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Schema)
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	client = &DB{db}
	return nil
}

func (db *DB) Close() {
	db.Close()
}
