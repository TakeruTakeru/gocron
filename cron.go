package gocron

import (
	"fmt"
	"math"
	"strings"
	"time"
)

const (
	maxMinutes = 59
	minMinutes = 0
	maxHours   = 23
	minHours   = 0
	maxDays    = 31
	minDays    = 1
	minMonth   = time.January
	minWeekday = time.Monday
)

type monthType = time.Month

type weekdayType = time.Weekday

type cronNumber interface {
	weekdayType | monthType | int
}

type cronTime interface {
	parse() string
}

type cronSchedule struct {
	minute cronTime
	hour   cronTime
	day    cronTime
	month  cronTime
	week   cronTime
	err    error
}

func (c cronSchedule) Parse() (string, error) {
	if c.err != nil {
		return "", c.err
	}
	return c.String(), nil
}

func (c cronSchedule) String() string {
	return fmt.Sprintf("%s %s %s %s %s", c.minute.parse(), c.hour.parse(), c.day.parse(), c.month.parse(), c.week.parse())
}

func (c *cronSchedule) Minutes(minutes ...int) *cronSchedule {
	if err := getThresholdError("minutes", maxMinutes, minMinutes, minutes); err != nil {
		c.err = err
	}
	values := newValues(minutes)
	c.minute = &values
	return c
}

func (c *cronSchedule) MinutesRange(from, to int) *cronSchedule {
	if err := getThresholdError("minutes", maxMinutes, minMinutes, []int{from, to}); err != nil {
		c.err = err
	}
	if err := getRangeError("minutes", from, to); err != nil {
		c.err = err
	}
	c.minute = newRange(from, to)
	return c
}

func (c *cronSchedule) MinutesInterval(parameter int) *cronSchedule {
	if err := getThresholdError("minutes", maxMinutes, minMinutes+1, []int{parameter}); err != nil {
		c.err = err
	}
	c.minute = newInterval(parameter)
	return c
}

func (c *cronSchedule) MinutesRangedInterval(from, to, parameter int) *cronSchedule {
	if err := getThresholdError("minutes", from, to, []int{parameter}); err != nil {
		c.err = err
	}
	if err := getThresholdError("minutes", math.MaxInt, minMinutes+1, []int{parameter}); err != nil {
		c.err = err
	}
	if err := getRangeError("minutes", from, to); err != nil {
		c.err = err
	}
	c.minute = newRangedInterval(from, to, parameter)
	return c
}

func (c *cronSchedule) Hours(hours ...int) *cronSchedule {
	if err := getThresholdError("hours", maxHours, minHours, hours); err != nil {
		c.err = err
	}
	values := newValues(hours)
	c.hour = &values
	return c
}

func (c *cronSchedule) HoursRange(from, to int) *cronSchedule {
	if err := getThresholdError("hours", maxHours, minHours, []int{from, to}); err != nil {
		c.err = err
	}
	if err := getRangeError("hours", from, to); err != nil {
		c.err = err
	}
	c.hour = newRange(from, to)
	return c
}

func (c *cronSchedule) HoursInterval(parameter int) *cronSchedule {
	if err := getThresholdError("hours", maxHours, minHours+1, []int{parameter}); err != nil {
		c.err = err
	}
	c.hour = newInterval(parameter)
	return c
}

func (c *cronSchedule) HoursRangedInterval(from, to, parameter int) *cronSchedule {
	if err := getThresholdError("hours", maxHours, minHours, []int{from, to}); err != nil {
		c.err = err
	}
	if err := getThresholdError("hours", maxHours, minHours+1, []int{parameter}); err != nil {
		c.err = err
	}
	if err := getRangeError("hours", from, to); err != nil {
		c.err = err
	}
	c.hour = newRangedInterval(from, to, parameter)
	return c
}

func (c *cronSchedule) Days(days ...int) *cronSchedule {
	if err := getThresholdError("days", maxDays, minDays, days); err != nil {
		c.err = err
	}
	values := newValues(days)
	c.day = &values
	return c
}

func (c *cronSchedule) DaysRange(from, to int) *cronSchedule {
	if err := getThresholdError("days", maxDays, minDays, []int{from, to}); err != nil {
		c.err = err
	}
	if err := getRangeError("days", from, to); err != nil {
		c.err = err
	}
	c.day = newRange(from, to)
	return c
}

func (c *cronSchedule) DaysInterval(parameter int) *cronSchedule {
	if err := getThresholdError("days", maxDays, minDays+1, []int{parameter}); err != nil {
		c.err = err
	}
	c.day = newInterval(parameter)
	return c
}

func (c *cronSchedule) DaysRangedInterval(from, to, parameter int) *cronSchedule {
	if err := getThresholdError("days", maxDays, minDays, []int{from, to}); err != nil {
		c.err = err
	}
	if err := getThresholdError("days", maxDays, minDays+1, []int{parameter}); err != nil {
		c.err = err
	}
	if err := getRangeError("days", from, to); err != nil {
		c.err = err
	}
	c.day = newRangedInterval(from, to, parameter)
	return c
}

