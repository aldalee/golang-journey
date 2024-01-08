package main

import "testing"

func TestGetDaysBetween2Date(t *testing.T) {
	days, err := GetDaysBetween2Date("2023-11-19", "2024-01-08")
	if err != nil {
		t.Errorf("Error occurred: %v", err)
	}
	if days != 50 {
		t.Errorf("Expected 50 days, but got %d", days)
	}
}

func TestGetDateAfterDays(t *testing.T) {
	newDate, err := GetDateAfterDays("2023-11-19", 50)
	if err != nil {
		t.Errorf("Error occurred: %v", err)
	}
	if newDate != "2024-01-08" {
		t.Errorf("Expected date 2024-01-08, but got %s", newDate)
	}
}
