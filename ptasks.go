/*This application is used to apply load on the CPU and is meant for testing on containers.
The program is written in golang.
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup

//func init() {
//runtime.GOMAXPROCS(runtime.NumCPU())
//}

func main() {
	res, err := http.Get("https://s3.amazonaws.com/sahgupta-booster-stress-test/s3_file.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	scanner.Scan()
	fmt.Println(scanner.Text())

	if scanner.Text() == "Start" || scanner.Text() == "start" {
		for {
			wg.Add(10000000)
			for i := 0; i < 10000000; i++ { //creates 10 million threads. this can be altered to put different load on the CPU
				go calc() //calls thread to calculate the value of pi
			}
			wg.Wait()
		}
	}
}

func calc() { //function to calculate the value of pi
	var N = 1000000000
	var sum float64
	var term float64
	var sign = 1.0
	for k := 0; k < N; k++ {
		term = float64((1.0) / (float64(2.0)*float64(k) + float64(1.0)))
		time.Sleep(2 * time.Millisecond)
		sum = sum + float64(sign)*term
		sign = -sign
	}
	fmt.Println("Pi=", float64(sum*4.0))
	wg.Done()
}
