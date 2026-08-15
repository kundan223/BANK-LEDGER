package service

import (
	"errors" )

var EmailAlreadyExistsError = errors.New("email already exists")

var SomeError = errors.New("some error occurred")