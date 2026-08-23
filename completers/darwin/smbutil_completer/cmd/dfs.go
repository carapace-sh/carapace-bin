package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dfsCmd = &cobra.Command{
	Use:   "dfs",
	Short: "Display Dfs referrals for a URL",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dfsCmd).Standalone()

	carapace.Gen(dfsCmd).PositionalCompletion(
		carapace.ActionValues(),
	)
}
