package util

import (
	"acttos.com/avatar/model/constant"
	"acttos.com/avatar/pkg/util/logger"
	"math"
	"time"
)

var TimeHelper *TimeUtil

func InitTimeHelper() {
	once.Do(func() {
		TimeHelper = &TimeUtil{}
	})
}

type TimeUtil struct{}

// returns the format with layout 'yyyyMMdd'
func (fs *TimeUtil) CompactDayFormatOfTime(aTime time.Time) string {
	return aTime.Format(constant.COMPACT_DAY_FORMAT_LAYOUT)
}

// returns the format with layout 'yyyy-MM-dd'
func (fs *TimeUtil) NormalDayFormatOfTime(aTime time.Time) string {
	return aTime.Format(constant.DEFAULT_DAY_FORMAT_LAYOUT)
}

// returns the format with layout 'yyyy-MM-dd HH:mm:ss'
func (fs *TimeUtil) NormalFormatOfTime(aTime time.Time) string {
	return aTime.Format(constant.DEFAULT_TIME_FORMAT_LAYOUT)
}

// 计算球桌游戏进行时长,向上取整,当前时间减去开始时间的[分钟数]
func (fs *TimeUtil) TimeCostMinutes(startAt time.Time, endAt time.Time) int32 {
	totalMinutes := int32(math.Ceil(endAt.Sub(startAt).Minutes()))
	return totalMinutes
}

func (fs *TimeUtil) DaysBetween(smallOne, bigOne, formatLayout string) (int64, error) {
	smallOneDate, err := time.ParseInLocation(formatLayout, smallOne, time.Local)
	if err != nil {
		logger.Monitor.Errorf("Bad format or time string. error:%+v", err)
		return -1, err
	}
	bigOneDate, err := time.ParseInLocation(formatLayout, bigOne, time.Local)
	if err != nil {
		logger.Monitor.Errorf("Bad format or time string. error:%+v", err)
		return -1, err
	}

	//return int64(bigOneDate.Sub(smallOneDate).Seconds() / 86400), nil
	return (bigOneDate.Unix() - smallOneDate.Unix()) / 86400, nil
}

func (fs *TimeUtil) IsTheSameDay(aTime, anotherTime time.Time) bool {
	aTimeDay := aTime.Format(constant.DEFAULT_DAY_FORMAT_LAYOUT)
	anotherTimeDay := anotherTime.Format(constant.DEFAULT_DAY_FORMAT_LAYOUT)
	return aTimeDay == anotherTimeDay
}

func (fs *TimeUtil) BeginningOfDayTime(aTime time.Time) time.Time {
	aTimeDay := aTime.Format(constant.DEFAULT_DAY_FORMAT_LAYOUT)
	result, _ := time.ParseInLocation("2006-01-02 15:04:05", aTimeDay+" 00:00:00", time.Local)
	return result
}

func (fs *TimeUtil) EndingOfDayTime(aTime time.Time) time.Time {
	aTimeDay := aTime.Format(constant.DEFAULT_DAY_FORMAT_LAYOUT)
	result, _ := time.ParseInLocation("2006-01-02 15:04:05", aTimeDay+" 23:59:59", time.Local)
	return result
}
