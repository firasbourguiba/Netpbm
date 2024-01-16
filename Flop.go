package netpbm
import (
	"fmt"
)
func (pbm *PBM) Flop(){
	for y := 0; y < pbm.width; y++ {

		for i, j := 0, pbm.width-1; i < j; i, j = i+1, j-1 {
			pbm.data[i][y], pbm.data[j][y] = pbm.data[j][y], pbm.data[i][y]
			fmt.Print(pbm.data[y][i], "--")
		}
	}
}