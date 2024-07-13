package postgres

import (
	"errors"
	"strings"
)

func (r *PostgresRepository) create(table string, columns []string, returningID bool, modelInstance interface{}) (int64, error) {
	query := "INSERT INTO " + table + " (" + strings.Join(columns, ",") + ") VALUES (:" + strings.Join(columns, ",:") + ")"
	if returningID {
		query += " RETURNING id"
	}

	if returningID {
		rows, err := r.db.NamedQuery(query, modelInstance)
		if err != nil {
			return 0, err
		}

		var dest struct {
			ID int64 `db:"id"`
		}
		rows.Next()
		if err := rows.StructScan(&dest); err != nil {
			return 0, err
		}

		return dest.ID, nil
	}

	_, err := r.db.NamedExec(query, modelInstance)
	if err != nil {
		return 0, err
	}

	return 0, nil
}

func (r *PostgresRepository) selectAll(dest interface{}, table string, column string, whereQuery string, orderBy string, queryArgs ...interface{}) error {
	query := "SELECT " + column + " FROM " + table
	if whereQuery != "" {
		query += " WHERE " + whereQuery
	}
	if orderBy != "" {
		query += " ORDER BY " + orderBy
	}

	if err := r.db.Select(dest, query, queryArgs...); err != nil {
		return err
	}

	return nil
}

func (r *PostgresRepository) selectOne(dest interface{}, table string, column string, whereQuery string, whereArgs ...interface{}) error {
	query := "SELECT " + column + " FROM " + table
	if whereQuery != "" {
		query += " WHERE " + whereQuery
	}

	if err := r.db.Get(dest, query, whereArgs...); err != nil {
		return err
	}

	return nil
}

func (r *PostgresRepository) updateByID(table string, columns []string, modelInstance interface{}) error {
	var setSlice []string
	for _, col := range columns {
		setSlice = append(setSlice, col+"=:"+col)
	}
	query := "UPDATE " + table + " SET " + strings.Join(setSlice, ",") + " WHERE id=:id"

	if _, err := r.db.NamedExec(query, modelInstance); err != nil {
		return err
	}

	return nil
}

func (r *PostgresRepository) updateByPK(table string, columns []string, pkColumns []string, modelInstance interface{}) error {
	var setSlice []string
	for _, col := range columns {
		setSlice = append(setSlice, col+"=:"+col)
	}
	var whereSlice []string
	for _, col := range pkColumns {
		whereSlice = append(setSlice, col+"=:"+col)
	}
	query := "UPDATE " + table + " SET " + strings.Join(setSlice, ",") + " WHERE " + strings.Join(whereSlice, " AND ")

	result, err := r.db.NamedExec(query, modelInstance)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no such row") // TODO: error handling
	}

	return nil
}

func (r *PostgresRepository) deleteByID(table string, id int64) error {
	result, err := r.db.Exec("DELETE FROM "+table+" WHERE id=$1", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no such row") // TODO: error handling
	}

	return nil
}

func (r *PostgresRepository) deleteWhere(table string, mustDeleteOne bool, whereQuery string, whereArgs ...interface{}) error {
	result, err := r.db.Exec("DELETE FROM "+table+" WHERE "+whereQuery, whereArgs...)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no such row") // TODO: error handling
	}
	if mustDeleteOne && (rowsAffected > 1) {
		return errors.New("more than one instances") // TODO: this should be before actually deleting them, alternatively do rollback?
	}

	return nil
}

func (r *PostgresRepository) count(table string, column string, whereQuery string, whereArgs ...interface{}) (int64, error) {
	var count int64
	if err := r.db.Get(&count, "SELECT count("+column+") FROM "+table+" WHERE "+whereQuery, whereArgs...); err != nil {
		return 0, err
	}

	return count, nil
}
