package main

import (
	//nolint:gosec
	"crypto/md5"
	"fmt"
	"strconv"
)

const input = "iwrupvqb"

func one() int {
	for i := 1; i < 1e10; i++ {
		//nolint:gosec
		sum := md5.Sum([]byte(input + strconv.Itoa(i)))
		if fmt.Sprintf("%x", sum)[0:5] == "00000" {
			return i
		}
	}

	return -1
}

func two() int {
	for i := 1; i < 1e10; i++ {
		//nolint:gosec
		sum := md5.Sum([]byte(input + strconv.Itoa(i)))
		if fmt.Sprintf("%x", sum)[0:6] == "000000" {
			return i
		}
	}

	return -1
}

func main() {
	//nolint
	fmt.Println(one())
	///nolint
	fmt.Println(two())
}
