package simpleStruct

import "fmt"

type Suck struct{
	name string
	killed int
}

func Sick() Suck {
	suck := Suck{
		name: "sashimi",
		killed: 4000,
	}
	fmt.Println(suck)
	return suck
}
