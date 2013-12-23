package main

import (
    "fmt"
    "io"
    "os"
    "strings"
)

type rot13Reader struct {
    r io.Reader
}

func (r13 *rot13Reader) Read(p []byte) (n int, err error) {
    bytecount, err := r13.r.Read(p)

    if err != nil {
        return bytecount, err
    }

    for i, v := range p {
        switch {
        case v >= 'A' && v <= 'Z':
            p[i] = 'A' + (v - 'A' + 13) % 26
        case v >= 'a' && v <= 'z':
            p[i] = 'a' + (v - 'a' + 13) % 26
        }
    }

    return bytecount, nil
}

func main() {
    s := strings.NewReader(
        "Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
    fmt.Println()
}