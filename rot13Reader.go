package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(b []byte) (int, error) {
	n, err := rot.r.Read(b)
	if err == nil {
		var firstLetter byte
		for i := 0; i < len(b); i++ {
			if b[i] >= 'A' && b[i] <= 'Z' {
				firstLetter = 'A'
			} else if b[i] >= 'a' && b[i] <= 'z' {
				firstLetter = 'a'
			} else {
				continue
			}
			b[i] = firstLetter + ((b[i] - firstLetter + 13) % 26)
		}
	}
	return n, err	
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
