package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// EBNF
//
//	letter      = "a" | "b" | "c" | "d" | "e" | "f" | "g"
//	            | "h" | "i" | "j" | "k" | "l" | "m" | "n"
//	            | "o" | "p" | "q" | "r" | "s" | "t" | "u"
//	            | "v" | "w" | "x" | "y" | "z" ;
//
//	digit       = "0" | "1" | "2" | "3" | "4" | "5" | "6"
//	            | "7" | "8" | "9" ;
//
//	wire        = letter { letter } ;
//	binop       = signal | wire
//	            , "OR" | "AND" | "LSHIFT" | "RSHIFT"
//	            , signal | wire ;
//	signal      = digit { digit } ;
//	unop        = "NOT", wire ;
//	instruction = binop| signal | unop, "->" wire ;
type signal uint16

type instruction interface {
	instruction()
}

type signalInstruction struct {
	signal signal
	wire   string
}

func (i signalInstruction) instruction() {}

func main() {
	logger := log.New(os.Stderr, "", log.Llongfile)

	input, err := os.Open("input.txt")
	if err != nil {
		logger.Fatalf("ERROR: %v", err)
	}

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	var instructions []instruction

	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		switch tokens[0][0] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			i, err := strconv.ParseUint(tokens[0], 10, 16)
			if err != nil {
				log.Fatalf("ERROR: unable to parse token %q as int: %v", tokens[0], err)
			}

			instructions = append(instructions, signalInstruction{signal(i), tokens[2]})
		}
	}

	for _, instruction := range instructions {
		//nolint:forbidigo
		fmt.Println(instruction)
	}
}
