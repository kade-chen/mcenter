package impl

import (
	"context"

	"github.com/kade-chen/library/exception"
	"github.com/kade-chen/mcenter/apps/roles"
)

func (s *service) CreateRole(ctx context.Context, req *roles.CreateRoleRequest) (*roles.Roles, error) {
	//0.new CreateRoleRequest
	rs, err := roles.NewRoles(req)
	if err != nil {
		return nil, err
	}
	//1.judgment wether the role name is sole
	// s.role.FindOne(ctx, bson.M{"_id": rs.Name})
	_, err = s.role.InsertOne(ctx, rs)
	if err != nil {
		return nil, exception.NewInternalServerError("inserted role(%s) document error, %s", rs.Name, err)
	}

	// if result != nil {
	// fmt.Println(result)
	// }
	return rs, nil
}
