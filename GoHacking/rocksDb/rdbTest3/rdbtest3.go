package main

import (
	"fmt"
	"log"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"

	"encoding/json"

	"github.com/tecbot/gorocksdb"
)

var (
	//device string = "en0"
	device  string = "enp3s0" // name of device
	snaplen int32  = 65535    // how many bytes ur going to collect
	//snaplen int32 = 353540
	promisc bool = false // promiscuous mode
	err     error
	timeout time.Duration = -1 * time.Second
	handle  *pcap.Handle
)

type Infor struct {
	SourceAddress      string `json:"SourceAddress"`
	DestinationAddress string `json:"DestinationAddress"`
	Protocol           string `json:"Protocol"`
	ICMPCode           string `json:"ICMPCode"`
	ICMPSequenceNumber string `json:"ICMPSequenceNumber"`
	PayloadDataLength  string `json:"PayloadDataLength"`
	PayloadData        string `json:"PayloadData"`
}

func main() {

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
		fmt.Println("ICMP Sequence Number: ", icmp_packet.Seq)
		fmt.Println("Payload data length", len(icmp_packet.Payload))
		fmt.Println("Payload data: ", icmp_packet.Payload)
		fmt.Println("Payload data to string: ", string(icmp_packet.Payload))
		fmt.Println("-------------------\n")
		//	fmt.Println(packet)

		informationVar := Infor{SourceAddress: ip_packet.SrcIP.String(), //working
			DestinationAddress: ip_packet.DstIP.String(),         //working
			Protocol:           ip_packet.Protocol.String(),      //working
			ICMPCode:           icmp_packet.TypeCode.String(),    //working
			ICMPSequenceNumber: string(icmp_packet.Seq),          //not
			PayloadDataLength:  string(len(icmp_packet.Payload)), //working
			PayloadData:        string(icmp_packet.Payload)}      //not

		opts := gorocksdb.NewDefaultOptions()
		opts.SetCreateIfMissing(true)
		db, err := gorocksdb.OpenDb(opts, "test3.db")
		if err != nil {
			fmt.Println(err)
		}

		wo := gorocksdb.NewDefaultWriteOptions()

		informationPars, err := json.Marshal(informationVar)

		if err != nil {
			fmt.Println(err)
			return
		}

		db.Put(wo, []byte("nekiAttack"), []byte(informationPars))

		db.Close()

	}

}
