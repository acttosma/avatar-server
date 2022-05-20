package entity

import (
	"acttos.com/avatar/pkg/util/logger"
	"acttos.com/avatar/pkg/util/mysql"
	"github.com/shopspring/decimal"
	"time"
)

type MemberCard struct {
	Id            int64           `gorm:"primary_key"`
	AccountId     int64           `gorm:"type:bigint;not null;index:idx_act_id;comment:会员卡所属的账户ID"`
	MerchantId    int64           `gorm:"type:bigint;not null;index:idx_mch_id;comment:会员卡所属的商户ID"`
	StoreId       int64           `gorm:"type:bigint;not null;index:idx_store_id;comment:店铺ID"`
	Mobile        string          `gorm:"type:varchar(20);not null;index:idx_mobile;comment:用户手机号码"`
	Amount        decimal.Decimal `gorm:"type:decimal(8,2);not null;default:0;comment:充值余额。与赠送余额相加得到真实余额"`
	GivenAmount   decimal.Decimal `gorm:"type:decimal(8,2);not null;default:0;comment:赠送余额。与充值余额相加得到真实余额"`
	DepositAmount decimal.Decimal `gorm:"type:decimal(8,2);not null;default:0;comment:押金金额"`
	FrozenAmt     decimal.Decimal `gorm:"type:decimal(8,2);not null;default:0;comment:充值余额的冻结部分,临时不可用"`
	FrozenGamt    decimal.Decimal `gorm:"type:decimal(8,2);not null;default:0;comment:赠送余额的冻结部分,临时不可用"`
	FrozenDamt    decimal.Decimal `gorm:"type:decimal(8,2);not null;default:0;comment:押金金额的冻结部分,临时不可用"`
	Score         int32           `gorm:"type:int;not null;default:0;comment:积分余额"`
	PlayCount     int32           `gorm:"type:int;not null;default:0;comment:消费次数"`
	UsedAmt       decimal.Decimal `gorm:"type:decimal(8,2);not null;default:0;comment:消费的充值金额"`
	UsedGamt      decimal.Decimal `gorm:"type:decimal(8,2);not null;default:0;comment:消费的赠送金额"`
	Status        int8            `gorm:"type:int;not null;default:0;comment:会员卡状态。。。"`

	BaseEntity
}

func (mc MemberCard) CreateTableIfNeeded() bool {
	db := mysql.Helper.Db
	err := db.AutoMigrate(mc)
	return err == nil
}

func (mc MemberCard) Add() (MemberCard, error) {
	db := mysql.Helper.Db
	err := db.Create(&mc).Error
	if err != nil {
		logger.Monitor.Errorf("method entity.MemberCard.Add, error:%+v", err)
		return mc, err
	}

	return mc, err
}

func (mc MemberCard) FindById(id int64) (MemberCard, error) {
	db := mysql.Helper.Db
	err := db.First(&mc, "id = ?", id).Error
	if err != nil {
		logger.Monitor.Errorf("Error when finding card with id:%d, error:%+v", id, err)
		return mc, err
	}

	return mc, nil
}

func (mc MemberCard) FindByAccountId(accountId int64) ([]MemberCard, error) {
	db := mysql.Helper.Db
	var mcs []MemberCard
	err := db.Model(&mc).Find(&mcs, "account_id = ?", accountId).Error
	if err != nil {
		logger.Monitor.Errorf("Error when finding card with accountId:%d, error:%+v", accountId, err)
		return nil, err
	}

	return mcs, nil
}

func (mc MemberCard) FindByMchIdAndStoreId(mchId, storeId, preId int64, idCmpSymbol string, size int) ([]MemberCard, error) {
	db := mysql.Helper.Db
	var mcs []MemberCard
	err := db.Order("id DESC").Limit(size).Find(&mcs, "merchant_id = ? AND store_id = ? AND id "+idCmpSymbol+" ?", mchId, storeId, preId).Error
	if err != nil {
		logger.Monitor.Errorf("Error when finding cards with mchId:%d, storeId:%d, preId:%d, error:%+v", mchId, storeId, preId, err)
		return nil, err
	}

	return mcs, nil
}

func (mc MemberCard) FindByMchIdStoreIdAndMobile(mchId, storeId int64, mobile string) (MemberCard, error) {
	db := mysql.Helper.Db
	err := db.First(&mc, "merchant_id = ? AND store_id = ? AND mobile = ?", mchId, storeId, mobile).Error
	if err != nil {
		logger.Monitor.Errorf("Error when finding cards with mchId:%d, storeId:%d, mobile:%s, error:%+v", mchId, storeId, mobile, err)
		return mc, err
	}

	return mc, nil
}

func (mc MemberCard) SearchByMchIdStoreIdAndMobile(mchId, storeId, preId int64, mobileLike, idCmpSymbol string, size int) ([]MemberCard, error) {
	db := mysql.Helper.Db
	var mcs []MemberCard
	err := db.Order("id DESC").Limit(size).Find(&mcs, "merchant_id = ? AND store_id = ? AND id "+idCmpSymbol+" ? AND mobile like ?", mchId, storeId, preId, mobileLike).Error
	if err != nil {
		logger.Monitor.Errorf("Error when searching cards with mchId:%d, storeId:%d, preId:%d, mobileLike:%s, error:%+v", mchId, storeId, preId, mobileLike, err)
		return nil, err
	}

	return mcs, nil
}

func (mc MemberCard) FindByAccountIdAndStoreId(accountId, storeId int64) (MemberCard, error) {
	db := mysql.Helper.Db
	err := db.First(&mc, "account_id = ? AND store_id = ?", accountId, storeId).Error
	if err != nil {
		logger.Monitor.Errorf("Error when finding card with accountId:%d, storeId:%d, error:%+v", accountId, storeId, err)
		return mc, err
	}

	return mc, nil
}

func (mc MemberCard) CountByMchIdAndStoreIdWithTimeZone(mchId, storeId int64, startTime, endTime time.Time) (int64, error) {
	db := mysql.Helper.Db
	var count int64
	err := db.Model(&mc).Where("merchant_id = ? AND store_id = ? AND created_at BETWEEN ? AND ?", mchId, storeId, startTime, endTime).Count(&count).Error
	if err != nil {
		logger.Monitor.Errorf("Error when counting member card number with mchId:%d, storeId:%d, error:%+v", mchId, storeId, err)
		return 0, err
	}

	return count, nil
}
