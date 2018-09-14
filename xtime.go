package xtime

import (
	"errors"
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
	strs []string //初始化可以解析的字符串数组
}

func (p *parseStrs2Stamp) Parse(str string) (int64, error) {
	for _, v := range p.strs {
		if tm, err := time.Parse(v, str); err == nil {
			return tm.Unix(), nil
		}
	}
	return 0, errors.New("要解析的格式失败:" + str)
}

// func ParseStrs2Stamp(strs []string) int64 {
// 	if str == "" {
// 		return 0
// 	}
// 	if tm, err := time.Parse("2006/1/2 15:04", str); err == nil {
// 		return tm.Unix()
// 	}
// 	if tm, err := time.Parse("2006/1/2", str); err == nil {
// 		return tm.Unix()
// 	}
// 	if tm, err := time.Parse("2006-01-02 15:04:05", str); err == nil {
// 		return tm.Unix()
// 	}
// 	if tm, err := time.Parse("2006-01-02", str); err == nil {
// 		return tm.Unix()
// 	}
// 	panic("时间戳解释失败:" + "'" + str + "'")
// }
