package main

import (
	"fmt"
	"time"
)

type Budget struct {
	CampaignID string
	Balance    float64 //USD
	Expires    time.Time
}

func (b Budget) TimeLeft() time.Duration {
	return b.Expires.Sub(time.Now().UTC())
}

func (b *Budget) Update(sum float64) { // If changing values, pass a pointer receiver - remember pass by value
	b.Balance += sum
}

func main() {
	b1 := Budget{"Campaign 001", 22.96, time.Now().Add(7 * 24 * time.Hour)}
	fmt.Printf("%#v\n", b1)
	fmt.Println("Time left ", b1.TimeLeft()) // Added after adding the function Timeleft
	b1.Update(18.14)
	fmt.Println("The new balance", b1.Balance) // prints 41.1

	fmt.Println("----------------------")

	b2 := Budget{
		CampaignID: "December Camp 1",
		Balance:    20.7,
	}
	fmt.Printf("%+v\n", b2)

	fmt.Println("-----------------------------")
	var b3 Budget
	fmt.Printf("%+v\n", b3)

}
