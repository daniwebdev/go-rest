package controllers

import (
	"encoding/json"
	"fmt"

	// "io/ioutil"
	"net/http"
)


func CreateArticle(w http.ResponseWriter, r *http.Request) {
	
	// _body, _ := ioutil.ReadAll(r.Body)
	_body := json.NewDecoder(r.Body)

	fmt.Println(_body)

	fmt.Fprintf(w, "%+v", _body["data"])
}