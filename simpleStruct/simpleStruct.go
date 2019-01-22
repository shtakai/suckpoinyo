package simpleStruct

import "fmt"

type Suck struct{
	name string
	killed int
}

func Sick(name string, killed int) Suck {
	suck := Suck{
		name: name,
		killed: killed,
	}
	fmt.Println(suck)
	return suck
}
