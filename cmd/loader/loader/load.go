package main

import (
	"developer/any/cmd/loader"
)

func main() {
	if err := loader.LoadGames(); err != nil {
		return
	}

	if err := loader.LoadRecords(); err != nil {
		return
	}
}
