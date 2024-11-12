package impl

import (
	"github.com/kade-chen/library/ioc"
	// "github.com/kade-chen/library/ioc/config/grpc"
	logs "github.com/kade-chen/library/ioc/config/log"
	ioc_mongo "github.com/kade-chen/library/ioc/config/mongo"
	"github.com/kade-chen/mcenter/apps/roles"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
)

var _ roles.Service = (*service)(nil)

func init() {
	ioc.Controller().Registry(&service{})
}

type service struct {
	ioc.ObjectImpl
	role               *mongo.Collection
	permission         *mongo.Collection
	user_binding_roles *mongo.Collection
	log                *zerolog.Logger
	roles.UnimplementedRPCServer
}

func (s *service) Init() error {
	db := ioc_mongo.DB()
	s.role = db.Collection("roles")
	s.permission = db.Collection("permission")
	s.user_binding_roles = db.Collection("user_binding_roles")
	s.log = logs.Sub(roles.AppName)

	// roles.RegisterRPCServer(grpc.Get().Server(), s)
	return nil
}

func (s *service) Name() string {
	return roles.AppName
}
