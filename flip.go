package netpbm

import "fmt"

func (pbm *PBM) Flip() {
	for y := 0; y < pbm.height; y++ {

		for i, j := 0, pbm.width-1; i < j; i, j = i+1, j-1 {
			pbm.data[y][i], pbm.data[y][j] = pbm.data[y][j], pbm.data[y][i]
			fmt.Print(pbm.data[y][i], "--")
		}
	}
}
