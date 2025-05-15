package parse

import (
	"testing"

	"github.com/mikelinesta/parse-time/time"
)

func TestParse(t *testing.T){
	parseTests := []struct {
		name string 
		exp time.Time
		shouldErr bool 
	}{	
		{
			"4pm",
			time.Time{
				TwentyFourHourTime: nil,
				TwelveHourTime: &time.TwelveHourTime{
					Hours: 4,
					Minutes:0,
					Period: "pm",
				},
			},
			false,	
		},
	}
}