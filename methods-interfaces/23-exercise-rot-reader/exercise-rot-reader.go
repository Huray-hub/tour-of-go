package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(b[]byte) (int, error) {
    n, err := rot.r.Read(b)

    if err!= nil {
        return 0, err
    }

    for i := 0; i < len(b); i++ {
        if b[i] <= 'Z' { 
            if b[i]+13 > 'Z' {
                b[i]-=13
            } else {
                b[i]+=13
            }
        } else {
            if b[i]+13 > 'z' {
                b[i]-=13
            } else {
                b[i]+=13
            }
        }
    }

    return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
