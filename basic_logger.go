package main

import (
	"log"
	"time"
)

func main() {
	timer := time.NewTicker(time.Second * 1)

	for {
		t := <-timer.C
		log.Println(t)
	}
}
