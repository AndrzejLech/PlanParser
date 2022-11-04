package utils

import (
	. "planScrapper/structs"
	"strings"
)

func CreateDays(subjects []Subject, nameOfDay []string) []Day {
	var days []Day
	var indexOfDays = 0
	var day Day

	for index, subject := range subjects {
		if index < 7 {
		} else if (index % 7) == 0 {
			day.Name = nameOfDay[indexOfDays]

			subject.Name = strings.TrimSpace(subject.Name)

			if day.IsBusy == true {
			} else if subject.Name == "-" {
				day.IsBusy = false
			} else {
				day.IsBusy = true
			}

			days = append(days, day)
			indexOfDays++
			day.Subjects = nil
			day.IsBusy = false
		}
		day.Subjects = append(day.Subjects, subject)
	}
	return days
}
