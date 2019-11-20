package main

import "fmt"

func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

func main() {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())

	j := 1
	someFunc := func() int {
		j++
		return j
	}
	someFunc()
	fmt.Println(j)
	someFunc()
	fmt.Println(j)
}