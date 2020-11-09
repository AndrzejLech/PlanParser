package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

type Subject struct {
	Hour     string `json:"hour"`
	Name     string `json:"name"`
	Lecturer string `json:"lecturer"`
	Class    string `json:"class"`
}

type Day struct {
	Name     string    `json:"name"`
	Subjects []Subject `json:"subjects"`
}

func GetPlan(context *gin.Context) {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	counterToFour := 1
	counterToTwo := 2
	var subjects []Subject
	var text string

	var names []string
	var lecturers []string
	var hours []string
	var classes []string

	var days []Day
	var nameOfday []string

	c.OnHTML("tbody", func(elementTBody *colly.HTMLElement) {
		elementTBody.ForEach("tr", func(_ int, elementTr *colly.HTMLElement) {

			elementTr.ForEach("td.test", func(_ int, elementTdTest *colly.HTMLElement) {
				text = elementTdTest.Text
				if text == "" {
					return
				}
				if counterToFour > 3 {
					if counterToFour == 4 {
						counterToFour = 0
					}
				} else if counterToFour == 1 {
					names = append(names, text)
				} else if counterToFour == 2 {
					lecturers = append(lecturers, text)
				}
				counterToFour++
			})

			elementTr.ForEach("td.test2", func(_ int, elementTdTest2 *colly.HTMLElement) {
				text = elementTdTest2.Text
				if text == "" {
					return
				}
				if counterToTwo == 1 {
					classes = append(classes, text)
				} else if counterToTwo == 2 {
					counterToTwo = 0
				}
				counterToTwo++
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
				nameOfday = append(nameOfday, text)
			})
		})
	})

	c.Visit("http://www.plan.pwsz.legnica.edu.pl/checkSpecjalnosc.php?specjalnosc=s3PAM")

	var subject Subject

	for index, name := range names {
		subject.Name = name
		subject.Hour = hours[index]
		subject.Lecturer = lecturers[index]
		subject.Class = classes[index]
		subjects = append(subjects, subject)
	}

	var indexOfDays = 0
	var day Day
	var index = 0

	for _, subject := range subjects {
		if (index % 7) == 0 {
			day.Name = nameOfday[indexOfDays]
			days = append(days, day)
			indexOfDays++
			day.Subjects = nilm
		}
		day.Subjects = append(day.Subjects, subject)
		index++
	}

	context.JSON(200, gin.H{
		"days": days,
	})
}

func main() {
	router := gin.Default()

	router.GET("/index", GetPlan)
	router.Run()
}
