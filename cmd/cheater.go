package main

import (
	"flag"

	log "github.com/sirupsen/logrus"
	"github.com/zhutiancheng1997/wechat_motion_cheat/model"
)

var (
	user     = flag.String("user", "2505913895@qq.com", "")
	password = flag.String("password", "aa2505913895", "")
	stepCnt  = flag.Int("step", 28888, "")
)

func main() {
	u := model.NewUser(*user, *password)
	err := u.CheatOnSteps(*stepCnt)
	if err != nil {
		log.Error("刷步失败，请检查...")
		return
	}
	log.Info("刷步成功！")
}
