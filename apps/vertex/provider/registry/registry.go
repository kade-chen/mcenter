package registry

import (
	"github.com/kade-chen/library/exception"
	"github.com/kade-chen/mcenter/apps/vertex/provider"
	_ "github.com/kade-chen/mcenter/apps/vertex/provider/gemini2_flash"
)

func ProviderInit() error {
	err := provider.Init()
	if err != nil {
		return exception.NewIocImplRegisterFailed("vertex provider init error, ERROR: %s", err.Error())
	}
	return nil
}
