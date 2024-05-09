package main

import (
	"context"
	"developer/any/cmd/loader"
)

func main() {
	ctx := context.Background()
	// if err := loader.LoadGames(ctx); err != nil {
	// 	return
	// }

	if err := loader.LoadRecords(ctx); err != nil {
		return
	}
}
