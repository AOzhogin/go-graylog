package mockserver

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/go-graylog"
)

// POST /users Create a new user account.
func (ms *Server) handleCreateUser(
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) (int, interface{}, error) {
	requiredFields := []string{
		"username", "email", "permissions", "full_name", "password"}
	allowedFields := []string{
		"startpage", "timezone", "session_timeout_ms", "roles"}
	body, sc, err := validateRequestBody(r.Body, requiredFields, allowedFields, nil)
	if err != nil {
		return sc, nil, err
	}

	user := &graylog.User{}
	if err := msDecode(body, user); err != nil {
		ms.Logger().WithFields(log.Fields{
			"body": body, "error": err,
		}).Info("Failed to parse request body as User")
		return 400, nil, err
	}

	if sc, err := ms.AddUser(user); err != nil {
		return sc, nil, err
	}
	ms.safeSave()
	return 201, nil, nil
}

// GET /users List all users
func (ms *Server) handleGetUsers(
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) (int, interface{}, error) {
	arr, err := ms.UserList()
	if err != nil {
		return 500, nil, err
	}
	users := &graylog.UsersBody{Users: arr}
	return 200, users, nil
}

// GET /users/{username} Get user details
func (ms *Server) handleGetUser(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (int, interface{}, error) {
	name := ps.ByName("username")
	user, err := ms.GetUser(name)
	if err != nil {
		return 500, nil, err
	}
	if user == nil {
		return 404, nil, fmt.Errorf(`no user found with name "%s"`, name)
	}
	return 200, user, nil
}

// PUT /users/{username} Modify user details.
func (ms *Server) handleUpdateUser(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (int, interface{}, error) {
	// required fields is nil
	acceptedFields := []string{
		"email", "permissions", "full_name", "password"}
	body, sc, err := validateRequestBody(r.Body, nil, nil, acceptedFields)
	if err != nil {
		return sc, nil, err
	}

	user := &graylog.User{Username: ps.ByName("username")}
	if err := msDecode(body, &user); err != nil {
		ms.Logger().WithFields(log.Fields{
			"body": body, "error": err,
		}).Info("Failed to parse request body as User")
		return 400, nil, err
	}

	if sc, err := ms.UpdateUser(user); err != nil {
		return sc, nil, err
	}
	ms.safeSave()
	return 200, nil, nil
}

// DELETE /users/{username} Removes a user account
func (ms *Server) handleDeleteUser(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (int, interface{}, error) {
	name := ps.ByName("username")
	if sc, err := ms.DeleteUser(name); err != nil {
		return sc, nil, err
	}
	ms.safeSave()
	return 204, nil, nil
}
