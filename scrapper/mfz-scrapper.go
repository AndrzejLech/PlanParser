package scrapper

import (
	"fmt"
	"github.com/gocolly/colly"
	. "planScrapper/structs"
)

func ScrapMFZ12() ([]Subject, []string) {
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
				if text == "" { return }

				switch counterToFour {
				case 3:
					names = append(names, text)
				case 4:
					lecturers = append(lecturers, text)
					counterToFour = 0
				}
				counterToFour++
			})

			elementTr.ForEach("td.test2", func(_ int, elementTdTest2 *colly.HTMLElement) {
				text = elementTdTest2.Text
				if text == "" { return }

				if counterToTwo == 2 {
					classes = append(classes, text)
					counterToTwo = 0
				}
				counterToTwo++
			})

			elementTr.ForEach("td.godzina", func(_ int, elementTdGodzina *colly.HTMLElement) {
				text = elementTdGodzina.Text
				if text == "" { return }
				hours = append(hours, text)
			})
			elementTr.ForEach("td.nazwaDnia", func(i int, elementTdNazwaDnia *colly.HTMLElement) {
				text = elementTdNazwaDnia.Text
				if text == "" { return }
				nameOfday = append(nameOfday, text)
			})
		})
	})

	c.Visit("http://www.plan.pwsz.legnica.edu.pl/checkSpecjalnoscStac.php?specjalnosc=s1MFZ")

	var subject Subject

	for index, name := range names {
		subject.Name = name
		subject.Hour = hours[index]
		subject.Lecturer = lecturers[index]
		subject.Room = classes[index]
		subjects = append(subjects, subject)
	}

	return subjects, nameOfday
}
