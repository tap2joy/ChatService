package utils

import (
	"fmt"
	"math"
)

// 格式化在线时长
// @param duration  时长，单位秒
func FormatOnlineTime(duration uint32) string {
	var days uint32
	var hours uint32
	var mins uint32
	var secs uint32

	if duration > 86400 {
		days = uint32(math.Floor(float64(duration) / 86400))
	}

	if duration > 3600 {
		hours = uint32(math.Floor(float64(duration-days*86400) / 3600))
	}

	if duration > 60 {
		mins = uint32(math.Floor(float64(duration-days*86400-hours*3600) / 60))
	}

	secs = duration - days*86400 - hours*3600 - mins*60
	retStr := fmt.Sprintf("%02dd %02dh %02dm %02ds", days, hours, mins, secs)
	return retStr
}
