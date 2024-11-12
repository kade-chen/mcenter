package roles

import "github.com/kade-chen/library/exception"

// New 新创建一个Role
func NewRoles(req *CreateRoleRequest) (*Roles, error) {
	//verity sturct wether valid
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("struct is invalid,err: %s", err.Error())
	}

	r := &Roles{
		Name:                req.Name,
		Title:               req.Title,
		Stage:               req.Stage,
		Description:         req.Description,
		IncludedPermissions: req.IncludedPermissions,
		Deleted:             req.Deleted,
	}
	return r, nil
}

// new_request sturstct
func NewCreateRoleRequest() *CreateRoleRequest {
	return &CreateRoleRequest{}
}

// NewRolesBindingUsers
func NewRoleUserBinding() *UserBindingRoles {
	return &UserBindingRoles{
		RoleId: []string{},
	}
}
