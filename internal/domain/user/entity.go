package user

import (
	"time"
)

type Entity struct {
	ID          int    `gorm:"primaryKey"`
	Username    string `gorm:"column:username;type:varchar(20)"`
	Email       string `gorm:"column:email;type:varchar(100)"`
	PhoneNumber string `gorm:"column:phone_number;type:varchar(15)"`
	Password    string `gorm:"column:password;type:varchar(255)"`
	Address     string `gorm:"column:address;type:text"`
	Dob         string `gorm:"column:dob;type:date"`
	FullName    string `gorm:"column:full_name;type:varchar(100)" validate:"required"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (*Entity) TableName() string {
	return "users"
}
