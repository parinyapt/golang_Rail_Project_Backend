package models

import (
	"time"
)

type Account struct {
	ID    uint   `gorm:"column:account_id"`
	UUID  string `gorm:"column:account_uuid"`
	Email string `gorm:"column:account_email"`
	Name  string `gorm:"column:account_name"`
}

func (Account) TableName() string {
	return "crp_account"
}

type OTP struct {
	ID          uint      `gorm:"column:otp_id"`
	UUID        string    `gorm:"column:otp_uuid"`
	AccountUUID string    `gorm:"column:otp_account_uuid"`
	Code        string    `gorm:"column:otp_code"`
	Status      string    `gorm:"column:otp_status"`
	CreatedAt   time.Time `gorm:"column:otp_create_at"`
	UpdatedAt   time.Time `gorm:"column:otp_update_at"`
}

func (OTP) TableName() string {
	return "crp_otp"
}
