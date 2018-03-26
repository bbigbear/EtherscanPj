package models

import "time"

type WalletNum struct {
	Id         int
	Timestamp  time.Time
	AddressNum int
}
