package netpbm

import (
	"fmt"
)

func (pbm *PBM) Set(x, y int, value bool) {
	if y >= 0 && x < pbm.height && x >= 0 && x < pbm.width {
		fmt.Print("ton new set   ", value)

		pbm.data[y][x] = value
		return
	}

}
