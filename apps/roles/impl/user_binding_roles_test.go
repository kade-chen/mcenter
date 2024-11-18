package impl_test

import (
	"testing"

	"github.com/kade-chen/library/ioc"
	"github.com/kade-chen/mcenter/apps/roles"
)

var (
// ctx  = context.Background()
// impl roles.Service
)

func Test_User(t *testing.T) {
	roles, err := impl.UserBindingRole(ctx, &roles.CreateUserBindingRoleRequest{UsernameId: "admin", RoleId: "Owner"})
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(roles)
	}
}

func init() {
	req := ioc.NewLoadConfigRequest()
	req.ConfigFile.Enabled = true
	req.ConfigFile.Path = "/Users/kade.chen/go-kade-project/mcenter/etc/config.toml"
	// impl = ioc.Controller().Get(roles.AppName).(roles.Service)
}
