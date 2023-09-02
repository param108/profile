package store

import "github.com/param108/profile/api/models"

var DefaultMaxResources = map[string]int{
	"image": 10,
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
