package impl_test

import (
	"context"
	"testing"

	"github.com/kade-chen/library/ioc"
	"github.com/kade-chen/mcenter/apps/stt"
	_ "github.com/kade-chen/mcenter/apps/stt/v2/impl"
)

var (
	ctx  = context.Background()
	impl stt.ServiceV2
)

func TestC(t *testing.T) {
	// impl.ListLocations(ctx)

	// impl.CreateRecognizer(ctx, stt.ENDPOINTS_SPEECH_GOOGLEAPIS_COM)
	impl.ListRecognizers(ctx, stt.ENDPOINTS_EUROPE_WEST4_SPEECH_GOOGLEAPIS_COM)
	// impl.GetModel(ctx, stt.ENDPOINTS_SPEECH_GOOGLEAPIS_COM)
}

func init() {
	req := ioc.NewLoadConfigRequest()
	req.ConfigFile.Enabled = true
	req.ConfigFile.Path = "../../../../etc/config.toml"
	ioc.DevelopmentSetup(req)
	impl = ioc.Controller().Get(stt.AppNameV2).(stt.ServiceV2)
}
