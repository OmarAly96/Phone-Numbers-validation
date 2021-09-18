package entity

import "errors"

var ErrPhoneAlreadyExists = errors.New("PhoneNumber already exists")

var ErrPhoneDoesNotExist = errors.New("PhoneNumber does not exist")

var ErrInvalidInput = errors.New("invalid input")
