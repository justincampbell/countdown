package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var (
	help string

	// Version is set via buildflags in the Makefile
	Version string
)

func init() {
	if Version == "" {
		Version = "DEV"
	}

	help = fmt.Sprintf(`countdown %s https://github.com/justincampbell/countdown
usage: countdown DURATION
example durations: 1m, 2m30s, 3h`, Version)
}

func main() {
	flag.Parse()
	arg := flag.Arg(0)

	switch arg {
	case "-h", "--help", "help":
		fmt.Println(help)
		os.Exit(0)
	case "":
		fmt.Println(help)
		os.Exit(1)
	}

	d, err := time.ParseDuration(arg)
	if err != nil {
		fmt.Println(err)
		fmt.Println(help)
		os.Exit(1)
	}

	interval := time.Tick(time.Second)
	end := time.After(d)
	endTime := time.Now().Add(d)

	countdown(interval, end, endTime)
	fmt.Printf("\n")
}

func countdown(interval <-chan time.Time, end <-chan time.Time, endTime time.Time) {
	printTimeRemaining(endTime)

	select {
	case _ = <-interval:
		countdown(interval, end, endTime)
	case <-end:
	}
}

func printTimeRemaining(end time.Time) {
	d := end.Sub(time.Now())
	fmt.Printf("\r%s", formatDuration(d))
}

func formatDuration(d time.Duration) string {
	if d < 0 {
		return "0:00"
	}

	r := int(d.Seconds() + 0.5)
	m := r / 60
	s := r % 60

	return fmt.Sprintf("%d:%02d", m, s)
}
