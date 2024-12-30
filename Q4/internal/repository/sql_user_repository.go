package repository

import (
	"Q4/internal/model"
	"database/sql"
)

type SQLUserRepository struct {
	DB *sql.DB
}

func NewSQLUserRepository(db *sql.DB) *SQLUserRepository {
	return &SQLUserRepository{
		DB: db,
	}
}

func (ur *SQLUserRepository) GetAllUsers() ([]model.User, error) {
	rows, err := ur.DB.Query("SELECT * FROM users;")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (ur *SQLUserRepository) GetUserByID(id int) (*model.User, error) {
	row := ur.DB.QueryRow("SELECT * FROM users WHERE id = ?", id)

	var user model.User
	if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *SQLUserRepository) CreateUser(user *model.User) error {
	_, err := ur.DB.Exec("INSERT INTO users (name, email) VALUES (?, ?);", user.Name, user.Email)
	return err
}

func (ur *SQLUserRepository) UpdateUser(user *model.User) error {
	_, err := ur.DB.Exec("UPDATE users SET name = ?, email = ? WHERE id = ?;", user.Name, user.Email, user.ID)
	return err
}

func (ur *SQLUserRepository) DeleteUser(id int) error {
	_, err := ur.DB.Exec("DELETE FROM users WHERE id = ?;", id)
	return err
}
