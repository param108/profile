package store

import (
	"time"

	"github.com/param108/profile/api/models"
)

// CreateOTP Create a new OTP entry
func (s *StoreImpl) CreateOTP(phone string, now time.Time, writer string) error {
	return s.db.CreateOTP(phone, now, writer)
}

// GetOTP Get an OTP entry without checking
func (s *StoreImpl) GetOTP(phone, writer string) (*models.SpOtp, error) {
	return s.db.GetOTP(phone, writer)
}

// CheckOTP Get an OTP after verification.
// Returns error if invalid or expiry etc
func (s *StoreImpl) CheckOTP(
	phone, code string,
	now time.Time, writer string) (*models.SpOtp, error) {
	return s.db.CheckOTP(phone, code, now, writer)
}

// DeleteAllOTPs Delete all otps of a writer.
func (s *StoreImpl) DeleteAllOTPs(writer string) error {
	return s.db.DeleteAllOTPs(writer)
}

// ExpireOTPs expire old otps
func (s *StoreImpl) ExpireOTPs(now time.Time, writer string) error {
	return s.db.ExpireOTPs(now, writer)
}
