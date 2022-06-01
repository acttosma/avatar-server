package entity

import (
	"avatarmeta.cc/avatar/pkg/util/logger"
	"avatarmeta.cc/avatar/pkg/util/mysql"
	"time"
)

type StubOfCollecting struct {
	Id        int64      `gorm:"primary_key"`
	AccountId int64      `gorm:"type:bigint;not null;index:idx_accountId"`
	StartedAt *time.Time `gorm:"type:datetime(3);index:idx_start_time;default:CURRENT_TIMESTAMP(3);comment:the collecting start time"`

	BaseEntity
}

func (af StubOfCollecting) CreateTableIfNeeded() bool {
	db := mysql.Helper.Db
	err := db.AutoMigrate(af)
	return err == nil
}

func (af StubOfCollecting) Add() (*StubOfCollecting, error) {
	db := mysql.Helper.Db
	err := db.Create(&af).Error
	if err != nil {
		logger.Monitor.Errorf("method entity.StubOfCollecting.Add, error:%+v", err)
		return nil, err
	}

	return &af, nil
}

func (af StubOfCollecting) FindByAccountId(actId, preId int64, idCmpSymbol string, size int) ([]StubOfCollecting, error) {
	db := mysql.Helper.Db
	var afs []StubOfCollecting
	err := db.Order("id DESC").Limit(size).Find(&afs, "account_id = ? AND id "+idCmpSymbol+" ?", actId, preId).Error
	if err != nil {
		return nil, err
	}

	return afs, nil
}

func (af StubOfCollecting) FindByAccountIdWithTimeZone(actId int64, startTime, endTime time.Time) ([]StubOfCollecting, error) {
	db := mysql.Helper.Db
	var afs []StubOfCollecting
	err := db.Find(&afs, "account_id = ? AND created_at BETWEEN ? AND ?", actId, startTime, endTime).Error
	if err != nil {
		return nil, err
	}

	return afs, nil
}
