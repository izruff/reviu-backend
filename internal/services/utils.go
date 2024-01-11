package services

import (
	"database/sql"
	"time"
)

// constructors for sql.Null* types

func NewInt(value int32) *sql.NullInt32 {
	return &sql.NullInt32{
		Int32: value,
		Valid: true,
	}
}

func NewString(value string) *sql.NullString {
	return &sql.NullString{
		String: value,
		Valid:  true,
	}
}

func NewBool(value bool) *sql.NullBool {
	return &sql.NullBool{
		Bool:  value,
		Valid: true,
	}
}

func NewTime(value time.Time) *sql.NullTime {
	return &sql.NullTime{
		Time:  value,
		Valid: true,
	}
}
