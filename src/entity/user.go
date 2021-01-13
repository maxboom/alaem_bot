package entity

// DBUserT Type
type DBUserT struct {
	ID           int
	Username     string
	IsAuthorized bool
}

// TableName method
func (DBUserT) TableName() string {
	return "users"
}
