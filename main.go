package main

import (
	"fmt"
	tsp "google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

//time.Format("02.01.2006 15:04:05")

const shortForm = "02.01.2006"

var date = "09.11.2020"

func printTime() (t time.Time) {
	t, _ = time.Parse(shortForm, date)
	return

}

func main() {

	t := printTime()

	ts := tsp.New(t)

	nt := ts.AsTime().Format("02.01.2006 15:04:05")

	fmt.Printf("Unix time: %v\nConvert to Proto time: %v\nConvert to Unix time: %v\n", t, ts, nt)
}
