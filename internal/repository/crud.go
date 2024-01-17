package repository

import (
	"errors"
	"strings"
)

type strToAny map[string]interface{}

func (q *PostgresQueries) create(table string, columns []string, returningID bool, modelInstance interface{}) (int64, error) {
	query := "INSERT INTO " + table + " (" + strings.Join(columns, ",") + ") VALUES (:" + strings.Join(columns, ",:") + ")"
	if returningID {
		query += " RETURNING id"
	}

	if returningID {
		rows, err := q.db.NamedQuery(query, modelInstance)
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

	_, err := q.db.NamedExec(query, modelInstance)
	if err != nil {
		return 0, err
	}

	return 0, nil
}

func (q *PostgresQueries) selectAll(dest interface{}, table string, column string, whereQuery string, whereArgs ...interface{}) error {
	if err := q.db.Select(dest, "SELECT "+column+" FROM "+table+" WHERE "+whereQuery, whereArgs...); err != nil {
		return err
	}

	return nil
}

func (q *PostgresQueries) selectOne(dest interface{}, table string, column string, whereQuery string, whereArgs ...interface{}) error {
	if err := q.db.Get(dest, "SELECT "+column+" FROM "+table+" WHERE "+whereQuery, whereArgs...); err != nil {
		return err
	}

	return nil
}

func (q *PostgresQueries) updateByID(table string, columns []string, modelInstance interface{}) error {
	var setSlice []string
	for _, col := range columns {
		setSlice = append(setSlice, col+"=:"+col)
	}
	query := "UPDATE " + table + " SET " + strings.Join(setSlice, ",") + " WHERE id=:id"

	if _, err := q.db.NamedExec(query, modelInstance); err != nil {
		return err
	}

	return nil
}

func (q *PostgresQueries) updateWhere(table string, mustUpdateOne bool, columns []string, modelInstance interface{}, whereQuery string, whereArgs ...interface{}) error {
	var setSlice []string
	for _, col := range columns {
		setSlice = append(setSlice, col+"=:"+col)
	}
	query := "UPDATE " + table + " SET " + strings.Join(setSlice, ",") + " WHERE " // TODO

	result, err := q.db.NamedExec(query, modelInstance)
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
	if mustUpdateOne && (rowsAffected > 1) {
		return errors.New("more than one instances") // TODO: this should be before actually updating them, alternatively do rollback?
	}

	return nil
}

func (q *PostgresQueries) deleteByID(table string, id int64) error {
	result, err := q.db.Exec("DELETE FROM "+table+" WHERE id=$1", id)
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

func (q *PostgresQueries) deleteWhere(table string, mustDeleteOne bool, whereQuery string, whereArgs ...interface{}) error {
	result, err := q.db.Exec("DELETE FROM "+table+" WHERE "+whereQuery, whereArgs...)
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

func (q *PostgresQueries) count(table string, column string, whereQuery string, whereArgs ...interface{}) (int64, error) {
	var count int64
	if err := q.db.Get(count, "SELECT count("+column+") FROM "+table+" WHERE "+whereQuery, whereArgs...); err != nil {
		return 0, err
	}

	return count, nil
}
