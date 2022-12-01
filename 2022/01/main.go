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

	cf, err := fetcher.NewCachingFetcher("https://adventofcode.com/2022/day/1/input", cookie, "input.txt")
	if err != nil {
		if errors.Is(err, fetcher.ErrSessionCookieRequired) {
			log.Fatalf("Unable to fetch input: AOC_SESSION_COOKIE environment variable is unset: %v\n", err)
		}
		log.Fatalf("Unable to create fetcher: %v\n", err)
	}

	r, err := cf.Fetch()
	if err != nil {
		log.Fatalf("Unable to fetch input: %v\n", err)
	}

	parsed, err := parse(r)
	if err != nil {
		log.Fatalf("Unable to parse input: %v\n", err)
	}

	fmt.Println(one(parsed))
	fmt.Println(twoNoSort(parsed))
}
