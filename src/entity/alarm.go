package entity

// DBAlaramT Type
type DBAlaramT struct {
	User   DBUserT `gorm:"foreignKey:UserID"`
	UserID int
	Text   *string
	Time   string
}

// TableName method
func (DBAlaramT) TableName() string {
	return "alarms"
}
