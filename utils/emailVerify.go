package utils

import (
	"errors"

	emailverifier "github.com/AfterShip/email-verifier"
)

func EmailVerify(email string) (bool, error) {
	verifier := emailverifier.NewVerifier()

	ret, err := verifier.Verify(email)
	if err != nil {
		return false, errors.New("Email verification failed")
	}

	if !ret.Syntax.Valid {
		return false, errors.New("Invalid email syntax")
	}

	return true, nil
}