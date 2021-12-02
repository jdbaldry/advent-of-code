package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const inputFile = "input.txt"

var logger = log.New(os.Stderr, "", log.Llongfile)

// EBNF
// letter      = "a" | "b" | "c" | "d" | "e" | "f" | "g"
//             | "h" | "i" | "j" | "k" | "l" | "m" | "n"
//             | "o" | "p" | "q" | "r" | "s" | "t" | "u"
//             | "v" | "w" | "x" | "y" | "z" ;
//
// digit       = "0" | "1" | "2" | "3" | "4" | "5" | "6"
//             | "7" | "8" | "9" ;
//
// wire        = letter { letter } ;
// binop       = signal | wire
//             , "OR" | "AND" | "LSHIFT" | "RSHIFT"
//             , signal | wire ;
// signal      = digit { digit } ;
// unop        = "NOT", wire ;
// instruction = binop| signal | unop, "->" wire ;

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
	input, err := os.Open(inputFile)
	if err != nil {
		logger.Fatalf("Unable to open file %s: %v", inputFile, err)
	}

	s := bufio.NewScanner(input)
	s.Split(bufio.ScanLines)

	var instructions []instruction
	for s.Scan() {
		tokens := strings.Split(s.Text(), " ")
		switch tokens[0][0] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			i, err := strconv.ParseUint(tokens[0], 10, 16)
			if err != nil {
				log.Fatalf("Unable to parse token %q as int: %v", tokens[0], err)
			}
			instructions = append(instructions, signalInstruction{signal(i), tokens[2]})
		}
	}

	for _, instruction := range instructions {
		fmt.Println(instruction)
	}
}
