package main

import (
	"fmt"
	"strconv"

	"github.com/tecbot/gorocksdb"
)

func main() {

	opts := gorocksdb.NewDefaultOptions()
	opts.SetCreateIfMissing(true)
	db, err := gorocksdb.OpenDbForReadOnly(opts, "test6.db", true)
	if err != nil {
		fmt.Println(err)
	}

	ro := gorocksdb.NewDefaultReadOptions()

	// vremeSutraPreteznoOblacno, err := db.Get(ro, []byte("ping0/Time"))

	// fmt.Println(string(vremeSutraPreteznoOblacno.Data()))

	var ping string
	count := 0

	for i := 0; i < 5; i++ {
		ping = ("ping" + strconv.Itoa(count))
		theTimeNow, err := db.Get(ro, []byte(ping+"/time"))
		if err != nil {
			fmt.Println("Error", err)
		}

		destinationAddress, err1 := db.Get(ro, []byte(ping+"/DestinationAddress"))
		if err1 != nil {
			fmt.Println("Error", err1)
		}

		protocol, err2 := db.Get(ro, []byte(ping+"/Protocol"))
		if err2 != nil {
			fmt.Println("Error", err2)
		}

		ICMPCode, err3 := db.Get(ro, []byte(ping+"/ICMPCode"))
		if err3 != nil {
			fmt.Println("Error", err3)
		}

		ICMPSequenceNumber, err4 := db.Get(ro, []byte(ping+"/ICMPSequenceNumber"))
		if err4 != nil {
			fmt.Println("Error", err4)
		}

		payloadDataLength, err5 := db.Get(ro, []byte(ping+"/PayloadDataLength"))
		if err5 != nil {
			fmt.Println("Error", err5)
		}

		payloadData, err6 := db.Get(ro, []byte(ping+"/PayloadData"))
		if err6 != nil {
			fmt.Println("Error", err6)
		}

		fmt.Println("-------------------")
		fmt.Println("Time: ", string(theTimeNow.Data()))
		fmt.Println("Destination Address: " + string(destinationAddress.Data()))
		fmt.Println("Protocol: ", string(protocol.Data()))

		fmt.Println("ICMP Code: ", string(ICMPCode.Data()))
		fmt.Println("ICMP Sequence Number: ", string(ICMPSequenceNumber.Data()))
		fmt.Println("Payload data length", string(payloadDataLength.Data()))
		fmt.Println("Payload data: " + string(payloadData.Data()))
		// fmt.Println("Payload data string format: ",
		// 	convert(icmp_packet.Payload))
		fmt.Println("-------------------\n")

		count++
	}

	db.Close()

}
