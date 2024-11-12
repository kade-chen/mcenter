package impl

import (
	"github.com/kade-chen/library/ioc"
	"github.com/kade-chen/mcenter/apps/policy"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/kade-chen/library/ioc/config/grpc"
	logs "github.com/kade-chen/library/ioc/config/log"
	ioc_mongo "github.com/kade-chen/library/ioc/config/mongo"
)

var _ policy.Service = (*service)(nil)

func init() {
	ioc.Controller().Registry(&service{})
}

type service struct {
	ioc.ObjectImpl
	policy.UnimplementedRPCServer
	log                *zerolog.Logger
	user_binding_roles *mongo.Collection
	role               *mongo.Collection
}

func (i *service) Name() string {
	return policy.AppName
}

func (s *service) Init() error {
	s.log = logs.Sub(policy.AppName)
	db := ioc_mongo.DB()
	s.role = db.Collection("roles")
	s.user_binding_roles = db.Collection("user_binding_roles")

	policy.RegisterRPCServer(grpc.Get().Server(), s)
	return nil
}
