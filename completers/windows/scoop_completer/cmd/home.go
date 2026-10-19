package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/scoop"
	"github.com/spf13/cobra"
)

var homeCmd = &cobra.Command{
	Use:   "home",
	Short: "opens the app homepage",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(homeCmd).Standalone()
	rootCmd.AddCommand(homeCmd)

	carapace.Gen(homeCmd).PositionalCompletion(
		scoop.ActionAvailableApps(),
	)
}
