package simpleStruct

import (
	"fmt"
	"strings"
)

type Suck struct{
	name string
	killed int
}

func Sick(suck *Suck) Suck {
	fmt.Println(&suck)
    moron(suck)
	fmt.Println(&suck)
	return *suck
}

func moron(suck *Suck) Suck {
	suck.killed *= 2
	suck.name = strings.ToUpper(suck.name)
	fmt.Println(&suck)
	return *suck
}

func SwapSick(suck1, suck2 Suck) (Suck, Suck){
	return suck2, suck1
}
