package main

import (
	"fmt"

	processor "github.com/open-korean-text/open-korean-text-go/processor"
)

func main() {
	result := processor.Normalize("만듀 먹것니? 먹겄서? 먹즤?")
	fmt.Println(result)
}
