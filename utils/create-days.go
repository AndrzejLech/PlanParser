package utils

import (
	"fmt"
	. "planScrapper/structs"
)

func CreateDays(subjects []Subject, nameOfDay []string) []Day {
	var days []Day
	var indexOfDays = 0
	var day Day

	for index, subject := range subjects {

		if index < 7 {
		} else if (index % 7) == 0 {
			day.Name = nameOfDay[indexOfDays]
			days = append(days, day)
			indexOfDays++
			day.Subjects = nil
		}
		if subject.Name != "-  " {
			fmt.Println("\"" + subject.Name + "\"")
			day.IsBusy = true
		} else {
			day.IsBusy = false
		}
		day.Subjects = append(day.Subjects, subject)
	}
	return days
}
