package netpbm

import "fmt"

func (pbm *PBM) Invert() {
	for y := 0; y < pbm.height; y++ {
		for x := 0; x < pbm.width; x++ {
			pbm.data[y][x] = !pbm.data[y][x]
			fmt.Print(pbm.data[y][x],"--")
		}
	}

}
