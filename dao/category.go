package dao

import (
	"context"
	"developer/any/dbmodels/models"

	"gorm.io/gorm"
)

type CategoryDAO struct {
	DB *gorm.DB
}

func NewCategoryDAO(db *gorm.DB) *CategoryDAO {
	return &CategoryDAO{db}
}

func (c *CategoryDAO) Create(ctx context.Context, category models.Category) (string, error) {
	//TODO: change back to just create
	if err := c.DB.FirstOrCreate(&category).Error; err != nil {
		return "", err
	}

	return category.Id, nil
}
