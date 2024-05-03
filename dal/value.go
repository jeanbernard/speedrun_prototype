package dal

import (
	"context"
	"developer/any/dao"
	dbmodels "developer/any/dbmodels/models"
	"developer/any/models"

	"gorm.io/gorm"
)

type Value struct {
	valueDAO *dao.ValueDAO
}

func NewValueDAL(db *gorm.DB) *Value {
	return &Value{valueDAO: dao.NewValueDAO(db)}
}

func (dal *Value) Create(ctx context.Context, resp models.ValuesContainer) error {
	var values []dbmodels.Value

	for id, val := range resp.Values {
		value := dbmodels.Value{
			Id:    id,
			Label: val.Label,
		}
		values = append(values, value)
	}

	return dal.valueDAO.Create(ctx, values)
}
