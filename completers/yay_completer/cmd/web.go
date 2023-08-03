package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/pacman"
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
		pacman.ActionPackageSearch().FilterArgs(),
	)
}
