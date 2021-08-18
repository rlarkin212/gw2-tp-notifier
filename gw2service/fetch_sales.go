package gw2service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/rlarkin212/gw2-tp-notifer/models"
	"github.com/rlarkin212/gw2-tp-notifer/util"
)

var gw2ApiToken = util.GetEnvVar("ApiKey")
var environment = util.GetEnvVar("Environment")

func FetchSales(baseUrl string) []models.Sale {
	url := fmt.Sprintf("%s/commerce/transactions/history/sells?access_token=%s", baseUrl, gw2ApiToken)
	currentTime := time.Now()
	offsetTime := currentTime.Add(-6 * time.Minute)

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err.Error())
	}

	rawSales := []models.Sale{}
	_ = unmarshallSales(res, &rawSales)

	sales := []models.Sale{}

	//if not dev get last 6 mins (api cache ezxpires every 5)
	if environment != "dev" {
		for _, sale := range rawSales {
			if inTimeSpan(offsetTime, currentTime, sale.Purchased) {
				sales = append(sales, sale)
			}
		}
	} else {
		sales = rawSales[:5]
	}

	return sales
}

func FetchSales2(baseUrl string, salesChan chan models.Sale) {
	url := fmt.Sprintf("%s/commerce/transactions/history/sells?access_token=%s", baseUrl, gw2ApiToken)
	currentTime := time.Now()
	offsetTime := currentTime.Add(-6 * time.Minute)

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err.Error())
	}

	rawSales := []models.Sale{}
	_ = unmarshallSales(res, &rawSales)

	//if not dev get last 6 mins (api cache ezxpires every 5)
	if environment != "dev" {
		for _, sale := range rawSales {
			if inTimeSpan(offsetTime, currentTime, sale.Purchased) {
				salesChan <- sale
			}
		}
	} else {
		for _, s := range rawSales[:5] {
			salesChan <- s
		}
	}
}

func unmarshallSales(res *http.Response, target *[]models.Sale) error {
	defer res.Body.Close()
	return json.NewDecoder(res.Body).Decode(target)
}

func inTimeSpan(after time.Time, before time.Time, control time.Time) bool {
	if control.After(after) && control.Before(before) {
		return true
	}

	return false
}
