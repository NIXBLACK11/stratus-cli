package api

import "errors"

func Login(email string, password string) (bool, error) {
	if email=="" || password=="" {
		return false, errors.New("Empty username or password")
	}

	
	return true, nil
}