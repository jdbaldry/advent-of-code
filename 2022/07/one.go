package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"path/filepath"
	"strconv"
	"strings"
)

var errUnexpectedLine = errors.New("line does not match expected format")

type filesystem map[string]fileinfo

type fileinfo struct {
	name      string
	size      int
	directory bool
	entries   []string
}

// cd follows dir from the cwd, updating the fsys table as it does.
// The new cwd is returned.
func cd(fsys filesystem, cwd, dir string) string {
	switch {
	case dir == "/", dir == ".." && cwd == "/":
		cwd = "/"
	case dir == "..":
		cwd = filepath.Dir(cwd)
	default:
		cwd = filepath.Join(cwd, dir)
	}

	if _, ok := fsys[cwd]; !ok {
		fsys[cwd] = fileinfo{
			name:      cwd,
			size:      0,
			directory: true,
			entries:   []string{},
		}
	}

	return cwd
}

func ls(fsys filesystem, cwd, size, name string) error {
	path := filepath.Join(cwd, name)

	listed := fileinfo{
		name:      path,
		size:      0,
		directory: false,
		entries:   []string{},
	}
	if size == "dir" {
		listed.directory = true
	} else {
		sizeInt, err := strconv.Atoi(size)
		if err != nil {
			return fmt.Errorf("unable to convert to integer: %w", err)
		}
		listed.size = sizeInt
	}

	dir := fsys[cwd]
	dir.entries = append(dir.entries, path)
	fsys[path] = listed
	fsys[cwd] = dir

	return nil
}

func walk(fsys filesystem, root string, visit func(fileinfo)) {
	file, ok := fsys[root]
	if !ok {
		return
	}

	for _, path := range fsys[root].entries {
		entry, ok := fsys[path]
		if !ok {
			return
		}

		if entry.directory {
			walk(fsys, path, visit)
		}

		visit(entry)
	}

	visit(file)
}

// du returns a list of directories with file size information.
// The computed file size information is recorded in the directory
// fileinfo.
func du(fsys filesystem) []fileinfo {
	var dirs []fileinfo

	for _, file := range fsys {
		if file.directory {
			var sum int

			walk(fsys, file.name, func(f fileinfo) {
				if !f.directory {
					sum += f.size
				}
			})

			file.size = sum
			fsys[file.name] = file

			dirs = append(dirs, file)
		}
	}

	return dirs
}

func constructFilesystem(r io.Reader, fsys filesystem) error {
	cwd := "/"

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		words := strings.Split(line, " ")
		switch len(words) {
		case 2: //nolint:gomnd
			if words[0] == "$" { // $ ls
				continue
			}

			if err := ls(fsys, cwd, words[0], words[1]); err != nil {
				return fmt.Errorf("%q: %w", cwd, err)
			}

		case 3: //nolint:gomnd
			cwd = cd(fsys, cwd, words[2])

		default:
			return fmt.Errorf("%q: %w", line, errUnexpectedLine)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("%w during scanning", err)
	}

	return nil
}

//nolint:varnamelen
func one(r io.Reader) (int, error) {
	fsys := filesystem{
		"/": fileinfo{
			name:      "/",
			size:      0,
			directory: true,
			entries:   []string{},
		},
	}

	if err := constructFilesystem(r, fsys); err != nil {
		return 0, err
	}

	var total int

	for _, dir := range du(fsys) {
		if maxSize := 100000; dir.size < maxSize {
			total += dir.size
		}
	}

	return total, nil
}
