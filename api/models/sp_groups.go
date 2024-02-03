package models

type CreateGroupRequest struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type AddGroupUserRequest struct {
	Phone   string `json:"phone"`
	GroupID string `json:"group_id"`
	Role    string `json:"role"`
}
