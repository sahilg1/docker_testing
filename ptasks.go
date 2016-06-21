/*This application is used to apply load on the CPU and is meant for testing on containers.
The program is written in golang.
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	for {
		if checkval() {
			wg.Add(1000)
			for i := 0; i < 1000; i++ { //creates 10 million threads. this can be altered to put different load on the CPU
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
		if checkval() {
			term = float64((1.0) / (float64(2.0)*float64(k) + float64(1.0)))
			time.Sleep(5 * time.Millisecond)
			sum = sum + float64(sign)*term
			sign = -sign
		} else {
			os.Exit(3)
		}
	}
	fmt.Println("Pi=", float64(sum*4.0))
	wg.Done()
}

//This function checks the value of the S3 bucket to check to stop or continue running
func checkval() bool {
	res, err := http.Get("https://s3.amazonaws.com/sahgupta-stress-test/s3file.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(res.Body)
	scanner.Scan()
	fmt.Println(scanner.Text())
	if scanner.Text() == "stop" || scanner.Text() == "Stop" {
		res.Body.Close()
		return false
	}
	res.Body.Close()
	return true
}
