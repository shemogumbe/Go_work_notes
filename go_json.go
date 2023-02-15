package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type Request struct { // this is a bank transaction
	Login  string  `json:"user"`
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
}

var data = `
{
	"user": "Shem Ogumbe",
	"type":"deposit",
	"amount": 1200.8
}
`

func main() {
	rdr := strings.NewReader(data) // simulates a file/socket

	dec := json.NewDecoder(rdr) // Decode json
	var req Request
	if err := dec.Decode(&req); err != nil {
		log.Fatalf("Error: can't decode - %s", err)
	}

	fmt.Printf("Got: %+v\n", req)

	fmt.Println("-----------------")
	// Create Response
	prevBalance := 1_1000_000.0 // Loaded from database
	resp := map[string]interface{}{
		"ok":      true,
		"balance": prevBalance + req.Amount,
	}

	//Encode Response
	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(resp); err != nil {
		log.Fatalf("error: can't decode - %s", err)
	}

}
