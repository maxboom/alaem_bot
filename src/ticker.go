package main

import (
	"callmebotapi"
	"fmt"
	"repositories"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	for {
		alarms := repositories.AlarmsRepositoryT{}.GetCurrentAlarms()
		fmt.Println(alarms)

		for _, alarm := range alarms {
			var text = "Привет! Я бот который тебя заебет!"

			if alarm.Text.Valid {
				text = alarm.Text.String
			}

			callmebotapi.CallMeBotT{}.CallUser(alarm.Username, text)
		}

		time.Sleep(1 * time.Minute)
	}
}
