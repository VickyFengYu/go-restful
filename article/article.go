package article

import (
	//Go comes with some excellent support for encoding and decoding these formats using the standard library package, encoding/json.
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// Article - struct for all articles
type Article struct {
	ID      string `json:"ID"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// Articles - a global Articles array
var Articles []Article

// IndexPage - function homePage
func IndexPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the IndexPage!") //Fprintf formats according to a format specifier
	fmt.Println("Endpoint Hit: indexPage")      //Println formats using the default formats
}

//ReturnAllArticles -  function Return All Articles
func ReturnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles) // NewEncoder returns a new encoder that writes to w.
}

// ReturnSingleArticle - function Return Single Article by id
func ReturnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // Vars returns the route variables for the current request, if any.
	fmt.Println("Route variables", vars)
	key := vars["id"]

	for _, article := range Articles {
		if article.ID == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

// CreateNewArticle - function Create New Article
func CreateNewArticle(w http.ResponseWriter, r *http.Request) {
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

// DeleteArticle - function Delete Article by id
func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, article := range Articles {
		if article.ID == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}

}
