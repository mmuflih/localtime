package localtime

/**
 * Created by M. Muflih Kholidin
 * https://github.com/mmuflih
 * mmuflic@gmail.com
 */

import (
	"strings"
	"time"
)

func DateTimeFormat(t time.Time) string {
	i := strings.Split(t.Format("02-January-2006-15:04:05"), "-")
	months := map[string]string{
		"January":   "Januari",
		"February":  "Februari",
		"March":     "Maret",
		"April":     "April",
		"May":       "Mei",
		"June":      "Juni",
		"July":      "Juli",
		"August":    "Agustus",
		"September": "September",
		"October":   "Oktober",
		"November":  "November",
		"December":  "Desember",
	}

	return i[0] + " " + months[i[1]] + " " + i[2] + " " + i[3]
}

func DateTimeFormatLong(t time.Time) string {
	i := strings.Split(t.Format("Monday-02-January-2006-15:04:05"), "-")
	months := map[string]string{
		"January":   "Januari",
		"February":  "Februari",
		"March":     "Maret",
		"April":     "April",
		"May":       "Mei",
		"June":      "Juni",
		"July":      "Juli",
		"August":    "Agustus",
		"September": "September",
		"October":   "Oktober",
		"November":  "November",
		"December":  "Desember",
	}

	days := map[string]string{
		"Wednesday": "Rabu",
		"Thursday":  "Kamis",
		"Friday":    "Jumat",
		"Saturday":  "Sabtu",
		"Sunday":    "Minggu",
		"Monday":    "Senin",
		"Tuesday":   "Selasa",
	}

	return days[i[0]] + ", " + i[1] + " " + months[i[2]] + " " + i[3] + " " + i[4]
}

func DateFormat(t time.Time) string {
	i := strings.Split(t.Format("02-January-2006"), "-")
	months := map[string]string{
		"January":   "Januari",
		"February":  "Februari",
		"March":     "Maret",
		"April":     "April",
		"May":       "Mei",
		"June":      "Juni",
		"July":      "Juli",
		"August":    "Agustus",
		"September": "September",
		"October":   "Oktober",
		"November":  "November",
		"December":  "Desember",
	}

	return i[0] + " " + months[i[1]] + " " + i[2]
}

func DateFormatLong(t time.Time) string {
	i := strings.Split(t.Format("Monday-02-January-2006"), "-")
	months := map[string]string{
		"January":   "Januari",
		"February":  "Februari",
		"March":     "Maret",
		"April":     "April",
		"May":       "Mei",
		"June":      "Juni",
		"July":      "Juli",
		"August":    "Agustus",
		"September": "September",
		"October":   "Oktober",
		"November":  "November",
		"December":  "Desember",
	}

	days := map[string]string{
		"Wednesday": "Rabu",
		"Thursday":  "Kamis",
		"Friday":    "Jumat",
		"Saturday":  "Sabtu",
		"Sunday":    "Minggu",
		"Monday":    "Senin",
		"Tuesday":   "Selasa",
	}

	return days[i[0]] + ", " + i[1] + " " + months[i[2]] + " " + i[3]
}
