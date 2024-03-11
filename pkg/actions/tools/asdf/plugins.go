package asdf

import (
	"time"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionPlugins completes plugins
//
//	1password-cli
//	R
func ActionPlugins() carapace.Action {
	return gh.ActionContents(gh.ContentOpts{Owner: "asdf-vm", Name: "asdf-plugins", Path: "plugins"}).Style(style.Default).Cache(24 * time.Hour)
}
