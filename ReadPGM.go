package netpbm

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PGM struct {
	data          [][]uint8
	width, height int
	magicNumber   string
	max           uint8
}


func ReadPGM(filename string) (*PGM, error) {
	//lire le fichier
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	//identifier le magic number
	var magicNumber string
	if !scanner.Scan() {
		return nil, fmt.Errorf("error pour lire  magic number")
	}
	if _, err := fmt.Sscanf(scanner.Text(), "%s", &magicNumber); err != nil {
		return nil, err
	}
	fmt.Println("Magic Number:", magicNumber)

	//lire les taille width et height
	if !scanner.Scan() {
		return nil, fmt.Errorf("error pour lire  les tailles")
	}

	var w, h int
	if _, err := fmt.Sscanf(scanner.Text(), "%d %d", &w, &h); err != nil {
		return nil, err
	}
	fmt.Println("Width:",w)
	fmt.Println("Height:", h)
	// lire le max de chaques pixels de 0 jusqua 255

if !scanner.Scan() {
    fmt.Println("Error scanning max value")
    return nil, fmt.Errorf("error pour lire  le max")
}
fmt.Println("apres scanner max")


fmt.Println("lavaleur de max:", scanner.Text())
	var max uint8
	if _, err := fmt.Sscanf(scanner.Text(), "%d", &max); err != nil {
		if max < 0 || max > 255 {
			return nil, fmt.Errorf("le max doit être compris entre 0 et 255")
		}
	} 
	// le nouveau pbm pour l'afficher
	pgm := &PGM{
		magicNumber: magicNumber,
		width:       w,
		height:      h,
		max:         max,
	}
	if err := LireData(scanner, pgm); err != nil {
		return nil, err
	}
	fmt.Println("Magic Number:", pgm.magicNumber)
fmt.Println("Width:", pgm.width)
fmt.Println("Height:", pgm.height)
fmt.Println("Max:", pgm.max)
fmt.Println("Data:")
for _, row := range pgm.data {
    for _, pixel := range row {
        fmt.Printf("%d ", pixel)
    }
    fmt.Println()
}
	return pgm, nil
   
	// la partie pour lire la data

}
func LireData(scanner *bufio.Scanner, pgm *PGM) error {
	pgm.data = make([][]uint8, pgm.height)
	for y := 0; y < pgm.height; y++ {
		if !scanner.Scan() {
			return fmt.Errorf("problem pour lire les taille %d: %v", y, scanner.Err())
		}
		//pour la partie P2 ASCii
		if pgm.magicNumber == "P2" {
			
			values := strings.Fields(scanner.Text())
			if len(values) != pgm.width {
				return fmt.Errorf("Nombre inattendu de valeurs dans la ligne %d, %d attendu, %d obtenu.", y, pgm.width, len(values))
			}
			row := make([]uint8, pgm.width)
			for x, val := range values {
				pixel, err := strconv.ParseUint(val, 10, 8)
				if err != nil {
					return fmt.Errorf("un error du ^pixel (%d, %d): %v", x, y, err)
				}
				row[x] = uint8(pixel)
			}
			pgm.data[y] = row
			//pour la partie P5 
			fmt.Println("Data Line:", scanner.Text())

		} else if pgm.magicNumber == "P5" {  
		data := scanner.Bytes()
		row := make([]uint8, pgm.width)
		copy(row, data[:pgm.width])
		pgm.data[y] = row
		} else {
			return fmt.Errorf("un problem de identifier le number magic ni P2 ni P5 : %s", pgm.magicNumber)
		}
	}
	
	return nil
}
