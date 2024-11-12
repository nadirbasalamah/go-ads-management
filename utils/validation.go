package utils

import (
	"fmt"
	"time"
)

func ValidateFile(ext string) bool {
	return ALLOWED_EXTENSIONS[ext]
}

func ValidateDate(startDateStr, endDateStr string) error {
	const dateFormat = "02-01-2006"

	startDate, err := time.Parse(dateFormat, startDateStr)
	if err != nil {
		return fmt.Errorf("invalid start_date: %v", err)
	}

	endDate, err := time.Parse(dateFormat, endDateStr)
	if err != nil {
		return fmt.Errorf("invalid end_date: %v", err)
	}

	if endDate.Before(startDate) {
		return fmt.Errorf("end_date must be on or after start_date")
	}

	return nil
}
