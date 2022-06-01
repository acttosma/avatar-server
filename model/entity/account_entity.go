package entity

import (
	"avatarmeta.cc/avatar/pkg/util/logger"
	"avatarmeta.cc/avatar/pkg/util/mysql"
)

type AccountSource int8
type AccountStatus int8

const (
	// Normal
	ACCOUNT_STATUS_NORMAL = 0
	// Forbidden
	ACCOUNT_STATUS_FORBIDDEN = 1
)

type Account struct {
	Id           int64         `gorm:"primary_key"`
	Name         string        `gorm:"type:varchar(20);not null;comment:name or nickname of account"`
	Mail         string        `gorm:"type:varchar(100);not null;index:uidx_mail,unique;comment:email address of account"`
	Mobile       string        `gorm:"type:varchar(20);not null;index:idx_mobile;comment:mobile number"`
	Password     string        `gorm:"type:varchar(50);comment:password,md5 digested"`
	PasswordSalt string        `gorm:"type:varchar(20);comment:password salt,mixed with password"`
	InviteCode   string        `gorm:"type:varchar(20);index:uidx_iCode,unique;comment:invite code of account,unique"`
	InviterId    int64         `gorm:"type:bigint;not null;index:idx_inviter_id;comment:the inviter account id"`
	Status       AccountStatus `gorm:"type:int;not null;default:0;comment:status of account"`

	BaseEntity
}

func (a Account) CreateTableIfNeeded() bool {
	db := mysql.Helper.Db
	err := db.AutoMigrate(a)
	return err == nil
}

func (a Account) Add() (*Account, error) {
	db := mysql.Helper.Db
	err := db.Create(&a).Error
	if err != nil {
		logger.Monitor.Errorf("method entity.Account.Add, error:%+v", err)
		return nil, err
	}

	return &a, nil
}

func (a Account) FindById(id int64) (*Account, error) {
	db := mysql.Helper.Db
	err := db.First(&a, "id = ?", id).Error
	if err != nil {
		logger.Monitor.Errorf("Error when finding account by id:%d, error:%+v", id, err)
		return nil, err
	}

	return &a, nil
}

func (a Account) FindByInviteCode(inviteCode string) (*Account, error) {
	db := mysql.Helper.Db
	err := db.First(&a, "inviter_code = ?", inviteCode).Error
	if err != nil {
		logger.Monitor.Errorf("Error when finding account by inviterCode:%s, error:%+v", inviteCode, err)
		return nil, err
	}

	return &a, nil
}

func (a Account) FindByMail(mail string) (*Account, error) {
	db := mysql.Helper.Db
	err := db.First(&a, "mail = ?", mail).Error
	if err != nil {
		logger.Monitor.Errorf("Error when finding account by mail:%s, error:%+v", mail, err)
		return nil, err
	}

	return &a, nil
}

func (a Account) Save() error {
	db := mysql.Helper.Db
	err := db.Model(&a).Updates(a).Error
	if err != nil {
		logger.Monitor.Errorf("Error when saving account:%+v, error:%+v", a, err)
		return err
	}

	return nil
}
