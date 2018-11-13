package dateTime

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

//获取昨天日期及开始结束时间
func Yesterday() (date string, beginTime int64, endTime int64) {
	date, beginTime, endTime = Date(-1)
	return
}

//获取日期及开始结束时间
func Date(days int) (date string, beginTime int64, endTime int64) {
	dateTime := time.Now().AddDate(0, 0, days)
	date = dateTime.Format("20060102")
	beginTime, endTime = DateBeginEndTime(date)
	return
}

//获取日期的开始和结束时间
func DateBeginEndTime(date string) (beginTime int64, endTime int64) {
	beginStr := fmt.Sprintf("%s000000", date)
	endStr := fmt.Sprintf("%s235959", date)
	beginTime, _ = strconv.ParseInt(beginStr, 10, 0)
	endTime, _ = strconv.ParseInt(endStr, 10, 0)
	return
}

func DateTimeToDateStr(dateTime int64) string {
	dateTimeStr := strconv.FormatInt(dateTime, 10)
	dateTimeStr = "2018-10-10 18:18:18"
	formatTime, err := time.Parse("20060102150405", dateTimeStr)
	if err != nil {
		return ""
	}
	return formatTime.Format("2006-01-02")
}

/**
 * 字符样式 64 位数字日期时间转为日期样式
 * 20060102000000 => 2006-01-02
 */
func DateTimeIntToDateStr(dateTime int64) string {
	dateTimeStr := strconv.FormatInt(dateTime, 10)
	dateTimeArr := strings.Split(dateTimeStr, "")

	return dateTimeArr[0] + dateTimeArr[1] + dateTimeArr[2] + dateTimeArr[3] + "-" + dateTimeArr[4] + dateTimeArr[5] + "-" + dateTimeArr[6] + dateTimeArr[7]
}

/**
 * 字符样式 64 位数字日期时间转为小时样式
 * 20060102000000 => 10:20
 */
func DateTimeToTimeStr(dateTime int64) string {
	dateTimeStr := strconv.FormatInt(dateTime, 10)
	dateTimeArr := strings.Split(dateTimeStr, "")

	return dateTimeArr[8] + dateTimeArr[9] + ":" + dateTimeArr[10] + dateTimeArr[11]
}

// 秒转换为时分秒
// 3600秒 => 10:20:30
func Sec2Time(secondsrc float32) string {
	var timeStr string = "00:00:00"
	var second float64 = float64(secondsrc)

	if second > 0 {
		var hour float64 = math.Floor(second / 3600)
		var minute float64 = math.Floor((second - 3600*hour) / 60)
		var seconds float64 = math.Floor(float64(int((second-3600*hour)-60*minute) % 60))
		var hourStr string = strconv.FormatFloat(float64(hour), 'f', 0, 64)
		if strings.Count(hourStr, "")-1 < 2 {
			hourStr = "0" + hourStr
		}
		var minuteStr string = strconv.FormatFloat(float64(minute), 'f', 0, 64)
		if strings.Count(minuteStr, "")-1 < 2 {
			minuteStr = "0" + minuteStr
		}
		var secondsStr string = strconv.FormatFloat(float64(seconds), 'f', 0, 64)
		if strings.Count(secondsStr, "")-1 < 2 {
			secondsStr = "0" + secondsStr
		}

		timeStr = hourStr + ":" + minuteStr + ":" + secondsStr
	}
	return timeStr
}
