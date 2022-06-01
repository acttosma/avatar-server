package entity

import (
	"avatarmeta.cc/avatar/pkg/util/logger"
	"avatarmeta.cc/avatar/pkg/util/mysql"
)

type UserInfo struct {
	Id        int64  `gorm:"primary_key"`
	AccountId int64  `gorm:"type:bigint;not null;index:udx_act_id,unique;comment:the account id"`
	Mail      string `gorm:"type:varchar(100);not null;index:idx_mail;comment:email address of account"`
	Nickname  string `gorm:"type:varchar(20);not null;comment:nickname"`
	Avatar    string `gorm:"type:varchar(200);not null;default:'';comment:avatar icon url"`
	Intro     string `gorm:"type:varchar(200);not null;default:'';comment:self introduction"`

	BaseEntity
}

func (a UserInfo) CreateTableIfNeeded() bool {
	db := mysql.Helper.Db
	err := db.AutoMigrate(a)
	return err == nil
}

func (a UserInfo) Add() (*UserInfo, error) {
	db := mysql.Helper.Db
	err := db.Create(&a).Error
	if err != nil {
		logger.Monitor.Errorf("method entity.UserInfo.Add, error:%+v", err)
		return nil, err
	}

	return &a, nil
}

func (a UserInfo) FindById(id int64) (*UserInfo, error) {
	db := mysql.Helper.Db
	err := db.First(&a, "id = ?", id).Error
	if err != nil {
		logger.Monitor.Errorf("Error when finding user info by id:%d, error:%+v", id, err)
		return nil, err
	}

	return &a, nil
}

func (a UserInfo) FindByAccountId(accountId int64) (*UserInfo, error) {
	db := mysql.Helper.Db
	err := db.First(&a, "account_id = ?", accountId).Error
	if err != nil {
		logger.Monitor.Errorf("Error when finding user info by accountId:%d, error:%+v", accountId, err)
		return nil, err
	}

	return &a, nil
}

func (a UserInfo) Save() (*UserInfo, error) {
	db := mysql.Helper.Db
	err := db.Model(&a).Updates(a).Error
	if err != nil {
		logger.Monitor.Errorf("Error when saving user info:%+v, error:%+v", a, err)
		return nil, err
	}

	return &a, nil
}
