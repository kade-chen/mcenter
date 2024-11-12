package impl

import (
	"context"

	"github.com/kade-chen/library/exception"
	"github.com/kade-chen/mcenter/apps/roles"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *service) UserBindingRole(ctx context.Context, req *roles.CreateUserBindingRoleRequest) (*roles.UserBindingRoles, error) {
	//0.query wether role for mongodb is exist in roles
	err1 := s.role.FindOne(ctx, bson.M{"_id": req.RoleId})
	if err1.Err() != nil {
		return nil, exception.NewInternalServerError("not this role:%s error:%v", req.RoleId, err1.Err().Error())
	}
	rub := roles.NewRoleUserBinding()
	rub.UsernameId = req.UsernameId
	rub.RoleId = append(rub.RoleId, req.RoleId)
	// 操作后返回更新后的文档。
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	//1.query wether user for mongodb is exist in roles_binding_user,if exist, update it
	err := s.user_binding_roles.FindOneAndUpdate(ctx, bson.M{"_id": rub.UsernameId}, bson.M{"$addToSet": bson.M{"role_id": bson.M{"$each": rub.RoleId}}}, opts).Decode(&rub)
	//2.if not exist user_binding_roles for mongodb,insert it
	if err != nil {
		s.log.Info().Msgf("not this user_binding_roles:%s error:%v", rub.UsernameId, err.Error())
		InsertOneResult, err := s.user_binding_roles.InsertOne(ctx, rub)
		if err != nil {
			return nil, exception.NewInternalServerError("insert user_binding_roles error:%s", err.Error())
		}
		s.log.Info().Msgf("inserted user_binding_roles username:%v", InsertOneResult.InsertedID)
	}

	// fmt.printf("Updated User: %v\n", rub)
	//3.create role user binding
	return rub, nil
}
