package main

import (
	"fmt"

	"github.com/rlarkin212/gw2-tp-notifer/gw2service"
)

const (
	gw2ApiBaseUrl = "https://api.guildwars2.com/v2"
)

func main() {
	sales := gw2service.FetchSales(gw2ApiBaseUrl)
	items := gw2service.FetchItems(gw2ApiBaseUrl, sales)

	for _, item := range items {
		fmt.Printf("%+v\n", item)
	}

	// for range time.Tick(time.Second * 3) {
	// 	go func() {
	// 		fmt.Print("yeet")
	// 	}()
	// }
}
