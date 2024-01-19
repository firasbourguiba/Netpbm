package netpbm

import (
	"fmt"
)

func (pgm *PGM) SetMagicNumberPGM(magicNumber string) {
	pgm.magicNumber = magicNumber
	//c'est tout waawawawaw
	fmt.Print("it work frero pour le pgm")
}
