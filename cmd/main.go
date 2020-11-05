package main

import (
	"bufio"
	"flag"
	"log"
	"os"

	"cmdlimit/internal"
)

const (
	defaultRate     = 1
	defaultInflight = 1
)

func main() {
	rate := flag.Int("rate", defaultRate, "command rate per a second")
	inflight := flag.Int("inflight", defaultInflight, "maximum parallel commands allowed")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		log.Fatal("you must provide a command")
	}
	command, args := args[0], args[1:]

	var input []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		in := scanner.Text()
		if in == "" { // for a manual input
			break
		}
		input = append(input, in)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	limiter := internal.NewLimiter(*rate, *inflight, command, args)
	limiter.Exec(input)
}
