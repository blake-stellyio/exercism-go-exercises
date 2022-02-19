package booking

import "time"

const (
    longForm = "1/2/2006 15:04:05"
    longForm2 = "January 2, 2006 15:04:05"
    longForm3 = "Monday, January 2, 2006 15:04:05"
    longForm4 = "2006-01-02 15:04:05"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	t, _ := time.Parse(longForm, date)
    return t
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
    otherTime, _ := time.Parse(longForm2, date)
    return time.Now().After(otherTime)   
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	t, _ := time.Parse(longForm3, date)
    h := t.Hour()
    if h >= 12 && h < 18 {
        return true
    } else {
    	return false
    }
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	t := Schedule(date)
    return t.Format("You have an appointment on Monday, January 2, 2006, at 15:04.")
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	a, _ := time.Parse(longForm4, "2022-09-15 00:00:00")
	return a
}
