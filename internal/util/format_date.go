package util

import (
	"time"
)

const layout = "2006-01-02T15:04:05.000Z"
const location = "America/Sao_Paulo"

func FormatDate() string {
	utcTime := time.Now()
	locationTimeZone, _ := time.LoadLocation(location)
	locationTime := utcTime.In(locationTimeZone)
	formattedDate := locationTime.Format(layout)
	return formattedDate
}

func FormatDateTime() time.Time {
	utcTime := time.Now()
	locationTimeZone, _ := time.LoadLocation(location)
	locationTime := utcTime.In(locationTimeZone)
	formattedDate := locationTime.Format(layout)
	parsedDate, _ := time.Parse(layout, formattedDate)
	return parsedDate
}

func FormatDateString(date time.Time) string {
	formattedDate := date.Format(layout)
	return formattedDate
}

func FormatDateStringToPointerTime(date string) *time.Time {
	parsedDate, _ := time.Parse(layout, date)
	return &parsedDate
}
