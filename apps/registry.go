package apps

import (
	//service registry for token
	_ "github.com/kade-chen/mcenter/apps/domain/impl"
	_ "github.com/kade-chen/mcenter/apps/policy/impl"
	_ "github.com/kade-chen/mcenter/apps/roles/impl"
	_ "github.com/kade-chen/mcenter/apps/stt/v1/impl"
	_ "github.com/kade-chen/mcenter/apps/token/impl"
	_ "github.com/kade-chen/mcenter/apps/tts/v1/impl"
	_ "github.com/kade-chen/mcenter/apps/user/impl"
	_ "github.com/kade-chen/mcenter/apps/vertex/impl"

	//api
	_ "github.com/kade-chen/mcenter/apps/policy/api"
	_ "github.com/kade-chen/mcenter/apps/stt/v1/api"
	_ "github.com/kade-chen/mcenter/apps/token/api"
	_ "github.com/kade-chen/mcenter/apps/tts/v1/api"
	_ "github.com/kade-chen/mcenter/apps/user/api"
	_ "github.com/kade-chen/mcenter/apps/vertex/api"

	//注册所有provider
	_ "github.com/kade-chen/mcenter/apps/token/provider/all"
)
