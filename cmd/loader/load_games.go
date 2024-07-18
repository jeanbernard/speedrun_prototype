package loader

import (
	"context"
	"developer/any/clients/speedrun"
	"developer/any/dal"
	database "developer/any/db"

	"github.com/rs/zerolog/log"
)

func LoadGames() error {
	ctx := context.Background()
	db := database.NewSQLiteDatabase()
	dal := dal.NewGameDAL(db.GetDb())

	resp, err := speedrun.GetGamesBulk()
	if err != nil {
		return err
	}

	if err := dal.CreateBulk(ctx, resp); err != nil {
		return err
	}

	cnt := 1
	log.Info().Int("Requests made:", cnt).Msg("requests")
	nextResp := resp
	for nextResp.Pagination.Size >= nextResp.Pagination.Max {
		for _, page := range nextResp.Pagination.Links {
			if page.Rel == "next" {
				nextResp, err = speedrun.GetGamesBulkPagination(page.URI)
				if err != nil {
					return err
				}

				if err := dal.CreateBulk(ctx, nextResp); err != nil {
					return err
				}
				cnt++
				log.Info().Int("Requests made:", cnt).Msg("requests")
			}
		}
	}

	// Used for only loading RE2 for testing
	// resp, err := speedrun.GetGames()
	// if err != nil {
	// 	return err
	// }

	// dal := dal.NewGameDAL(db.GetDb())
	// if err := dal.Create(ctx, resp); err != nil {
	// 	return err
	// }

	return nil
}
