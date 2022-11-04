package utils

import (
	. "planScrapper/structs"
	"strings"
)

func CreateDays(subjects []Subject, nameOfDay []string) []Day {
	var days []Day
	var indexOfDays = 0
	var day = NewDay("", []Subject{})
	subjects = append([]Subject{{"", "", "", ""}}, subjects...)

	for index, subject := range subjects {
		subject.Name = strings.TrimSpace(subject.Name)

		if day.IsBusy {
		} else if subject.Name == "-" {
			day.IsBusy = false
		} else {
			day.IsBusy = true
		}

		if ((index+1) % 7) == 0 {
			day.Name = nameOfDay[indexOfDays]
			days = append(days, day)

			indexOfDays++
			day = NewDay("", []Subject{})
		}

		day.Subjects = append(day.Subjects, subject)
	}
	return days
}
