package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	log_analysis()

}
func log_analysis() {
	var first string

	// Taking input from user
	fmt.Println("Enter Index From:")

	fmt.Scanln(&first)
	fmt.Println("Enter Index To: ")
	var second string
	fmt.Scanln(&second)
	first_inx, _ := strconv.Atoi(first)
	second_inx, _ := strconv.Atoi(second)

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
	// count := 0
	var freq = make(map[string]int)
	// var txns []string
	var avg_indexfreq = make(map[string]int)
	inx := 0
	for s.Scan() {

		matched, err := regexp.Match("block", []byte(s.Text()))
		if err != nil {
			fmt.Print("\n found", err)

		}
		if matched == true {
			a := strings.Split(s.Text(), " ")
			for _, j := range a {

				if strings.HasPrefix(j, "txs=") {
					total_txns, _ := strconv.Atoi(strings.Split(j, "=")[1])
					if _, ok := freq[a[2]]; ok {
						freq[a[2]] = freq[a[2]] + total_txns
					} else {
						freq[a[2]] = total_txns

					}
				}

				if strings.HasPrefix(j, "index=") {

					inx, _ = strconv.Atoi(strings.Split(j, "=")[1])
				}

				if first_inx <= inx && inx <= second_inx {

					fmt.Println("enter1")
					if strings.HasPrefix(j, "txs=") {

						total_txns, _ := strconv.Atoi(strings.Split(j, "=")[1])
						if _, ok := avg_indexfreq[a[2]]; ok {

							avg_indexfreq[a[2]] = avg_indexfreq[a[2]] + total_txns
						} else {

							avg_indexfreq[a[2]] = total_txns
						}

					}
				}
			}

		}
	}
	max := 0
	var max_txn = make(map[string]int)
	for i, j := range freq {
		if j >= max {
			max = j
			max_txn = make(map[string]int)
			max_txn[i] = j
		}
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}

	for i, j := range max_txn {

		fmt.Printf("Maximum transaction is %v at %v\n", j, i)
	}

	sum := 0
	avg := 0
	for _, j := range avg_indexfreq {

		sum = sum + j

	}
	avg = sum / len(avg_indexfreq)
	fmt.Println("Average of transaction per second from  index %v to %v is %v", first_inx, second_inx, avg)
}
