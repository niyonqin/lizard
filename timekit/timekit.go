package timekit

import "time"

// DurationToMillis converts duration to milliseconds.
func DurationToMillis(d time.Duration) int64 {
	return int64(d / time.Millisecond)
}

// NowInMillis returns timestamp in milliseconds.
func NowInMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// NowInSecs returns timestamp in seconds.
func NowInSecs() int64 {
	return time.Now().Unix()
}


// UTCNowTime returns current time in utc.
func UTCNowTime() time.Time {
	return time.Now().UTC()
}

// SecondToUTCTime returns second to time
func SecondToUTCTime(t int64) time.Time{
	return time.Unix(t, 0).UTC()
}

// SecondToUTCTime returns utc time to second
func UTCTimeToSecond(t time.Time) int64 {
	if t.IsZero() {
		return 0
	} else {
		return t.UTC().Unix()
	}
}

// TodayZeroUTCTime return 当前北京时间的当天的0点0时0分的 utc的时间
func TodayZeroUTCTime() time.Time {
	now := UTCNowTime()
	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	bjZone := time.FixedZone("Beijing", secondsEastOfUTC)
	bjTime := now.In(bjZone)

	zeroBjTime := time.Date(bjTime.Year(), bjTime.Month(), bjTime.Day(), 0, 0, 0, 0, bjZone)
	return zeroBjTime.In(time.UTC)
}

// NowToYMDhms returns current time which is str. eg: 20060102150405
func NowToYMDhms() string{
	return UTCNowTime().Format("20060102150405")
}

// NowToYMD returns current time which is str. eg: 20060102
func NowToYMD() string{
	return UTCNowTime().Format("20060102")
}

// NowToStandardDate e.g. 2017-01-09 03:12:01
func NowToStandardDate() string{
	return UTCNowTime().Format("2006-01-02 15:04:05")
}

// SecondToStandardDate  e.g.  1604996388 to  2017-01-09 03:12:01
func SecondToStandardDate(s int64) string{
	tmpTime := time.Unix(s, 0)
	return tmpTime.UTC().Format("2006-01-02 15:04:05")
}

// Diff return t1 - t2
func Diff(t1, t2 time.Time) float64 {
	return t1.Sub(t2).Seconds()
}

// e.g. 2020-11-10 00:00:00 +0800 CST
func SecondToZeroLocTime(t int64) time.Time {
	tmpTime := time.Unix(t, 0)
	year, month, day := tmpTime.Date()
	dateTime := time.Date(year, month, day, 0, 0, 0, 0, tmpTime.Location())
	return dateTime
}

// MaxTime returns maxTime
func MaxTime() time.Time {
	t, _ := time.Parse("2006-01-02 15:04:05", "9999-12-31 00:00:00")
	return t
}

// IsBetween judges targetTime is between startTime and endTime
func IsBetween(targetTime, startTime, endTime time.Time) bool {
	targetStamp := UTCTimeToSecond(targetTime)
	startStamp := UTCTimeToSecond(startTime)
	endStamp := UTCTimeToSecond(endTime)
	return targetStamp <= endStamp && targetStamp >= startStamp
}

// MillsSinceTime returns millSecond since start time
func MillsSinceTime(start time.Time) int64 {
	return time.Since(start).Nanoseconds() / int64(time.Millisecond)
}

// AddDateToStr
func AddDateToStr(t time.Time, year,month,day int) string{
	return t.AddDate(year, month, day).Format("2006-01-02")
}

// SubDateToStr
func SubDateToStr(t time.Time, year,month,day int) string{
	return t.AddDate(-year, -month, -day).Format("2006-01-02")
}
