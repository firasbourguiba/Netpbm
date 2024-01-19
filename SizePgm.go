package netpbm
import	(
	"fmt"

)
func (pgm *PGM) Size() (int,int){
	fmt.Println("Magic Number:", pgm.magicNumber)
	fmt.Println("Width:", pgm.width)
	return pgm.height, pgm.width
}