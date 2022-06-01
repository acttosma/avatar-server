package entity

import (
	"avatarmeta.cc/avatar/pkg/util/logger"
	"avatarmeta.cc/avatar/pkg/util/mysql"
	"github.com/shopspring/decimal"
	"time"
)

type ScoreFlowType int8

const (
	SCORE_FLOW_TYPE_INVITER             ScoreFlowType = 1
	SCORE_FLOW_TYPE_EXCHANGE_AVATAR     ScoreFlowType = 2
	SCORE_FLOW_TYPE_SUB_EXCHANGE_AVATAR ScoreFlowType = 3
	SCORE_FLOW_TYPE_UNKNOWN             ScoreFlowType = -99
)

type ScoreFlow struct {
	Id          int64           `gorm:"primary_key"`
	AccountId   int64           `gorm:"type:bigint;not null;index:idx_accountId"`
	Amount      decimal.Decimal `gorm:"type:decimal(10,6);not null;default:0.0"`
	Type        ScoreFlowType   `gorm:"type:int;not null;default:0;comment:1-INVITER,2-EXCHANGE_AVATAR,3-SUB_EXCHANGE_AVATAR"`
	Description string          `gorm:"type:varchar(200);not null;default:''"`
	Status      int8            `gorm:"type:int;not null;default:0"`

	BaseEntity
}

func (af ScoreFlow) CreateTableIfNeeded() bool {
	db := mysql.Helper.Db
	err := db.AutoMigrate(af)
	return err == nil
}

func (af ScoreFlow) Add() (*ScoreFlow, error) {
	db := mysql.Helper.Db
	err := db.Create(&af).Error
	if err != nil {
		logger.Monitor.Errorf("method entity.ScoreFlow.Add, error:%+v", err)
		return nil, err
	}

	return &af, nil
}

func (af ScoreFlow) FindByAccountId(actId, preId int64, idCmpSymbol string, size int) ([]ScoreFlow, error) {
	db := mysql.Helper.Db
	var afs []ScoreFlow
	err := db.Order("id DESC").Limit(size).Find(&afs, "account_id = ? AND id "+idCmpSymbol+" ?", actId, preId).Error
	if err != nil {
		return nil, err
	}

	return afs, nil
}

func (af ScoreFlow) FindByAccountIdWithTimeZone(actId int64, startTime, endTime time.Time) ([]ScoreFlow, error) {
	db := mysql.Helper.Db
	var afs []ScoreFlow
	err := db.Find(&afs, "account_id = ? AND created_at BETWEEN ? AND ?", actId, startTime, endTime).Error
	if err != nil {
		return nil, err
	}

	return afs, nil
}
