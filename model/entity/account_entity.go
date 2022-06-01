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
	Id                int64         `gorm:"primary_key"`
	Mail              string        `gorm:"type:varchar(100);not null;index:udx_mail,unique;comment:email address of account"`
	Password          string        `gorm:"type:varchar(50);comment:password,md5 digested"`
	PasswordSalt      string        `gorm:"type:varchar(20);comment:password salt,mixed with password"`
	TradePassword     string        `gorm:"type:varchar(50);comment:trade password,md5 digested"`
	TradePasswordSalt string        `gorm:"type:varchar(20);comment:trade password salt,mixed with password"`
	InviteCode        string        `gorm:"type:varchar(20);index:uidx_iCode,unique;comment:invite code of account,unique"`
	InviterId         int64         `gorm:"type:bigint;not null;index:idx_inviter_id;comment:the inviter account id"`
	Status            AccountStatus `gorm:"type:int;not null;default:0;comment:status of account"`

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

// Add account record and balance & userinfo record, with transaction supported
func (a Account) Register() (*Account, error) {
	db := mysql.Helper.Db
	tx := db.Begin()
	err := tx.Create(&a).Error
	if err != nil {
		logger.Monitor.Errorf("method entity.Account.Register, error:%+v", err)
		tx.Rollback()
		return nil, err
	}
	// Create extra tables
	balance := UserBalance{
		AccountId: a.Id,
	}
	err = tx.Create(&balance).Error
	if err != nil {
		logger.Monitor.Errorf("method entity.Account.Register.Create balance, error:%+v", err)
		tx.Rollback()
		return nil, err
	}
	info := UserInfo{
		AccountId: a.Id,
		Mail:      a.Mail,
	}
	err = tx.Create(&info).Error
	if err != nil {
		logger.Monitor.Errorf("method entity.Account.Register.Create user info, error:%+v", err)
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

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
	err := db.First(&a, "invite_code = ?", inviteCode).Error
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

func (a Account) Save() (*Account, error) {
	db := mysql.Helper.Db
	err := db.Model(&a).Updates(a).Error
	if err != nil {
		logger.Monitor.Errorf("Error when saving account:%+v, error:%+v", a, err)
		return nil, err
	}

	return &a, nil
}
