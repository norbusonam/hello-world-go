package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	title   string
	content []string
}

func main() {

	bookshelf := map[string]Book{}
	book1 := Book{
		"book1",
		[]string{
			"b1 p1",
			"b1 p2",
			"b1 p3",
			"b1 p4",
			"b1 p5",
		},
	}
	bookshelf[book1.title] = book1
	book2 := Book{
		"book2",
		[]string{
			"b2 p1",
			"b2 p2",
			"b2 p3",
			"b2 p4",
			"b2 p5",
		},
	}
	bookshelf[book2.title] = book2

	router := mux.NewRouter()

	router.HandleFunc("/book/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		book, ok := bookshelf[vars["title"]]
		// check book was found
		if !ok {
			fmt.Fprint(w, "book not found!")
			return
		}
		page, err := strconv.Atoi(vars["page"])
		// check string was converted to int successfully
		if err != nil {
			fmt.Fprint(w, "page was not a number!")
			return
		}
		// check page in bounds
		if page <= 0 || page > len(book.content) {
			fmt.Fprint(w, "that page doesnt exist!")
			return
		}
		pageContent := book.content[page-1]
		fmt.Fprint(w, pageContent)
	}).Methods()

	http.ListenAndServe(":80", router)

}
