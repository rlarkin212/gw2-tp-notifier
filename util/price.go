package util

import (
	"fmt"
	"strconv"
)

func PriceToGold(price int64) string {
	p := strconv.Itoa(int(price))
	formattedPrice := fmt.Sprintf("%06s", p)

	gold := formattedPrice[0:2]
	silver := formattedPrice[2:4]
	copper := formattedPrice[4:6]

	formatted := fmt.Sprintf("%s G %s S %s C", gold, silver, copper)

	return formatted
}
