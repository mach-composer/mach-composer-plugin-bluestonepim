package main

import (
	"github.com/mach-composer/mach-composer-plugin-bluestonepim/internal"
	"github.com/mach-composer/mach-composer-plugin-sdk/plugin"
)

func main() {
	p := internal.NewBluestonePimPlugin()
	plugin.ServePlugin(p)
}
