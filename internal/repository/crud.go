package repository

import (
	"errors"
	"strings"
)

func (q *PostgresQueries) create(
	table string,
	suppliedColumns []string,
	returningID bool,
	modelInstance interface{},
) (int32, error) {

	query := "INSERT INTO " + table + " (" + strings.Join(suppliedColumns, ",") +
		") VALUES (:" + strings.Join(suppliedColumns, ",:") + ")"
	if returningID {
		query += " RETURNING id"
	}

	result, err := q.db.NamedExec(query, modelInstance)
	if err != nil {
		return 0, err
	}

	if returningID {
		id, err := result.LastInsertId()
		if err != nil {
			return 0, err
		}

		return int32(id), nil
	}

	return 0, nil
}

func (q *PostgresQueries) findAll(
	dest []interface{},
	table string,
	whereQuery string,
	whereArgs ...interface{},
) error {

	if err := q.db.Select(dest, "SELECT * FROM "+table+" WHERE "+whereQuery, whereArgs...); err != nil {
		return err
	}

	return nil
}

func (q *PostgresQueries) findOne(
	dest interface{},
	table string,
	whereQuery string,
	whereArgs ...interface{},
) error {

	if err := q.db.Get(dest, "SELECT * FROM "+table+" WHERE "+whereQuery, whereArgs...); err != nil {
		// TODO: fix this, according to docs it returns error if result set is empty
		return err
	}

	return nil
}

// TODO: should detect only properties with non-null value
func (q *PostgresQueries) updateByID(
	table string,
	suppliedColumns []string,
	modelInstance interface{},
) error {

	var setSlice []string
	for _, col := range suppliedColumns {
		setSlice = append(setSlice, col+"=:"+col)
	}
	query := "UPDATE " + table + " SET " + strings.Join(setSlice, ",") + " WHERE id=:id"

	_, err := q.db.NamedExec(query, modelInstance)
	if err != nil {
		return err
	}

	return nil
}

func (q *PostgresQueries) deleteByID(
	table string,
	id int32,
) error {

	result, err := q.db.Exec("DELETE FROM"+table+"WHERE id = $1", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no such user") // TODO: error handling
	}

	return nil
}
