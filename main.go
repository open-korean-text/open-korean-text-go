package main

import (
	"fmt"

	processor "openkoreantext.org/processor"
)

func main() {
	result := processor.Normalize("만듀 먹것니? 먹겄서? 먹즤?")
	fmt.Println(result)
}
