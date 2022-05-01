package main

import (
	"bufio"
	"flag"
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
	//	var output string
	first := flag.Int("from", 0, "from index")
	second := flag.Int("to", 0, "from index")
	flag.Parse()
	// fmt.Println("Flag -o : ", *first)
	// fmt.Println("Positional Args : ", *second)

	first_inx := *first
	second_inx := *second
	// fmt.Println("value", first, second)
	if first_inx < 0 {
		first_inx = 0
	}
	if second_inx < 0 {
		second_inx = 0
	}
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
				if first_inx > 0 && second_inx > 0 && first_inx < second_inx {

					if first_inx <= inx && inx <= second_inx {

						if strings.HasPrefix(j, "txs=") {

							total_txns, _ := strconv.Atoi(strings.Split(j, "=")[1])

							if _, ok := avg_indexfreq[a[2]]; ok {
								avg_indexfreq[a[2]] = avg_indexfreq[a[2]] + total_txns
							} else {

								avg_indexfreq[a[2]] = total_txns
							}

						}
					}
				} else if second_inx == 0 && first_inx != 0 && first_inx <= inx {
					if strings.HasPrefix(j, "txs=") {

						total_txns, _ := strconv.Atoi(strings.Split(j, "=")[1])

						if _, ok := avg_indexfreq[a[2]]; ok {
							avg_indexfreq[a[2]] = avg_indexfreq[a[2]] + total_txns
						} else {

							avg_indexfreq[a[2]] = total_txns
						}

					}
				} else if first_inx == 0 && second_inx != 0 && second_inx >= inx {

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
	var indexes string

	if len(avg_indexfreq) != 0 {
		for _, j := range avg_indexfreq {
			sum = sum + j
		}
		avg = sum / len(avg_indexfreq)
		if second_inx == 0 {
			indexes = "last"
		} else {
			indexes = fmt.Sprintf("%v", second_inx)

		}
		fmt.Printf("Average of transaction per second from  index %v to %v is %v\n", first_inx, indexes, avg)
	} else {
		for _, j := range freq {
			sum = sum + j
		}
		avg = sum / len(freq)
		fmt.Printf("No Specific Indexes are found\n")
		fmt.Printf("Average of transaction per second %v\n", avg)

	}
}
