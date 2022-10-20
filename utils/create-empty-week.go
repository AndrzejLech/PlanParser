package utils

import . "planScrapper/structs"

func CreateEmptyWeek() Week {
	var emptySubject = Subject{
		Hour:     "-",
		Name:     "-",
		Lecturer: "-",
		Room:     "-",
		}
		var emptySubjects []Subject

	m := 0
	for m <= 6 {
		emptySubjects = append(emptySubjects, emptySubject)
		m++
	}

	var week = Week{}

	n := 0
	for n <= 6 {
		week.Days = append(week.Days, Day{
			Name:     "-",
			Subjects: emptySubjects,
			})
		n++
	}

	return week
}
