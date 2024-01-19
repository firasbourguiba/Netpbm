package netpbm

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func (pgm *PGM) Save(filename string) error {
	//on va donner un nom pour chaque sdave pour eviter d'avoir des fichier de meme nom par le datede leur naissance
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	newFilename := fmt.Sprintf("%s_%s", strings.TrimSuffix(filename, ".pgm"), timestamp)

	//le save
	file, err := os.Create(newFilename)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	fmt.Fprintf(file, "%s\n%d %d\n", pgm.magicNumber, pgm.width, pgm.height)

	if pgm.magicNumber == "P2" {
		// ASCII format (P2)
		for _, row := range pgm.data {
			for _, pixel := range row {
				_, err := fmt.Fprintf(file, "%d ", pixel)
				if err != nil {
					return err
				}
			}
			_, err := fmt.Fprintln(file)
			if err != nil {
				return err
			}
		}
	} else if pgm.magicNumber == "P5" {
	// Binary format (P5)
	for _, row := range pgm.data {
		_, err := file.Write(row)
		if err != nil {
			return err
		}
	}
	  } else {
			// faut 0 
			if _, err := file.Write([]byte{0}); err != nil {
					return err
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

