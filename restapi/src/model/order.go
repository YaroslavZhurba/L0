package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Delivery struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}

type Payment struct {
	Transaction  string `json:"transaction"`
	RequestId    string `json:"request_id"`
	Currency     string `json:"currency"`
	Provider     string `json:"provider"`
	Amount       int    `json:"amount"`
	PaymentDt    int    `json:"payment_dt"`
	Bank         string `json:"bank"`
	DeliveryCost int    `json:"delivery_cost"`
	GoodsTotal   int    `json:"goods_total"`
	CustomFee    int    `json:"custom_fee"`
}

type Item struct {
	CharId      int    `json:"chrt_id"`
	TrackNumber string `json:"track_number"`
	Price       int    `json:"price"`
	RId         string `json:"rid"`
	Name        string `json:"name"`
	Sale        int    `json:"sale"`
	Size        string `json:"size"`
	TotalPrice  int    `json:"total_price"`
	NmId        int    `json:"nm_id"`
	Brand       string `json:"brand"`
	Status      int    `json:"status"`
}

type OrderJson struct {
	OrderUid          string   `json:"order_uid"`
	TrackNumber       string   `json:"track_number"`
	Entry             string   `json:"entry"`
	Payment           Payment  `json:"payment"`
	Delivery          Delivery `json:"delivery"`
	Items             []Item   `json:"items"`
	Locale            string   `json:"locale"`
	InternalSignature string   `json:"internal_signature"`
	CustomerId        string   `json:"customer_id"`
	DeliveryService   string   `json:"delivery_service"`
	Shardkey          string   `json:"shardkey"`
	SmId              int      `json:"sm_id"`
	DateCreated       string   `json:"date_created"`
	OofShard          string   `json:"oof_shard"`

}

func (a OrderJson) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *OrderJson) Scan(value interface{}) error {
    b, ok := value.([]byte)
    if !ok {
        return errors.New("type assertion to []byte failed")
    }

    return json.Unmarshal(b, &a)
}

type Order struct {
	Id        int       `json:"id"`
	OrderJson OrderJson `json:"order_json"`
}

func OrderDefault() (*Order){
	orderDefault := Order{
		Id: 1,
		OrderJson: OrderJson {
			OrderUid: "b563feb7b2b84b6test",
			TrackNumber: "WBILMTESTTRACK",
			Entry: "WBIL",
			Delivery: Delivery {
				Name: "Test Testov",
				Phone: "+9720000000",
				Zip: "2639809",
				City: "Kiryat Mozkin",
				Address: "Ploshad Mira 15",
				Region: "Kraiot",
				Email: "test@gmail.com"},
			Payment: Payment{
				Transaction: "b563feb7b2b84b6test",
				RequestId: "",
				Currency: "USD",
				Provider: "wbpay",
				Amount: 1817,
				PaymentDt: 1637907727,
				Bank: "alpha",
				DeliveryCost: 1500,
				GoodsTotal: 317,
				CustomFee: 0},
			Items: []Item{
				{
					CharId: 9934930,
					TrackNumber: "WBILMTESTTRACK",
					Price: 453,
					RId: "ab4219087a764ae0btest",
					Name: "Mascaras",
					Sale: 30,
					Size: "0",
					TotalPrice: 317,
					NmId: 2389212,
					Brand: "Vivienne Sabo",
					Status: 202}},
			Locale: "en",
			InternalSignature: "",
			CustomerId: "test",
			DeliveryService: "meest",
			Shardkey: "9",
			SmId: 99,
			DateCreated: "2021-11-26T06:22:19Z",
			OofShard: "1"}}
	return &orderDefault
}

