package main

import "time"
import "fmt"

func main() {
	ticker := time.NewTimer(time.Millisecond * 500)
	go func(){
		for t := range ticker.C {
			fmt.Println("Ticket at", t)
		}
	}()

	time.Sleep(time.Millisecond * 3200)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
