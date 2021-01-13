package repositories

import (
	"database/sql"

	"gorm.io/gorm"

	"entity"
	"fmt"
)

// AlarmsRepositoryI Interface
type AlarmsRepositoryI interface {
	GetCurrentAlarms() []entity.DBAlaramT
	AddAlarm(db *gorm.DB, user entity.DBUserT, time string, text string) entity.DBAlaramT
}

// AlarmsRepositoryT Struct
type AlarmsRepositoryT struct {
}

// GetCurrentAlarms Method
func (AlarmsRepositoryT) GetCurrentAlarms() []entity.DBAlaramT {
	db, err := sql.Open("mysql", "root:password@tcp(mysql)/project")

	// fmt.Println("asdad")

	// var test string
	// db.QueryRow("SELECT NOW()").Scan(&test)
	// fmt.Println(test)

	results, err := db.Query("select u.username, a.text from alarms a inner join users u on a.user_id = u.id where time = DATE_FORMAT(NOW(), '%k:%i')")
	// results, err := db.Query("select u.username, a.text from alarms a inner join users u on a.user_id = u.id LIMIT 1")

	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}

	defer db.Close()

	alarms := make([]entity.DBAlaramT, 0)

	for results.Next() {
		var alarm entity.DBAlaramT
		// for each row, scan the result into our tag composite object
		err = results.Scan(&alarm.User, &alarm.Text)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		alarms = append(alarms, alarm)
	}

	return alarms
}

// AddAlarm Method
func (AlarmsRepositoryT) AddAlarm(db *gorm.DB, user entity.DBUserT, time string, text string) entity.DBAlaramT {
	var alarm = entity.DBAlaramT{
		User: user,
		Time: time,
		Text: &text,
	}

	db.Create(&alarm)

	return alarm
}
