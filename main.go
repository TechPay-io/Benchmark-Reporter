package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

// const time = time.ParseDuration("1s").Milliseconds()

func main() {
	var delay = time.Duration(1000) * time.Millisecond

	// loop the function until terminated
	// for {
	// 	// update the price
	// 	log_analysis()
	// 	// wait for termination or delay
	// 	// select {
	// 	// case <-pro.sigClose:
	// 	// 	// stop signal received
	// 	// 	return
	// 	// case <-time.After(delay):
	// 	// 	// we repeat the function
	// 	// }
	// }
	for {
		wg.Add(1)
		go log_analysis()
		wg.Wait()
		time.After(delay)
	}
}
func log_analysis() {
	if _, err := os.Stat("./photon.log"); err != nil {
		log.Printf("File doesn't exist %v", err)
	}

	f, err := os.Open("./photon.log")
	if err != nil {
		log.Printf("File can't be open %v", err)

	}
	// fmt.Println("file data", file)
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	s := bufio.NewScanner(f)
	i := 0
	// count := 0
	var freq = make(map[string]int)
	// var txns []string
	for s.Scan() {
		a := strings.Split(s.Text(), " ")
		if a[8] == "block" {
			txns := strings.Split(a[len(a)-3], "=")
			total_txns, _ := strconv.Atoi(txns[1])
			if _, ok := freq[a[2]]; ok {
				freq[a[2]] = freq[a[2]] + total_txns
			} else {
				freq[a[2]] = total_txns
			}
		}

		i++
	}
	max := 0
	var max_txn = make(map[string]int)
	for i, j := range freq {
		if j >= max {
			max = j
			max_txn = make(map[string]int)
			max_txn[i] = j
		}
		// fmt.Println("New Block", i, j)
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
	for i, j := range max_txn {

		fmt.Printf("Maximum transaction is %v at %v", j, i)
	}
	wg.Done()
}
