package controllers

import (
	"github.com/gin-gonic/gin"
	. "planScrapper/scrapper"
	. "planScrapper/structs"
	. "planScrapper/utils"
)

func GetMFZ12(context *gin.Context) {
	var subjects, nameOfDay = ScrapMFZ12()

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

	days = RemovePastDays(days)

	days = days[:len(days)-1]

	context.JSON(200, days)
}