package main

import (
	"fmt"
	"log"

	"github.com/blevesearch/bleve"
)

func main() {
	index, err := bleve.Open("peopleDemo4.bleve")
	if err != nil {
		log.Fatal(err)
	}

	// query := bleve.NewTermQuery("riki siki")
	// request := bleve.NewSearchRequest(query)

	query := bleve.NewQueryStringQuery("desc:desc 2")
	request := bleve.NewSearchRequest(query)
	results, err := index.Search(request)

	if err != nil {
		log.Fatal(err)
	}

	for _, p := range results.Hits {
		fmt.Println(p.Fields)
	}
}
