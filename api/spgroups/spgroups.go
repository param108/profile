package spgroups

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/param108/profile/api/models"
	"github.com/param108/profile/api/store"
	"github.com/param108/profile/api/utils"
)

// CreatePostSPGroupHandler create a group
func CreatePostSPGroupHandler(db store.Store) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		// First check if the user is an admin
		userID := r.Header.Get("SP_USERID")
		if len(userID) == 0 {
			utils.WriteError(rw, http.StatusForbidden, "forbidden")
			return
		}

		writer := os.Getenv("WRITER")
		user, err := db.GetSPUserByID(userID, writer)
		if err != nil {
			utils.WriteError(rw, http.StatusBadRequest, "unknown user")
			return
		}

		if user.Role != "admin" {
			utils.WriteError(rw, http.StatusForbidden, "forbidden")
			return
		}

		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			utils.WriteError(rw, http.StatusBadRequest, "couldnt read:"+err.Error())
			return
		}

		req := &models.CreateGroupRequest{}
		err = json.Unmarshal(data, req)
		if err != nil {
			utils.WriteError(rw, http.StatusBadRequest, "couldnt unmarshall:"+err.Error())
			return
		}

		if len(req.Name) < 5 || len(req.Desc) < 5 {
			utils.WriteError(rw, http.StatusBadRequest, "invalid data")
			return
		}

		group := &models.SpGroup{
			Name:        req.Name,
			Description: req.Desc,
			Writer:      writer,
		}

		group, err = db.AddSPGroup(group, writer)
		if err != nil {
			log.Print(fmt.Sprintf("failed to add group: %s", err.Error()))
			utils.WriteError(rw, http.StatusInternalServerError, "failed")
			return
		}

		// Now add creator as admin
		grpUser := &models.SpGroupUser{
			SpUserID:  userID,
			SpGroupID: group.ID,
			Role:      "admin",
			Writer:    writer,
		}

		_, err = db.AddSPUserToGroup(grpUser, writer)
		if err != nil {
			// FIXME remember to clean up
			utils.WriteError(rw, http.StatusInternalServerError, "failed add user")
			log.Print(fmt.Sprintf("failed to add user to new group: %s", err.Error()))
			return
		}
		utils.WriteData(rw, http.StatusOK, group)
	}
}

func CreateGetSPGroupsHandler(db store.Store) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		// First check if the user is an admin
		userID := r.Header.Get("SP_USERID")
		if len(userID) == 0 {
			utils.WriteError(rw, http.StatusForbidden, "forbidden")
			return
		}

		writer := os.Getenv("WRITER")

		groups, err := db.GetSPGroupsForUser(userID, writer)
		if err != nil {
			utils.WriteError(rw, http.StatusBadRequest, "bad input")
			return
		}

		utils.WriteData(rw, http.StatusOK, groups)
	}
}

func CreateAddGroupUserHandler(db store.Store) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		// First check if the user is an admin
		userID := r.Header.Get("SP_USERID")
		if len(userID) == 0 {
			utils.WriteError(rw, http.StatusForbidden, "forbidden")
			return
		}

		writer := os.Getenv("WRITER")

		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			utils.WriteError(rw, http.StatusBadRequest, "couldnt read:"+err.Error())
			return
		}

		req := &models.AddGroupUserRequest{}
		err = json.Unmarshal(data, req)
		if err != nil {
			utils.WriteError(rw, http.StatusBadRequest, "couldnt unmarshall:"+err.Error())
			return
		}

		if len(req.Phone) != 10 || len(req.GroupID) == 0 || len(req.Role) == 0 {
			utils.WriteError(rw, http.StatusBadRequest, "missing mandatory fields")
			return
		}

		// check if the user is a superadmin

		reqUser, err := db.GetSPUserByID(userID, writer)
		if err != nil {
			utils.WriteError(rw, http.StatusBadRequest, "unknown user")
		}

		if reqUser.Role != "admin" {
			// check if the requester is admin of the group
			grpUser, err := db.GetSPGroupUser(userID, req.GroupID, writer)
			if err != nil {
				log.Printf("groupuser not found:%s", err.Error())
				utils.WriteError(rw, http.StatusBadRequest, "invalid input")
				return
			}

			if grpUser.Role != "admin" {
				utils.WriteError(rw, http.StatusForbidden, "forbidden")
				return
			}
		}

		// get userid for phone
		user, err := db.FindOrCreateSPUser(req.Phone, writer)
		if err != nil {
			log.Printf("failed to create user: %s", err.Error())
			utils.WriteError(rw, http.StatusInternalServerError, "failed to create user")
			return
		}

		// If the user doesnt exist just save the phone number as the name
		if user.Name == "" {
			user.Name = req.Phone
			db.UpdateSPUser(user)
		}

		newGrpUser := models.SpGroupUser{
			SpGroupID: req.GroupID,
			SpUserID:  user.ID,
			Role: req.Role,
		}

		_, err = db.AddSPUserToGroup(&newGrpUser, writer)
		if err != nil {
			log.Printf("failed to add user to grp: %s", err)
			utils.WriteError(rw, http.StatusInternalServerError, "failed to add user")
			return
		}

		utils.WriteData(rw, http.StatusOK, newGrpUser)
	}
}

func CreateGetSPGroupUsersHandler(db store.Store) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		// First check if the user is an admin
		userID := r.Header.Get("SP_USERID")
		if len(userID) == 0 {
			utils.WriteError(rw, http.StatusForbidden, "forbidden")
			return
		}

		writer := os.Getenv("WRITER")

		v := mux.Vars(r)

		groupID := strings.TrimSpace(v["group_id"])

		if len(groupID) == 0 {
			utils.WriteError(rw, http.StatusBadRequest, "group_id mandatory")
			return
		}

		reqUser, err := db.GetSPUserByID(userID, writer)
		if err != nil {
			utils.WriteError(rw, http.StatusBadRequest, "unknown user")
			return
		}

		// check if this user is part of this group
		// OR
		// user is superadmin
		if reqUser.Role != "admin" {
			// check if the requester is admin of the group
			_, err := db.GetSPGroupUser(userID, groupID, writer)
			if err != nil {
				log.Printf("groupuser not found:%s", err.Error())
				utils.WriteError(rw, http.StatusBadRequest, "invalid input")
				return
			}
		}

		users, err := db.GetSPGroupUsers(groupID, writer)
		if err != nil {
			log.Printf("groupusers not found:%s", err.Error())
			utils.WriteError(rw, http.StatusBadRequest, "invalid input")
			return
		}

		utils.WriteData(rw, http.StatusOK, users)
	}
}
