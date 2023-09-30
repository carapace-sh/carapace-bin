package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var styleCmd = &cobra.Command{
	Use:   "style",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(styleCmd).Standalone()

	carapace.Gen(styleCmd).PositionalCompletion(
		carapace.ActionStyleConfig(),
	)
}
