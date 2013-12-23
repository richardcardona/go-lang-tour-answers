package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
    image := make([][]uint8, dy)
    for row, _ := range image {
        image[row] = make([]uint8, dx)
        for pixel, _ := range image[row] {
            image[row][pixel] = uint8(row^pixel)
        }
    }
    return image
}

func main() {
    pic.Show(Pic)
}