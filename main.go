package main

import (
	//Go comes with some excellent support for encoding and decoding these formats using the standard library package, encoding/json.
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	article "github.com/VickyFengYu/go-restful/article"

	"github.com/gorilla/mux"
)

//Article - defined in article package
type Article = article.Article

// Articles - a global Articles array
var Articles []Article

func indexPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the IndexPage!") //Fprintf formats according to a format specifier
	fmt.Println("Endpoint Hit: indexPage")      //Println formats using the default formats
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles) // NewEncoder returns a new encoder that writes to w.
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // Vars returns the route variables for the current request, if any.
	fmt.Println("route variables", vars)
	key := vars["id"]

	for _, article := range Articles {
		if article.ID == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a new Article struct
	// append this to our Articles array.
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)
	// update our global Articles array to include
	// our new Article
	Articles = append(Articles, article)

	json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, article := range Articles {
		if article.ID == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}

}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", indexPage)
	myRouter.HandleFunc("/articles", returnAllArticles)
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	Articles = []Article{
		Article{ID: "1", Title: "Article Title 1", Desc: "Article Description 1", Content: "Article Content 1"},
		Article{ID: "2", Title: "Article Title 2", Desc: "Article Description 2", Content: "Article Content 2"},
	}
	handleRequests()
}
