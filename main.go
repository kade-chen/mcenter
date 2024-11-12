package main

import (
	_ "github.com/kade-chen/mcenter/apps" //registry all impl/api

	"github.com/kade-chen/mcenter/cmd"
)

func main() {
	cmd.Execute()
}
