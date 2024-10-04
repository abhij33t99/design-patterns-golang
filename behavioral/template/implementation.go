package behavioral

import (
	"fmt"
	"math"
	"math/rand"
)

type SMS struct {
	Otp
}

func (s *SMS) OTP(len int) string {
	return generateRandomOtp(len)

}

func (s *SMS) saveOTPCache(otp string) {
	fmt.Printf("SMS: saving otp: %s to cache\n", otp)
}

func (s *SMS) getMessage(otp string) string {
	return "SMS OTP for login is " + otp
}

func (s *SMS) sendNotification(message string) error {
	fmt.Printf("SMS: sending sms: %s\n", message)
	return nil
}

func generateRandomOtp(len int) string {
	maxLimit := int(math.Pow10(len))
	minLimit := int(math.Pow10(len - 1))
	randInt := int(rand.Float64() * float64(maxLimit))
	if randInt < minLimit {
		randInt += minLimit
	}
	return string(randInt)
}
