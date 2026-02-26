package models

type Response struct {
	Status  string `json:"status"`
	Time    string `json:"time"`
	Version string `json:"version"`
	Code    int    `json:"code"`
	Content string `json:"content"`
}
