package model

type Model struct {
}

type ResponseHealthCheck struct {
	Endpoint    string `json:"endpoint"`
	Environment string `json:"dev"`
	Version     string `json:"version"`
}

type Input struct {
	Title     string   `json:"title"`
	Published int      `json:"published"`
	Pages     int      `json:"pages"`
	Genres    []string `json:"genres"`
	Rating    float64  `json:"rating"`
}
