package action

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

// ActionCocs completes codes of conduct
func ActionCocs(cmd *cobra.Command) carapace.Action {
	return gh.ActionCocs()
}
