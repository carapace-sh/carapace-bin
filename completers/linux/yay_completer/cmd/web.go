package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/yay"
	"github.com/spf13/cobra"
)

var webCmd = &cobra.Command{
	Use:     "web",
	Aliases: []string{"W"},
	Short:   "Web operations",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(webCmd).Standalone()

	carapace.Gen(webCmd).PositionalAnyCompletion(
		yay.ActionPackageSearch().FilterArgs(),
	)
}
