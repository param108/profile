package postgres

import (
	"github.com/param108/profile/api/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (db *PostgresDB) LockResource(req *models.Resource) (*models.Resource, error) {
	err := db.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "writer"}, {Name: "t"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"value": gorm.Expr(`resources.value + 1`)}),
	}).Create(req).Error

	if err != nil {
		return nil, err
	}

	res := []*models.Resource{}
	err = db.db.Where("user_id = ? and writer = ? and t = ?",
		req.UserID, req.Writer, req.T).Find(&res).Error
	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, errors.New("not found")
	}

	return res[0], nil
}

func (db *PostgresDB) UnlockResource(req *models.Resource) (*models.Resource, error) {
	err := db.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "writer"}, {Name: "t"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"value": gorm.Expr("resources.value - 1")}),
	}).Create(req).Error

	if err != nil {
		return nil, err
	}

	res := []*models.Resource{}
	err = db.db.Where("user_id = ? and writer = ? and t = ?",
		req.UserID, req.Writer, req.T).Find(&res).Error
	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, errors.New("not found")
	}

	return res[0], nil
}

func (db *PostgresDB) GetResources(userID, writer string) ([]*models.Resource, error) {
	resources := []*models.Resource{}
	if err := db.db.Where("user_id = ? and writer = ?", userID, writer).Find(&resources).Error; err != nil {
		return nil, err
	}

	return resources, nil
}
