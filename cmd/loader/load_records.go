package loader

import (
	"context"
	"developer/any/clients/speedrun"
	"developer/any/dal"
	database "developer/any/db"
	"fmt"
)

func LoadRecords() error {
	ctx := context.Background()
	db := database.NewSQLiteDatabase()
	categoryDAL := dal.NewCategoryDAL(db.GetDb())
	runDAL := dal.NewRunDAL(db.GetDb())

	// we need to call SQLite to get the games
	// so we can load their respective records
	gameDAL := dal.NewGameDAL(db.GetDb())
	games, err := gameDAL.GetAll(ctx)
	if err != nil {
		return err
	}

	// loop through every game
	// using game_id, get records and insert to other tables:
	// runs, variables, categories
	for _, game := range games {
		records, err := speedrun.GetRecords(game)
		if err != nil {
			return err
		}
		fmt.Printf("%v %v\n", game.Name, game.Id)

		for _, record := range records.Data {
			fmt.Println(record)

			// category
			categoryId, err := categoryDAL.Create(ctx, game.Id, record.Category)
			if err != nil {
				return err
			}
			fmt.Println("CATEGORY: ", record.Category.Data.Name)

			// runs
			for _, run := range record.Runs {
				if err := runDAL.Create(ctx, game.Id, categoryId, run.Run); err != nil {
					return err
				}
			}
		}
		fmt.Printf("Completed load for %v \n", game.Name)
	}
	return nil
}
