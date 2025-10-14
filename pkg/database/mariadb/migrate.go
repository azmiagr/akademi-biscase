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
		&entity.ClassType{},
		&entity.Class{},
		&entity.CartItem{},
		&entity.Review{},
	)

	if err != nil {
		return err
	}

	return nil
}
