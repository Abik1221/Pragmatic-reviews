package entity

type Actor struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}

type Video struct {
	Title       string `json:"titile`
	Description string `json:"description"`
	Url         string `json:"url"`
	Actor       Actor  `json:"actor"`
}
