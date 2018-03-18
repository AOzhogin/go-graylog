package mockserver

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/suzuki-shunsuke/go-graylog"
)

type membersBody struct {
	Role  string         `json:"role"`
	Users []graylog.User `json:"users"`
}

// GET /roles/{rolename}/members Retrieve the role's members
func (ms *Server) handleRoleMembers(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (int, interface{}, error) {
	name := ps.ByName("rolename")
	ok, err := ms.HasRole(name)
	if err != nil {
		return 500, nil, err
	}
	if !ok {
		return 404, nil, fmt.Errorf("No role found with name %s", name)
	}
	arr, err := ms.RoleMembers(name)
	if err != nil {
		return 500, nil, err
	}
	users := &membersBody{Users: arr, Role: name}
	return 200, users, nil
}

// PUT /roles/{rolename}/members/{username} Add a user to a role
func (ms *Server) handleAddUserToRole(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (int, interface{}, error) {
	roleName := ps.ByName("rolename")
	userName := ps.ByName("username")
	ok, err := ms.HasRole(roleName)
	if err != nil {
		return 500, nil, err
	}
	if !ok {
		return 404, nil, fmt.Errorf("No role found with name %s", roleName)
	}

	user, err := ms.GetUser(userName)
	if err != nil {
		return 500, nil, err
	}
	if user == nil {
		return 404, nil, fmt.Errorf("User %s has not been found.", userName)
	}
	user.Roles = addToStringArray(user.Roles, roleName)
	ms.AddUser(user)
	return 200, nil, nil
}

// DELETE /roles/{rolename}/members/{username} Remove a user from a role
func (ms *Server) handleRemoveUserFromRole(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (int, interface{}, error) {
	roleName := ps.ByName("rolename")
	userName := ps.ByName("username")
	ok, err := ms.HasRole(roleName)
	if err != nil {
		return 500, nil, err
	}
	if !ok {
		return 404, nil, fmt.Errorf("No role found with name %s", roleName)
	}

	user, err := ms.GetUser(userName)
	if err != nil {
		return 500, nil, err
	}
	if user == nil {
		return 404, nil, fmt.Errorf("User %s has not been found.", userName)
	}
	user.Roles = removeFromStringArray(user.Roles, roleName)
	if sc, err := ms.UpdateUser(user); err != nil {
		return sc, nil, err
	}
	return 200, nil, nil
}
