package impl

import (
	"context"
	"fmt"

	"github.com/kade-chen/library/exception"
	"github.com/kade-chen/mcenter/apps/roles"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/kade-chen/mcenter/apps/policy"
)

func (s *service) Check_Permission(ctx context.Context, req *policy.User_Permission_Request) (*policy.Permission_Message, error) {
	var RoleUserBinding roles.UserBindingRoles
	var Role roles.Roles
	var username string
	for _, v := range req.User {
		username = v.Spec.Username
	}
	// result = roles.NewRoleUserBinding()
	//1.query wether user binding roles is exist //2.if exist, query wether user has permission
	err := s.user_binding_roles.FindOne(ctx, bson.M{"_id": username}).Decode(&RoleUserBinding)
	if err != nil {
		//2.if not exist,return error
		s.log.Error().Msgf("this %s not found", username)
		return nil, exception.NewPermissionAuthenticationFailed("this %s not found", username)
	} else {
		//3.permission check
		for _, v := range RoleUserBinding.RoleId {
			// 3.1 if v is "User Admin",then pass
			if v == req.Roles_Admin || v == "Owner" {
				// 4.create user
				s.log.Info().Msgf("The roles is %s for %s", v, username)
				goto condition
				// break
			} else {
				//3.2 else if v == "User Admin" ,query wether user has permission
				s.role.FindOne(ctx, bson.M{"_id": v}).Decode(&Role)
				for _, v := range Role.IncludedPermissions {
					if v == req.Permission_Name {
						// 4.create user
						s.log.Info().Msgf("The roles is %s for %s", v, username)
						goto condition
						// break
					}
				}
			}
		}
		s.log.Error().Msgf("this %s not permission, Lack of permission %s", username, req.Permission_Name)
		return nil, exception.NewPermissionAuthenticationFailed("this %s not permission, Lack of permission %s", username, req.Permission_Name)
	}
condition:
	return &policy.Permission_Message{Message: fmt.Sprintf("this %s has permission", username)}, nil
}
