package roles

import "context"

const (
	AppName = "roles"
)

type Service interface {
	CreateRole(context.Context, *CreateRoleRequest) (*Roles, error)
	RPCServer
}
