// Generate a real filesystem from the puzzle input rooted at a temporary directory.
// Then we can use a bash script to solve the problem.
package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

//nolint:unused,deadcode
const example = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k
`

//nolint:cyclop,funlen
func main() {
	flag.Parse()

	input, err := os.Open(flag.Arg(0))
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}

	root, err := os.MkdirTemp(os.TempDir(), "no-space-left-on-device")
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}

	cwd := root

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()

		words := strings.Split(line, " ")
		switch len(words) {
		case 2: //nolint:gomnd
			if words[0] == "$" || words[0] == "dir" { // $ ls || dir <DIR>
				continue
			}

			size, err := strconv.Atoi(words[0])
			if err != nil {
				log.Fatalf("ERROR: %v", err)
			}

			f, err := os.Create(filepath.Join(cwd, words[1]))
			if err != nil {
				log.Fatalf("ERROR: %v", err)
			}

			if err := f.Truncate(int64(size)); err != nil {
				log.Fatal(err)
			}

		case 3: //nolint:gomnd
			cwd = filepath.Join(cwd, words[2])

			if err := os.Mkdir(cwd, os.ModePerm); err != nil {
				if errors.Is(err, fs.ErrExist) {
					continue
				}

				log.Fatalf("ERROR: %v", err)
			}

		default:
			log.Fatalf("ERROR: unexpected line: %q", line)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("ERROR: %v", err)
	}

	//nolint:forbidigo
	fmt.Println(root)
}
