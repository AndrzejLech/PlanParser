package controllers

import (
	"github.com/gin-gonic/gin"
	. "planScrapper/scrapper"
	. "planScrapper/structs"
	. "planScrapper/utils"
	"strings"
	"time"
)

func GetMFZ12All(context *gin.Context) {
	var subjects, nameOfDay = ScrapMFZ12()

	days := createDays(subjects, nameOfDay)
	days = RemovePastDays(days)
	days = days[:len(days)-1]

	context.JSON(200, days)
}

func GetMFZ12PerWeek(context *gin.Context) {
	var subjects, nameOfDay = ScrapMFZ12()
	var days = createDays(subjects, nameOfDay)
	var weeksArray []Week
	var previousWeekday = 0
	var week = createEmptyWeek()

	for index, element := range days {
		newWeekday := getWeekDay(element.Name)

		if previousWeekday < newWeekday {
			if newWeekday == 0 {
				newWeekday = 6
			} else {
				week.Day[newWeekday-1] = days[index]
			}
		} else {
			weeksArray = append(weeksArray, week)
			week = createEmptyWeek()
		}

		previousWeekday = newWeekday
	}

	context.JSON(200, weeksArray)
}

func getWeekDay(unformattedDay string) int {
	var year, month, day = splitUnformated(unformattedDay)
	location, _ := time.LoadLocation("Europe/Warsaw")

	date := time.Date(
		year,
		time.Month(month),
		day,
		0,
		0,
		0,
		0,
		location,
	)
	return int(date.Weekday())
}

func splitUnformated(unformattedDay string) (int, int, int) {
	splited := strings.Split(unformattedDay, " ")
	splited = strings.Split(splited[1], "-")

	return int(StringToInt64(splited[0])), int(StringToInt64(splited[1])), int(StringToInt64(splited[2]))
}

func createDays(subjects []Subject, nameOfDay []string) []Day {
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
		day.Subjects = append(day.Subjects, subject)
	}
	return days
}

func createEmptyWeek() Week {
	var emptySubject = Subject{
		Hour:     "-",
		Name:     "-",
		Lecturer: "-",
		Room:     "-",
	}
	var emptySubjects []Subject

	n := 0
	m := 0
	for m <= 7 {
		emptySubjects = append(emptySubjects, emptySubject)
		m++
	}

	var week = Week{}

	for n <= 6 {
		week.Day = append(week.Day, Day{
			Name:     "-",
			Subjects: emptySubjects,
		})
		n++
	}

	return week
}
