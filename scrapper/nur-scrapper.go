package scrapper

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
	. "planScrapper/structs"
	. "planScrapper/utils"
)


func GetNur(context *gin.Context) {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	counterToSix := 1
	counterToFive := 1

	var subjects []Subject
	var text string

	var names []string
	var lecturers []string
	var hours []string
	var classes []string

	var days []Day
	var nameOfDay []string

	c.OnHTML("tbody", func(elementTBody *colly.HTMLElement) {
		elementTBody.ForEach("tr", func(_ int, elementTr *colly.HTMLElement) {

			elementTr.ForEach("td.test", func(_ int, elementTdTest *colly.HTMLElement) {
				text = elementTdTest.Text
				if text == "" {
					return
				}
				if counterToSix == 1 {
					names = append(names, text)
				} else if counterToSix == 2 {
					lecturers = append(lecturers, text)
				} else if counterToSix == 6 {
					counterToSix = 0
				}
				counterToSix++
			})

			elementTr.ForEach("td.test2", func(_ int, elementTdTest2 *colly.HTMLElement) {
				text = elementTdTest2.Text
				if text == "" {
					return
				}
				if counterToFive == 1 {
					classes = append(classes, text)
				} else if counterToFive == 5 {
					counterToFive = 0
				}
				counterToFive++
			})

			elementTr.ForEach("td.godzina", func(_ int, elementTdGodzina *colly.HTMLElement) {
				text = elementTdGodzina.Text
				if text == "" {
					return
				}
				hours = append(hours, text)
			})

			elementTr.ForEach("td.nazwaDnia", func(i int, elementTdNazwaDnia *colly.HTMLElement) {
				text = elementTdNazwaDnia.Text
				if text == "" {
					return
				}
				nameOfDay = append(nameOfDay, text)
			})
		})
	})

	err := c.Visit("http://www.plan.pwsz.legnica.edu.pl/checkSpecjalnoscStac.php?specjalnosc=s1MP")

	if err != nil {
		return
	}

	var subject Subject

	for index, name := range names {
		subject.Name = name
		subject.Hour = hours[index]
		subject.Lecturer = lecturers[index]
		subject.Room = classes[index]
		subjects = append(subjects, subject)
	}

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
