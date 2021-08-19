package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/rlarkin212/gw2-tp-notifer/gw2service"
	"github.com/rlarkin212/gw2-tp-notifer/util"
)

const (
	gw2ApiBaseUrl = "https://api.guildwars2.com/v2"
)

var tgApi = util.GetEnvVar("TgApiKey")

func main() {
	// bot, err := tgbot.NewBotAPI(tgApi)
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	port := httpPort()
	http.HandleFunc("/", home)
	http.ListenAndServe(port, nil)

	for range time.Tick(time.Minute * 6) {
		getSales()
		log.Println("called GetSales")
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("yeet"))
}

func httpPort() string {
	port := "5000"

	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	return fmt.Sprintf(":%s", port)
}

func getSales() {
	sales := gw2service.FetchSales(gw2ApiBaseUrl)
	items := gw2service.FetchItems(gw2ApiBaseUrl, sales)
	log.Println(items)

	//telegramservice.SendMessage(bot, items)
}
