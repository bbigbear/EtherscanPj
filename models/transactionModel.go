package models

type Transaction struct {
	Id                int
	Address           string
	BlockNumber       string
	TimeStamp         string
	Hash              string
	Nonce             string
	BlockHash         string
	TransactionIndex  string
	From              string
	To                string
	Value             string
	Gas               string
	GasPrice          string
	IsError           string
	Txreceipt_status  string
	Input             string
	ContractAddress   string
	CumulativeGasUsed string
	GasUsed           string
	Confirmations     string
}
