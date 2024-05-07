package dao

import (
	"context"
	"developer/any/dbmodels/models"

	"gorm.io/gorm"
)

type GameDAO struct {
	DB *gorm.DB
}

func NewGameDAO(db *gorm.DB) *GameDAO {
	return &GameDAO{db}
}

func (g *GameDAO) Create(ctx context.Context, game models.Game) error {
	if err := g.DB.FirstOrCreate(&game).Error; err != nil {
		return err
	}
	return nil
}

func (g *GameDAO) CreateBulk(ctx context.Context, game []models.Game) error {
	return g.DB.Create(&game).Error
}

func (g *GameDAO) GetAll(ctx context.Context) ([]models.Game, error) {
	var games []models.Game
	err := g.DB.Find(&games).Error
	return games, err
}
