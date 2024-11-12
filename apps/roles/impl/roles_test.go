package impl_test

import (
	"context"
	"testing"

	"github.com/kade-chen/library/ioc"
	_ "github.com/kade-chen/mcenter/apps/roles/impl"

	// _ "github.com/kade-chen/mcenter/apps"
	"github.com/kade-chen/mcenter/apps/roles"
)

var (
	ctx  = context.Background()
	impl roles.Service
)

// ----------------User----------------
func Test_User_Admin(t *testing.T) {
	req := roles.NewCreateRoleRequest()
	req.Name = "User Admin"
	req.Title = "roles/user.userAdmin"
	req.Description = "user.userAdmin.v1"
	req.Deleted = false
	req.IncludedPermissions = []string{
		"user.*",
		"user.list",
		"user.get",
		"user.create",
		"user.delete",
	}
	roles, err := impl.CreateRole(ctx, req)
	t.Log(roles, err)
}

func Test_User_View(t *testing.T) {
	req := roles.NewCreateRoleRequest()
	req.Name = "User view"
	req.Title = "roles/user.userView"
	req.Description = "user.userView.v1"
	req.Deleted = false
	req.IncludedPermissions = []string{
		"user.list",
	}
	roles, err := impl.CreateRole(ctx, req)
	t.Log(roles, err)
}

func Test_User_create(t *testing.T) {
	req := roles.NewCreateRoleRequest()
	req.Name = "User create"
	req.Title = "roles/user.userView"
	req.Description = "user.userView.v1"
	req.Deleted = false
	req.IncludedPermissions = []string{
		"user.create",
	}
	roles, err := impl.CreateRole(ctx, req)
	t.Log(roles, err)
}

// ----------------User----------------

// ----------------TextToSpeechV1----------------
func Test_TextToSpeechV1_Admin(t *testing.T) {
	req := roles.NewCreateRoleRequest()
	req.Name = "Text To Speech V1 Admin"
	req.Title = "roles/TextToSpeechV1.admin"
	req.Description = "TextToSpeechV1.admin.v1"
	req.Deleted = false
	req.IncludedPermissions = []string{
		"TextToSpeechV1.*",
		"TextToSpeechV1.list",
		"TextToSpeechV1.get",
		"TextToSpeechV1.create",
		"TextToSpeechV1.delete",
	}
	roles, err := impl.CreateRole(ctx, req)
	t.Log(roles, err)
}

func Test_TextToSpeechV1_View(t *testing.T) {
	req := roles.NewCreateRoleRequest()
	req.Name = "Text To Speech V1 View"
	req.Title = "roles/TextToSpeechV1.view"
	req.Description = "TextToSpeechV1.view.v1"
	req.Deleted = false
	req.IncludedPermissions = []string{
		"TextToSpeechV1.get",
		"TextToSpeechV1.list",
	}
	roles, err := impl.CreateRole(ctx, req)
	t.Log(roles, err)
}

// ----------------TextToSpeechV1----------------


// ----------------SpeechToTextV1----------------
func Test_SpeechToTextV1_Admin(t *testing.T) {
	req := roles.NewCreateRoleRequest()
	req.Name = "Speech To Text V1 Admin"
	req.Title = "roles/SpeechToTextV1.admin"
	req.Description = "SpeechToTextV1.admin.v1"
	req.Deleted = false
	req.IncludedPermissions = []string{
		"SpeechToTextV1.*",
		"SpeechToTextV1.list",
		"SpeechToTextV1.get",
		"SpeechToTextV1.create",
		"SpeechToTextV1.delete",
	}
	roles, err := impl.CreateRole(ctx, req)
	t.Log(roles, err)
}

func Test_SpeechToTextV1_View(t *testing.T) {
	req := roles.NewCreateRoleRequest()
	req.Name = "Speech To Text V1 View"
	req.Title = "roles/SpeechToTextV1.view"
	req.Description = "SpeechToTextV1.view.v1"
	req.Deleted = false
	req.IncludedPermissions = []string{
		"SpeechToTextV1.get",
		"SpeechToTextV1.list",
	}
	roles, err := impl.CreateRole(ctx, req)
	t.Log(roles, err)
}

// ----------------SpeechToTextV1----------------

func init() {
	req := ioc.NewLoadConfigRequest()
	req.ConfigFile.Enabled = true
	req.ConfigFile.Path = "/Users/kade.chen/go-kade-project/mcenter/etc/config.toml"
	ioc.DevelopmentSetup(req)
	impl = ioc.Controller().Get(roles.AppName).(roles.Service)
}
