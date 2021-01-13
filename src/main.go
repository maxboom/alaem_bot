package main

import (
	"fmt"
	"log"
	"regexp"
	"repositories"
	"strings"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:password@tcp(mysql)/project?charset=utf8mb4&parseTime=True&loc=Local"
	db, error := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if error != nil {
		log.Fatal(error)
		return
	}

	poller := &tb.LongPoller{Timeout: 1 * time.Second}

	// auth := tb.NewMiddlewarePoller(poller, func(update *tb.Update) bool {
	// 	user := repositories.UserRepositoryT{}.GetUser(db, update.Message.Sender.Username)

	// 	var regex = regexp.MustCompile("^\\/(start|test)$")

	// 	if user.IsAuthorized == false && regex.MatchString(update.Message.Text) == false {
	// 		fmt.Println("False")
	// 		return false
	// 	}

	// 	return true
	// })

	bot, error := tb.NewBot(tb.Settings{
		Token:     "1488442278:AAEsGoCz5v_8DLrXsSQVKVNdlyBjBKdYnn8",
		Poller:    poller,
		ParseMode: tb.ModeHTML,
	})

	if error != nil {
		log.Fatal(error)
		return
	}

	menu := &tb.ReplyMarkup{
		ResizeReplyKeyboard: true,
		OneTimeKeyboard:     true,
		ReplyKeyboardRemove: true,
	}

	register := menu.Text("⚙ Настроить бота")
	auth := menu.Text("→ Пройти авторизацию")

	bot.Handle("/start", func(messsage *tb.Message) {
		sendStart(bot, register, *messsage, db)
	})

	bot.Handle(&register, func(messsage *tb.Message) {
		sendRegisterMessage(*bot, messsage, auth)
	})

	bot.Handle(&auth, func(messsage *tb.Message) {
		user := repositories.UserRepositoryT{}.GetUser(db, messsage.Sender.Username)
		// user.IsAuthorized = callmebotapi.CallMeBotT{}.CallUser(user.Username, "Привет! Ты успешно прошел авторизацию! Продолжи настройку в боте.")

		user.IsAuthorized = true

		fmt.Println(user)

		if user.IsAuthorized == false {
			bot.Send(messsage.Sender, "К сожалению авторзация не прошла успешно. \nПопробуй еще раз...")
			sendRegisterMessage(*bot, messsage, auth)
		} else {
			user = repositories.UserRepositoryT{}.UpdateUser(db, user)
			sendStart(bot, register, *messsage, db)
		}

	})

	bot.Handle("/add", func(message *tb.Message) {
		handleTimer(message, db)
	})

	bot.Start()
}

func sendRegisterMessage(bot tb.Bot, messsage *tb.Message, auth tb.Btn) {
	inline := &tb.ReplyMarkup{}

	call := inline.URL("ℹ Сделать запрос на разрешение доступа", "https://api2.callmebot.com/txt/auth.php")

	inline.Inline(
		inline.Row(call),
	)

	replyMenu := &tb.ReplyMarkup{
		ResizeReplyKeyboard: true,
		OneTimeKeyboard:     true,
		ReplyKeyboardRemove: true,
	}

	replyMenu.Reply(
		replyMenu.Row(auth),
	)

	bot.Send(messsage.Sender, "Сейчас я сделаю запрос досупа к совершению звонков...", replyMenu)
	bot.Send(
		messsage.Sender,
		"<i>Я не запрашиваю доступ</i><strong> к твоим звонкам или контантам!</strong><i>Всегда читай список запрашиваимых данныйх.</i>",
		inline,
	)
}

func sendStart(bot *tb.Bot, register tb.Btn, messsage tb.Message, db *gorm.DB) {
	user := repositories.UserRepositoryT{}.GetUser(db, messsage.Sender.Username)

	replyMenu := &tb.ReplyMarkup{
		ResizeReplyKeyboard: true,
		OneTimeKeyboard:     true,
		ReplyKeyboardRemove: true,
	}

	if user.IsAuthorized == false {
		replyMenu.Reply(replyMenu.Row(register))

		bot.Send(messsage.Sender, "Привет, я бот который будет напоминать тебе о важном!", replyMenu)
	} else {
		bot.Send(messsage.Sender, "Что бы добавить оповещение напиши: \n <strong>/add 12:00 текст_звонка</strong>")
	}
}

func handleTimer(messsage *tb.Message, db *gorm.DB) {
	user := repositories.UserRepositoryT{}.GetUser(db, messsage.Sender.Username)

	regex := regexp.MustCompile("\\/add\\s(?P<time>[\\d]{2}:[\\d]{2})\\s(?P<end>(?:.)+$)")
	match := regex.MatchString(strings.ToLower(messsage.Text))

	if true == match {
		matchSubmatch := regex.FindStringSubmatch(strings.ToLower(messsage.Text))
		result := make(map[string]string)
		for i, name := range regex.SubexpNames() {
			if i != 0 && name != "" {
				result[name] = matchSubmatch[i]
			}
		}

		repositories.AlarmsRepositoryT{}.AddAlarm(db, user, result["time"], result["end"])

		// repositories.AlarmsRepositoryT{}.AddAlarm(getSQLConnection(), user, result["time"], result["end"])

		// sendMessage := telegramapi.SendMessageT{ChatID: update.Message.Chat.ID, Message: "Добавлено оповещение на " + result["time"]}

		// telegramapi.BotT{}.SendMessage(botData, sendMessage)
	}
}
