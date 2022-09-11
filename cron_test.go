package gocron

import (
	"testing"
	"time"
)

func TestParser(t *testing.T) {
	tests := []struct {
		name string
		args *cronSchedule
		want string
	}{
		{
			name: "All Values Cases",
			args: Schedule().Weeks(time.Monday, time.Sunday).Months(time.May, time.April).Days(1, 30).Hours(1, 10).Minutes(30, 50),
			want: "30,50 1,10 1,30 5,4 1,0",
		},
		{
			name: "All Range Cases",
			args: Schedule().WeeksRange(time.Monday, time.Friday).MonthsRange(time.April, time.August).DaysRange(1, 30).HoursRange(1, 10).MinutesRange(30, 50),
			want: "30-50 1-10 1-30 4-8 1-5",
		},
		{
			name: "All Interval Cases",
			args: Schedule().WeeksInterval(2).MonthsInterval(5).DaysInterval(20).HoursInterval(10).MinutesInterval(30),
			want: "*/30 */10 */20 */5 */2",
		},
		{
			name: "Complex Cases 1",
			args: Schedule().DaysRangedInterval(1, 31, 2).Weeks(time.Monday).Hours(0, 1, 2, 3).MinutesInterval(10),
			want: "*/10 0,1,2,3 1-31/2 * 1",
		},
		{
			name: "Complex Cases 2",
			args: Schedule().MonthsRange(time.May, time.September).WeeksRangedInterval(time.Monday, time.Friday, 2),
			want: "* * * 5-9 1-5/2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := tt.args.Parse(); got != tt.want || err != nil {
				if err != nil {
					t.Errorf("Caused Error %v", err)
				} else {
					t.Errorf("Parser Result = %v, but you wanted %v", got, tt.want)
				}
			}
		})
	}
}

func TestParserError(t *testing.T) {
	tests := []struct {
		name string
		args *cronSchedule
	}{
		// minutes
		{
			name: "Invalid Minutes: max",
			args: Schedule().Minutes(0, 60),
		},
		{
			name: "Invalid Minutes: min",
			args: Schedule().Minutes(-1, 59),
		},
		{
			name: "Invalid MinutesRange: max",
			args: Schedule().MinutesRange(0, 60),
		},
		{
			name: "Invalid MinutesRange: min",
			args: Schedule().MinutesRange(-1, 59),
		},
		{
			name: "Invalid MinutesInterval: max",
			args: Schedule().MinutesInterval(60),
		},
		{
			name: "Invalid MinutesInterval: min",
			args: Schedule().MinutesInterval(0),
		},
		{
			name: "Invalid MinutesRange",
			args: Schedule().MinutesRange(59, 0),
		},
		{
			name: "Invalid MinutesRangedInterval",
			args: Schedule().MinutesRangedInterval(59, 0, 2),
		},
		// hours
		{
			name: "Invalid Hours: max",
			args: Schedule().Hours(0, 24),
		},
		{
			name: "Invalid Hours: min",
			args: Schedule().Hours(-1, 23),
		},
		{
			name: "Invalid HoursRange: max",
			args: Schedule().HoursRange(0, 24),
		},
		{
			name: "Invalid HoursRange: min",
			args: Schedule().HoursRange(-1, 23),
		},
		{
			name: "Invalid HoursInterval: max",
			args: Schedule().HoursInterval(24),
		},
		{
			name: "Invalid HoursInterval: min",
			args: Schedule().HoursInterval(0),
		},
		{
			name: "Invalid HoursRange",
			args: Schedule().HoursRange(23, 0),
		},
		{
			name: "Invalid HoursRangedInterval",
			args: Schedule().HoursRangedInterval(23, 0, 2),
		},
		// days
		{
			name: "Invalid Days: max",
			args: Schedule().Days(1, 32),
		},
		{
			name: "Invalid Days: min",
			args: Schedule().Days(0, 31),
		},
		{
			name: "Invalid DaysRange: max",
			args: Schedule().DaysRange(1, 32),
		},
		{
			name: "Invalid DaysRange: min",
			args: Schedule().DaysRange(0, 31),
		},
		{
			name: "Invalid DaysInterval: max",
			args: Schedule().DaysInterval(32),
		},
		{
			name: "Invalid DaysInterval: min",
			args: Schedule().DaysInterval(0),
		},
		{
			name: "Invalid DaysRange",
			args: Schedule().DaysRange(31, 1),
		},
		{
			name: "Invalid DaysRangedInterval",
			args: Schedule().DaysRangedInterval(31, 1, 2),
		},
		// Months
		{
			name: "Invalid MonthsRange",
			args: Schedule().MonthsRange(time.February, time.January),
		},
		{
			name: "Invalid MonthsRangedInterval",
			args: Schedule().MonthsRangedInterval(time.February, time.January, 2),
		},
		// Weeks
		{
			name: "Invalid WeeksRange",
			args: Schedule().WeeksRange(time.Saturday, time.Sunday),
		},
		{
			name: "Invalid WeeksRangedInterval",
			args: Schedule().WeeksRangedInterval(time.Saturday, time.Sunday, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := tt.args.Parse(); err == nil {
				t.Errorf("Something Went Wrong... %s", got)
			}
		})
	}
}
