package entity

import (
	"avatarmeta.cc/avatar/pkg/util/logger"
	"avatarmeta.cc/avatar/pkg/util/mysql"
	"github.com/shopspring/decimal"
)

type UserBalance struct {
	Id        int64           `gorm:"primary_key"`
	AccountId int64           `gorm:"type:bigint;not null;index:udx_act_id,unique;comment:the account id"`
	Amt       decimal.Decimal `gorm:"type:decimal(10,6);not null;default:0.0"`
	FrozenAmt decimal.Decimal `gorm:"type:decimal(10,6);not null;default:0.0"`
	Score     decimal.Decimal `gorm:"type:decimal(10,6);not null;default:0.0"`

	BaseEntity
}

func (a UserBalance) CreateTableIfNeeded() bool {
	db := mysql.Helper.Db
	err := db.AutoMigrate(a)
	return err == nil
}

func (a UserBalance) Add() (*UserBalance, error) {
	db := mysql.Helper.Db
	err := db.Create(&a).Error
	if err != nil {
		logger.Monitor.Errorf("method entity.UserBalance.Add, error:%+v", err)
		return nil, err
	}

	return &a, nil
}

func (a UserBalance) FindById(id int64) (*UserBalance, error) {
	db := mysql.Helper.Db
	err := db.First(&a, "id = ?", id).Error
	if err != nil {
		logger.Monitor.Errorf("Error when finding user info by id:%d, error:%+v", id, err)
		return nil, err
	}

	return &a, nil
}

func (a UserBalance) FindByAccountId(accountId int64) (*UserBalance, error) {
	db := mysql.Helper.Db
	err := db.First(&a, "account_id = ?", accountId).Error
	if err != nil {
		logger.Monitor.Errorf("Error when finding user info by accountId:%d, error:%+v", accountId, err)
		return nil, err
	}

	return &a, nil
}

func (a UserBalance) Save() (*UserBalance, error) {
	db := mysql.Helper.Db
	err := db.Model(&a).Updates(a).Error
	if err != nil {
		logger.Monitor.Errorf("Error when saving user info:%+v, error:%+v", a, err)
		return nil, err
	}

	return &a, nil
}
