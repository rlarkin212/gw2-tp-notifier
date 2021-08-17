package gw2service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rlarkin212/gw2-tp-notifer/models"
)

func FetchItems(baseUrl string, sales []models.Sale) []models.Item {
	var buf bytes.Buffer

	buf.Write([]byte(fmt.Sprintf("%s/items?ids=", baseUrl)))
	for _, sale := range sales {
		buf.Write([]byte(fmt.Sprintf("%d,", sale.ItemID)))
	}
	url := buf.String()

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err.Error())
	}

	items := []models.Item{}
	_ = unmarshallItems(res, &items)

	return items
}

func unmarshallItems(res *http.Response, target *[]models.Item) error {
	defer res.Body.Close()
	return json.NewDecoder(res.Body).Decode(target)
}
