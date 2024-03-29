package postgres

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/param108/profile/api/models"
	"gorm.io/gorm"
)

const OTP_EXPIRY_MINUTES = 20
const OTP_MAX_RETRIES = 3

func generateCode() string {
	seq := []int{}
	for i := 0; i < 6; i++ {
		seq = append(seq, rand.Intn(10))
	}

	return fmt.Sprintf("%d%d%d%d%d%d", seq[0], seq[1], seq[2],
		seq[3], seq[4], seq[5])
}

func (db *PostgresDB) CreateOTP(phone string, now time.Time, writer string) error {

	// if there is an existing valid otp don't create a new one.

	oldSpOtps := []*models.SpOtp{}

	err := db.db.Where("phone = ? AND writer = ?", phone, writer).Find(&oldSpOtps).Error
	if err == nil && len(oldSpOtps) > 0 {
		// the inserted time is local time of the server
		// so we can check with time.Now()
		if oldSpOtps[0].Expiry.After(time.Now()) && oldSpOtps[0].Retries < OTP_MAX_RETRIES {
			return nil
		}
	}

	// no valid old expiry so create a new one.
	spOtp := &models.SpOtp{
		Phone:  phone,
		Writer: writer,
		Expiry: now.Add(OTP_EXPIRY_MINUTES * time.Minute),
		Code:   generateCode(),
	}

	// best effort delete any previous otp
	db.db.Table("sp_otps").Where("phone = ? AND writer = ?", phone, writer).Delete(&models.SpOtp{})

	if err := db.db.Create(spOtp).Error; err != nil {
		return err
	}

	return nil
}

func (db *PostgresDB) GetOTP(phone, writer string) (*models.SpOtp, error) {
	spOtps := []*models.SpOtp{}

	if err := db.db.Table("sp_otps").Find(&spOtps, "phone = ? and writer = ?",
		phone, writer).Error; err != nil {
		return nil, err
	}

	if len(spOtps) == 0 {
		return nil, errors.New("Not Found")
	}

	return spOtps[0], nil
}

func (db *PostgresDB) CheckOTP(
	phone, code string,
	now time.Time, writer string) (*models.SpOtp, error) {
	spOtps := []*models.SpOtp{}

	if err := db.db.Table("sp_otps").Find(&spOtps, "phone = ? and writer = ?",
		phone, writer).Error; err != nil {
		return nil, err
	}

	if len(spOtps) == 0 {
		return nil, errors.New("Not Found")
	}

	if spOtps[0].Retries == OTP_MAX_RETRIES {
		return nil, errors.New("Expired")
	}

	spOtps[0].Retries += 1

	if err := db.db.Save(spOtps[0]).Error; err != nil {
		return nil, err
	}

	ret := spOtps[0]

	if ret.Expiry.Before(now) {
		return nil, errors.New("Expired")
	}

	if code != ret.Code {
		return nil, errors.New("Expired")
	}

	return spOtps[0], nil
}

func (db *PostgresDB) ExpireOTPs(now time.Time, writer string) error {
	err := db.db.Table("sp_otps").Where(
		"expiry < ? and writer = ?",
		now, writer).Delete(&models.SpOtp{}).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	return err
}

func (db *PostgresDB) DeleteAllOTPs(writer string) error {
	return db.db.Table("sp_otps").Where(
		"writer = ?",
		writer).Delete(&models.SpOtp{}).Error
}
