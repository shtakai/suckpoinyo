package simpleInt

func SimpleInt(i int) int {
	suck(i)
	return i
}

func suck(n int) {
	//_, _ = fmt.Fprintf(os.Stdout, "n: %v   *n: %v \n", n, *n)
	n = n * 10
}
