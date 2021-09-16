package main

import (
	"os/exec"
	"log"
	"fmt"
	"strconv"
	"strings"
)

func ping(count int, site string, channel chan string) {
	var c string = strconv.Itoa(count)

	out, err := exec.Command("ping", "-c", c,  site, "-q").Output()
	if err != nil {
		log.Fatal(err)
	}
	var result string = string(out)

	channel <- result
}

/*
Trims all unneccessary info from commandline ping -q call, so that we only see
ping stats in the form min/avg/max/stddev in ms
*/
func printer(raw string) {
	sliceBegin := strings.Index(raw, "=")
	sliceEnd := strings.Index(raw[sliceBegin + 2:], " ")
	fmt.Println(raw[sliceBegin + 2:sliceBegin + sliceEnd + 2]) // take position after "= "
}

func main() {
	// these two will eventually be decided as user input
	var countInput int = 5
	var siteInput string = "google.com"
	var numRoutines int = 5

	channel := make(chan string)

	// create numRoutines subroutines pinging siteInput (countInput times each)
	for r := 0; r < numRoutines; r++ {
		go ping(countInput, siteInput, channel)
	}

	// read subroutines channels and print out ping result
	for i := 0; i < numRoutines; i++ {
		printer(<-channel)
		//fmt.Println(<-channel)
	}
}
