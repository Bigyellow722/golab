package timeop

import (
	"time"
)

const (
	minute int64 = 60
	hour   int64 = 3600
)

func gettimeDiffer(start_time, end_time string, base int64) int64 {
	var timetmp int64
	t1, err := time.ParseInLocation("2006-01-02 15:04:05", start_time, time.Local)
	t2, err := time.ParseInLocation("2006-01-02 15:04:05", end_time, time.Local)
	if err == nil && t1.Before(t2) {
		diff := t2.Unix() - t1.Unix()
		timetmp = diff / base
		return timetmp
	} else {
		return timetmp
	}
}

func Gethourdiffer(start_time, end_time string) int64 {
	return gettimeDiffer(start_time, end_time, hour)
}

func Getminutediffer(start_time, end_time string) int64 {
	return gettimeDiffer(start_time, end_time, minute)
}
