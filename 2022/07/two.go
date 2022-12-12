package main

import (
	"io"
)

// removalCandidate finds the directory with the smallest size that
// if removed would free enough disk space.
func removalCandidate(fsys filesystem, used int) (fileinfo, int) {
	need, total := 30000000, 70000000

	candidate := fsys["/"]

	for _, file := range fsys {
		if file.directory {
			if available := total - used; file.size+available > need {
				if file.size < candidate.size {
					candidate = file
				}
			}
		}
	}

	return candidate, candidate.size
}

//nolint:varnamelen
func two(r io.Reader) (int, error) {
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

	du(fsys)

	_, size := removalCandidate(fsys, fsys["/"].size)

	return size, nil
}
