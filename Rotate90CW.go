package netpbm
func (pgm *PGM) Rotate90CW(){
   //inversee les lignes et les colones  pour 90wc
	pgm.width, pgm.height = pgm.height, pgm.height
	//criée un nouveau data 
	newdata := make([][]uint8, pgm.width)
	for i:= range newdata{
		newdata[i] = make([]uint8, pgm.height)
	}
	//remplir 
	for y:=0;y<pgm.height;y++ {
		for x:=0;x<pgm.width;x++ {
			newdata[x][pgm.width-1-y] = pgm.data[y][x]
		}
	}
	pgm.data = newdata
	//verifier
	println(pgm.data)
}