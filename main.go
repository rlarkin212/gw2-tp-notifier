package main

import (
	"fmt"
	"log"
	"time"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/rlarkin212/gw2-tp-notifer/gw2service"
	"github.com/rlarkin212/gw2-tp-notifer/telegramservice"
	"github.com/rlarkin212/gw2-tp-notifer/util"
)

const (
	gw2ApiBaseUrl = "https://api.guildwars2.com/v2"
	iso8601       = "2006-01-02T15:04:05-0700"
)

var tgApi = util.GetEnvVar("TgApiKey")

func main() {
	bot, err := tgbot.NewBotAPI(tgApi)
	if err != nil {
		log.Fatal(err.Error())
	}

	for range time.Tick(time.Minute * 6) {
		getSales(bot)
		fmt.Printf("called GetSales @ %s", time.Now().UTC().Format(iso8601))
	}
}

func getSales(bot *tgbot.BotAPI) {
	sales := gw2service.FetchSales(gw2ApiBaseUrl)
	items := gw2service.FetchItems(gw2ApiBaseUrl, sales)

	telegramservice.SendMessage(bot, items)
}
