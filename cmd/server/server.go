package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zhutiancheng1997/wechat_motion_cheat/handler"
)

func main() {
	router := gin.Default()

	// Registering the /api/push_step endpoint
	handler := handler.NewStepHandler()
	router.POST("/api/push_step", handler.PushStep)
	router.GET("/api/get_push_step_cnt", handler.GetPushStepCount)

	// Listen and serve on port 3888
	router.Run(":3888")
}
