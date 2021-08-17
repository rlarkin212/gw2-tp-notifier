package gw2service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rlarkin212/gw2-tp-notifer/models"
)

func FetchItems(baseUrl string, sales []models.Sale) []models.Item {
	items := []models.Item{}

	for _, sale := range sales {
		url := fmt.Sprintf("%s/items?id=%d", baseUrl, sale.ItemID)

		res, err := http.Get(url)
		if err != nil {
			log.Fatal(err.Error())
		}

		item := models.Item{
			Price: sale.Price,
		}
		_ = unmarshallItem(res, &item)

		items = append(items, item)
	}

	return items
}

func unmarshallItem(res *http.Response, target *models.Item) error {
	defer res.Body.Close()
	return json.NewDecoder(res.Body).Decode(target)
}
