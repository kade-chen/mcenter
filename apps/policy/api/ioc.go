package api

import (
	"github.com/kade-chen/library/ioc"
	"github.com/kade-chen/library/ioc/config/gorestful"
	logs "github.com/kade-chen/library/ioc/config/log"
	"github.com/kade-chen/mcenter/apps/policy"
	"github.com/kade-chen/mcenter/middlewares"
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/rs/zerolog"
)

func init() {
	ioc.Api().Registry(&policyHandler{})
}

type policyHandler struct {
	ioc.ObjectImpl
	log *zerolog.Logger
	// user_binding_roles *mongo.Collection
	// role               *mongo.Collection
	policy policy.Service
}

func (u *policyHandler) Init() error {
	// db := ioc_mongo.DB()
	u.log = logs.Sub(policy.AppName)
	// u.role = db.Collection("roles")
	// u.user = ioc.Controller().Get(user.AppName).(user.Service)
	u.policy = ioc.Controller().Get(policy.AppName).(policy.Service)
	// u.user_binding_roles = db.Collection("user_binding_roles")
	u.Registry()
	return nil
}

func (u *policyHandler) Name() string {
	return policy.AppName
}

func (u *policyHandler) Version() string {
	return "v1"
}

func (u *policyHandler) Registry() {
	tags := []string{"Permission Verification"}
	ws := gorestful.InitRouter(u)
	ws.Route(ws.GET("/").To(u.check_permission).
		Doc("check permission user").                         //描述
		Metadata(restfulspec.KeyOpenAPITags, tags).           //标签
		Notes("check permission use").                        //描述
		Reads(policy.User_Permission_Request{}).              //请求参数
		Writes(policy.User_Permission_Request{}).             //返回值
		Returns(200, "Ok", policy.User_Permission_Request{}). //返回值
		Filter(middlewares.NewTokenAuther().Auth_Login))

}
