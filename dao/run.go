package dao

import (
	"context"
	dbmodels "developer/any/dbmodels/models"

	"gorm.io/gorm"
)

type RunDAO struct {
	DB *gorm.DB
}

func NewRunDAO(db *gorm.DB) *RunDAO {
	return &RunDAO{db}
}

func (r *RunDAO) Create(ctx context.Context, run dbmodels.Run) (string, error) {
	//TODO: change back to just create
	if err := r.DB.FirstOrCreate(&run).Error; err != nil {
		return "", err
	}

	return run.Id, nil
}

func (r *RunDAO) GetAll(ctx context.Context) ([]dbmodels.Run, error) {
	var runs []dbmodels.Run
	err := r.DB.Find(&runs).Error
	return runs, err
}

func (r *RunDAO) GetAllVariables(ctx context.Context) ([]dbmodels.Run, error) {
	var runs []dbmodels.Run
	err := r.DB.Model(&dbmodels.Run{}).Preload("Variables").Find(&runs).Error
	return runs, err
}
