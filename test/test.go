package main

import (
	"fmt"

	netpbm "github.com/firasbourguiba/Netpbm1"
)

func main() {

	image, err := netpbm.ReadPBM("test.pbm")
	if err != nil {
		fmt.Println("Error reading PBM file:", err)
		return
	}

	fmt.Println(image)

	//fmt.Println("Magic Number:", image.magicNumber)
	// fmt.Println("Width:", image.width)
	// fmt.Println("Height:", image.height)
	// fmt.Println("Data:")
	// for _, row := range image.data {
	// 	for _, pixel := range row {
	// 		if pixel {
	// 			fmt.Print("1 ")
	// 		} else {
	// 			fmt.Print("0 ")
	// 		}
	// 	}
	// 	fmt.Println()
	// }
}

/*func readPBM(file *os.File) {
	panic("unimplemented")
}*/
