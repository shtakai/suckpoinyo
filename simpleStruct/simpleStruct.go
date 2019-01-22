package simpleStruct

import "fmt"

type Suck struct{
	name string
	killed int
}

func Sick() {
	suck := Suck{
		name: "sashimi",
		killed: 4000,
	}
	fmt.Println(suck)
}
