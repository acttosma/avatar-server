package entity

import (
	"acttos.com/avatar/pkg/util/logger"
	"acttos.com/avatar/pkg/util/mysql"
	"github.com/shopspring/decimal"
	"time"
)

type AvatarProduct struct {
	Id           int64           `gorm:"primary_key"`
	Name         int64           `gorm:"type:varchar(20);not null;comment:name of avatar"`
	Price        decimal.Decimal `gorm:"type:decimal(8,2);not null;default:0;comment:price of avatar"`
	AMTOutput    decimal.Decimal `gorm:"type:decimal(8,2);not null;default:0;comment:the output of avatar per period"`
	DailyHours   decimal.Decimal `gorm:"type:decimal(8,2);not null;default:0;comment:the hours of day which avatar works,for display"`
	DailySeconds decimal.Decimal `gorm:"type:decimal(8,2);not null;default:0;comment:the seconds of day which avatar work,for calculating"`
	ValidDays    decimal.Decimal `gorm:"type:decimal(8,2);not null;default:0;comment:the valid days of avatar"`
	LimitCount   int32           `gorm:"type:int;not null;default:0;comment:the limit count per avatar the user can hold in the current time"`
	Status       int8            `gorm:"type:int;not null;default:0;comment:the status of avatar"`

	BaseEntity
}

func (mc AvatarProduct) CreateTableIfNeeded() bool {
	db := mysql.Helper.Db
	err := db.AutoMigrate(mc)
	return err == nil
}

func (mc AvatarProduct) Add() (AvatarProduct, error) {
	db := mysql.Helper.Db
	err := db.Create(&mc).Error
	if err != nil {
		logger.Monitor.Errorf("method entity.AvatarProduct.Add, error:%+v", err)
		return mc, err
	}

	return mc, err
}

func (mc AvatarProduct) FindById(id int64) (AvatarProduct, error) {
	db := mysql.Helper.Db
	err := db.First(&mc, "id = ?", id).Error
	if err != nil {
		logger.Monitor.Errorf("Error when finding card with id:%d, error:%+v", id, err)
		return mc, err
	}

	return mc, nil
}

func (mc AvatarProduct) FindByAccountId(accountId int64) ([]AvatarProduct, error) {
	db := mysql.Helper.Db
	var mcs []AvatarProduct
	err := db.Model(&mc).Find(&mcs, "account_id = ?", accountId).Error
	if err != nil {
		logger.Monitor.Errorf("Error when finding card with accountId:%d, error:%+v", accountId, err)
		return nil, err
	}

	return mcs, nil
}

func (mc AvatarProduct) FindByMchIdAndStoreId(mchId, storeId, preId int64, idCmpSymbol string, size int) ([]AvatarProduct, error) {
	db := mysql.Helper.Db
	var mcs []AvatarProduct
	err := db.Order("id DESC").Limit(size).Find(&mcs, "merchant_id = ? AND store_id = ? AND id "+idCmpSymbol+" ?", mchId, storeId, preId).Error
	if err != nil {
		logger.Monitor.Errorf("Error when finding cards with mchId:%d, storeId:%d, preId:%d, error:%+v", mchId, storeId, preId, err)
		return nil, err
	}

	return mcs, nil
}

func (mc AvatarProduct) FindByMchIdStoreIdAndMobile(mchId, storeId int64, mobile string) (AvatarProduct, error) {
	db := mysql.Helper.Db
	err := db.First(&mc, "merchant_id = ? AND store_id = ? AND mobile = ?", mchId, storeId, mobile).Error
	if err != nil {
		logger.Monitor.Errorf("Error when finding cards with mchId:%d, storeId:%d, mobile:%s, error:%+v", mchId, storeId, mobile, err)
		return mc, err
	}

	return mc, nil
}

func (mc AvatarProduct) SearchByMchIdStoreIdAndMobile(mchId, storeId, preId int64, mobileLike, idCmpSymbol string, size int) ([]AvatarProduct, error) {
	db := mysql.Helper.Db
	var mcs []AvatarProduct
	err := db.Order("id DESC").Limit(size).Find(&mcs, "merchant_id = ? AND store_id = ? AND id "+idCmpSymbol+" ? AND mobile like ?", mchId, storeId, preId, mobileLike).Error
	if err != nil {
		logger.Monitor.Errorf("Error when searching cards with mchId:%d, storeId:%d, preId:%d, mobileLike:%s, error:%+v", mchId, storeId, preId, mobileLike, err)
		return nil, err
	}

	return mcs, nil
}

func (mc AvatarProduct) FindByAccountIdAndStoreId(accountId, storeId int64) (AvatarProduct, error) {
	db := mysql.Helper.Db
	err := db.First(&mc, "account_id = ? AND store_id = ?", accountId, storeId).Error
	if err != nil {
		logger.Monitor.Errorf("Error when finding card with accountId:%d, storeId:%d, error:%+v", accountId, storeId, err)
		return mc, err
	}

	return mc, nil
}

func (mc AvatarProduct) CountByMchIdAndStoreIdWithTimeZone(mchId, storeId int64, startTime, endTime time.Time) (int64, error) {
	db := mysql.Helper.Db
	var count int64
	err := db.Model(&mc).Where("merchant_id = ? AND store_id = ? AND created_at BETWEEN ? AND ?", mchId, storeId, startTime, endTime).Count(&count).Error
	if err != nil {
		logger.Monitor.Errorf("Error when counting member card number with mchId:%d, storeId:%d, error:%+v", mchId, storeId, err)
		return 0, err
	}

	return count, nil
}
