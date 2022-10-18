package structs
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

type Week struct {
	Day []Day `json:"day"`
}