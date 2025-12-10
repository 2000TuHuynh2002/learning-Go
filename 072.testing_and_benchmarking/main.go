// For the sake of demonstration, this code is in package `main`, but it could be
// any package. Testing code typically lives in the same package as the code it
// tests.
package main

import ()

// We’ll be testing this simple implementation of an integer minimum. Typically,
// the code we’re testing would be in a source file named something like `intutils.go`,
// and the test file for it would then be named `intutils_test.go`.
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {}
