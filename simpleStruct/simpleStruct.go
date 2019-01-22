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
	swap(&suck1, &suck2)
	return suck1, suck2
}

func swap(s1, s2 *Suck){
    tmpName, tmpKilled := s2.name, s2.killed
    s2.name = s1.name
    s2.killed = s1.killed
    s1.name = tmpName
    s1.killed = tmpKilled

}
