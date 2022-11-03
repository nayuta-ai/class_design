package record

type Record struct {
	Distance float64
	Time Time
}

type Time struct {
	Hours int64
	Minutes float64
}
