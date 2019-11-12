package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/blevesearch/bleve"
)

//Data catching from pcap structure
type PcapData struct {
	UUID               string `json:"uuid"`
	Time               string `json:"time"`
	DestinationAddress string `json:"destinationAddress"`
	Protocol           string `json:"protocol"`
	ICMPCode           string `json:"iCMPCode"`
	ICMPSequenceNumber string `json:"iCMPSequenceNumber"`
	PayloadDataLength  string `json:"payloadDataLength"`
	PayloadData        string `json:"payloadData"`
}

var mainIndex bleve.Index = nil
var indexMutex = &sync.Mutex{}

// Opening a new index(and returning same)
func openNewIndex() bleve.Index {
	fmt.Println("openNewIndex 0")

	if mainIndex != nil {
		fmt.Println("openNewIndex 1a")
		return mainIndex
	}

	// Name(path) of the index
	indexPath := "goProjectDemo.bleve"
	mapping := bleve.NewIndexMapping()
	index, err := bleve.New(indexPath, mapping)

	if err != nil {
		index, err = bleve.Open(indexPath)

		if err != nil {
			log.Fatal("Error/opening index/bleve", err)
		}
	}

	mainIndex = index
	fmt.Println("openNewIndex 1b")
	return index
}
