package models

import "time"

type Sale struct {
	ID        int64     `json:"id"`
	ItemID    int64     `json:"item_id"`
	Price     int64     `json:"price"`
	Quantity  int64     `json:"quantity"`
	Created   time.Time `json:"created"`
	Purchased time.Time `json:"purchased"`
}
