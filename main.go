package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	token := ""
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			if strings.Contains(update.Message.Text, "petunjuk") {
				message, err := os.ReadFile("bot-response/guide.txt")
				if err != nil {
					log.Fatal(err)
				}
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, string(message))
				msg.ReplyToMessageID = update.Message.MessageID

				bot.Send(msg)

			} else if strings.Contains(update.Message.Text, "siapa kamu") {
				message, err := os.ReadFile("bot-response/tentang-bot.txt")
				if err != nil {
					log.Fatal(err)
				}
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, string(message))
				msg.ReplyToMessageID = update.Message.MessageID

				bot.Send(msg)

			} else if strings.Contains(update.Message.Text, "siapa saya") {
				message := fmt.Sprintf("Kamu adalah %s nama sendiri kok lupa. \n\nnote : kalau jawaban diatas salah berarti id telegram anda alay", update.Message.From.UserName)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, string(message))
				msg.ReplyToMessageID = update.Message.MessageID

				bot.Send(msg)

			} else if strings.Contains(update.Message.Text, "kenapa saya jelek") {
				message, err := os.ReadFile("bot-response/kenapa-jelek.txt")
				if err != nil {
					log.Fatal(err)
				}
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, string(message))
				msg.ReplyToMessageID = update.Message.MessageID

				bot.Send(msg)

			} else if strings.Contains(update.Message.Text, "kenapa saya bau") {
				message, err := os.ReadFile("bot-response/kenapa-bau.txt")
				if err != nil {
					log.Fatal(err)
				}
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, string(message))
				msg.ReplyToMessageID = update.Message.MessageID

				bot.Send(msg)

			} else if strings.Contains(update.Message.Text, "solusi agar tidak jelek") {
				message, err := os.ReadFile("bot-response/solusi-jelek.txt")
				if err != nil {
					log.Fatal(err)
				}
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, string(message))
				msg.ReplyToMessageID = update.Message.MessageID

				bot.Send(msg)

			} else if strings.Contains(update.Message.Text, "solusi agar tidak bau") {
				message, err := os.ReadFile("bot-response/solusi-bau.txt")
				if err != nil {
					log.Fatal(err)
				}
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, string(message))
				msg.ReplyToMessageID = update.Message.MessageID

				bot.Send(msg)

			} else if strings.Contains(update.Message.Text, "berikan saya pantun") {
				message, err := os.ReadFile("bot-response/baca-pantun.txt")
				if err != nil {
					log.Fatal(err)
				}
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, string(message))
				msg.ReplyToMessageID = update.Message.MessageID

				bot.Send(msg)

			} else if strings.Contains(update.Message.Text, "game") {
				message, err := os.ReadFile("bot-response/main-game.txt")
				if err != nil {
					log.Fatal(err)
				}
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, string(message))
				msg.ReplyToMessageID = update.Message.MessageID

				bot.Send(msg)

			} else {
				welcomeMessage := fmt.Sprintf("Halo Sobat Kep! \n\nSaya botkep siap melayani anda wahai paduka %s \nSilahkan ketik pesan 'petunjuk' untuk mendapat layanan hamba ini", update.Message.From.UserName)

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, welcomeMessage)
				msg.ReplyToMessageID = update.Message.MessageID

				bot.Send(msg)
			}
		}
	}
}
