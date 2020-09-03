package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

func readProblems(fi *string) map[string]int {
	f, err := os.Open(*fi)
	if err != nil {
		fmt.Println("Couldn't read the problems.")
		os.Exit(1)
	}
	ques := make(map[string]int)

	r := csv.NewReader(f)
	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		x, _ := strconv.Atoi(rec[1])
		ques[rec[0]] = x
	}
	return ques
}

func main() {
	var t = flag.Int("t", 10, "The duration of the test in seconds.")
	var f = flag.String("f", "probs.csv", "The csv file with the problems in the form : ques,ans")
	flag.Parse()
	ques := readProblems(f)

	correct := 0
	taken := 0
	var ans int

	c := make(chan int)
	go func() {
		time.Sleep(time.Duration(*t) * time.Second)
		c <- 1
	}()

loop:
	for q, a := range ques {
		select {
		case <-c:
			fmt.Println("Time is up.")
			break loop
		default:
			fmt.Println(q)
			_, err := fmt.Scanf("%d\n", &ans)
			if err != nil {
				fmt.Println(err)
			}
			if ans == a {
				correct++
			}
			taken++
		}
	}
	fmt.Printf("You attempted %v out of %v questions\n", taken, len(ques))
	fmt.Printf("Your score is %v out of %v\n", correct, len(ques))

}
