package password_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/kade-chen/library/ioc"
	_ "github.com/kade-chen/mcenter/apps"
	"github.com/kade-chen/mcenter/apps/token"
	"github.com/kade-chen/mcenter/apps/token/provider"
)

var (
	impl provider.TokenIssuer
	ctx  = context.Background()
)

func TestIssueToken(t *testing.T) {
	// req := token.NewPasswordIssueTokenRequest("admin", "123456")
	tk, err := impl.IssueToken(ctx, &token.IssueTokenRequest{Username: "admin", Password: "123456"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tk.Json())
}

func TestMain1(t *testing.T) {
	fmt.Println(time.Unix(1731983085, 0))
	expiredAt := time.Unix(1731983085, 0).Add(time.Duration(1) * time.Hour * 24)
	fmt.Println(expiredAt)
	fmt.Println(time.Unix(1731983085, 0).Sub(expiredAt).Hours())
}

func init() {
	// os.Setenv("DEBUG", "true") //this debug conflicts with another vs debug
	req := ioc.NewLoadConfigRequest()
	req.ConfigFile.Enabled = true
	req.ConfigFile.Path = "/Users/kade.chen/go-kade-project/github/mcenter/etc/config.toml"
	ioc.DevelopmentSetup(req)
	impl = provider.GetTokenIssuer(token.GRANT_TYPE_PASSWORD)
}
