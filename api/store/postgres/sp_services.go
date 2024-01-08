package postgres

import "github.com/param108/profile/api/models"

func (db *PostgresDB) GetSPServices(writer string) ([]*models.SpService, error) {
	services := []*models.SpService{}
	err := db.db.Where("writer = ?").Order("category ASC").Find(&services).Error
	if err != nil {
		return nil, err
	}

	return services, nil
}
