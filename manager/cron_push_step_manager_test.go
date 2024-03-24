package manager

import (
	"github.com/zhutiancheng1997/wechat_motion_cheat/common/db"
	"github.com/zhutiancheng1997/wechat_motion_cheat/config"
	"github.com/zhutiancheng1997/wechat_motion_cheat/dao"
	"testing"
)

func TestCronPushStepManager(t *testing.T) {
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

	d := &dao.ZeppLifeConfigDao{MysqlCli: db.GetMysqlClient()}
	manager := &CronPushStepManager{dao: d}
	manager.run()
}
