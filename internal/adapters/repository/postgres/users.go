package postgres

import (
	"errors"
	"strconv"

	"github.com/izruff/reviu-backend/internal/core/domain"
)

func (r *PostgresRepository) CreateUser(newUser *domain.User) (int64, error) {
	userID, err := r.create("users", []string{"email", "password_hash", "mod_role", "username"}, true, newUser)
	// TODO: error handling when form is incomplete or user already exist
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (r *PostgresRepository) GetUserByID(id int64) (*domain.User, error) {
	user := &domain.User{}
	if err := r.selectOne(user, "users", "*", "id=$1", id); err != nil {
		return nil, err // TODO: error handling when user does not exist
	}

	return user, nil
}

func (r *PostgresRepository) GetUserIDByEmail(email string) (int64, error) {
	var userID int64
	if err := r.selectOne(&userID, "users", "id", "email=$1", email); err != nil {
		return 0, err // TODO: error handling when user does not exist
	}

	return userID, nil
}

func (r *PostgresRepository) GetUserIDByUsername(username string) (int64, error) {
	var userID int64
	if err := r.selectOne(&userID, "users", "id", "username=$1", username); err != nil {
		return 0, err // TODO: error handling when user does not exist
	}

	return userID, nil
}

func (r *PostgresRepository) GetUsersWithOptions(options *domain.SearchUsersOptions) ([]domain.User, error) {
	var whereQuery, orderBy string
	var queryArgs []interface{}
	argsIndex := 1

	if options.Query == "" {
		return nil, errors.New("unexpected error: query is empty") // TODO: this should be allowed for browsing
	}

	if options.SortBy == "popularity" || options.SortBy == "" {
		orderBy = "username <-> $" + strconv.Itoa(argsIndex)
		queryArgs = append(queryArgs, options.Query)
		argsIndex++
	} else if options.SortBy == "popularity" {
		orderBy = "" // TODO
	} else {
		return nil, errors.New("unexpected error: invalid option for sort-by")
	}

	if options.MustMatch == "left" {
		whereQuery = "username ILIKE $" + strconv.Itoa(argsIndex)
		queryArgs = append(queryArgs, options.Query+"%")
		argsIndex++
	} else if options.MustMatch == "substring" {
		whereQuery = "username ILIKE $" + strconv.Itoa(argsIndex)
		queryArgs = append(queryArgs, "%"+options.Query+"%")
		argsIndex++
	} else if options.MustMatch == "none" || options.MustMatch == "" {
		whereQuery = "username % $" + strconv.Itoa(argsIndex)
		queryArgs = append(queryArgs, options.Query)
		argsIndex++
	} else {
		return nil, errors.New("unexpected error: invalid option for must-match")
	}

	users := []domain.User{}
	if err := r.selectAll(&users, "users", "*", whereQuery, orderBy, queryArgs...); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *PostgresRepository) UpdateUserByID(updatedUser *domain.User) error {
	if !updatedUser.ID.Valid {
		return errors.New("ID not provided")
	}

	var columns []string
	if updatedUser.Email.Valid {
		columns = append(columns, "email")
	}
	if updatedUser.PasswordHash.Valid {
		columns = append(columns, "password_hash")
	}
	if updatedUser.ModRole.Valid {
		columns = append(columns, "mod_role")
	}
	if updatedUser.Username.Valid {
		columns = append(columns, "username")
	}
	if updatedUser.Nickname.Valid {
		columns = append(columns, "nickname")
	}
	if updatedUser.About.Valid {
		columns = append(columns, "about")
	}
	// TODO: error handling if nothing is updated

	if err := r.updateByID("users", columns, updatedUser); err != nil {
		return err // TODO: error handling when user does not exist
	}

	return nil
}

// TODO: not sure if this is possible when user has already posted
func (r *PostgresRepository) DeleteUserByID(id int64) error {
	if err := r.deleteByID("users", id); err != nil {
		return err // TODO: error handling when user does not exist
	}

	return nil
}
