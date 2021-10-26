package users

import (
	"gin-react-auth/app/database/mysql"
	"gin-react-auth/utils/errors"
)

var (
	queryInsertUser     = "INSERT INTO users(first_name, last_name, email, password) VALUES (?, ?, ?, ?);"
	queryGetUserByEmail = "SELECT id, first_name, last_name, email, password FROM users WHERE email = ?;"
	queryGetUserById    = "SELECT id, first_name, last_name, email, password FROM users WHERE id = ?;"
)

func (user *User) Save() *errors.RestErr {
	statement, err := mysql.Client.Prepare(queryInsertUser)

	if err != nil {
		return errors.BadRequestError("query prep error")
	}

	defer statement.Close()

	insertResult, saveErr := statement.Exec(user.FirstName, user.LastName, user.Email, user.Password)
	if saveErr != nil {
		return errors.InternalServerError("data saving error")
	}

	userId, saveErr := insertResult.LastInsertId()
	if saveErr != nil {
		return errors.InternalServerError("finding inserted id error")
	}
	user.ID = userId
	return nil
}

func (user *User) GetByEmail() *errors.RestErr {
	statement, err := mysql.Client.Prepare(queryGetUserByEmail)
	if err != nil {
		return errors.BadRequestError("database error")
	}

	defer statement.Close()

	result := statement.QueryRow(user.Email)
	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password); getErr != nil {
		return errors.InternalServerError("database error")
	}

	return nil
}

func (user *User) GetById() *errors.RestErr {
	statement, err := mysql.Client.Prepare(queryGetUserById)
	if err != nil {
		return errors.BadRequestError("database error")
	}

	defer statement.Close()

	result := statement.QueryRow(user.ID)
	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password); getErr != nil {
		return errors.InternalServerError("database error")
	}

	return nil
}
