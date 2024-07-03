// repository/user_repository.go
package repository

import (
	"crud_structured/model"
	"database/sql"
)

type UserRepository interface {
	GetAll() ([]model.User, error)
	GetByID(id int) (*model.User, error)
	Create(user *model.User) error
	Update(user *model.User) error
	Delete(id int) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) GetAll() ([]model.User, error) {
	rows, err := r.db.Query("SELECT id, name, status FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Status); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *userRepository) GetByID(id int) (*model.User, error) {
	var user model.User
	err := r.db.QueryRow("SELECT id, name, status FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Create(user *model.User) error {
	result, err := r.db.Exec("INSERT INTO users (name, status) VALUES (?, ?)", user.Name, user.Status)
	if err != nil {
		return err
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = int(lastInsertId)
	return nil
}

func (r *userRepository) Update(user *model.User) error {
	_, err := r.db.Exec("UPDATE users SET name = ?, status = ? WHERE id = ?", user.Name, user.Status, user.ID)
	return err
}

func (r *userRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = ?", id)
	return err
}
