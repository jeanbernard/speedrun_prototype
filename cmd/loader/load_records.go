package loader

import (
	"context"
	"developer/any/clients/speedrun"
	"developer/any/dal"
	database "developer/any/db"
	"developer/any/models"
	"fmt"
)

func LoadRecords() error {
	ctx := context.Background()
	db := database.NewSQLiteDatabase()
	categoryDAL := dal.NewCategoryDAL(db.GetDb())
	runDAL := dal.NewRunDAL(db.GetDb())
	variableDAL := dal.NewVariableDAL(db.GetDb())
	valueDAL := dal.NewValueDAL(db.GetDb())

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
			// category
			categoryId, err := categoryDAL.Create(ctx, game.Id, record.Category)
			if err != nil {
				return err
			}
			fmt.Println("CATEGORY: ", record.Category.Data.Name)

			//fmt.Printf("DATA!: %v\n", record.Variables.Data)
			// variables & values

			vars := map[string]models.VariableData{}

			if len(record.Variables.Data) != 0 {
				for _, variable := range record.Variables.Data {

					vars[variable.Id] = variable

					// variables
					if err := variableDAL.Create(ctx, variable); err != nil {
						return err
					}

					// values
					if err := valueDAL.Create(ctx, variable.Values); err != nil {
						return err
					}
				}
			} else {
				fmt.Printf("Game: %v doesn't have variables \n", game.Name)
			}

			// create variable/value mapping
			for k, v := range record.Runs[0].Run.Values {
				fmt.Println(vars[k].Name, vars[k].Values.Values[v.(string)].Label)
			}

			fmt.Println(record.Runs[0].Run.Times.PrimaryTime)
			fmt.Println("=================")

			// Runs
			// fmt.Println(runId)
			// runs - only bringing back ONE run for now
			// there could a tie/s between run/s but let's focus on one
			_, err = runDAL.Create(ctx, game.Id, categoryId, record.Runs[0].Run)
			if err != nil {
				return err
			}

		}
	}

	return nil
}
