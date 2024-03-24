package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
	"github.com/zhutiancheng1997/wechat_motion_cheat/common"
	"github.com/zhutiancheng1997/wechat_motion_cheat/model"
	"sync/atomic"
)

type StepHandler struct {
	count uint64
	cron  *cron.Cron
}

func NewStepHandler() *StepHandler {
	handler := &StepHandler{}
	c := cron.New()
	_, _ = c.AddFunc("0 0 * * *", func() {
		atomic.StoreUint64(&handler.count, 0)
	})
	handler.cron = c
	c.Start()
	return handler
}

func (h *StepHandler) initSmt() {

}

func (h *StepHandler) PushStep(c *gin.Context) {
	atomic.AddUint64(&h.count, 1)
	var req = &model.PushStepFEReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		log.WithError(err).Error("pushStep bind json error")
		common.EchoErrorResponse(c, common.ErrorParams)
		return
	}
	tid := uuid.New().String()
	logger := log.WithField("req", req).WithField("trace_id", tid)
	u := model.NewUserWithTrace(req.User, req.Password, tid)
	err := u.CheatOnSteps(req.Step)
	if err != nil {
		logger.WithError(err).Error("cheat step err")
		common.EchoErrorResponse(c, common.NewErrorCode(common.ErrorToCode(err), tid))
		return
	}
	logger.Info("cheat on step success")
	common.EchoSuccessResponse(c, gin.H{"tid": tid})
}

func (h *StepHandler) GetPushStepCount(c *gin.Context) {
	common.EchoSuccessResponse(c, gin.H{
		"cnt": h.count,
	})
}
