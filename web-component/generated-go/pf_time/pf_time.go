package pf_time

import (
	"time"
)

type Time struct {
	Year       int
	Month      int
	Day        int
	Hour       int
	Minute     int
	Second     int
	Nanosecond int
	Location   string
}

var PIPEFISH_FUNCTION_CONVERTER = map[string](func(t uint32, v any) any){
	"Time": func(t uint32, v any) any {
		return Time{v.([]any)[0].(int), v.([]any)[1].(int), v.([]any)[2].(int), v.([]any)[3].(int), v.([]any)[4].(int), v.([]any)[5].(int), v.([]any)[6].(int), v.([]any)[7].(string)}
	},
}

var PIPEFISH_VALUE_CONVERTER = map[string]any{
	"Time": (*Time)(nil),
}

func Add(t Time, duration int) Time {
	return goToPf(pfToGo(t).Add(time.Duration(duration)))
}

func AddDate(t Time, years int, months int, days int) Time {
	return goToPf(pfToGo(t).AddDate(years, months, days))
}

func After(t Time, u Time) bool {
	return pfToGo(t).After(pfToGo(u))
}

func Before(t Time, u Time) bool {
	return pfToGo(t).Before(pfToGo(u))
}

func Compare(t Time, u Time) int {
	return pfToGo(t).Compare(pfToGo(u))
}

func Equal(t Time, u Time) bool {
	return pfToGo(t).Equal(pfToGo(u))
}

func Format(t Time, layout string) string {
	return pfToGo(t).Format(layout)
}

func IsDst(t Time) bool {
	return pfToGo(t).IsDST()
}

func ISOWeek(t Time) (int, int) {
	return pfToGo(t).ISOWeek()
}

func IsZero(t Time) bool {
	return pfToGo(t).IsZero()
}

func Local(t Time) Time {
	return goToPf(pfToGo(t).Local())
}

func Parse(layout string, value string) any {
	newTime, err := time.Parse(layout, value)
	if err != nil {
		return err
	}
	return goToPf(newTime)
}

func ParseDuration(s string) any {
	dur, err := time.ParseDuration(s)
	if err != nil {
		return err
	}
	return int(dur)
}

func Round(t Time, duration int) Time {
	return goToPf(pfToGo(t).Round(time.Duration(duration)))
}

func Sub(t Time, u Time) int {
	return int(pfToGo(t).Sub(pfToGo(u)))
}

func TimeIn(t Time, location string) any {
	newLoc, err := time.LoadLocation(location)
	if err != nil {
		return err
	}
	return goToPf(pfToGo(t).In(newLoc))
}

func TimeToUnix(t Time) int {
	return int(pfToGo(t).Unix())
}

func TimeToUnixMicro(t Time) int {
	return int(pfToGo(t).UnixMicro())
}

func TimeToUnixMilli(t Time) int {
	return int(pfToGo(t).UnixMilli())
}

func TimeToUnixNano(t Time) int {
	return int(pfToGo(t).UnixNano())
}

func Truncate(t Time, duration int) Time {
	return goToPf(pfToGo(t).Truncate(time.Duration(duration)))
}

func UnixToTime(sec int, nsec int) Time {
	return goToPf(time.Unix(int64(sec), int64(nsec)))
}

func UnixMicroToTime(usec int) Time {
	return goToPf(time.UnixMicro(int64(usec)))
}

func UnixMilliToTime(msec int) Time {
	return goToPf(time.UnixMilli(int64(msec)))
}

func UTC(t Time) Time {
	return goToPf(pfToGo(t).UTC())
}

func GoWeekday(t Time) int {
	return int((pfToGo(t)).Weekday())
}

func YearDay(t Time) int {
	return (pfToGo(t)).YearDay()
}

func GoGetClock() Time {
	goNow := time.Now()
	return Time{goNow.Year(), int(goNow.Month()), goNow.Day(), goNow.Hour(), goNow.Minute(), goNow.Second(), goNow.Nanosecond(), goNow.Location().String()}
}

func pfToGo(t Time) time.Time {
	locObj, _ := time.LoadLocation(t.Location)
	goMonth := time.Month(t.Month)
	return time.Date(t.Year, goMonth, t.Day, t.Hour, t.Minute, t.Second, t.Nanosecond, locObj)
}

func goToPf(t time.Time) Time {
	return Time{t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location().String()}
}
