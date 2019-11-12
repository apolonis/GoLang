package main

import (
	"fmt"
	"log"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/pborman/uuid"
)

//Variable's for router
var (
	device  string = "enp3s0" // name of device
	snaplen int32  = 65535    // how many bytes app is going to collect
	promisc bool   = false    // promiscuous mode
	err     error
	timeout time.Duration = -1 * time.Second
	handle  *pcap.Handle
)

//function for converting byte type to string
func convert(b []byte) string {

	s := make([]string, len(b))

	for i := range b {
		s[i] = strconv.Itoa(int(b[i]))
	}

	return strings.Join(s, ",")
}

/*Function capture (using pcap to catch the trafic by ping and
using bleve index to store it)*/
func capture() {

	//opening a new index(take index opened in main)
	indexMutex.Lock()
	index := openNewIndex()
	indexMutex.Unlock()

	//Openning pcap live (using variable's of router)
	handle, err = pcap.OpenLive(device, snaplen, promisc, timeout)

	if err != nil {
		log.Fatal("Error/opening device ", err)
	}

	defer handle.Close()

	//Ip of the host
	var filter string = "src host 192.168.1.106 and icmp"

	err = handle.SetBPFFilter(filter)

	if err != nil {
		log.Fatal("Error/filtering ", err)
	}

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	//Go through all packeges we collect
	for packet := range packetSource.Packets() {
		fmt.Println("capture")

		//Needed the time when packet is recived
		theTimeNow := time.Now()

		iplayer := packet.Layer(layers.LayerTypeIPv4)
		ippacket, _ := iplayer.(*layers.IPv4)
		icmplayer := packet.Layer(layers.LayerTypeICMPv4)
		icmppacket := icmplayer.(*layers.ICMPv4)

		//this is only for testing if the packet is recived or not
		if icmppacket.TypeCode.String() == "EchoRequest" {
			if len(icmppacket.Payload) > 0 {
				log.Println("Info: EchoRequest Recived")
			} else {
				log.Println("Warn: Empty EchoRequest Recived")

			}
		}

		//Generating uuid of PcapData(packet)
		id := uuid.NewUUID()

		//Converting to string all arguments
		//Creating an object of a struct
		pcapData := PcapData{
			UUID:               id.String(),
			Time:               theTimeNow.String(),
			DestinationAddress: ippacket.DstIP.String(),
			Protocol:           ippacket.Protocol.String(),
			ICMPCode:           icmppacket.TypeCode.String(),
			ICMPSequenceNumber: strconv.Itoa(int(icmppacket.Seq)),
			PayloadDataLength:  strconv.Itoa(int(len(icmppacket.Payload))),
			PayloadData:        convert(icmppacket.Payload),
		}
		//Here we're indexing data we capture from pcap
		err := index.Index(pcapData.UUID, pcapData)

		if err != nil {
			log.Fatal("Error/indexing", err)
		}

		fmt.Println("Indexed Document")
		// time.Sleep(1 * time.Millisecond)
		runtime.Gosched()
	}

}
