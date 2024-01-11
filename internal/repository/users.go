package repository

import "github.com/izruff/reviu-backend/internal/models"

func (q *PostgresQueries) CreateUser(newUser *models.User) (int32, error) {
	userID, err := q.create("users", []string{"email", "password_hash", "mod_role", "username"}, true, newUser)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (q *PostgresQueries) GetUserByID(id int32) (*models.User, error) {
	user := &models.User{}
	// TODO: see findOne TODO
	if err := q.selectOne(user, "users", "*", "id=:id", id); err != nil {
		return nil, err
	}

	return user, nil
}

func (q *PostgresQueries) GetUserIDByEmail(email string) (int32, error) {
	var userID int32
	if err := q.selectOne(userID, "users", "id", "email=:email", email); err != nil {
		return 0, err
	}

	return userID, nil
}

func (q *PostgresQueries) GetUserIDByUsername(username string) (int32, error) {
	var userID int32
	if err := q.selectOne(userID, "users", "id", "username=:username", username); err != nil {
		return 0, err
	}

	return userID, nil
}

func (q *PostgresQueries) UpdateUserByID(id int32, updatedUser *models.User) error {
	return nil // TODO
}

func (q *PostgresQueries) DeleteUserByID(id int32) error {
	return nil // TODO
}
