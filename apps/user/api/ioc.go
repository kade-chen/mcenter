package api

import (
	"github.com/kade-chen/library/ioc"
	"github.com/kade-chen/library/ioc/config/gorestful"
	logs "github.com/kade-chen/library/ioc/config/log"
	ioc_mongo "github.com/kade-chen/library/ioc/config/mongo"
	"github.com/kade-chen/mcenter/apps/policy"
	"github.com/kade-chen/mcenter/apps/user"
	"github.com/kade-chen/mcenter/middlewares"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	ioc.Api().Registry(&userHandler{})
}

type userHandler struct {
	ioc.ObjectImpl
	user               user.Service
	log                *zerolog.Logger
	user_binding_roles *mongo.Collection
	role               *mongo.Collection
	policy             policy.Service
}

func (u *userHandler) Init() error {
	db := ioc_mongo.DB()
	u.log = logs.Sub(user.AppName)
	u.role = db.Collection("roles")
	u.user = ioc.Controller().Get(user.AppName).(user.Service)
	u.policy = ioc.Controller().Get(policy.AppName).(policy.Service)
	u.user_binding_roles = db.Collection("user_binding_roles")
	u.Registry()
	return nil
}

func (u *userHandler) Name() string {
	return user.AppName
}

func (u *userHandler) Version() string {
	return "v1"
}

func (u *userHandler) Registry() {
	tags := []string{"User"}
	ws := gorestful.InitRouter(u)
	ws.Route(ws.POST("/").To(u.create_user).
		Doc("create user").
		Metadata(restfulspec.KeyOpenAPITags, tags). //标签
		Notes("create user").
		Writes(user.User{}).
		Reads(user.CreateUserRequest{}).
		Returns(200, "Ok", user.User{}).
		Filter(middlewares.NewTokenAuther().Auth_Login))

	ws.Route(ws.GET("/").To(u.list_user).
		Doc("list user").
		Metadata(restfulspec.KeyOpenAPITags, tags). //标签
		Notes("list user").
		Reads(user.QueryUserRequest{}).
		Writes(user.UserSet{}).
		Returns(200, "Ok", user.UserSet{}).
		Filter(middlewares.NewTokenAuther().Auth_Login))

	ws.Route(ws.DELETE("/").To(u.delete_user).
		Doc("delete user").
		Metadata(restfulspec.KeyOpenAPITags, tags). //标签
		Notes("delete user").
		Reads(user.DeleteUserRequest{}).
		Writes(user.UserSet{}).
		Returns(200, "Ok", user.UserSet{}).
		Filter(middlewares.NewTokenAuther().Auth_Login))

}
