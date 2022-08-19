package main

import (
	"fmt"
	"time"
)

func main() {

	for {
		fmt.Printf("*** Hello World ****: %v+\n", time.Now())
		time.Sleep(time.Second)
	}

}
