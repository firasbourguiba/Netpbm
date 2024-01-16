package netpbm

import "fmt"

func (pbm *PBM) AT(x, y int) bool {
	if y >= 0 && y < pbm.height && x >= 0 && x < pbm.width {
		fmt.Print("ton chois  ", pbm.data[y][x])
		return pbm.data[y][x]
	}
	return false
}
