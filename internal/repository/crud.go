package repository

import (
	"errors"
	"strings"
)

func (q *PostgresQueries) create(table string, columns []string, returningID bool, modelInstance interface{}) (int32, error) {
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
			ID int32 `db:"id"`
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

func (q *PostgresQueries) selectAll(dest []interface{}, table string, column string, whereQuery string, whereArgs ...interface{}) error {
	if err := q.db.Select(dest, "SELECT "+column+" FROM "+table+" WHERE "+whereQuery, whereArgs...); err != nil {
		return err
	}

	return nil
}

func (q *PostgresQueries) selectOne(dest interface{}, table string, column string, whereQuery string, whereArgs ...interface{}) error {
	if err := q.db.Get(dest, "SELECT "+column+" FROM "+table+" WHERE "+whereQuery, whereArgs...); err != nil {
		// TODO: fix this, according to docs it returns error if result set is empty
		return err
	}

	return nil
}

// TODO: should detect only properties with non-null value
func (q *PostgresQueries) updateByID(table string, columns []string, modelInstance interface{}) error {
	var setSlice []string
	for _, col := range columns {
		setSlice = append(setSlice, col+"=:"+col)
	}
	query := "UPDATE " + table + " SET " + strings.Join(setSlice, ",") + " WHERE id=:id"

	_, err := q.db.NamedExec(query, modelInstance)
	if err != nil {
		return err
	}

	return nil
}

func (q *PostgresQueries) deleteByID(table string, id int32) error {
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
