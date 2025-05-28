package utils

import (
	"fmt"
	"time"
)

// Format time WIB (Waktu Indonesia Barat)
var WIB = time.FixedZone("WIB", 7*3600)

/* ==== Formating time to string ==== */
// Format time to string WIB
func FormatWIB(t time.Time) string {
	return t.In(WIB).Format("02 Januari 2006 15:04")
}

// Format time only
func FormatTimeINA(t time.Time) string {
	t = t.In(WIB)
	return fmt.Sprintf("%02d:%02d", t.Hour(), t.Minute())
}

// Format time to string with the format (dd-mm-yyyy / 28-05-2025)
func FormatDateINA(t time.Time) string {
	return t.In(WIB).Format("02-01-2006")
}

// Format datetime to string with the format (dd-mm-yyyy HH:mm / 28-05-2025 14:30)
func FormatDateTimeINA(t time.Time) string {
	return t.In(WIB).Format("02-01-2006 15:04")
}

/* ==== Formating string to time ==== */
// Parse string (dd-mm-yyyy / 28-05-2025) to time.Time
func ParseDateNumeric(dateStr string) (time.Time, error) {
	layout := "02-01-2006"
	return time.ParseInLocation(layout, dateStr, WIB)
}

// Parse string (dd-mm-yyyy HH:mm / 28-05-2025 14:30) to time.Time
func ParseDateTimeNumeric(dateStr string) (time.Time, error) {
	layout := "02-01-2006 15:04"
	return time.ParseInLocation(layout, dateStr, WIB)
}

// Get current time WIB
func NowWIB() time.Time {
	return time.Now().In(WIB)
}
