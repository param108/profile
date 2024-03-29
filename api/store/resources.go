package store

import "github.com/param108/profile/api/models"

const ImageType = "image"

var DefaultMaxResources = map[string]int{
	ImageType: 10000000, // 10 Megabytes
}

var DEFAULT_MAX_RESOURCE = 10

// getMaxResources Max resources for a type
func getMaxResources(t string) int {
	val, ok := DefaultMaxResources[t]
	if ok {
		return val
	}

	return DEFAULT_MAX_RESOURCE
}

// LockResource incr a resource count
func (s *StoreImpl) LockResource(userID, t, writer string) (*models.Resource, error) {
	req := models.Resource{
		UserID: userID,
		T:      t,
		Writer: writer,
		Value:  1,
		Max:    getMaxResources(t),
	}
	return s.db.LockResource(&req)
}

// UnlockResource decr a resource count
func (s *StoreImpl) UnlockResource(userID, t, writer string) (*models.Resource, error) {
	req := models.Resource{
		UserID: userID,
		T:      t,
		Writer: writer,
		Value:  0,
		Max:    getMaxResources(t),
	}
	return s.db.UnlockResource(&req)
}

// GetResources get resource count
func (s *StoreImpl) GetResources(userID, writer string) ([]*models.Resource, error) {
	return s.db.GetResources(userID, writer)
}

func (s *StoreImpl) SetResources(
	userID, t string, value int, writer string) (*models.Resource, error) {
	req := models.Resource{
		UserID: userID,
		T:      t,
		Writer: writer,
		Value:  value,
		Max:    getMaxResources(t),
	}

	return s.db.SetResources(&req, value, writer)
}
