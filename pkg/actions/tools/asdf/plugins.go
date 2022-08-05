package asdf

import (
	"time"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/gh"
	"github.com/rsteube/carapace/pkg/style"
)

// ActionPlugins completes plugins
//   1password-cli
//   R
func ActionPlugins() carapace.Action {
	return gh.ActionContents(gh.ContentOpts{Owner: "asdf-vm", Name: "asdf-plugins", Path: "plugins"}).Style(style.Default).Cache(24 * time.Hour)
}
