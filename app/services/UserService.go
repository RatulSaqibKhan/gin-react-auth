package services

import (
	"gin-react-auth/app/domain/users"
	"gin-react-auth/utils/errors"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	// encrypt password
	pwdSlice, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)

	if err != nil {
		return nil, errors.BadRequestError("Failed to encrypt password!")
	}

	user.Password = string(pwdSlice[:])

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUser(user users.User) (*users.User, *errors.RestErr) {
	result := &users.User{Email: user.Email}

	if err := result.GetByEmail(); err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password)); err != nil {
		return nil, errors.BadRequestError("failed to decrypt password")
	}

	resultWithoutPassword := &users.User{ID: result.ID, FirstName: result.FirstName, Email: result.Email}

	return resultWithoutPassword, nil
}

func GetUserByID(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{ID: userId}

	if err := result.GetById(); err != nil {
		return nil, err
	}
	return result, nil
}
