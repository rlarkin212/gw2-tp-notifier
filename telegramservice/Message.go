package telegramservice

import (
	"fmt"
	"strconv"
	"strings"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/rlarkin212/gw2-tp-notifer/models"
	"github.com/rlarkin212/gw2-tp-notifer/util"
)

var tgChatId, _ = strconv.ParseInt(util.GetEnvVar("TgChatId"), 10, 64)

func SendMessage(bot *tgbot.BotAPI, items []models.Item) {
	for _, item := range items {
		var b strings.Builder

		b.WriteString(fmt.Sprintf("Item: %s", item.Name))
		b.WriteString(fmt.Sprintf("\nPrice %s", util.PriceToGold(item.Sale.Price)))
		b.WriteString(fmt.Sprintf("\nQty: %d", item.Sale.Quantity))
		b.WriteString(fmt.Sprintf("\n%s", item.Icon))

		str := b.String()

		msg := tgbot.NewMessage(tgChatId, str)
		bot.Send(msg)
	}
}
