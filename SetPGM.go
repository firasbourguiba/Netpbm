package netpbm

func(pgm *PGM) SetPGM(x,y int, value uint8){
	//verifier si le coordonner donner appartient a la intervalle 
	if x < 0 || x >= pgm.width ||y <0 || y >= pgm.height {
		return
	}
	// change le valeur
	pgm.data[y][x]=value
}