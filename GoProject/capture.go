package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/tecbot/gorocksdb"
)

func captureTest() {
	i := 0
	for {
		fmt.Println("Capturing i : ", i)
		time.Sleep(3 * time.Second)
		i++
	}
}

func convert(b []byte) string {
	s := make([]string, len(b))
	for i := range b {
		s[i] = strconv.Itoa(int(b[i]))
	}
	return strings.Join(s, ",")
}
func capture() {
	// var *gorocksdb.DB db = openDataBase()
	db, err := openDataBase()

	if err != nil {
		fmt.Println(err)
	}

	var ping string
	var count int
	count = 0

	wo := gorocksdb.NewDefaultWriteOptions()

	handle, err = pcap.OpenLive(device, snaplen, promisc, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	var filter string = "src host 192.168.1.106 and icmp"
	err = handle.SetBPFFilter(filter)
	if err != nil {
		log.Fatal(err)
	}

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	//For is to go through all packeges we collect
	for packet := range packetSource.Packets() {

		//Needed the time when packet is recived
		theTimeNow := time.Now()

		iplayer := packet.Layer(layers.LayerTypeIPv4)
		ippacket, _ := iplayer.(*layers.IPv4)
		icmplayer := packet.Layer(layers.LayerTypeICMPv4)
		icmppacket := icmplayer.(*layers.ICMPv4)

		//this is only for testing if the packet is full or not
		if icmppacket.TypeCode.String() == "EchoRequest" {
			if len(icmppacket.Payload) > 0 {
				log.Println("Info: EchoRequest Recived")
			} else {
				log.Println("Warn: Empty EchoRequest Recived")

			}
		}

		ping = ("ping" + strconv.Itoa(count))

		fmt.Println("\n !!!!!!!!!! DEBUG: nil pointer !!!!!!!!!!!! \n")

		db.Put(wo, []byte(ping+"/Time"), []byte(theTimeNow.String()))

		db.Put(wo, []byte(ping+"/DestinationAddress"), []byte(ippacket.DstIP.String()))

		db.Put(wo, []byte(ping+"/Protocol"), []byte(ippacket.Protocol.String()))

		db.Put(wo, []byte(ping+"/ICMPCode"), []byte(icmppacket.TypeCode.String()))

		db.Put(wo, []byte(ping+"/ICMPSequenceNumber"), []byte(strconv.Itoa(int(icmppacket.Seq))))

		db.Put(wo, []byte(ping+"/PayloadDataLength"), []byte(strconv.Itoa(int(len(icmppacket.Payload)))))

		db.Put(wo, []byte(ping+"/PayloadData"), []byte(convert(icmppacket.Payload)))

		count++

	}
	db.Close()

}
