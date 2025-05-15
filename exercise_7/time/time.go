package time

type Time struct {
	TwentyFourHourTime *TwentyFourHourTime 
	TwelveHourTime *TwelveHourTime
}

type TwentyFourHourTime struct {
	Hours int 
	Minutes int
}

type TwelveHourTime struct {
	Hours int 
	Minutes int
	Period string
}

