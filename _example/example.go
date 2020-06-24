package main

import (
	"fmt"

	"github.com/terryliu/maths"
)
func main() {
	s :=maths.NewSwatch()
	s.PutList([]float64{47.42,46.79,49,49.82,48.01,49.01,50.23,50,50.25,49.36,51.16,50.32,50,51.88,51.93,50.86,51.31,50.45,51.43,51.73,50.93,50.7,49.8,50.37})
	fmt.Printf("M24:\t%+v\n",s.Ma())
	fmt.Printf("Upper:\t%+v\n",s.Upper())
	fmt.Printf("Lower:\t%+v\n",s.Lower())


	/**
	M24:    50.115
	Upper:  52.78775464480081
	Lower:  47.44224535519919

	*/
}



