package entity

import "time"

type BaseEntity struct {
	CreatedAt *time.Time `gorm:"type:datetime(3);index:idx_ctime;default:CURRENT_TIMESTAMP(3);comment:creation time"`
	UpdatedAt *time.Time `gorm:"type:datetime(3);index:idx_utime;default:CURRENT_TIMESTAMP(3);comment:modification time"`
}

// Comment this function after deployed
func InitTablesIfNeeded() {
	var (
		account       Account
		avatar        Avatar
		avatarProduct AvatarProduct
	)

	account.CreateTableIfNeeded()
	avatar.CreateTableIfNeeded()
	avatarProduct.CreateTableIfNeeded()

}
