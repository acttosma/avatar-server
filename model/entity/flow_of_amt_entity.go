package entity

import (
	"avatarmeta.cc/avatar/pkg/util/logger"
	"avatarmeta.cc/avatar/pkg/util/mysql"
	"github.com/shopspring/decimal"
	"time"
)

type AmtFlowType int8

const (
	AMT_FLOW_TYPE_BUY             AmtFlowType = 1
	AMT_FLOW_TYPE_SELL            AmtFlowType = 2
	AMT_FLOW_TYPE_EXCHANGE_AVATAR AmtFlowType = 3
	AMT_FLOW_TYPE_RENEW_AVATAR    AmtFlowType = 4
	AMT_FLOW_TYPE_COLLECT         AmtFlowType = 5
	AMT_FLOW_TYPE_UNKNOWN         AmtFlowType = -99
)

type FlowOfAmt struct {
	Id          int64           `gorm:"primary_key"`
	AccountId   int64           `gorm:"type:bigint;not null;index:idx_accountId"`
	Amount      decimal.Decimal `gorm:"type:decimal(10,6);not null;default:0.0"`
	Type        AmtFlowType     `gorm:"type:int;not null;default:0;comment:1-BUY,2-SELL,3-EXCHANGE_AVATAR,4-RENEW_AVATAR,5-COLLECT"`
	Description string          `gorm:"type:varchar(200);not null;default:''"`
	Status      int8            `gorm:"type:int;not null;default:0"`

	BaseEntity
}

func (af FlowOfAmt) CreateTableIfNeeded() bool {
	db := mysql.Helper.Db
	err := db.AutoMigrate(af)
	return err == nil
}

func (af FlowOfAmt) Add() (*FlowOfAmt, error) {
	db := mysql.Helper.Db
	err := db.Create(&af).Error
	if err != nil {
		logger.Monitor.Errorf("method entity.FlowOfAmt.Add, error:%+v", err)
		return nil, err
	}

	return &af, nil
}

func (af FlowOfAmt) FindByAccountId(actId, preId int64, idCmpSymbol string, size int) ([]FlowOfAmt, error) {
	db := mysql.Helper.Db
	var afs []FlowOfAmt
	err := db.Order("id DESC").Limit(size).Find(&afs, "account_id = ? AND id "+idCmpSymbol+" ?", actId, preId).Error
	if err != nil {
		return nil, err
	}

	return afs, nil
}

func (af FlowOfAmt) FindByAccountIdWithTimeZone(actId int64, startTime, endTime time.Time) ([]FlowOfAmt, error) {
	db := mysql.Helper.Db
	var afs []FlowOfAmt
	err := db.Find(&afs, "account_id = ? AND created_at BETWEEN ? AND ?", actId, startTime, endTime).Error
	if err != nil {
		return nil, err
	}

	return afs, nil
}
