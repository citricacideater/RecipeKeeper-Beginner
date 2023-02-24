package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type Recipe struct {
	Id          string            `json:"Id"`
	RecipeName  string            `json:"Recipe Name"`
	Source      string            `json:"Source"`
	PrepTime    string            `json:"Preperation Time"`
	CookTime    string            `json:"Cook Time"`
	ServingSize int               `json:"Serving Size"`
	Ingredients map[string]string `json:"Ingredients"`
	Directions  map[int]string    `json:"Directions"`
	Tags        []string          `json:"Tags"`
}

var data []Recipe

var templates = template.Must(template.ParseGlob("assets/tmpl/*.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, data any) {
	err := templates.ExecuteTemplate(w, tmpl+".html",data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func fetchAllRecipes() {
	info, err := http.Get("https://api.npoint.io/fff0f131782057b16a12")
	if err != nil {
		log.Fatal(err, "ERROR 500: Failed to find url")
	}
	defer info.Body.Close()
	body, err := ioutil.ReadAll(info.Body)
	if err != nil {
		log.Fatal(err, "Failed to read json")
	}

	json.Unmarshal(body, &data)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w,"index",data)
}

func recipePage(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/recipe/"):]
	var single Recipe

	for _, v := range data {
		if v.Id == id {
			single = v
			break
		}
	}
	renderTemplate(w,"recipe",single)

}

func handleRequest() {
	fetchAllRecipes()
	server := http.NewServeMux()
	style := http.FileServer(http.Dir("static/css"))

	server.Handle("/static/css/", http.StripPrefix("/static/css/", style))

	server.HandleFunc("/", homePage)
	server.HandleFunc("/recipe/", recipePage)

	log.Fatal(http.ListenAndServe(":8000", server))
}

func main() {
	handleRequest()
}
