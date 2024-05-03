package dal

import (
	"context"
	"developer/any/dao"

	dbmodels "developer/any/dbmodels/models"
	"developer/any/models"

	"gorm.io/gorm"
)

type Run struct {
	RunDAO *dao.RunDAO
}

func NewRunDAL(db *gorm.DB) *Run {
	return &Run{RunDAO: dao.NewRunDAO(db)}
}

func (dal *Run) Create(ctx context.Context, gameId, categoryId string, resp models.Run) (string, error) {
	run := dbmodels.Run{
		Id:         resp.Id,
		GameID:     gameId,
		CategoryID: categoryId,
		Level:      &resp.Level,
		Runtime:    resp.Times.PrimaryTime,
		VideoURI:   resp.Videos.Links[0].URI,
	}

	runId, err := dal.RunDAO.Create(ctx, run)
	if err != nil {
		return runId, err
	}

	return runId, nil
}
