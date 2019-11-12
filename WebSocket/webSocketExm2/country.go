package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Response struct {
	Name      string      `json:"name"`
	Countries []Countries `json:"data"`
}

type Countries struct {
	Country string `json:"country"`
	Code    string `json:"code"`
}

func getCountries() ([]Countries, error) {

	fmt.Printf("DEBUG IN GET COUNTRIES \n")

	response, err := http.Get("https://api.coingecko.com/api/v3/events/countries")

	if err != nil {
		fmt.Print("DEBUG: error getting countries from api", err.Error())
		//os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("DEBUG: error reading all countries from req body", err)
	}
	//	fmt.Println(string(responseData))
	var responseObject Response
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		log.Fatal("DEBUG: error unmarshaling in get countries", err)
	}
	// fmt.Println(len(responseObject.Countries))

	// for i := 0; i < len(responseObject.Countries); i++ {
	// 	fmt.Println(responseObject.Countries[i])
	// }
	return responseObject.Countries, err

}
