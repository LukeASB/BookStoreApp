package model

type ResponseHealthCheck struct {
	Endpoint    string `json:"endpoint"`
	Environment string `json:"dev"`
	Version     string `json:"version"`
}
