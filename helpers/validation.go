package helpers

import (
	"errors"
)

func ValidateUsername(username string) error {
	if username == "" || username == "null" {
		return errors.New("Username tidak boleh kosong")
	}
	return nil
}
