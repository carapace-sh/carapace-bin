package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pi"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"uninstall"},
	Short:   "Remove extension source from settings",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(removeCmd).Standalone()
	removeCmd.Flags().BoolP("approve", "a", false, "Trust project-local files for this command")
	removeCmd.Flags().BoolP("local", "l", false, "Remove from project settings (.pi/settings.json)")
	removeCmd.Flags().Bool("no-approve", false, "Ignore project-local files for this command")
	rootCmd.AddCommand(removeCmd)

	carapace.Gen(removeCmd).PositionalAnyCompletion(
		pi.ActionPackages(),
	)
}