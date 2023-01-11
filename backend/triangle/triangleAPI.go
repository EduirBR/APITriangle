package triangle

func TriangleBuildAPI(high int) []string {

	line2 := ""
	TringleArray := []string{}
	//construcción del triangulo
	min := (high / 2) //delimita el lado izquierdo
	max := (high / 2) //delimita el lado derecho
	for i := 0; i < (high/2)+1; i++ {
		for j := 0; j < high; j++ {
			if j < min || j > max {

				line2 += " "
			} else {

				line2 += "*"
			}
			if j == high-1 { //ejecuta al llegar al final de la linea

				TringleArray = append(TringleArray, line2)
				line2 = ""
				min-- //reduce los espacios por la izquierda
				max++ //reduce los espacios por la derecha
			}
		}
	}

	return TringleArray
}
