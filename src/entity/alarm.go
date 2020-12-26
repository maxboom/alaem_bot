package entity

import "database/sql"

// DBAlaramT Type
type DBAlaramT struct {
	Username string
	Text     sql.NullString
}
