package models

type Monitior struct {
	Id       int64
	Userid   string
	Contract string
	Address  string
	Time     string
	Value    string
}

type ResultMoitor struct {
	Code int
	Msg  string
	Id   string
}

type GetMonitor struct {
	Userid   string
	Contract string
	Address  string
	duration string
	value    string
	Id       string
}
