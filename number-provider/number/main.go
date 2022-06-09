package main

type Number struct {
	Number number
}

var numbers map[string]*Number

func main() {
	numbers = make(map[string]*Number)
}
