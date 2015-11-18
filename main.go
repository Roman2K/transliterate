package main

import (
	"io"
	"log"
	"os"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func main() {
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	r := transform.NewReader(os.Stdin, t)
	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}
