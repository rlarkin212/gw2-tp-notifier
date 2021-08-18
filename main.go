package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/rlarkin212/gw2-tp-notifer/gw2service"
	"github.com/rlarkin212/gw2-tp-notifer/telegramservice"
	"github.com/rlarkin212/gw2-tp-notifer/util"
)

const (
	gw2ApiBaseUrl = "https://api.guildwars2.com/v2"
)

var tgApi = util.GetEnvVar("TgApiKey")

func main() {
	bot, err := tgbot.NewBotAPI(tgApi)
	if err != nil {
		log.Fatal(err)
	}

	http.ListenAndServe(":5000", nil)
	for range time.Tick(time.Minute * 6) {
		getSales(bot)
		fmt.Println("called GetSales")
	}
}

func getSales(bot *tgbot.BotAPI) {
	sales := gw2service.FetchSales(gw2ApiBaseUrl)
	items := gw2service.FetchItems(gw2ApiBaseUrl, sales)

	telegramservice.SendMessage(bot, items)
}
