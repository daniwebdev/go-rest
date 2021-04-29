package controllers

import (
	"encoding/json"
	"fmt"
	"go-rest/helpers"
	"go-rest/models"
	"log"

	"net/http"
)

var _article models.Article
var _response models.Response
var _articleList []models.Article

func CreateArticle(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&_article)

	if err != nil {
		panic(err)
	}

	insert, err := helpers.DB().Query("INSERT INTO articles (title, description, content) VALUES ('" + _article.Title + "', '" + _article.Desc + "', '" + _article.Content + "')")

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

	fmt.Println(_article)
	_response.Status = 1
	_response.Message = "Berhasil Disimpan"

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(_response)
}

func GetArticle(w http.ResponseWriter, r *http.Request) {
	get, err := helpers.DB().Query("SELECT * FROM articles")
	if err != nil {
		panic(err.Error())
	}

	for get.Next() {
		if err := get.Scan(&_article.Id, &_article.Title, &_article.Desc, &_article.Content); err != nil {
			log.Fatal(err.Error())
		} else {
			_articleList = append(_articleList, _article)
		}
	}

	_response.Status = 1
	_response.Message = "Berhasil Disimpan"
	_response.Data = _articleList

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(_response)
}
