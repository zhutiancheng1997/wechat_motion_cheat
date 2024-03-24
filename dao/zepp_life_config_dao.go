package dao

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/zhutiancheng1997/wechat_motion_cheat/common/db"
	"github.com/zhutiancheng1997/wechat_motion_cheat/model"
)

type ZeppLifeConfigDao struct {
	MysqlCli *db.DB
}

const (
	selectQuery = `SELECT * FROM zepp_life_user WHERE delete_at = 0 LIMIT ?,?`
)

func (d *ZeppLifeConfigDao) UpdatePushStatus(userID string, opt *model.ZeppLifeDBConfigUserOpt) error {
	subSql := ""
	args := []interface{}{}
	if opt.PushStatus != nil {
		subSql += "push_status = ?"
		args = append(args, *opt.PushStatus)
	}
	if opt.CurrentStepCnt != nil {
		subSql += "current_step_cnt = ?"
		args = append(args, *opt.CurrentStepCnt)
	}
	if opt.PushStatusUpdateAt != nil {
		subSql += "push_status_update_at = ?"
		args = append(args, *opt.PushStatusUpdateAt)
	}
	args = append(args, userID)
	query := fmt.Sprintf("UPDATE zepp_life_db_config_user SET %s WHERE id = ?", subSql)
	stmt, err := d.MysqlCli.Prepare(query)
	if err != nil {
		log.WithError(err).Error("failed to prepare update statement")
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(args...)
	if err != nil {
		log.WithError(err).Error("failed to prepare update statement")
		return err
	}
	return nil
}

func (d *ZeppLifeConfigDao) GetAllUsers(offset, limit int) ([]*model.ZeppLifeDBConfigUser, error) {
	rows, err := d.MysqlCli.Query(selectQuery, offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := make([]*model.ZeppLifeDBConfigUser, 0)
	for rows.Next() {
		var user model.ZeppLifeDBConfigUser
		err = rows.Scan(
			&user.ID,
			&user.CreatorID,
			&user.User,
			&user.Password,
			&user.PushType,
			&user.TargetStepCnt,
			&user.PushStatus,
			&user.PushStatusUpdateAt,
			&user.CreateAt,
			&user.UpdateAt,
			&user.DeleteAt,
		)
		if err != nil {
			log.WithError(err).Error("GetAllUsers err")
			continue
		}
		users = append(users, &user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
