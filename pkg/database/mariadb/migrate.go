package mariadb

import (
	"akademi-business-case/entity"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&entity.Role{},
		&entity.User{},
		&entity.OtpCode{},
		&entity.Cart{},
		&entity.Product{},
		&entity.CartItem{},
	)

	if err != nil {
		return err
	}

	return nil
}
