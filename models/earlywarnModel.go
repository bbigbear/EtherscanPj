package models

type EarlyWarn struct {
	Id       int
	Status   string
	Snum     float64
	Spercent float64
	Dnum     float64
	Dpercent float64
	Hour     int
	Hnum     float64
	Hpercent float64
	Tel      string
}
