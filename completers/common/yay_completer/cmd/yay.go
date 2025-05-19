package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/yay"
	"github.com/spf13/cobra"
)

var yayCmd = &cobra.Command{
	Use:     "yay",
	Aliases: []string{"Y"},
	Short:   "YAY specific operations",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(yayCmd).Standalone()

	yayCmd.Flags().BoolP("clean", "c", false, "Remove unneeded dependencies")
	yayCmd.Flags().Bool("gendb", false, "Generates development package DB used for updating")

	carapace.Gen(yayCmd).PositionalAnyCompletion(
		yay.ActionPackageSearch().FilterArgs(),
	)
}
