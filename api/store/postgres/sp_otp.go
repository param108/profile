package postgres

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/param108/profile/api/models"
)

const OTP_EXPIRY_MINUTES = 20

func generateCode() string {
	seq := []int{}
	for i := 0; i < 6; i++ {
		seq = append(seq, rand.Intn(10))
	}

	return fmt.Sprintf("%d%d%d%d%d%d", seq[0], seq[1], seq[2],
		seq[3], seq[4], seq[5])
}

func (db *PostgresDB) CreateOTP(phone string, now time.Time, writer string) error {
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
	return db.db.Table("sp_otps").Where(
		"expiry < ? and writer = ?",
		now, writer).Delete(&models.SpOtp{}).Error
}

func (db *PostgresDB) DeleteAllOTPs(writer string) error {
	return db.db.Table("sp_otps").Where(
		"writer = ?",
		writer).Delete(&models.SpOtp{}).Error
}
