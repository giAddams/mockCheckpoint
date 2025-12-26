package main

import (
	"fmt"
	piscine "mockCheckpoint/10/countalpha/piscine"
)

func main() {
	fmt.Println(piscine.CountAlpha("Hello world"))
	fmt.Println(piscine.CountAlpha("H e l l o"))
	fmt.Println(piscine.CountAlpha("H1e2l3l4o"))
}
