package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"os"
	"strconv"
)

const input = "iwrupvqb"

var logger = log.New(os.Stderr, "", log.Llongfile)

func one() int {
	for i := 1; i < 1e10; i++ {
		sum := md5.Sum([]byte(input + strconv.Itoa(i)))
		if fmt.Sprintf("%x", sum)[0:5] == "00000" {
			return i
		}
	}
	return -1
}

func two() int {
	for i := 1; i < 1e10; i++ {
		sum := md5.Sum([]byte(input + strconv.Itoa(i)))
		if fmt.Sprintf("%x", sum)[0:6] == "000000" {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(one())
	fmt.Println(two())
}
