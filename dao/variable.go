package dao

import (
	"context"
	"developer/any/dbmodels/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type VariableDAO struct {
	DB *gorm.DB
}

func NewVariableDAL(db *gorm.DB) *VariableDAO {
	return &VariableDAO{db}
}

func (v *VariableDAO) Create(ctx context.Context, variable models.Variable) error {
	// TODO: change this back to Create
	return v.DB.Clauses(clause.OnConflict{DoNothing: true}).Create(variable).Error
}
