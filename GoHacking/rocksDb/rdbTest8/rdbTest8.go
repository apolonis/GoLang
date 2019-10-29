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

//This app is to collect data with ping ip, and to save it to database ('rocksdb')

//Variable for connection on a router
var (
	device  string = "enp3s0" // name of device
	snaplen int32  = 65535    // how many bytes ur going to collect
	promisc bool   = false    // promiscuous mode
	err     error
	timeout time.Duration = -1 * time.Second
	handle  *pcap.Handle
)

//Structure like class in oop with the atributes we need and collect from ping
type Infor struct {
	Time               string `json:"Time"`
	DestinationAddress string `json:"DestinationAddress"`
	Protocol           string `json:"Protocol"`
	ICMPCode           string `json:"ICMPCode"`
	ICMPSequenceNumber string `json:"ICMPSequenceNumber"`
	PayloadDataLength  string `json:"PayloadDataLength"`
	PayloadData        string `json:"PayloadData"`
}

//This function is for converting the []byte type to string
func convert(b []byte) string {
	s := make([]string, len(b))
	for i := range b {
		s[i] = strconv.Itoa(int(b[i]))
	}
	return strings.Join(s, ",")
}

func main() {

	var ping string
	//Count is for attack variable to change after every ping(attack)
	count := 0
	var pingRead string
	//Connection to database and creating if not exist
	opts := gorocksdb.NewDefaultOptions()
	opts.SetCreateIfMissing(true)
	db, err := gorocksdb.OpenDb(opts, "test8.db")
	if err != nil {
		fmt.Println(err)
	}

	wo := gorocksdb.NewDefaultWriteOptions()
	ro := gorocksdb.NewDefaultReadOptions()

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

		ip_layer := packet.Layer(layers.LayerTypeIPv4)
		ip_packet, _ := ip_layer.(*layers.IPv4)
		icmp_layer := packet.Layer(layers.LayerTypeICMPv4)
		icmp_packet := icmp_layer.(*layers.ICMPv4)

		//Constructor and creating new struct type (class)
		informationVar := Infor{
			Time:               theTimeNow.String(),
			DestinationAddress: ip_packet.DstIP.String(),
			Protocol:           ip_packet.Protocol.String(),
			ICMPCode:           icmp_packet.TypeCode.String(),
			ICMPSequenceNumber: strconv.Itoa(int(icmp_packet.Seq)),
			PayloadDataLength:  strconv.Itoa(int(len(icmp_packet.Payload))), //Converting to string all arguments
			PayloadData:        convert(icmp_packet.Payload)}

		ping = ("ping" + strconv.Itoa(count))

		db.Put(wo, []byte(ping+"/Time"), []byte(informationVar.Time))

		db.Put(wo, []byte(ping+"/DestinationAddress"), []byte(informationVar.DestinationAddress))

		db.Put(wo, []byte(ping+"/Protocol"), []byte(informationVar.Protocol))

		db.Put(wo, []byte(ping+"/ICMPCode"), []byte(informationVar.ICMPCode))

		db.Put(wo, []byte(ping+"/ICMPSequenceNumber"), []byte(informationVar.ICMPSequenceNumber))

		db.Put(wo, []byte(ping+"/PayloadDataLength"), []byte(informationVar.PayloadDataLength))

		db.Put(wo, []byte(ping+"/PayloadData"), []byte(informationVar.PayloadData))

		//count++

		//		ro := gorocksdb.NewDefaultReadOptions()

		// vremeSutraPreteznoOblacno, err := db.Get(ro, []byte("ping0/Time"))

		// fmt.Println(string(vremeSutraPreteznoOblacno.Data()))

		//	count := 0

		//For is to go through all packeges we collect

		pingRead = ("ping" + strconv.Itoa(count))
		theTimeNow1, err := db.Get(ro, []byte(pingRead+"/time"))
		if err != nil {
			fmt.Println("Error", err)
		}

		destinationAddress1, err1 := db.Get(ro, []byte(pingRead+"/DestinationAddress"))
		if err1 != nil {
			fmt.Println("Error", err1)
		}

		protocol1, err2 := db.Get(ro, []byte(pingRead+"/Protocol"))
		if err2 != nil {
			fmt.Println("Error", err2)
		}

		ICMPCode1, err3 := db.Get(ro, []byte(pingRead+"/ICMPCode"))
		if err3 != nil {
			fmt.Println("Error", err3)
		}

		ICMPSequenceNumber1, err4 := db.Get(ro, []byte(pingRead+"/ICMPSequenceNumber"))
		if err4 != nil {
			fmt.Println("Error", err4)
		}

		payloadDataLength1, err5 := db.Get(ro, []byte(pingRead+"/PayloadDataLength"))
		if err5 != nil {
			fmt.Println("Error", err5)
		}

		payloadData1, err6 := db.Get(ro, []byte(pingRead+"/PayloadData"))
		if err6 != nil {
			fmt.Println("Error", err6)
		}

		fmt.Println("-------------------")
		fmt.Println("Time: ", string(theTimeNow1.Data()))
		fmt.Println("Destination Address: " + string(destinationAddress1.Data()))
		fmt.Println("Protocol: ", string(protocol1.Data()))

		fmt.Println("ICMP Code: ", string(ICMPCode1.Data()))
		fmt.Println("ICMP Sequence Number: ", string(ICMPSequenceNumber1.Data()))
		fmt.Println("Payload data length", string(payloadDataLength1.Data()))
		fmt.Println("Payload data: " + string(payloadData1.Data()))
		// fmt.Println("Payload data string format: ",
		// 	convert(icmp_packet.Payload))
		fmt.Println("-------------------\n")

		count++
	}

	db.Close()

}
