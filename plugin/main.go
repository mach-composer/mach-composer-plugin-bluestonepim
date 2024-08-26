package plugin

import (
	"github.com/mach-composer/mach-composer-plugin-sdk/plugin"

	"github.com/mach-composer/mach-composer-plugin-bluestonepim/internal"
)

// Serve serves the plugin
func Serve() {
	p := internal.NewBluestonePimPlugin()
	plugin.ServePlugin(p)
}
