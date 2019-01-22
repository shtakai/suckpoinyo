package simpleStruct

import (
	"fmt"
	"strings"
)

type Suck struct{
	name string
	killed int
}

func Sick(suck Suck) Suck {
	fmt.Println(suck)
	suck.killed *= 2
	suck.name = strings.ToUpper(suck.name)
	fmt.Println(suck)
	return suck
}
