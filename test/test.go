package main

import (
	"fmt"

	netpbm "github.com/firasbourguiba/Netpbm1"
	//netpbmpgm "github.com/firasbourguiba/Netpbm/PGM"
)

func main() {

	/*image, err := netpbm.ReadPBM("test.pbm")
	if err != nil {
		fmt.Println("Error reading PBM file:", err)

		return
	}
	fmt.Println("Magic sfzfsfsfsfNumber:")
	fmt.Println("Width:sfsfsfsff")


	fmt.Println(image)
	image.AT(5, 8)
	fmt.Print("teste save")
	image.Save("test.pbm")
	fmt.Print("teste Invert \n")
	image.Invert()
	fmt.Print(" \n")
	fmt.Print("teste filp \n")
	image.Flip()
	fmt.Print(" \n")
	fmt.Print("teste flop \n")
	image.Flop()*/
	PPGM, err := netpbm.ReadPGM("test.pgm")
	if err != nil {
		fmt.Println("Error reading PGM file:", err)

		return
	}
	fmt.Println(PPGM)
	
}
