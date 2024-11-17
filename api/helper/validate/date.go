package validate

import "time"

// This function checks date is valid
func IsDate(date string) bool {
	_, err := time.Parse("01/02/2006", date)
	return err == nil
}
