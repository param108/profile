package store

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestOTP(t *testing.T) {
	writer := "843c3e36-00da-48b7-a361-016c760034be"
	testDB.DeleteAllOTPs(writer)
	st := time.Now()
	phone1 := "9876587654"
	validCode := ""
	t.Run("can create otp", func(t *testing.T) {
		err := testDB.CreateOTP(phone1, st, writer)
		assert.Nil(t, err, "Failed to create OTP")

		otp, err := testDB.GetOTP(phone1, writer)
		assert.Nil(t, err, "couldnt find OTP")

		assert.Equal(t, phone1, otp.Phone, "invalid phone found")
		assert.NotEmpty(t, otp.Code, "code empty")

		calcExpiry := st.Add(20 * time.Minute)
		if calcExpiry.After(otp.Expiry) {
			// difference should be less than 1 second.
			assert.True(t, calcExpiry.Sub(otp.Expiry) < 1*time.Second,
				"too much error")
		} else {
			// difference should be less than 1 second.
			assert.True(t, otp.Expiry.Sub(calcExpiry) < 1*time.Second,
				"too much error")
		}
		validCode = otp.Code
	})

	t.Run("check invalid otp", func(t *testing.T) {
		invalidCode := "0000000" // 7 digits so it wont match anything.
		_, err := testDB.CheckOTP(phone1, invalidCode, st.Add(2*time.Minute), writer)
		assert.NotNil(t, err, "Check succeeded for invalid otp")
	})

	t.Run("check expired otp", func(t *testing.T) {
		_, err := testDB.CheckOTP(phone1, validCode, st.Add(25*time.Minute), writer)
		assert.NotNil(t, err, "Check succeeded for expired otp")
	})

	t.Run("check invalid phone", func(t *testing.T) {
		// Invalid phone will not increment retries
		_, err := testDB.CheckOTP("9999999999", validCode, st.Add(2*time.Minute), writer)
		assert.NotNil(t, err, "Check succeeded for invalid phone")
	})

	t.Run("check expiry after 3 tries for valid code", func(t *testing.T) {
		// This is the 3rd retry with proper phone number
		testDB.CheckOTP(phone1, validCode, st.Add(25*time.Minute), writer)

		// even though everything is valid, this will fail
		// because we have run out of retries
		_, err := testDB.CheckOTP(phone1, validCode, st.Add(2*time.Minute), writer)
		assert.NotNil(t, err, "succeeded after 3  retries")
	})

	t.Run("check valid otp", func(t *testing.T) {
		err := testDB.CreateOTP(phone1, st, writer)
		assert.Nil(t, err, "Failed to create OTP")

		otp, err := testDB.GetOTP(phone1, writer)
		assert.Nil(t, err, "couldnt find OTP")

		_, err = testDB.CheckOTP(phone1, otp.Code, st.Add(2*time.Minute), writer)
		assert.Nil(t, err, "Check failed for valid otp")
	})

}
