package entity

import (
	"acttos.com/avatar/pkg/util/logger"
	"acttos.com/avatar/pkg/util/mysql"
)

type AccountSource int8
type AccountStatus int8

const (
	// 自有手机号码注册类型
	ACCOUNT_TYPE_MOBILE AccountSource = 0
	// 微信用户源类型
	ACCOUNT_TYPE_WECHAT AccountSource = 1
	// 支付宝用户源类型
	ACCOUNT_TYPE_ALIPAY AccountSource = 2

	// 状态正常
	ACCOUNT_STATUS_NORMAL = 0
	// 被封禁
	ACCOUNT_STATUS_FORBIDDEN = 1
)

type Account struct {
	Id           int64         `gorm:"primary_key"`
	OpenId       string        `gorm:"type:varchar(50);not null;index:idx_openid;comment:用户在第三方平台上的开放ID"`
	Mobile       string        `gorm:"type:varchar(20);not null;index:uidx_mobile,unique;comment:用户手机号码,唯一,登录用"`
	Password     string        `gorm:"type:varchar(50);comment:密码,摘要信息,http传递的密码信息,与salt共同摘要后为password的值"`
	PasswordSalt string        `gorm:"type:varchar(20);comment:密码的盐,在用户注册、修改密码时会自动生成字串"`
	Avatar       string        `gorm:"type:varchar(200);not null;default:'';comment:用户头像"`
	Type         int8          `gorm:"type:int;not null;default:0;comment:用户类型 0-普通用户"`
	Source       AccountSource `gorm:"type:int;not null;default:0;comment:用户来源 0-普通用户,1-微信用户,2-支付宝用户……"`
	Status       AccountStatus `gorm:"type:int;not null;default:0;comment:用户状态 0-正常,1-封禁..."`

	BaseEntity
}

func (a Account) CreateTableIfNeeded() bool {
	db := mysql.Helper.Db
	err := db.AutoMigrate(a)
	return err == nil
}

func (a Account) Add() (Account, error) {
	db := mysql.Helper.Db
	err := db.Create(&a).Error
	if err != nil {
		logger.Monitor.Errorf("method entity.Account.Add, error:%+v", err)
		return a, err
	}

	return a, nil
}

func (a Account) AddAccountAndMemberCard(mchId, storeId int64) (Account, MemberCard, bool) {
	db := mysql.Helper.Db
	tx := db.Begin()
	err := tx.Create(&a).Error
	if err != nil {
		logger.Monitor.Errorf("Error when adding account, error:%+v", err)
		tx.Rollback()
		return a, MemberCard{}, false
	}
	card := MemberCard{
		AccountId:  a.Id,
		MerchantId: mchId,
		StoreId:    storeId,
		Mobile:     a.Mobile,
	}
	err = tx.Create(&card).Error
	if err != nil {
		logger.Monitor.Errorf("Error when adding member card, error:%+v", err)
		tx.Rollback()
		return a, card, false
	}
	err = tx.Commit().Error
	if err != nil {
		logger.Monitor.Errorf("Error when commiting transaction, error:%+v", err)
		tx.Rollback()
		return a, card, false
	}

	return a, card, true
}

func (a Account) FindById(id int64) (Account, error) {
	db := mysql.Helper.Db
	err := db.First(&a, "id = ?", id).Error
	if err != nil {
		logger.Monitor.Errorf("Error when finding account by id:%d, error:%+v", id, err)
		return a, err
	}

	return a, nil
}

func (a Account) FindByMobile(mobile string) (Account, error) {
	db := mysql.Helper.Db
	err := db.First(&a, "mobile = ?", mobile).Error
	if err != nil {
		return a, err
	}

	return a, nil
}

func (a Account) FindByOpenId(openId string) (Account, error) {
	db := mysql.Helper.Db
	err := db.First(&a, "open_id = ?", openId).Error
	if err != nil {
		logger.Monitor.Errorf("Error occurs when finding account by openId:%s, with error:%+v", openId, err)
		return a, err
	}

	return a, nil
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
