package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
	"time"
)

type Subject struct {
	Hour     string `json:"hour"`
	Name     string `json:"name"`
	Lecturer string `json:"lecturer"`
	Room     string `json:"room"`
}

type Day struct {
	Name     string    `json:"name"`
	Subjects []Subject `json:"subjects"`
}

type DoubleDay struct {
	Name     string    `json:"name:"`
	Subject1 []Subject `json:"subject1"`
	Subject2 []Subject `json:"subject2"`
}

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

	days = days[:len(days)-1]

	context.JSON(200, days)
}

func GetInfOneOne(context *gin.Context) {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	counterToFour := 1
	counterToTwo := 1
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
				if counterToFour == 1 {
					names = append(names, text)
				} else if counterToFour == 2 {
					lecturers = append(lecturers, text)
				} else if counterToFour == 4 {
					counterToFour = 0
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

	err := c.Visit("http://www.plan.pwsz.legnica.edu.pl/checkSpecjalnoscStac.php?specjalnosc=s4PAM")

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
			day.Name = nameOfday[indexOfDays]
			days = append(days, day)
			indexOfDays++
			day.Subjects = nil
		}
		day.Subjects = append(day.Subjects, subject)
	}

	days = days[:len(days)-1]

	context.JSON(200, days)
}

func GetInfOneTwo(context *gin.Context) {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	counterToFour := 1
	counterToTwo := 1
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
				if counterToFour == 3 {
					names = append(names, text)
				} else if counterToFour == 4 {
					lecturers = append(lecturers, text)
					counterToFour = 0
				}
				counterToFour++
			})

			elementTr.ForEach("td.test2", func(_ int, elementTdTest2 *colly.HTMLElement) {
				text = elementTdTest2.Text
				if text == "" {
					return
				}
				if counterToTwo == 2 {
					classes = append(classes, text)
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

	err := c.Visit("http://www.plan.pwsz.legnica.edu.pl/checkSpecjalnoscStac.php?specjalnosc=s4PAM")

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
			day.Name = nameOfday[indexOfDays]
			days = append(days, day)
			indexOfDays++
			day.Subjects = nil
		}
		day.Subjects = append(day.Subjects, subject)
	}

	days = days[:len(days)-1]

	context.JSON(200, days)
}

func GetSzymin(context *gin.Context) {

	var indexOfDays = 0
	var day DoubleDay
	var days []DoubleDay

	subjects1, nameOfDay := getS3GK()

	subjects2 := getS2ZIIN()

	for index, subject := range subjects1 {
		if (index % 7) == 0 {
			day.Name = nameOfDay[indexOfDays]
			days = append(days, day)
			indexOfDays++
			day.Subject1 = nil
			day.Subject2 = nil
		}
		day.Subject1 = append(day.Subject1, subject)
		day.Subject2 = append(day.Subject2, subjects2[index])
	}

	days = days[:len(days)-1]

	context.JSON(200, days)
}

func getS3GK() ([]Subject, []string) {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	counterToFour := 1
	counterToTwo := 1
	var subjects []Subject
	var text string

	var nameOfDay []string

	var names []string
	var lecturers []string
	var hours []string
	var classes []string

	c.OnHTML("tbody", func(elementTBody *colly.HTMLElement) {
		elementTBody.ForEach("tr", func(_ int, elementTr *colly.HTMLElement) {

			elementTr.ForEach("td.test", func(_ int, elementTdTest *colly.HTMLElement) {
				text = elementTdTest.Text
				if text == "" {
					return
				}
				if counterToFour == 1 {
					names = append(names, text)
				} else if counterToFour == 2 {
					lecturers = append(lecturers, text)
				} else if counterToFour == 4 {
					counterToFour = 0
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
				nameOfDay = append(nameOfDay, text)
			})
		})
	})

	c.Visit("http://www.plan.pwsz.legnica.edu.pl/checkSpecjalnosc.php?specjalnosc=s3GK")

	var subject Subject

	for index, name := range names {
		subject.Name = name
		subject.Hour = hours[index]
		subject.Lecturer = lecturers[index]
		subject.Room = classes[index]
		subjects = append(subjects, subject)
	}

	return subjects, nameOfDay
}

func getS2ZIIN() []Subject {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	counterToFour := 1
	counterToTwo := 1
	var subjects []Subject
	var text string

	var names []string
	var lecturers []string
	var hours []string
	var classes []string

	var nameOfday []string

	c.OnHTML("tbody", func(elementTBody *colly.HTMLElement) {
		elementTBody.ForEach("tr", func(_ int, elementTr *colly.HTMLElement) {

			elementTr.ForEach("td.test", func(_ int, elementTdTest *colly.HTMLElement) {
				text = elementTdTest.Text
				if text == "" {
					return
				}
				if counterToFour == 1 {
					names = append(names, text)
				} else if counterToFour == 2 {
					lecturers = append(lecturers, text)
				} else if counterToFour == 4 {
					counterToFour = 0
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

	c.Visit("http://www.plan.pwsz.legnica.edu.pl/checkSpecjalnosc.php?specjalnosc=s2ZIIN")

	var subject Subject

	for index, name := range names {
		subject.Name = name
		subject.Hour = hours[index]
		subject.Lecturer = lecturers[index]
		subject.Room = classes[index]
		subjects = append(subjects, subject)
	}

	return subjects
}

func main() {
	router := gin.New()
	router.LoadHTMLGlob("image.html")

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	router.GET("/", GetImage)
	router.GET("/inf-1-1", GetInfOneOne)
	router.GET("/inf-1-2", GetInfOneTwo)
	router.GET("/nur", GetNur)
	router.GET("/szymin", GetSzymin)
	err := router.Run()

	if err != nil {
		return
	}
}
