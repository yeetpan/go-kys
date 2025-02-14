package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time")

	presentTime := time.Now()

	//standard GO time-> 01-02-2006 15:04:05 Monday
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	//Create Date
	createdDate := time.Date(2020, time.April, 07, 04, 21, 22, 02, time.Local)
	fmt.Println(createdDate)

	// go build builds an exe (or an executable for any OS )
}
