package main

import (
	"log"

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

//Opening a new index(and returning same)
func openNewIndex() bleve.Index {
	indexPath := "goProjectDemo.bleve"
	mapping := bleve.NewIndexMapping()
	index, err := bleve.New(indexPath, mapping)

	if err != nil {
		index, err = bleve.Open(indexPath)

		if err != nil {
			log.Fatal(err)
		}
	}
	return index
}
