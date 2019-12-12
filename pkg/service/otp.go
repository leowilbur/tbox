package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/leowilbur/tbox/pkg/model"
	"github.com/leowilbur/tbox/pkg/utils"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/resty.v1"
)

var (
	baseSMSURI = "https://5db83e44177b350014ac77c6.mockapi.io"
)

// OTPGenerate generate otp for input phone number
func OTPGenerate(db *sql.DB, req model.OTP) error {
	var existOTP = model.OTP{
		PhoneNumber: req.PhoneNumber,
	}

	switch err := getOTPByPhoneNumber(db, &existOTP); err {
	case sql.ErrNoRows:
		break
	case nil:
		// User could ask for new otp message after 30s
		if time.Now().Unix() < existOTP.StampAt.Add(30*time.Second).Unix() {
			return errors.New("Please wait to request new OTP message! ")
		}
		break
	default:
		return err
	}

	// make random OTP with 6 characters numbers
	req.OTP = utils.MakeRandString(6, []rune("0123456789"))

	if false { // Assume that it send message success
		var restClient = resty.New().SetHostURL(baseSMSURI)
		resp, err := restClient.NewRequest().
			SetContext(context.Background()).
			SetBody(map[string]interface{}{
				"phone_number": req.PhoneNumber,
				"content":      fmt.Sprintf("Your OTP number is : %s that valid within 60 seconds ", req.OTP),
			}).
			Post("v1/sms")

		if err != nil || resp.IsError() {
			return errors.New("Can not send SMS")
		}
	}

	if _, err := db.Exec("INSERT INTO authorizations(phone_number,otp,stamp_at) VALUES(?,?,?)", req.PhoneNumber, req.OTP, time.Now()); err != nil {
		return errors.New("Error when execute database sql")
	}

	return nil
}

// OTPValidate validate otp for input phone number
func OTPValidate(db *sql.DB, req model.OTP) error {
	var existOTP = model.OTP{
		PhoneNumber: req.PhoneNumber,
	}

	switch err := getOTPByPhoneNumber(db, &existOTP); err {
	case sql.ErrNoRows:
		return errors.New("Phone number is not found! ")
	case nil:
		break
	default:
		return err
	}

	if existOTP.OTP != req.OTP {
		return errors.New("OTP input is not valid! ")
	}

	// OTP only valid within 60s
	if time.Now().Unix() > existOTP.StampAt.Add(60*time.Second).Unix() {
		return errors.New("OTP input is expired! ")
	}

	return nil
}

// get OTP data by input phone number
func getOTPByPhoneNumber(db *sql.DB, req *model.OTP) error {
	row, err := db.Query("SELECT otp,stamp_at FROM authorizations WHERE phone_number = ? ORDER BY id DESC LIMIT 1", req.PhoneNumber)
	if err != nil {
		return errors.New("Error when execute database sql")
	}

	if row.Next() {
		return row.Scan(&req.OTP, &req.StampAt)
	}
	return sql.ErrNoRows
}
