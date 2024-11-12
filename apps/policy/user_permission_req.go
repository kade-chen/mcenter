package policy

import "github.com/kade-chen/mcenter/apps/user"

// new user_permission_request receive username rolesAdmin permissionName parameters
func NewUser_Permission_Request(username, rolesAdmin, permissionName string) *User_Permission_Request {
	return &User_Permission_Request{
		User: []*user.User{
			{Spec: &user.CreateUserRequest{
				Username: username,
			}}},
		Permission_Name: permissionName,
		Roles_Admin:     rolesAdmin,
	}
}
