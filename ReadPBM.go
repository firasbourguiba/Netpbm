package netpbm

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type PBM struct {
	data          [][]bool
	width, height int
	magicNumber   string
}

func parsePixel(value string) (bool, error) {
	switch value {
	case "0":
		return false, nil
	case "1":
		return true, nil
	default:
		return false, fmt.Errorf("unexpected pixel value: %s", value)
	}
}

func ReadPBM(filename string) (*PBM, error) {
	//lire le fichier
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	//identifier le magic number
	var magicNumber string
	if !scanner.Scan() {
		return nil, fmt.Errorf("error reading magic number")
	}
	if _, err := fmt.Sscanf(scanner.Text(), "%s", &magicNumber); err != nil {
		return nil, err
	}

	//lire les taille width et height
	if !scanner.Scan() {
		return nil, fmt.Errorf("error reading width and height")
	}

	var w, h int
	if _, err := fmt.Sscanf(scanner.Text(), "%d %d", &w, &h); err != nil {
		return nil, err
	}
	// le nouveau pbm pour l'afficher
	pbm := &PBM{
		magicNumber: magicNumber,
		width:       w,
		height:      h,
	}

	//affiche l'image ( data h w )
	if pbm.magicNumber == "P1" {
		pbm.data = make([][]bool, pbm.height)
		for y := 0; y < pbm.height; y++ {
			if !scanner.Scan() {
				return nil, fmt.Errorf("error reading row %d: %v", y, scanner.Err())
			}
			row := make([]bool, pbm.width)
			// Split the line into individual values
			values := strings.Fields(scanner.Text())
			if len(values) != pbm.width {
				return nil, fmt.Errorf("unexpected number of values in row %d, expected %d, got %d", y, pbm.width, len(values))
			}
			for x, val := range values {
				pixel, err := parsePixel(val)
				if err != nil {
					return nil, fmt.Errorf("error parsing pixel at (%d, %d): %v", x, y, err)
				}
				row[x] = pixel
			}
			pbm.data[y] = row
		}

	} else if pbm.magicNumber == "P4" {
		var data []byte
		data, err = ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}
		pbm.data = make([][]bool, pbm.height)
		for y := 0; y < pbm.height; y++ {
			row := make([]bool, pbm.width)
			for x := 0; x < pbm.width; x++ {
				row[x] = data[y*pbm.width+x] == '1'
			}
			pbm.data[y] = row
		}
	} else {
		return nil, fmt.Errorf("invalid magic number : %s", pbm.magicNumber)
	}

	fmt.Println("Magic Number:", pbm.magicNumber)
	fmt.Println("Width:", pbm.width)
	fmt.Println("Height:", pbm.height)
	fmt.Println("Data:")
	for _, row := range pbm.data {
		for _, pixel := range row {
			if pixel {
				fmt.Print("1 ")
			} else {
				fmt.Print("0 ")
			}
		}
		fmt.Println()

	}
	return pbm, nil
}
