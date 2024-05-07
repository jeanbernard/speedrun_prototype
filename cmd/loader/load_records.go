package loader

import (
	"context"
	"developer/any/clients/speedrun"
	"developer/any/dal"
	database "developer/any/db"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func LoadRecords() error {
	ctx := context.Background()
	db := database.NewSQLiteDatabase()
	categoryDAL := dal.NewCategoryDAL(db.GetDb())
	runDAL := dal.NewRunDAL(db.GetDb())

	// logger
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

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
		log.Info().Str("gameId:", game.Id).Str("game: ", game.Name).Msg("loading game...")
		records, err := speedrun.GetRecords(game)
		if err != nil {
			log.Error().AnErr("speedrun.GetRecords ", err).Msg("error")
			return err
		}

		for _, record := range records.Data {
			// category
			categoryId, err := categoryDAL.Create(ctx, game.Id, record.Category)
			if err != nil {
				return err
			}
			log.Info().Str("categoryId", record.Category.Data.Id).Str("category:", record.Category.Data.Name).Msg("loaded category")

			// runs
			for _, run := range record.Runs {
				if err := runDAL.Create(ctx, game.Id, categoryId, run.Run); err != nil {
					return err
				}
			}
		}
		log.Info().Str("gameId:", game.Id).Str("game:", game.Name).Msg("loaded game!")
	}
	return nil
}
