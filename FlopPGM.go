package netpbm
func (pgm *PGM) FlopPGM(){
	
        for i := 0; i < pgm.height/2; i++ {
            top := i * pgm.width
            bottom := (pgm.height - 1 - i) * pgm.width
			for j := 0; j < pgm.width; j++ {
                pgm.data[top+j], pgm.data[bottom+j] = pgm.data[bottom+j], pgm.data[top+j]
            }
		}
}