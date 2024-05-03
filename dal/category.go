package dal

import (
	"context"
	"developer/any/dao"
	dbmodels "developer/any/dbmodels/models"
	"developer/any/models"

	"gorm.io/gorm"
)

type Category struct {
	categoryDAO *dao.CategoryDAO
}

func NewCategoryDAL(db *gorm.DB) *Category {
	return &Category{categoryDAO: dao.NewCategoryDAO(db)}
}

func (dal *Category) Create(ctx context.Context, gameId string, resp models.Category) (string, error) {
	category := dbmodels.Category{
		GameID: gameId,
		Id:     resp.Data.Id,
		Name:   resp.Data.Name,
		Type:   resp.Data.Type,
	}

	categoryId, err := dal.categoryDAO.Create(ctx, category)
	if err != nil {
		return categoryId, err
	}

	return categoryId, nil
}
