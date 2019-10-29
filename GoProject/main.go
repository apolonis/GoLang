package main

import (
	"fmt"
	"time"

	"github.com/google/gopacket/pcap"
	"github.com/tecbot/gorocksdb"
)

var (
	device  string = "enp3s0" // name of device
	snaplen int32  = 65535    // how many bytes ur going to collect
	promisc bool   = false    // promiscuous mode
	err     error
	timeout time.Duration = -1 * time.Second
	handle  *pcap.Handle
)

func openDataBase() (*gorocksdb.DB, error) {
	opts := gorocksdb.NewDefaultOptions()
	opts.SetCreateIfMissing(true)
	db, err := gorocksdb.OpenDb(opts, "Example.db")
	return db, err
}

func main() {
	// go captureTest()
	// go analyzeTest()
	go capture()
	go analyze()
	fmt.Scanln()
}
