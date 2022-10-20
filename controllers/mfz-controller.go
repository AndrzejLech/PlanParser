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

		fmt.Print(
			"date: " + strconv.Itoa(previousDate.Day()) +
				" weekday: " + ChangeWeekdayNumberToName(previousWeekday) +
				" start week: " + strconv.Itoa(previousDate.AddDate(0, 0, -(previousWeekday)).Day()) + "\n",
		)
		fmt.Print(
			"date: " + strconv.Itoa(date.Day()) +
				" weekday: " + ChangeWeekdayNumberToName(newWeekday) +
				" start week: " + strconv.Itoa(date.AddDate(0, 0, -(newWeekday)).Day()) + "\n",
		)

		if date.AddDate(0, 0, -newWeekday).Day() == previousDate.AddDate(0, 0, -previousWeekday).Day() || index == 0 {
			week.Days[newWeekday] = days[index]
		} else {
			weeks = append(weeks, week)
			week = CreateEmptyWeek()
			week.Days[newWeekday] = days[index]
		}

		previousDate = date
		previousWeekday = newWeekday
	}

	context.JSON(200, weeks)
}
