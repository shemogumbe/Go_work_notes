package main

import "fmt"

func main() {
	stocks := map[string]float64{
		"AMZN": 2087.89,
		"GOOG": 2540.00,
		"MSFT": 287.70, // Must have last trailig comma
	}

	//Length - number of keys
	fmt.Println("Lenghth :", len(stocks))

	//Get value - Available
	fmt.Println("Amazon stock ", stocks["AMZN"])

	// Find value - Unavailable
	fmt.Println("Tesla stock", stocks["TSLA"])

	//Assign new key an dvalue
	stocks["TSLA"] = 822.12
	fmt.Println(stocks)

	// Delete key value
	delete(stocks, "AMZN")

	// Loop using range function
	for key := range stocks {
		fmt.Println(key)
	}

	fmt.Println("------------------")

	for key, value := range stocks {
		fmt.Println(key, value)
	}

}
