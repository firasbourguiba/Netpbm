package netpbm

import (
	"os"
	"strconv"
	"strings"
)

type Image struct {
    magicNumber string
    width       int
    height      int
    data        [][]bool
}

func ReadPBM(file *os.File) (Image, error) {
    var image Image

    // Read magic number
    var magicNumber [2]byte
    _, err := file.Read(magicNumber[:])
    if err != nil {
        return image, err
    }
    image.magicNumber = string(magicNumber[:])

    // Read size
    var sizeLine [10]byte
    _, err = file.Read(sizeLine[:])
    if err != nil {
        return image, err
    }
    sizeParts := strings.Split(string(sizeLine[:]), " ")
    image.width, err = strconv.Atoi(sizeParts[0])
    if err != nil {
        return image, err
    }
    image.height, err = strconv.Atoi(sizeParts[1])
    if err != nil {
        return image, err
    }

    // Create data slice
    image.data = make([][]bool, image.height)
    for i := 0; i < image.height; i++ {
        image.data[i] = make([]bool, image.width)
    }

    // Read image data
    for i := 0; i < image.height; i++ {
        for j := 0; j < image.width; j++ {
            var pixelByte [1]byte
            _, err := file.Read(pixelByte[:])
            if err != nil {
                return image, err
            }
            image.data[i][j] = pixelByte[0] == '#'
        }
    }

    return image, nil
}


