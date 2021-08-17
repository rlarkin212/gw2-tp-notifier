package main

import (
	"fmt"
	"log"
	"strconv"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/rlarkin212/gw2-tp-notifer/gw2service"
	"github.com/rlarkin212/gw2-tp-notifer/util"
)

const (
	gw2ApiBaseUrl = "https://api.guildwars2.com/v2"
)

var tgApi = util.GetEnvVar("TgApiKey")
var tgChatId, _ = strconv.ParseInt(util.GetEnvVar("TgChatId"), 10, 64)

func main() {
	bot, err := tgbot.NewBotAPI(tgApi)
	if err != nil {
		log.Fatal(err)
	}

	sales := gw2service.FetchSales(gw2ApiBaseUrl)
	items := gw2service.FetchItems(gw2ApiBaseUrl, sales)

	for _, item := range items {
		fmt.Printf("%+v\n", item)

		str := fmt.Sprintf("%s sold for %d", item.Name, item.Price)
		msg := tgbot.NewMessage(tgChatId, str)

		bot.Send(msg)
	}

	// for range time.Tick(time.Second * 3) {
	// 	go func() {
	// 		fmt.Print("yeet")
	// 	}()
	// }
}
