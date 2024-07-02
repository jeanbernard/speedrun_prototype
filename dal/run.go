package dal

import (
	"context"
	"developer/any/dao"
	"errors"

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

// TODO: Deal with multiple video URIs and Player IDs
// Also deal with some videos being "Text" only. lol.
func (dal *Run) Create(ctx context.Context, gameId, categoryId string, resp models.Run) error {
	if resp.Videos == nil {
		return errors.New("does not contain videos")
	}

	if len(resp.Videos.Links) == 0 {
		return errors.New("does not contain links for videos")
	}

	run := dbmodels.Run{
		Id:         resp.Id,
		GameID:     gameId,
		CategoryID: categoryId,
		Level:      &resp.Level,
		Runtime:    resp.Times.PrimaryTime,
		VideoURI:   resp.Videos.Links[0].URI,
		Values:     resp.Values,
	}

	if err := dal.RunDAO.Create(ctx, run); err != nil {
		return err
	}
	return nil
}

func (dal *Run) GetAll(ctx context.Context) ([]dbmodels.Run, error) {
	return dal.RunDAO.GetAll(ctx)
}
