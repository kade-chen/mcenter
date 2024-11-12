package impl_test

import (
	"context"
	"testing"

	"github.com/kade-chen/library/ioc"
	"github.com/kade-chen/mcenter/apps/policy"
	_ "github.com/kade-chen/mcenter/apps/policy/impl"
)

var (
	ctx  = context.Background()
	impl policy.Service
)

func Test_Check_permission(t *testing.T) {
	req := policy.NewUser_Permission_Request("admin", "User Admin", "user.create")
	// req.Roles_Admin = "User Admin"
	// req.Permission_Name = "user.create"
	// req.User = []*user.User{
	// 	{
	// 		Spec: &user.CreateUserRequest{
	// 			Username: "admin",
	// 		},
	// 	},
	// }
	a, err := impl.Check_Permission(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(a)
}

func init() {
	req := ioc.NewLoadConfigRequest()
	req.ConfigFile.Enabled = true
	req.ConfigFile.Path = "/Users/kade.chen/go-kade-project/mcenter/etc/config.toml"
	ioc.DevelopmentSetup(req)
	impl = ioc.Controller().Get(policy.AppName).(policy.Service)
}
