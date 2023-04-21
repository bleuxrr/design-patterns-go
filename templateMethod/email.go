package main

import (
	"fmt"
)

type email struct {
	otp
}

func (s *email) genRandomOTP(length int) string {
	randomOTP := "1234"
	fmt.Printf("email: generating random otp %s\n", randomOTP)
	return randomOTP
}

func (s *email) saveOTPCache(otp string) {
	fmt.Printf("email: saving otp %s to cache\n", otp)
}

func (s *email) getMessage(otp string) string {
	return "email otp for login is " + otp
}

func (s *email) sendNotification(message string) error {
	fmt.Printf("email sending email: %s\n", message)
	return nil
}

func (s *email) publishMetric() {
	fmt.Printf("email: publishing metrics\n")
}
