package models

type Article struct {
	Title string `json:title`
	Desc string `json:desc`
	Content string `json:content`
	Id string `json:id`
}

type Response struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Data []Article
}