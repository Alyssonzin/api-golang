package repository

import (
	"api-golang/model"
	"database/sql"
	"fmt"
)

type UserRepository struct {
	connection *sql.DB
}

func NewUserRepository(conn *sql.DB) UserRepository {
	return UserRepository{
		connection: conn,
	}
}

func (ur *UserRepository) GetAllUsers() ([]model.User, error) {
	query := "SELECT * FROM users"
	rows, error := ur.connection.Query(query)

	if error != nil {
		fmt.Println(error)
		return []model.User{}, error
	}

	var userList []model.User
	var userObj model.User

	for rows.Next() {
		error = rows.Scan(
			&userObj.ID,
			&userObj.Name,
			&userObj.Email,
			&userObj.CreatedAt,
			&userObj.UpdatedAt,
		)

		if error != nil {
			fmt.Println(error)
			return []model.User{}, error
		}

		userList = append(userList, userObj)
	}

	rows.Close()

	return userList, nil

}
