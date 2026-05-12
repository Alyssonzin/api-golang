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

func (ur *UserRepository) CreateUser(user model.User) (int, error) {

	var id int
	query, err := ur.connection.Prepare("INSERT INTO users" +
		"(name, email)" +
		" VALUES ($1, $2) RETURNING id")

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(user.Name, user.Email).Scan(&id)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return id, nil
}
