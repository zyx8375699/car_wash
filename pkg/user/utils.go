package user

import (
	"fmt"
	"regexp"
)

const (
	PHONEREG = `^1([38][0-9]|14[57]|5[^4])\d{8}$`
)

func validatePhone(phone int64) bool {
	phoneReg, err := regexp.Compile(PHONEREG)
	if err != nil {
		return false
	}
	return phoneReg.MatchString(fmt.Sprintf("%d", phone))
}
