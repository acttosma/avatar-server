package entity

import "time"

type BaseEntity struct {
	CreatedAt *time.Time `gorm:"type:datetime(3);index:idx_ctime;default:CURRENT_TIMESTAMP(3);comment:创建时间"`
	UpdatedAt *time.Time `gorm:"type:datetime(3);index:idx_utime;default:CURRENT_TIMESTAMP(3);comment:修改时间"`
}

// 首次部署后,此方法代码可以注释掉,因为不再需要
func InitTablesIfNeeded() {
	var (
		account    Account
		memberCard MemberCard
	)

	account.CreateTableIfNeeded()
	memberCard.CreateTableIfNeeded()

}
