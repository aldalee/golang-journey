package main

import "time"

// GetDaysBetween2Date get the number of days difference between two dates
func GetDaysBetween2Date(date1, date2 string) (int, error) {
	d1, err := time.Parse("2006-01-02", date1)
	if err != nil {
		return 0, err
	}
	d2, err := time.Parse("2006-01-02", date2)
	if err != nil {
		return 0, err
	}
	return int(d2.Sub(d1).Hours() / 24), nil
}

// GetDateAfterDays get the date in a few days
func GetDateAfterDays(startDate string, days int) (string, error) {
	start, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return "", err
	}
	newDate := start.AddDate(0, 0, days)
	return newDate.Format("2006-01-02"), nil
}
