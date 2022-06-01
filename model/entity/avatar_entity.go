package entity

import (
	"avatarmeta.cc/avatar/pkg/util/logger"
	"avatarmeta.cc/avatar/pkg/util/mysql"
	"github.com/shopspring/decimal"
	"time"
)

type Avatar struct {
	Id           int64           `gorm:"primary_key"`
	Name         int64           `gorm:"type:varchar(20);not null;comment:name of avatar"`
	Price        decimal.Decimal `gorm:"type:decimal(10,6);not null;default:0;comment:price of avatar"`
	AMTOutput    decimal.Decimal `gorm:"type:decimal(10,6);not null;default:0;comment:the output of avatar per period"`
	DailyHours   decimal.Decimal `gorm:"type:decimal(10,6);not null;default:0;comment:the hours of day which avatar works,for display"`
	DailySeconds decimal.Decimal `gorm:"type:decimal(10,6);not null;default:0;comment:the seconds of day which avatar work,for calculating"`
	ValidDays    decimal.Decimal `gorm:"type:decimal(10,6);not null;default:0;comment:the valid days of avatar"`
	LimitCount   int32           `gorm:"type:int;not null;default:0;comment:the limit count per avatar the user can hold in the current time"`
	Status       int8            `gorm:"type:int;not null;default:0;comment:the status of avatar"`

	BaseEntity
}

func (mc Avatar) CreateTableIfNeeded() bool {
	db := mysql.Helper.Db
	err := db.AutoMigrate(mc)
	return err == nil
}

func (mc Avatar) Add() (Avatar, error) {
	db := mysql.Helper.Db
	err := db.Create(&mc).Error
	if err != nil {
		logger.Monitor.Errorf("method entity.Avatar.Add, error:%+v", err)
		return mc, err
	}

	return mc, err
}

func (mc Avatar) FindById(id int64) (Avatar, error) {
	db := mysql.Helper.Db
	err := db.First(&mc, "id = ?", id).Error
	if err != nil {
		logger.Monitor.Errorf("Error when finding card with id:%d, error:%+v", id, err)
		return mc, err
	}

	return mc, nil
}

func (mc Avatar) FindByAccountId(accountId int64) ([]Avatar, error) {
	db := mysql.Helper.Db
	var mcs []Avatar
	err := db.Model(&mc).Find(&mcs, "account_id = ?", accountId).Error
	if err != nil {
		logger.Monitor.Errorf("Error when finding card with accountId:%d, error:%+v", accountId, err)
		return nil, err
	}

	return mcs, nil
}

func (mc Avatar) FindByMchIdAndStoreId(mchId, storeId, preId int64, idCmpSymbol string, size int) ([]Avatar, error) {
	db := mysql.Helper.Db
	var mcs []Avatar
	err := db.Order("id DESC").Limit(size).Find(&mcs, "merchant_id = ? AND store_id = ? AND id "+idCmpSymbol+" ?", mchId, storeId, preId).Error
	if err != nil {
		logger.Monitor.Errorf("Error when finding cards with mchId:%d, storeId:%d, preId:%d, error:%+v", mchId, storeId, preId, err)
		return nil, err
	}

	return mcs, nil
}

func (mc Avatar) FindByMchIdStoreIdAndMobile(mchId, storeId int64, mobile string) (Avatar, error) {
	db := mysql.Helper.Db
	err := db.First(&mc, "merchant_id = ? AND store_id = ? AND mobile = ?", mchId, storeId, mobile).Error
	if err != nil {
		logger.Monitor.Errorf("Error when finding cards with mchId:%d, storeId:%d, mobile:%s, error:%+v", mchId, storeId, mobile, err)
		return mc, err
	}

	return mc, nil
}

func (mc Avatar) SearchByMchIdStoreIdAndMobile(mchId, storeId, preId int64, mobileLike, idCmpSymbol string, size int) ([]Avatar, error) {
	db := mysql.Helper.Db
	var mcs []Avatar
	err := db.Order("id DESC").Limit(size).Find(&mcs, "merchant_id = ? AND store_id = ? AND id "+idCmpSymbol+" ? AND mobile like ?", mchId, storeId, preId, mobileLike).Error
	if err != nil {
		logger.Monitor.Errorf("Error when searching cards with mchId:%d, storeId:%d, preId:%d, mobileLike:%s, error:%+v", mchId, storeId, preId, mobileLike, err)
		return nil, err
	}

	return mcs, nil
}

func (mc Avatar) FindByAccountIdAndStoreId(accountId, storeId int64) (Avatar, error) {
	db := mysql.Helper.Db
	err := db.First(&mc, "account_id = ? AND store_id = ?", accountId, storeId).Error
	if err != nil {
		logger.Monitor.Errorf("Error when finding card with accountId:%d, storeId:%d, error:%+v", accountId, storeId, err)
		return mc, err
	}

	return mc, nil
}

func (mc Avatar) CountByMchIdAndStoreIdWithTimeZone(mchId, storeId int64, startTime, endTime time.Time) (int64, error) {
	db := mysql.Helper.Db
	var count int64
	err := db.Model(&mc).Where("merchant_id = ? AND store_id = ? AND created_at BETWEEN ? AND ?", mchId, storeId, startTime, endTime).Count(&count).Error
	if err != nil {
		logger.Monitor.Errorf("Error when counting member card number with mchId:%d, storeId:%d, error:%+v", mchId, storeId, err)
		return 0, err
	}

	return count, nil
}
