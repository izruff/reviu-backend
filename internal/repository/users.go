package repository

import "github.com/izruff/reviu-backend/internal/models"

func (q *PostgresQueries) CreateUser(newUser *models.User) (int32, error) {
	return 0, nil // TODO
}

func (q *PostgresQueries) GetUserByID(id int32) (*models.User, error) {
	return &models.User{}, nil // TODO
}

func (q *PostgresQueries) GetUserIDByEmail(email string) (int32, error) {
	return 0, nil // TODO
}

func (q *PostgresQueries) GetUserIDByUsername(username string) (int32, error) {
	return 0, nil // TODO
}

func (q *PostgresQueries) UpdateUserByID(id int32, updatedUser *models.User) error {
	return nil // TODO
}

func (q *PostgresQueries) DeleteUserByID(id int32) error {
	return nil // TODO
}
