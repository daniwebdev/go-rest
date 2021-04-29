package models

type Article struct {
	Title string `json:title`
	Desc string `json:description`
	Content string `json:content`
	Id string `json:id`
}