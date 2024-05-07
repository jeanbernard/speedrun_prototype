package main

import (
	"developer/any/cmd/loader"
	"fmt"
)

func main() {
	if err := loader.LoadGames(); err != nil {
		fmt.Println(err.Error())
		return
	}

	if err := loader.LoadRecords(); err != nil {
		fmt.Println(err.Error())
		return
	}
}
