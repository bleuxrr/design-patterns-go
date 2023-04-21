package main

import (
	"fmt"
)

type sms struct {
	otp
}

func (s *sms) genRandomOTP(length int) string {
	randomOTP := "1234"
	fmt.Printf("sms: generating random otp %s\n", randomOTP)
	return randomOTP
}

func (s *sms) saveOTPCache(otp string) {
	fmt.Printf("sms: saving otp %s to cache\n", otp)
}

func (s *sms) getMessage(otp string) string {
	return "sms otp for login is " + otp
}

func (s *sms) sendNotification(message string) error {
	fmt.Printf("sms sending sms: %s\n", message)
	return nil
}

func (s *sms) publishMetric() {
	fmt.Printf("sms: publishing metrics\n")
}
