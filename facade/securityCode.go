package main

import "fmt"

type securityCode struct {
	code int
}

func newSecurityCode(code int) *securityCode {
	return &securityCode{
		code: code,
	}
}

func (s *securityCode) checkCode(incomingCode int) error {
	if s.code != incomingCode {
		return fmt.Errorf("security code is incorrect")
	}
	fmt.Println("securityCode verified")
	return nil
}
