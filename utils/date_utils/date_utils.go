package date_utils

import "time"

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
)

// extracting this as a function will allow us to perform operations on it going forward
// rather than just getting the time as a string at all times
func GetNow() time.Time {
	// returns the current time for the server running the app
	// use standard time zone, so it doesn't matter where your app instance is being ran (Virginia, UK, etc.)
	// this basically ignores time zones in this way
	// that way if you needed to show them the time, you'd just need to add/subtract time based on their calculated time zone
	return time.Now().UTC()
}

func GetNowString() string {
	// these values are all placeholders, per-the documentation form .format() (read it)
	return GetNow().Format(apiDateLayout)
}
