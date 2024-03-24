package db

import (
	"github.com/zhutiancheng1997/wechat_motion_cheat/config"
	"testing"
)

func TestInitMysql(t *testing.T) {
	err := InitDB(&config.MySQLConfig{
		User:     "root",
		Host:     "127.0.0.1",
		Port:     3306,
		Password: "12345678",
		Schema:   "wx_step",
	})
	if err != nil {
		t.Fatal(err)
	}
}
