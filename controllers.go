package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

func GetImage(context *gin.Context) {
	context.HTML(200, "image.html", gin.H{})
}

func GetNur(context *gin.Context) {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	counterToEight := 1
	counterToFour := 1
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
				if counterToEight == 1 {
					names = append(names, text)
				} else if counterToEight == 2 {
					lecturers = append(lecturers, text)
				} else if counterToEight == 8 {
					counterToEight = 0
				}
				counterToEight++
			})

			elementTr.ForEach("td.test2", func(_ int, elementTdTest2 *colly.HTMLElement) {
				text = elementTdTest2.Text
				if text == "" {
					return
				}
				if counterToFour == 1 {
					classes = append(classes, text)
				} else if counterToFour == 4 {
					counterToFour = 0
				}
				counterToFour++
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

	days = removePastDays(days)

	days = days[:len(days)-1]

	context.JSON(200, days)
}

func GetInfOneOne(context *gin.Context) {
	subjects, nameOfDay := getS4PAM11All()

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

	days = removePastDays(days)

	days = days[:len(days)-1]

	context.JSON(200, days)
}

func GetInfOneTwo(context *gin.Context) {
	subjects, nameOfDay := getS4PAM12All()

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

	days = removePastDays(days)

	days = days[:len(days)-1]

	context.JSON(200, days)
}

func GetSzymin(context *gin.Context) {

	var indexOfDays = 0
	var day DoubleDay
	var days []DoubleDay

	subjects1, nameOfDay := getS3GK()

	subjects2 := getS2ZIIN1()

	fmt.Println(len(subjects1), len(subjects2))

	for index, subject1 := range subjects1 {
		if index < 7 {

		} else if (index % 7) == 0 {
			day.Name = nameOfDay[indexOfDays]
			days = append(days, day)
			indexOfDays++
			day.Subject1 = nil
			day.Subject2 = nil
		}
		day.Subject1 = append(day.Subject1, subject1)
		day.Subject2 = append(day.Subject2, subjects2[index])
	}

	days = days[:len(days)-1]

	context.JSON(200, days)
}

func GetMamlina(context *gin.Context) {
	var indexOfDays = 0
	var day DoubleDay
	var days []DoubleDay

	subjects1, nameOfDay := getS4PAM11()

	subjects2, _ := getS1ZIP11()

	fmt.Println(len(subjects1), len(subjects2))

	for index, subject1 := range subjects1 {
		if index < 7 {

		} else if (index % 7) == 0 {
			day.Name = nameOfDay[indexOfDays]
			days = append(days, day)
			indexOfDays++
			day.Subject1 = nil
			day.Subject2 = nil
		}
		day.Subject1 = append(day.Subject1, subject1)
		day.Subject2 = append(day.Subject2, subjects2[index])
	}

	days = days[:len(days)-1]

	context.JSON(200, days)
}
