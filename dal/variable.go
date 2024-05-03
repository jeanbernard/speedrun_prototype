package dal

import (
	"context"
	"developer/any/dao"

	dbmodels "developer/any/dbmodels/models"
	"developer/any/models"

	"gorm.io/gorm"
)

type Variable struct {
	variableDAO *dao.VariableDAO
}

func NewVariableDAL(db *gorm.DB) *Variable {
	return &Variable{variableDAO: dao.NewVariableDAL(db)}
}

func (dal *Variable) Create(ctx context.Context, resp models.VariableData) error {
	variable := dbmodels.Variable{
		Id:   resp.Id,
		Name: resp.Name,
	}
	return dal.variableDAO.Create(ctx, variable)
}
