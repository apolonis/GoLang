package main

import (
	"fmt"
	"os"

	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcapgo"
)

func main() {
	dumpFile, _ := os.Create("dump.pcap")
	defer dumpFile.Close()

	packetWriter := pcapgo.NewWriter(dumpFile)
	packetWriter.WriteFileHeader(
		65535,
		layers.LinkTypeEthernet,
	)
	fmt.Println(dumpFile)
	fmt.Println(packetWriter)

}
