package dal

import (
	"context"
	"developer/any/dao"
	dbmodels "developer/any/dbmodels/models"
	"developer/any/models"

	"gorm.io/gorm"
)

type Game struct {
	gameDAO *dao.GameDAO
}

func NewGameDAL(db *gorm.DB) *Game {
	return &Game{gameDAO: dao.NewGameDAO(db)}
}

func (dal *Game) Create(ctx context.Context, resp models.GamesResponse) error {
	game := dbmodels.Game{
		Id:   resp.GamesData.Id,
		Name: resp.GamesData.Names["international"],
	}
	return dal.gameDAO.Create(ctx, game)
}

func (dal *Game) CreateBulk(ctx context.Context, resp models.GamesBulkResponse) error {
	var games []dbmodels.Game

	for _, game := range resp.Data {
		games = append(games, dbmodels.Game{
			Id:   game.Id,
			Name: game.Names["international"],
		})
	}

	return dal.gameDAO.CreateBulk(ctx, games)
}

func (dal *Game) GetAll(ctx context.Context) ([]dbmodels.Game, error) {
	return dal.gameDAO.GetAll(ctx)
}
