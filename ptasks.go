/*This application is used to apply load on the CPU and is meant for testing on containers.
The program is written in golang.
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

//func init() {
//runtime.GOMAXPROCS(runtime.NumCPU())
//}

func main() {
	for {
		wg.Add(10)
		for i := 0; i < 10; i++ { //creates 10 million threads. this can be altered to put different load on the CPU
			go calc() //calls thread to calculate the value of pi
		}

		wg.Wait()
	}

}

func calc() { //function to calculate the value of pi
	var N = 1000000000
	var sum float64
	var term float64
	var sign = 1.0
	for k := 0; k < N; k++ {
		term = float64((1.0) / (float64(2.0)*float64(k) + float64(1.0)))
		time.Sleep(5 * time.Millisecond)
		sum = sum + float64(sign)*term
		sign = -sign
	}
	fmt.Println("Pi=", float64(sum*4.0))
	wg.Done()
}
