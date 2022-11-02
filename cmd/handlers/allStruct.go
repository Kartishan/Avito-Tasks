package handlers

import (
	"Avito-Tasks/cmd/data"
	"log"
)

const version = "1.0.0"

type Сonfig struct {
	Port int
	Env  string
	Db   struct {
		Dsn string
	}
}

type Application struct {
	Config Сonfig
	Logger *log.Logger
	Models data.Models
}
type Transaction struct {
	TransactionId   int64   `json:"transaction_id"`
	SenderId        int64   `json:"sender_id"`
	ReceiverId      int64   `json:"receiver_id"`
	OperationId     int64   `json:"operation_id"`
	TransactionTime string  `json:"transaction_time"`
	Price           float64 `json:"price"`
}
