package main

import (
	"callmebotapi"
	"database/sql"
	"entity"
	"fmt"
	"regexp"
	"repositories"
	"strings"
	"telegramapi"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const token = "bot1488442278:AAEsGoCz5v_8DLrXsSQVKVNdlyBjBKdYnn8"

func main() {
	botData := telegramapi.BotSettingsT{Token: token}
	getUpdatesRequest := telegramapi.GetUpdatesRequestT{Offset: 0}

	for {
		fmt.Println(fmt.Sprintf("%d", getUpdatesRequest.Offset))

		getUpdates := telegramapi.BotT{}.GetUpdates(botData, getUpdatesRequest)

		fmt.Println(getUpdates)

		for _, update := range getUpdates.Result {
			getUpdatesRequest.Offset = update.UpdateID
			getUpdatesRequest.Offset++

			fmt.Println(fmt.Sprintf("%d", update.UpdateID))

			user := getOrCreate(getSQLConnection(), update.Message.From.Username)

			fmt.Println(user)

			if !user.IsAuthorized {
				if strings.ToLower(update.Message.Text) == "/start" {
					sendMessage := telegramapi.SendMessageT{ChatID: update.Message.Chat.ID, Message: "Привет! \nЯ бот котороый тебя заебет! \nДля успешной работы бота перейди по линке: https://api2.callmebot.com/txt/auth.php \nНажми /test Что бы сделать тестовый вызов "}

					telegramapi.BotT{}.SendMessage(botData, sendMessage)
				}

				if strings.ToLower(update.Message.Text) == "/test" {
					isAuthorized := callmebotapi.CallMeBotT{}.CallUser(update.Message.From.Username, "Привет! Я бот котороый тебя заебет!")

					if isAuthorized != true {
						sendMessage := telegramapi.SendMessageT{ChatID: update.Message.Chat.ID, Message: "Ты не авторизировался \nДля успешной работы бота перейди по линке: https://api2.callmebot.com/txt/auth.php \nНажми /test Что бы сделать тестовый вызов "}

						telegramapi.BotT{}.SendMessage(botData, sendMessage)
					} else {
						user := persistUser(getSQLConnection(), update.Message.From.Username)
						fmt.Println(user)
						sendMessage := telegramapi.SendMessageT{ChatID: update.Message.Chat.ID, Message: "Отлично! Авторизация успешна!\nПерейдем к настройке оповещений"}

						telegramapi.BotT{}.SendMessage(botData, sendMessage)
					}
				}
			} else {
				if strings.ToLower(update.Message.Text) == "/start" {
					sendMessage := telegramapi.SendMessageT{ChatID: update.Message.Chat.ID, Message: "Что бы дабавить новое оповещение напиши /add 12:00 текст_звонка"}

					telegramapi.BotT{}.SendMessage(botData, sendMessage)
				} else if strings.Contains(strings.ToLower(update.Message.Text), "/add") == true {
					regex := regexp.MustCompile("\\/add\\s(?P<time>[\\d]{2}:[\\d]{2})\\s(?P<end>(?:.)+$)")
					match := regex.MatchString(strings.ToLower(update.Message.Text))

					if true == match {
						matchSubmatch := regex.FindStringSubmatch(strings.ToLower(update.Message.Text))
						result := make(map[string]string)
						for i, name := range regex.SubexpNames() {
							if i != 0 && name != "" {
								result[name] = matchSubmatch[i]
							}
						}

						repositories.AlarmsRepositoryT{}.AddAlarm(getSQLConnection(), user, result["time"], result["end"])

						sendMessage := telegramapi.SendMessageT{ChatID: update.Message.Chat.ID, Message: "Добавлено оповещение на " + result["time"]}

						telegramapi.BotT{}.SendMessage(botData, sendMessage)
					}
				}

			}

		}

		time.Sleep(10 * time.Second)

	}
}

func getOrCreate(db *sql.DB, username string) entity.DBUserT {
	user := entity.DBUserT{Username: username, IsAuthorized: false}

	defer func() {
		if recover := recover(); recover != nil {
			print("will Work")
		}
	}()

	user = repositories.UserRepositoryT{}.GetUser(db, username)

	defer db.Close()

	return user
}

func persistUser(db *sql.DB, username string) entity.DBUserT {
	defer func() {
		if recover := recover(); recover != nil {
			print("will Work")
		}
	}()

	repositories.UserRepositoryT{}.CreateUser(db, username)

	defer db.Close()

	return getOrCreate(db, username)
}

func getSQLConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(mysql)/project")

	if err != nil {
		panic(err.Error())
	}

	return db
}
