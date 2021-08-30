package helpers

import (
	"errors"
	"strings"
)

const PASSWORD_LEN = 4

func CheckNewUserCreds(username, password, email, handle string) error {
	if len(username) < 3 {
		return errors.New("username must contain at least 3 characters")
	}
	if len(password) < PASSWORD_LEN {
		return errors.New("password must contain at least 4 characters")
	}
	if ok := strings.Contains(email, "@"); !ok {
		return errors.New("invalid email")
	}
	if ok := strings.Contains(handle, "@"); ok {
		return errors.New("handle cannot contain '@'")
	}
	return nil
}

func CheckLoginInput(email, password string) error {
	if ok := strings.Contains(email, "@"); !ok {
		return errors.New("invalid email")
	}
	if len(password) < PASSWORD_LEN {
		return errors.New("password can't be smaller than 4 characters")
	}
	return nil
}
