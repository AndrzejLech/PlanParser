package main

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
	Name     string    `json:"name"`
	Subject1 []Subject `json:"subjects1"`
	Subject2 []Subject `json:"subjects2"`
}
