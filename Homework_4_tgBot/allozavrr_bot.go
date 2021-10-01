package main

import (
	"os"
	"log"
	"regexp"
	"strconv"
	"time"
	"github.com/Syfaro/telegram-bot-api"
)
//Создаем бота
func allozavrr_bot() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	homeworks := [...]string{"Homework_1", "Homework_1_zoOapp", "Homework_2", "Homework_3_Docker", "Homework_4_tgBot"}
	links := [...]string{"https://github.com/allozavrr/Andersen_DevOps_Course/tree/main/Homework_1", "https://github.com/allozavrr/Andersen_DevOps_Course/tree/main/Homework_1_zoOapp", "https://github.com/allozavrr/Andersen_DevOps_Course/tree/main/Homework_2", "https://github.com/allozavrr/Andersen_DevOps_Course/tree/main/Homework_3_Docker", "https://github.com/allozavrr/Andersen_DevOps_Course/tree/main/Homework_4_tgBot", "https://github.com/allozavrr/Andersen_DevOps_Course/tree/main/Homework_5"}
	var commandPat = regexp.MustCompile(`^homework[0-9]`)

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)
//Устанавливаем время обновления
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
//Получаем обновления от бота
	updates, err := bot.GetUpdatesChan(u)
//Игнорируем неверные сообщения
	for update := range updates {
		if update.Message == nil { 
			continue
		}
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
//Отправляем сообщение
		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			switch cmd := update.Message.Command(); {
			case cmd == "start":
				msg.Text = "Нажми /help чтобы увидеть меню"
			case cmd == "help":
				msg.ParseMode = "html"
				msg.Text = "Поддерживаемые команды бота:\n /home - репозиторий с проектами\n /homeworks - все выполненные репозитории\n /homework# - (# номер задания) линк на данное задание"
			case cmd == "home":
				msg.Text = "https://github.com/allozavrr/Andersen_DevOps_Course"
			case cmd == "homeworks":
				msg.ParseMode = "html"
				msg.Text = "Список выполненных заданий:\n"
				for i := 0; i < len(homeworks); i++ {
					msg.Text = msg.Text + strconv.Itoa(i+1) + ". " + homeworks[i] + "  -  " + "<a href='" + links[i] + "'> <b>link</b></a>" + "\n"
				}
			case commandPat.MatchString(cmd):
				cmd := cmd[4:5]
				taskIndex, _ := strconv.Atoi(cmd)
				if taskIndex <= len(links) {
					for i := taskIndex; i <= (taskIndex); i++ {
						msg.Text = links[i-1]
					}
				} else {
					msg.Text = "Что-то не так. Попробуйте ещё раз"
				}
			default:
				msg.Text = "Чтоа"
			}
			bot.Send(msg)
		}
	}
}

func main() {
	time.Sleep(1 * time.Minute)

	time.Sleep(1 * time.Minute)
	
	//Вызываем бота
	allozavrr_bot()
}
