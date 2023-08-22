package models

import (
	"encoding/json"
	"time"
)

type Order struct {
	OrderUID    string          `json:"order_uid"`
	DateCreated time.Time       `json:"date_created"`
	Data        json.RawMessage `json:"data"`
}
