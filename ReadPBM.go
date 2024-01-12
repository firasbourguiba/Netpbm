package netpbm

import (
	"fmt"
	"io/ioutil"
	"os"
)

type PBM struct {
	data          [][]bool
	width, height int
	magicNumber   string
}

func ReadPBM(filename string) (*PBM, error) {
	//lire le ficheier
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	//identifier le magicnumber
	var magicNumber string
	if _, err = fmt.Fscan(file, &magicNumber); err != nil {
		return nil, err
	}
	//lire les taille width et height
	var w, h int
	if _, err = fmt.Fscanf(file, "%d %d", &w, &h); err != nil {
		return nil, err
	}
	// le nouveau pbm pour l'afficher
	pbm := &PBM{
		magicNumber: magicNumber,
		width:       w,
		height:      h,
	}
	//affiche l'image ( data h w )
	if pbm.magicNumber == "P1" {
		pbm.data = make([][]bool, pbm.height)
		for y := 0; y < pbm.height; y++ {
			row := make([]bool, pbm.width)
			for x := 0; x < pbm.width; x++ {
				if _, err = fmt.Fscan(file, &row); err != nil {
					return nil, err
				}
				pbm.data[y] = row
			}
		}
	} else if pbm.magicNumber == "P4" {
		var data []byte
		if data, err = ioutil.ReadAll(file); err != nil {
			return nil, err
		}
		pbm.data = make([][]bool, pbm.height)
		for y := 0; y < pbm.height; y++ {
			row := make([]bool, pbm.width)
			for x := 0; x < pbm.width; x++ {

				row[x] = data[y*pbm.width+x] == '1'
			}
			pbm.data[y] = row
		}
	} else {
		return nil, fmt.Errorf("invalid magic number : %s", pbm.magicNumber)
	}
	return pbm, nil

}
