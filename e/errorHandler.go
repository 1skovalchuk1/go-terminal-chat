package e

import "fmt"

func Print(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
func Panic(err error) {
	if err != nil {
		panic(err)
	}
}
