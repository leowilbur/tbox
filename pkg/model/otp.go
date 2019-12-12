package model

import "time"

// OTP model
type OTP struct {
	PhoneNumber string    `json:"phoneNumber" db:"phone_number" valid:"required"`
	OTP         string    `json:"otp" db:"otp"`
	StampAt     time.Time `json:"-" db:"stamp_at"`
}
