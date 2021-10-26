package users

import (
	"gin-react-auth/utils/errors"
	"strings"
)

type User struct {
	ID        int64  `json:"ID"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (user *User) Validate() *errors.RestErr {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(user.Email)
	user.Password = strings.TrimSpace(user.Password)

	if user.Email == "" {
		return errors.BadRequestError("Invalid email address!")
	}

	if user.Password == "" {
		return errors.BadRequestError("Invalid password!")
	}

	return nil
}
