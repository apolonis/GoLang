package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/document"
)

//Function to get all documents from search results
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

//Function analyze (searching all data from bleve)
func analyze() {

	//opening a new index
	index := openNewIndex()

	for {
		fmt.Println("analyze")

		//Searching by key
		query := bleve.NewQueryStringQuery("protocol:ICMPv4")
		request := bleve.NewSearchRequest(query)
		results, err := index.Search(request)

		if err != nil {
			log.Fatal(err)
		}
		//curent struct for formating documents
		type Doc struct {
			ID     string   `json:"id"`
			Fields PcapData `json:"fields"`
		}
		//getting all documents from search results
		docs, err := getDocsFromSearchResults(index, results)

		var stringInterfaceMap Doc

		for _, d := range docs {
			//Unmarshaling []byte type documents into doc type struct
			err := json.Unmarshal(d, &stringInterfaceMap)

			if err != nil {
				log.Fatal("ERROR", err)

			}
			//Taking doc type field ('fields') and creating new variable
			var pcapData PcapData = stringInterfaceMap.Fields
			fmt.Println(pcapData)

			time.Sleep(2 * time.Second)

		}
	}

}
