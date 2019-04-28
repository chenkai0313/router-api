package app

import "database/sql"

func NullString(data sql.NullString) (res string) {
	if data.Valid {
		return data.String
	} else {
		return ""
	}
}

func NullInt64(data sql.NullInt64) (res int64) {
	if data.Valid {
		return data.Int64
	} else {
		return 0
	}
}