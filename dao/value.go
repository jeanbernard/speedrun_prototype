package dao

import (
	"context"
	"developer/any/dbmodels/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ValueDAO struct {
	DB *gorm.DB
}

func NewValueDAO(db *gorm.DB) *ValueDAO {
	return &ValueDAO{db}
}

func (v *ValueDAO) Create(ctx context.Context, values []models.Value) error {
	// TODO: change this back to Create
	return v.DB.Clauses(clause.OnConflict{DoNothing: true}).Create(values).Error
}
