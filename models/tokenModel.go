package models

import "time"

type Token struct {
	Id              int       `json:"id"`
	Timestamp       time.Time `json:"timeStamp"`
	BlockNumber     int       `json:"blockNumber"`
	TransactionHash string    `json:"transactionHash"`
	ContractAddress string    `json:"contractAddress"`
	Value           string    `json:"value"`
	FromAddress     string    `json:"fromAddress"`
	ToAddress       string    `json:"toAddress"`
}
