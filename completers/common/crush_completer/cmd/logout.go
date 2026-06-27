package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/crush"
	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	Use:     "logout [platform]",
	Short:   "Logout Crush from a platform",
	Aliases: []string{"signout"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logoutCmd).Standalone()

	logoutCmd.Flags().BoolP("force", "f", false, "Skip logout confirmation prompt")
	rootCmd.AddCommand(logoutCmd)

	carapace.Gen(logoutCmd).PositionalCompletion(
		crush.ActionPlatforms(),
	)
}
