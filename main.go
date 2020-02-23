package main

import (
	//Go comes with some excellent support for encoding and decoding these formats using the standard library package, encoding/json.

	"log"
	"net/http"

	// article "github.com/VickyFengYu/go-restful/article"

	article "./article"

	"github.com/gorilla/mux"
)

// Article - defined in article package
type Article = article.Article

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", article.IndexPage)
	myRouter.HandleFunc("/articles", article.ReturnAllArticles)
	myRouter.HandleFunc("/article", article.CreateNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", article.DeleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", article.ReturnSingleArticle)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	article.Articles = []Article{
		Article{ID: "1", Title: "Article Title 1", Desc: "Article Description 1", Content: "Article Content 1"},
		Article{ID: "2", Title: "Article Title 2", Desc: "Article Description 2", Content: "Article Content 2"},
	}
	handleRequests()
}
