package loader

import (
	"context"
	"developer/any/clients/speedrun"
	"developer/any/dal"
	database "developer/any/db"
)

func LoadGames() error {
	ctx := context.Background()
	db := database.NewSQLiteDatabase()

	// resp, err := speedrun.GetGamesBulk()
	// if err != nil {
	// 	return err
	// }

	// dal := dal.NewGameDAL(db.GetDb())
	// if err := dal.CreateBulk(ctx, resp); err != nil {
	// 	return err
	// }

	resp, err := speedrun.GetGames()
	if err != nil {
		return err
	}

	dal := dal.NewGameDAL(db.GetDb())
	if err := dal.Create(ctx, resp); err != nil {
		return err
	}

	return nil
}
