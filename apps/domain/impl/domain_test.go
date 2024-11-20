package impl_test

import (
	"context"
	"testing"

	"github.com/kade-chen/library/ioc"
	"github.com/kade-chen/mcenter/apps/domain"

	_ "github.com/kade-chen/mcenter/apps/domain/impl"
)

var (
	ctx  context.Context
	impl domain.Service
)

func TestCreateDomain(t *testing.T) {
	req := domain.NewCreateDomainRequest()
	req.Name = "kade-domain"
	req.Description = "test"
	ins, _ := impl.CreateDomain(ctx, req)
	t.Log(ins.ToJson())
}

func TestDescribeDomain(t *testing.T) {
	// req := domain.NewDescribeDomainRequestByName(domain.DEFAULT_DOMAIN)
	req := domain.NewDescribeDomainRequestByName("kade-domain")
	ins, err := impl.DescribeDomain(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	// t.Log(ins)
	t.Log(ins.ToJson())
}

func init() {
	req := ioc.NewLoadConfigRequest()
	req.ConfigFile.Enabled = true
	req.ConfigFile.Path = "/Users/kade.chen/go-kade-project/github/mcenter/etc/config.toml"
	ioc.DevelopmentSetup(req)
	impl = ioc.Controller().Get(domain.AppName).(domain.Service)
}
