package maths

import "fmt"

func Par(num int) bool {
	//Verifica si es impar
	if num%2 != 0 {
		return true
	} else {
		fmt.Println(num, "invalid entry")
		return false
	}

}
func Digit(num int) bool{
i:=0
	for num>0{
		num /= 10
		i++
	}
	if i>2{
		return true
	}
	return false
}
