package main

import (
	"io"
	"log"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(p []byte) (int, error) {
	n, err := r.r.Read(p)
	if err != nil {
		return n, err
	}
	for i := 0; i < n; i++ {
		switch e := p[i]; {
		case (e >= 'A' && e < 'N') || (e >= 'a' && e < 'n'):
			p[i] = e + 13
		case (e >= 'M' && e < 'Z') || (e >= 'm' && e < 'z'):
			p[i] = e - 13
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	_, err := io.Copy(os.Stdout, &r)
	if err != nil {
		log.Fatal(err)
	}
}
