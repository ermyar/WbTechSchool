package json

import "time"

type OrderJSON struct {
	Order_uid          string       `json:"order_uid" faker:"uuid_digit"`
	Track_number       string       `json:"track_number" `
	Entry              string       `json:"entry"`
	Locate             string       `json:"locate"`
	Customer_id        string       `json:"customer_id"`
	Delivery_service   string       `json:"delivery_service"`
	Shardkey           string       `json:"shardkey"`
	Sm_id              int          `json:"sm_id"`
	Date_created       time.Time    `json:"Date_created"`
	Oof_shard          string       `json:"oof_shard"`
	Internal_signature string       `json:"internal_signature"`
	Delivery           DeliveryJSON `json:"delivery"`
	Items              []ItemJSON   `json:"items"`
	Payment            PaymentJSON  `json:"payment"`
}

type DeliveryJSON struct {
	Name    string `json:"name" faker:"name"`
	Phone   string `json:"phone" faker:"phone_number"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email" faker:"email"`
}

type ItemJSON struct {
	Chrt_id      int    `json:"chrt_id"`
	Track_number string `json:"track_number"`
	Price        int    `json:"price"`
	Rid          string `json:"rid"`
	Name         string `json:"name" faker:"word,unique"`
	Sale         int    `json:"sale" faker:"boundary_start=0, boundary_end=100"`
	Size         int    `json:"size"`
	Total_price  int    `json:"total_price"`
	Nm_id        int    `json:"nm_id"`
	Brand        string `json:"brand"`
	Status       int    `json:"status"`
}

type PaymentJSON struct {
	Transaction   string `json:"transaction"`
	Request_id    string `json:"request_id"`
	Currency      string `json:"currency" faker:"currency"`
	Provider      string `json:"provider"`
	Amount        int    `json:"amount"`
	Payment_dt    int    `json:"payment_dt" faker:"unix_time"`
	Bank          string `json:"bank"`
	Delivery_cost int    `json:"delivery_cost"`
	Goods_total   int    `json:"goods_total"`
	Custom_fee    int    `json:"custom_fee"`
}
