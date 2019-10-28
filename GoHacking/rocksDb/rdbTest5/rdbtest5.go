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
	SourceAddress      string `json:"SourceAddress"`
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
	//This variable is created to be a common key in database
	var attack string
	//Count is for attack variable to change after every ping(attack)
	count := 0

	//Connection to database and creating if not exist
	opts := gorocksdb.NewDefaultOptions()
	opts.SetCreateIfMissing(true)
	db, err := gorocksdb.OpenDb(opts, "test5.db")
	if err != nil {
		fmt.Println(err)
	}

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

		ethernet_layer := packet.Layer(layers.LayerTypeEthernet)
		ethernet_frame := ethernet_layer.(*layers.Ethernet)

		ip_layer := packet.Layer(layers.LayerTypeIPv4)
		ip_packet, _ := ip_layer.(*layers.IPv4)
		icmp_layer := packet.Layer(layers.LayerTypeICMPv4)
		icmp_packet := icmp_layer.(*layers.ICMPv4)

		if icmp_packet.TypeCode.String() == "EchoRequest" {
			if len(icmp_packet.Payload) > 0 {
				log.Println("Info: EchoRequest Recived")
			} else {
				log.Println("Warn: Empty EchoRequest Recived")
				ethernet_frame_copy := *ethernet_frame
				ip_packet_copy := *ip_packet
				icmp_packet_copy := *icmp_packet

				ethernet_frame_copy.SrcMAC = ethernet_frame.DstMAC
				ethernet_frame_copy.DstMAC = ethernet_frame.SrcMAC

				ip_packet_copy.SrcIP = ip_packet.DstIP
				ip_packet_copy.DstIP = ip_packet.SrcIP

				icmp_packet_copy.TypeCode = layers.ICMPv4TypeEchoReply

				var buffer gopacket.SerializeBuffer
				var options gopacket.SerializeOptions
				options.ComputeChecksums = true

				buffer = gopacket.NewSerializeBuffer()
				gopacket.SerializeLayers(buffer, options,
					&ethernet_frame_copy,
					&ip_packet_copy,
					&icmp_packet_copy,
					gopacket.Payload(icmp_packet_copy.Payload),
				)
				new_message := buffer.Bytes()
				err = handle.WritePacketData(new_message)
				if err != nil {
					log.Fatal(err)
				}
			}
		}

		fmt.Println("-------------------")
		fmt.Println("Source address: " + ip_packet.SrcIP.String())
		fmt.Println("Destination Address: " + ip_packet.DstIP.String())
		fmt.Println("Protocol: ", ip_packet.Protocol)

		fmt.Println("ICMP Code: ", icmp_packet.TypeCode)
		fmt.Println("ICMP Sequence Number: ", strconv.Itoa(int(icmp_packet.Seq)))
		fmt.Println("Payload data length", len(icmp_packet.Payload))
		fmt.Println("Payload data: ", icmp_packet.Payload)
		fmt.Println("Payload data string format: ",
			convert(icmp_packet.Payload))

		fmt.Println("-------------------\n")

		//Constructor and creating new struct type (class)
		informationVar := Infor{SourceAddress: ip_packet.SrcIP.String(),
			DestinationAddress: ip_packet.DstIP.String(),
			Protocol:           ip_packet.Protocol.String(),
			ICMPCode:           icmp_packet.TypeCode.String(),
			ICMPSequenceNumber: strconv.Itoa(int(icmp_packet.Seq)),
			PayloadDataLength:  strconv.Itoa(int(len(icmp_packet.Payload))), //Converting to string all arguments
			PayloadData:        convert(icmp_packet.Payload)}

		attack = ("attack" + strconv.Itoa(count))

		db.Put(wo, []byte(attack), []byte("aloha"))

		db.Put(wo, []byte(attack+"/DestinationAddress"), []byte(informationVar.DestinationAddress))

		db.Put(wo, []byte(attack+"/Protocol"), []byte(informationVar.Protocol))

		db.Put(wo, []byte(attack+"/ICMPCode"), []byte(informationVar.ICMPCode))

		db.Put(wo, []byte(attack+"/ICMPSequenceNumber"), []byte(informationVar.ICMPSequenceNumber))

		db.Put(wo, []byte(attack+"/PayloadDataLength"), []byte(informationVar.PayloadDataLength))

		db.Put(wo, []byte(attack+"/PayloadData"), []byte(informationVar.PayloadData))

		count++

	}
	db.Close()
}
