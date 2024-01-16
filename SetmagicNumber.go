package netpbm

import (
	"fmt"
)

func (pbm *PBM) SetMagicNumber(magicNumber string) {
	pbm.magicNumber = magicNumber
	//c'est tout waawawawaw
	fmt.Print("it work frero")
}
