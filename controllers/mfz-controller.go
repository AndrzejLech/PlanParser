package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	. "planScrapper/scrapper"
	. "planScrapper/structs"
	. "planScrapper/utils"
	"strconv"
	"time"
)

func GetMFZ12All(context *gin.Context) {
	var subjects, nameOfDay = ScrapMFZ12()

	days := CreateDays(subjects, nameOfDay)
	days = RemovePastDays(days)
	days = days[:len(days)-1]

	context.JSON(200, days)
}

func GetMFZ12PerWeek(context *gin.Context) {
	var subjects, nameOfDay = ScrapMFZ12()
	var days = CreateDays(subjects, nameOfDay)
	var weeks []Week
	var previousWeekday = 0
	var previousDate time.Time
	var week = CreateEmptyWeek()

	for index, element := range days {
		newWeekday, date := GetWeekDay(element.Name)
		newWeekday = GetWeekdaysStraight(newWeekday)

		monday := previousDate.AddDate(0, 0, -previousWeekday)

		if date.AddDate(0, 0, -newWeekday).Day() == previousDate.AddDate(0, 0, -previousWeekday).Day() || index == 0 {
			week.Days[newWeekday] = days[index]
		} else {

			for numberOfElement, element := range week.Days {
				dateOfDay := monday.AddDate(0,0, numberOfElement)
				week.Days[numberOfElement].Name = ChangeWeekdayNumberToName(numberOfElement) + " " + strconv.Itoa(dateOfDay.Day()) + "." + strconv.Itoa(int(dateOfDay.Month())) + "." + strconv.Itoa(dateOfDay.Year())
				fmt.Print(element.Name + "\n")
			}

			weeks = append(weeks, week)

			week = CreateEmptyWeek()
			week.Days[newWeekday] = days[index]
		}

		previousDate = date
		previousWeekday = newWeekday
	}

	context.JSON(200, weeks)
}
