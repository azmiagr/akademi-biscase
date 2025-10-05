package entity

type Role struct {
	RoleID int    `json:"role_id" gorm:"type:int;primaryKey;autoIncrement"`
	Role   string `json:"role" gorm:"type:varchar(25);not null"`

	Users []User `gorm:"foreignKey:RoleID"`
}
