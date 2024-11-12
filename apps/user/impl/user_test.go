package impl_test

import (
	"context"
	"testing"

	"github.com/kade-chen/library/ioc"
	"github.com/kade-chen/mcenter/apps/user"
	_ "github.com/kade-chen/mcenter/apps/user/impl"
)

var (
	ctx  = context.Background()
	impl user.Service
)

// create user
func TestCreateUser(t *testing.T) {
	u, err := impl.CreateUser(ctx, &user.CreateUserRequest{
		Username: "kade1",
		Password: "123456",
		Domain:   "kade-domain",
		Type:     user.TYPE_SUB,
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(u.ToJson())
}

// query user
func TestQueryUser(t *testing.T) {
	req := user.NewQueryUserRequest(nil)
	var myType user.TYPE = user.TYPE_SUB
	req.Type = &myType
	a, err := impl.ListUser(ctx, req)
	if err != nil {
		t.Error(err)
	}
	// fmt.Println(a)
	t.Log(a.Tojson1())
}

// delete user
func TestDeleteUser(t *testing.T) {
	req := user.NewDeleteUserRequest()
	req.UserIds = []string{"CreteSubUser@kade-domain", "cc@cc"}
	a, err := impl.DeleteUser(ctx, req)
	if err != nil {
		t.Error(err)
	}
	// fmt.Println(a)
	t.Log(a.Tojson())
}

func init() {
	req := ioc.NewLoadConfigRequest()
	req.ConfigFile.Enabled = true
	req.ConfigFile.Path = "/Users/kade.chen/go-kade-project/mcenter/etc/config.toml"
	ioc.DevelopmentSetup(req)
	impl = ioc.Controller().Get(user.AppName).(user.Service)
}
