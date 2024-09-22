package entity

import "database/sql"

type USERS struct {
	ID         int32
	NAMA       string
	AGE        int
	EMAIL      sql.NullString
	RATING     sql.NullFloat64
	CREATED_AT sql.NullTime
	BIRTH_DATE sql.NullTime
	MARRIED    sql.NullBool
}
