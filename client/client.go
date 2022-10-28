package main

import (
	"fmt"
	// "time"
	"strconv"
	// "log"
	// nats "github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)

const (
	URL = "nats://127.0.0.1:4222"
)

func getJsonOK(id int) (string) {
	return `{
		"id" :` + strconv.Itoa(id) + `,
		"order_json" : {
		  "order_uid": "b563feb7b2b84b6test",
		  "track_number": "WBILMTESTTRACK",
		  "entry": "WBIL",
		  "delivery": {
			"name": "Test Testov",
			"phone": "+9720000000",
			"zip": "2639809",
			"city": "Kiryat Mozkin",
			"address": "Ploshad Mira 15",
			"region": "Kraiot",
			"email": "test@gmail.com"
		  },
		  "payment": {
			"transaction": "b563feb7b2b84b6test",
			"request_id": "",
			"currency": "USD",
			"provider": "wbpay",
			"amount": 1817,
			"payment_dt": 1637907727,
			"bank": "alpha",
			"delivery_cost": 1500,
			"goods_total": 317,
			"custom_fee": 0
		  },
		  "items": [
			{
			  "chrt_id": 9934930,
			  "track_number": "WBILMTESTTRACK",
			  "price": 453,
			  "rid": "ab4219087a764ae0btest",
			  "name": "Mascaras",
			  "sale": 30,
			  "size": "0",
			  "total_price": 317,
			  "nm_id": 2389212,
			  "brand": "Vivienne Sabo",
			  "status": 202
			}
		  ],
		  "locale": "en",
		  "internal_signature": "",
		  "customer_id": "test",
		  "delivery_service": "meest",
		  "shardkey": "9",
		  "sm_id": 99,
		  "date_created": "2021-11-26T06:22:19Z",
		  "oof_shard": "1"
		}
	  }`
}

func getJsonBad() (string) {
	return `{{{{{
	  }`
}

func main() {

	sc, err := stan.Connect("test-cluster", "pub-1")
	if err != nil {
		panic(err)
	}
	defer sc.Close()

	var channelName string = "channel-1"
	var str string
	fmt.Println("Started")

	for ;; {
		fmt.Scanln(&str)
		if str == "exit" {
			break
		}
		if str == "chan" {
			fmt.Scanln(&channelName)
			fmt.Println("New channel to publish is " + channelName)
		}
		if str == "pub-good" {
			var id int
			fmt.Println("Print id, to generate json:")
			fmt.Scanln(&id)
			sc.Publish(channelName, []byte(getJsonOK(id)))
			// fmt.Println(getJsonOK(id))
		}
		if str == "pub-bad" {
			sc.Publish(channelName, []byte(getJsonBad()))
		}
	}

}
