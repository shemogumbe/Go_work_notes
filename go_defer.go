package main

import "fmt"

func main() {
	worker()
}

func worker() {
	r1, err := acquire("A")
	if err != nil {
		fmt.Println("Error :", err)
		return
	}

	defer release(r1)
	fmt.Println("Worker")
}

func acquire(name string) (string, error) {
	return name, nil
}

func release(name string) {
	fmt.Printf("Cleaning up %s\n", name)
}
