package main

import (
	"flag"
	"log"
	"os"
	"regexp"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const ReVideoIDIndx = 6

func createRegexObj() *regexp.Regexp {
	return regexp.MustCompile(`((?:https?:)?\/\/)?((?:www|m)\.)?((?:youtube(-nocookie)?\.com|youtu.be)\/)((?:watch\?v=|shorts\/|v\/|embed\/)?([\w\-]+))(\S+)?`)
}

func youtubeVideoLink(videoID string) string {
	return "https://www.youtube.com/watch?v=" + videoID
}

func main() {
	token := os.Getenv("TG_TOKEN")
	if token == "" {
		log.Fatal("Token is not set")
	}
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}
	{
		debug := flag.Bool("debug", false, "")
		flag.Parse()
		bot.Debug = *debug
	}

	log.Println("Bot is running...")

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 10

	re := createRegexObj()
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		matches := re.FindAllStringSubmatch(update.Message.Text, -1)

		if len(matches) > 0 {
			var cleanedLinks []string = make([]string, 1, len(matches))
			cleanedLinks[0] = "Cleaned YouTube Links:"
			for _, match := range matches {
				if match[ReVideoIDIndx] != "" {
					cleanedLinks = append(cleanedLinks, youtubeVideoLink(match[ReVideoIDIndx]))
				}
			}
			if len(cleanedLinks) > 0 {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, strings.Join(cleanedLinks, "\n"))
				msg.ReplyToMessageID = update.Message.MessageID
				bot.Send(msg)
			}
		}
	}
}
