package manager

import (
	"github.com/google/uuid"
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
	"github.com/zhutiancheng1997/wechat_motion_cheat/dao"
	"github.com/zhutiancheng1997/wechat_motion_cheat/model"
	"math/rand"
	"time"
)

const (
	d = 8
)

var (
	mg *CronPushStepManager
)

//todo init

type CronPushStepManager struct {
	dao *dao.ZeppLifeConfigDao `inject:""`
	cron.Cron
}

func InitCronPushStepManager() { //todo
	c := cron.New()
	c.AddFunc("0 */3 * * *", mg.run)
}

func (c *CronPushStepManager) run() {
	jobID := uuid.New().String()
	log.WithField("job_id", jobID).Info("start run")
	limit, offset := 500, 0
	for {
		us, err := c.loadDBUser(offset, limit)
		if err != nil {
			log.WithError(err).Error("load db user error")
			break
		}
		if len(us) == 0 {
			break
		}
		offset += limit
		for _, u := range us {
			tid := uuid.New().String()
			user := model.NewUserWithTrace(u.User, u.Password, tid)
			var currentStepCnt = 0
			if u.TargetStepCnt > 0 {
				currentStepCnt = c.getCurrentStepCnt(u.TargetStepCnt)
			} else {
				currentStepCnt = c.getCurrentStepCnt(28888)
			}

			err := user.CheatOnSteps(currentStepCnt)
			if err != nil {
				log.WithField("trace_id", tid).WithError(err).Error("chet on step error")

				// todo  log & update
			}
			//todo log update

			// rand sleep
			time.Sleep(time.Second * time.Duration(rand.Int63n(3)))
		}
	}
	log.WithField("job_id", jobID).Info("end run")
}

func (c *CronPushStepManager) getCurrentStepCnt(targetStepCnt int) int {
	m := float64(time.Now().Hour()) / 24.0
	if m > 0.75 {
		m = 1
	}

	return int(float64(targetStepCnt) * m)
}

func (c *CronPushStepManager) loadDBUser(offset, limit int) ([]*model.ZeppLifeDBConfigUser, error) {
	return c.dao.GetAllUsers(offset, limit)
}
