package atchelper

import (
	"bytes"
	"io"
	"strings"
)

// Checker is helper (os.Stdin -> f -> os.Stdout)
func Checker(input string, f func(io.Reader, io.Writer)) string {
	reader := strings.NewReader(input)
	writer := bytes.NewBufferString("")
	f(reader, writer)
	return writer.String()
}
