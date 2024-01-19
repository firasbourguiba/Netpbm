package netpbm

import "fmt"

func (pgm *PGM) InvertPgm() {
	for y := 0; y < pgm.height; y++ {
		for x := 0; x < pgm.width; x++ {
			pgm.data[y][x] = pgm.max - pgm.data[y][x]
			fmt.Print(pgm.data[y][x],"--")
		}
	}

}
