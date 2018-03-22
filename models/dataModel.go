package models

import "time"

type Data struct {
	Id              int
	Timestamp       time.Time
	BlockNumber     string
	TransactionHash string
	ContractAddress string
	Value           string
	FromAddress     string
	ToAddress       string
	Percent         string
	Status          string
}
