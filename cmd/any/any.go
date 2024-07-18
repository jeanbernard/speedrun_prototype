package main

import (
	"context"
	"developer/any/dal"
	database "developer/any/db"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	var buildTime float64

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.Float64Flag{
				Name:        "build",
				Value:       0,
				Usage:       "your build time",
				Destination: &buildTime,
			},
		},
		Action: func(cCtx *cli.Context) error {
			ctx := context.Background()
			err := serveRun(ctx, buildTime)
			if err != nil {
				return err
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func serveRun(ctx context.Context, buildTime float64) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	db := database.NewSQLiteDatabase()

	runDAL := dal.NewRunDAL(db.GetDb())
	run, err := runDAL.GetRandomRun(ctx, buildTime)
	if err != nil {
		return err
	}

	fmt.Println(run)

	return nil
}
