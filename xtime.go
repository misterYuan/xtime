package xtime

import (
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

// //获取往后推几个月的月份
// type YearMonthDays struct {
// 	Year  int
// 	Month int
// 	Days  int
// }

// func GetMonthNo9BeforeNo(no int) *YearMonthDays {
// 	ym := new(YearMonthDays)
// 	t := time.Now().AddDate(0, -no, 0)
// 	ym.Month, _ = strconv.Atoi(t.Format("1"))
// 	ym.Year, _ = strconv.Atoi(t.Format("2006"))
// 	ym.setDays()
// 	return ym
// }

// // 获取指定月份的天数,如果为当月获取当月到当天的天数
// func (ymd *YearMonthDays) setDays() {
// 	if ok, days := ymd.isCurrentMonth(); ok {
// 		ymd.Days = days
// 		return
// 	}
// 	if ymd.Month != 2 {
// 		if ymd.Month == 4 || ymd.Month == 6 || ymd.Month == 9 || ymd.Month == 11 {
// 			ymd.Days = 30
// 		} else {
// 			ymd.Days = 31
// 		}
// 	} else {
// 		if ((ymd.Year%4) == 0 && (ymd.Year%100) != 0) || (ymd.Year%400) == 0 {
// 			ymd.Days = 29
// 		} else {
// 			ymd.Days = 28
// 		}
// 	}
// }

// // 判断是否为当月,是当月返回当月已过的天数
// func (y *YearMonthDays) isCurrentMonth() (bool, int) {
// 	now := time.Now()
// 	year := now.Year()
// 	month, _ := strconv.Atoi(now.Format("1"))
// 	if year == y.Year && month == y.Month {
// 		return true, now.Day()
// 	}
// 	return false, 0
// }

// // 获取年月开始的一天的0点时间戳
// func (y *YearMonthDays) GetYearMonthStartDayZeroStamp() int64 {
// 	ym := strconv.Itoa(y.Year) + strconv.Itoa(y.Month)
// 	t, err := time.Parse("20061", ym)
// 	if err != nil {
// 		log.Println(err)
// 		panic(err.Error())
// 	}
// 	return t.Unix()
// }

// // 获取年月最后一天的结束点时间戳
// func (y *YearMonthDays) GetYearMonthEndDayEndStamp() int64 {
// 	ymd := strconv.Itoa(y.Year) + "-" + strconv.Itoa(y.Month) + "-" + strconv.Itoa(y.Days)
// 	t, err := time.Parse("2006-1-2", ymd)
// 	if err != nil {
// 		log.Println(err)
// 		panic(err.Error())
// 	}
// 	t = t.AddDate(0, 0, 1).Add(-1)
// 	return t.Unix()
// }

// // 根据年的月份获取起止时间戳
// func (y *YearMonthDays) GetMonthStartEndStamp() []int64 {
// 	return []int64{y.GetYearMonthStartDayZeroStamp(), y.GetYearMonthEndDayEndStamp()}
// }
