package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/document"
)

type Person struct {
	UID  string `json:"uid"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

func getDocsFromSearchResults(
	index bleve.Index,
	results *bleve.SearchResult,
) ([][]byte, error) {
	docs := make([][]byte, 0)

	for _, val := range results.Hits {
		id := val.ID
		doc, _ := index.Document(id)

		rv := struct {
			ID     string                 `json:"id"`
			Fields map[string]interface{} `json:"fields"`
		}{
			ID:     id,
			Fields: map[string]interface{}{},
		}
		for _, field := range doc.Fields {
			var newval interface{}
			switch field := field.(type) {
			case *document.TextField:
				newval = string(field.Value())
			case *document.NumericField:
				n, err := field.Number()
				if err == nil {
					newval = n
				}
			case *document.DateTimeField:
				d, err := field.DateTime()
				if err == nil {
					newval = d.Format(time.RFC3339Nano)
				}
			}
			existing, existed := rv.Fields[field.Name()]
			if existed {
				switch existing := existing.(type) {
				case []interface{}:
					rv.Fields[field.Name()] = append(existing, newval)
				case interface{}:
					arr := make([]interface{}, 2)
					arr[0] = existing
					arr[1] = newval
					rv.Fields[field.Name()] = arr
				}
			} else {
				rv.Fields[field.Name()] = newval
			}
		}
		j2, _ := json.MarshalIndent(rv, "", "    ")
		docs = append(docs, j2)
	}

	return docs, nil
}

func main() {
	// open a new index
	indexPath := "peopleDemo7.bleve"
	mapping := bleve.NewIndexMapping()
	index, err := bleve.New(indexPath, mapping)

	if err != nil {
		// log.Fatal(err)
		index, err = bleve.Open(indexPath)

		if err != nil {
			log.Fatal(err)
		}
	}

	persons := []Person{
		Person{"11", "Viktroija Beckham", "Desc 1"},
		Person{"22", "Riki Siki", "Desc 2"},
	}

	// index some data
	for _, p := range persons {
		err := index.Index(p.UID, p)

		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Indexed Document")

	// search
	query := bleve.NewQueryStringQuery("desc:desc 2")
	request := bleve.NewSearchRequest(query)
	results, err := index.Search(request)

	if err != nil {
		log.Fatal(err)
	}

	type Doc struct {
		Id     string `json:"id"`
		Fields Person `json:"fields"`
	}

	fmt.Println(results)
	docs, err := getDocsFromSearchResults(index, results)
	var stringInterfaceMap Doc

	for _, d := range docs {
		err := json.Unmarshal(d, &stringInterfaceMap)

		if err != nil {
			log.Fatal(err)
		}

		var p Person = stringInterfaceMap.Fields
		fmt.Println(p)
	}
}
