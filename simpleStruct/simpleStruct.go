package simpleStruct

import "fmt"

type Suck struct{
	name string
	killed int
}

func Sick(name string) Suck {
	suck := Suck{
		name: name,
		killed: 40000,
	}
	fmt.Println(suck)
	return suck
}
