package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/pacman"
	"github.com/spf13/cobra"
)

var webCmd = &cobra.Command{
	Use:     "web",
	Aliases: []string{"W"},
	Short:   "",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(webCmd).Standalone()

	carapace.Gen(webCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return pacman.ActionPackageSearch().Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
