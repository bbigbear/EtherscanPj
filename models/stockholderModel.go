package models

type Stockholder struct {
	ID      int64  `orm:"column(ID)"`
	NAME    string `orm:"column(NAME)"`
	TEL     string `orm:"column(TEL)"`
	NUM     string `orm:"column(NUM)"`
	ADDRESS string `orm:"column(ADDRESS)"`
}