func (c *cronSchedule) Months(month ...monthType) *cronSchedule {
	values := newValues(month)
	c.month = &values
	return c
}

func (c *cronSchedule) MonthsRange(from, to monthType) *cronSchedule {
	if err := getRangeError("months", int(from), int(to)); err != nil {
		c.err = err
	}
	c.month = newRange(from, to)
	return c
}

func (c *cronSchedule) MonthsInterval(parameter int) *cronSchedule {
	if err := getThresholdError("months", math.MaxInt, int(minMonth), []int{parameter}); err != nil {
		c.err = err
	}
	c.month = newInterval(parameter)
	return c
}

func (c *cronSchedule) MonthsRangedInterval(from, to monthType, parameter int) *cronSchedule {
	if err := getThresholdError("months", math.MaxInt, int(minMonth), []int{parameter}); err != nil {
		c.err = err
	}
	if err := getRangeError("months", int(from), int(to)); err != nil {
		c.err = err
	}
	c.month = newRangedInterval(int(from), int(to), parameter)
	return c
}

func (c *cronSchedule) Weeks(weekday ...weekdayType) *cronSchedule {
	values := newValues(weekday)
	c.week = &values
	return c
}

func (c *cronSchedule) WeeksRange(from, to weekdayType) *cronSchedule {
	if err := getRangeError("weeks", int(from), int(to)); err != nil {
		c.err = err
	}
	c.week = newRange(from, to)
	return c
}

func (c *cronSchedule) WeeksInterval(parameter int) *cronSchedule {
	if err := getThresholdError("weeks", math.MaxInt, int(minWeekday), []int{parameter}); err != nil {
		c.err = err
	}
	c.week = newInterval(parameter)
	return c
}

func (c *cronSchedule) WeeksRangedInterval(from, to weekdayType, parameter int) *cronSchedule {
	if err := getThresholdError("weeks", math.MaxInt, int(minWeekday), []int{parameter}); err != nil {
		c.err = err
	}
	if err := getRangeError("weeks", int(from), int(to)); err != nil {
		c.err = err
	}
	c.week = newRangedInterval(int(from), int(to), parameter)
	return c
}

func Schedule() *cronSchedule {
	return &cronSchedule{
		minute: defaultCronTime{},
		hour:   defaultCronTime{},
		day:    defaultCronTime{},
		month:  defaultCronTime{},
		week:   defaultCronTime{},
	}
}

type cronValues[T cronNumber] []T

func (v cronValues[T]) isValid() bool {
	return len(v) > 0
}

func (v cronValues[T]) parse() string {
	if !v.isValid() {
		return "*"
	}
	stringValues := make([]string, len(v))
	for i, value := range v {
		stringValues[i] = fmt.Sprintf("%d", value)
	}
	return strings.Join(stringValues, ",")
}

type cronRange[T cronNumber] struct {
	from T
	to   T
}

func (r cronRange[T]) isValid() bool {
	return r.from > 0 || r.to > 0 && r.to > r.from
}

func (r cronRange[T]) parse() string {
	if !r.isValid() {
		return "*"
	}
	return fmt.Sprintf("%d-%d", r.from, r.to)
}

type cronInterval[T cronNumber] struct {
	parameter T
}

func (i cronInterval[T]) isValid() bool {
	return i.parameter > 1
}

func (i cronInterval[T]) parse() string {
	if !i.isValid() {
		return "*"
	}
	parameter := i.parameter
	return fmt.Sprintf("*/%d", parameter)
}

type cronRangedInterval[T cronNumber] struct {
	molecule  cronRange[T]
	parameter T
}

func (i cronRangedInterval[T]) isValid() bool {
	return i.molecule.isValid() && i.parameter > 1
}

func (i cronRangedInterval[T]) parse() string {
	if !i.isValid() {
		return "*"
	}
	molecule := i.molecule
	parameter := i.parameter
	moleculeStr := molecule.parse()
	return fmt.Sprintf("%s/%d", moleculeStr, parameter)
}

type defaultCronTime struct {
}

func (d defaultCronTime) parse() string {
	return "*"
}

func newValues[T cronNumber](v []T) cronValues[T] {
	return append([]T{}, v...)
}

func newRange[T cronNumber](from, to T) *cronRange[T] {
	ranges := &cronRange[T]{
		from, to,
	}
	return ranges
}

func newInterval[T cronNumber](parameter T) *cronInterval[T] {
	ranges := &cronInterval[T]{
		parameter,
	}
	return ranges
}

func newRangedInterval[T cronNumber](from, to, parameter T) *cronRangedInterval[T] {
	ranges := newRange(from, to)
	interval := &cronRangedInterval[T]{
		molecule:  *ranges,
		parameter: parameter,
	}
	return interval
}
