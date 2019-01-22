package simpleStruct

import "fmt"

type Suck struct{
	name string
	killed int
}

func Sick(suck Suck) Suck {
	fmt.Println(suck)
	return suck
}
