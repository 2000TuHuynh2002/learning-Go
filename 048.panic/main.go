package main

import (
	"os"
	"path/filepath"
)

func main() {
	// A common use of panic is to abort if a function returns an error value that
	// we don’t know how to (or want to) handle. Here’s an example of `panic`king
	// if we get an unexpected error when creating a new file.
	path := filepath.Join(os.TempDir(), "file")
	_, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	// We’ll use panic throughout this site to check for unexpected errors. This is
	// the only program on the site designed to panic.
	panic("a problem")
}
