package models

import "time"

type Notifcation struct {
	Id                int
	Timestamp         time.Time
	Target            string
	Style             string
	Num               string
	Percent           string
	Hash              string
	TransactionStatus string
}
