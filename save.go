package netpbm

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func (pbm *PBM) Save(filename string) error {
	//on va donner un nom pour chaque sdave pour eviter d'avoir des fichier de meme nom par le datede leur naissance
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	newFilename := fmt.Sprintf("%s_%s", strings.TrimSuffix(filename, ".pbm"), timestamp)

	//le save
	file, err := os.Create(newFilename)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	fmt.Fprintf(file, "%s\n%d %d\n", pbm.magicNumber, pbm.width, pbm.height)
	if pbm.magicNumber == "P1" {
		for y := 0; y < pbm.height; y++ {
			for x := 0; x < pbm.width; x++ {
				// si true 1 false 0 
				if pbm.data[y][x] {
					fmt.Fprint(file, "1 ")
				} else {
					fmt.Fprint(file, "0 ")
				}
			}
			// ligne par ligne
			fmt.Fprintln(file)
		}
	} else if pbm.magicNumber == "P4" {
		// Write the binary data directly without spaces
		for y := 0; y < pbm.height; y++ {
			for x := 0; x < pbm.width; x++ {
				if pbm.data[y][x] {
					//vrai 1
					if _, err := file.Write([]byte{1}); err != nil {
						return err
					}
				} else {
					// faut 0 
					if _, err := file.Write([]byte{0}); err != nil {
						return err
					}
				}
			}
		}
	}
	if err != nil {
		return err
	}
	err = writer.Flush()
	if err != nil {
		return err
	}

	fmt.Printf("Fichier enregistré avec le nom: %s\n", newFilename)

	return nil
}
