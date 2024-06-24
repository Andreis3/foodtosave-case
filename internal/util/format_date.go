package util

import (
	"time"
)

const layout = "2006-01-02T15:04:05.000Z"
const location = "America/Sao_Paulo"

func FormatDateTime() time.Time {
	utcTime := time.Now()
	locationTimeZone, _ := time.LoadLocation(location)
	locationTime := utcTime.In(locationTimeZone)
	formattedDate := locationTime.Format(layout)
	parsedDate, _ := time.Parse(layout, formattedDate)
	return parsedDate
}
