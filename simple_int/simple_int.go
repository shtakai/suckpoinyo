package simpleInt

import (
	"fmt"
	"os"
)

func SimpleInt(i int) int {
	return suck(&i) * 10
}

func suck(n *int) int {
	_, _ = fmt.Fprintf(os.Stdout, "n: %v   *n: %v \n", n, *n)
	return *n
}
