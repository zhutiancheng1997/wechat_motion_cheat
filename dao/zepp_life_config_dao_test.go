package dao

import (
	"github.com/zhutiancheng1997/wechat_motion_cheat/common/db"
	"github.com/zhutiancheng1997/wechat_motion_cheat/config"
	"testing"
)

func TestGetAll(t *testing.T) {
	err := db.InitDB(&config.MySQLConfig{
		User:     "root",
		Host:     "127.0.0.1",
		Port:     3306,
		Password: "12345678",
		Schema:   "wx_step",
	})

	if err != nil {
		t.Fatal(err)
	}

	dao := ZeppLifeConfigDao{MysqlCli: db.GetMysqlClient()}
	users, err := dao.GetAllUsers(1, 0)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(users)

}
