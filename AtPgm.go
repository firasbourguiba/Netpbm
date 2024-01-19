package netpbm

func (pgm *PGM) At(x, y int) uint8 {
	//affiche  d'apres les coordonner choisir par l'utilisateur
	if x < 0 || x >= pgm.width || y < 0 || y >= pgm.height {

		return 0
	}

	return pgm.data[y][x]
}
