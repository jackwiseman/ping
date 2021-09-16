package main

import (
	"os/exec"
	"log"
	"fmt"
	"strconv"
)

func ping(count int, site string, channel chan string) {
	var c string = strconv.Itoa(count)

	out, err := exec.Command("ping", "-c", c,  site).Output()
	if err != nil {
		log.Fatal(err)
	}
	var result string = string(out)

	channel <- result
//	fmt.Println(string(out))
}

func main() {
	// these two will eventually be decided as user input
	var countInput int = 5
	var siteInput string = "google.com"
	var numRoutines int = 3

	channel := make(chan string)

	// create numRoutines subroutines pinging siteInput (countInput times each)
	for r := 0; r < numRoutines; r++ {
		go ping(countInput, siteInput, channel)
	}

	// read subroutines channels and print out ping result
	for i := 1; i <= numRoutines; i++ {
		fmt.Println(<-channel)
		fmt.Println("\n")
	}
}
