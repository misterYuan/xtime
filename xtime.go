package xtime

import (
	"errors"
	"math"
	"time"
)

// 今天的0点时间戳
func TodayZeroStamp() int64 {
	return TodayZeroTime().Unix()
}

// 今天的0点
func TodayZeroTime() time.Time {
	ts := time.Now().Format("20060102")
	t, _ := time.Parse("20060102", ts)
	return t
}

// 今天结束点
func TodayEndTime() time.Time {
	return TodayZeroTime().AddDate(0, 0, 1).Add(-1)
}

// 今天结束点时间戳
func TodayEndStamp() int64 {
	return TodayEndTime().Unix()
}

// 指定时间一天开始的时间
func Time2DayZeroTime(t time.Time) time.Time {
	ts := t.Format("20060102")
	tr, _ := time.Parse("20060102", ts)
	return tr
}

// 指定时间一天结束的时间
func Time2DayEndTime(t time.Time) time.Time {
	ts := Time2DayZeroTime(t)
	return ts.AddDate(0, 0, 1).Add(-1)
}

// 依次解析字符串格式数组为时间戳(解析成功停止解析)
type stampParser interface {
	Parse(string) (int64, error)
}

func NewStampParser(strs []string) stampParser {
	return &parseStrs2Stamp{strs}
}

type parseStrs2Stamp struct {
	strs []string //可以解析的字符串数组
}

func (p *parseStrs2Stamp) Parse(str string) (int64, error) {
	for _, v := range p.strs {
		if tm, err := time.Parse(v, str); err == nil {
			return tm.Unix(), nil
		}
	}
	return 0, errors.New("要解析的格式失败:" + str)
}

/*
两个时间戳相差的天数
*/
func DaysBetweenStamp(s1, s2 int64) int {
	return int(math.Abs(time.Unix(s1, 0).Sub(time.Unix(s2, 0)).Hours() / 24))
}
