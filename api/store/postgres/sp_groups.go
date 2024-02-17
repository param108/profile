package postgres

import (
	"errors"

	"github.com/param108/profile/api/models"
)

func (db *PostgresDB) AddSPGroup(group *models.SpGroup, writer string) (*models.SpGroup, error) {
	group.Writer = writer
	err := db.db.Create(group).Error
	if err != nil {
		return nil, err
	}

	return group, nil
}

func (db *PostgresDB) GetSPGroupsForUser(userID, writer string) ([]*models.SpGroupSend, error) {
	groups := []*models.SpGroupSend{}
	err := db.db.Model(&models.SpGroupUser{}).Where(
		"sp_group_users.sp_user_id = ? and sp_group_users.writer = ? and sp_group_users.deleted = false",
		userID, writer).Joins(
		"left join sp_groups on sp_group_users.sp_group_id = sp_groups.id",
	).Joins(
		"left join sp_users on sp_group_users.sp_user_id = sp_users.id",
	).Select("sp_groups.id as id, " +
		" sp_groups.name as name," +
		" sp_groups.description as description," +
		" sp_groups.created_at as created_at," +
		" sp_groups.updated_at as updated_at," +
		" sp_groups.writer as writer," +
		" sp_group_users.sp_user_id as user_id," +
		" sp_users.name as username," +
		" sp_group_users.role as role",
	).Scan(&groups).Error
	if err != nil {
		return nil, err
	}

	return groups, nil
}

func (db *PostgresDB) AddSPUserToGroup(
	spGroupUser *models.SpGroupUser, writer string) (*models.SpGroupUser, error) {
	spGroupUser.Writer = writer

	// only two options.
	if spGroupUser.Role != "admin" {
		spGroupUser.Role = "user"
	}

	err := db.db.Create(spGroupUser).Error
	if err != nil {
		return nil, err
	}

	return spGroupUser, nil
}

func (db *PostgresDB) GetSPGroupUser(
	userID, groupID, writer string) (*models.SpGroupUser, error) {

	grpUsers := []*models.SpGroupUser{}
	if err := db.db.Where("sp_user_id = ? and sp_group_id = ? and writer = ?",
		userID, groupID, writer).Find(&grpUsers).Error; err != nil {
		return nil, err
	}

	if len(grpUsers) == 0 {
		return nil, errors.New("not found")
	}

	return grpUsers[0], nil
}

func (db *PostgresDB) GetSPGroupUsers(
	groupID, writer string) ([]*models.SpUser, error) {
	users := []*models.SpUser{}

	if err := db.db.Where(
		"sp_group_users.sp_group_id = ? and sp_group_users.writer = ? and sp_group_users.deleted = false",
		groupID, writer).Joins(
		"left join sp_users on sp_users.id = sp_group_user.sp_user_id").Select(
		"sp_users.id as id," +
			"sp_users.name as name," +
			"sp_users.photo_url as photo_url," +
			"sp_users.created_at as created_at," +
			"sp_users.writer as writer," +
			"sp_users.profile_updated as profile_updated").Find(
		&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
