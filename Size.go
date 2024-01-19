package netpbm
import	(
	"fmt"

)
func (pbm *PBM) SizePGM() (int,int){
	fmt.Println("Magic Number:", pbm.magicNumber)
	fmt.Println("Width:", pbm.width)
	return pbm.height, pbm.width
}