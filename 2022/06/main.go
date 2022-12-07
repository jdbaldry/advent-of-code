package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	log "github.com/golang/glog"
	"github.com/jdbaldry/advent-of-code/pkg/fetcher"
)

func main() {
	flag.Parse()

	cookie := os.Getenv("AOC_SESSION_COOKIE")

	//nolint:varnamelen
	cf, err := fetcher.NewCachingFetcher("https://adventofcode.com/2022/day/6/input", cookie, "input.txt")
	if err != nil {
		log.Fatalf("Unable to create fetcher: %v\n", err)
	}

	//nolint:varnamelen
	r, err := cf.Fetch()
	if err != nil {
		if errors.Is(err, fetcher.ErrSessionCookieRequired) {
			log.Fatalf("Unable to fetch input: AOC_SESSION_COOKIE environment variable is unset: %v\n", err)
		}

		log.Fatalf("Unable to fetch input: %v\n", err)
	}

	got, err := one(r)
	if err != nil {
		log.Fatalf("Unable to solve part one: %v\n", err)
	}
	//nolint:forbidigo
	fmt.Println(got)

	if _, err := r.Seek(0, 0); err != nil {
		log.Fatalf("Unable to seek to start of input: %v\n", err)
	}

	got, err = two(r)
	if err != nil {
		log.Fatalf("Unable to solve part two: %v\n", err)
	}
	//nolint:forbidigo
	fmt.Println(got)
}
