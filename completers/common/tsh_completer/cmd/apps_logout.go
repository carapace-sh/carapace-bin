package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tsh"
	"github.com/spf13/cobra"
)

var apps_logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Remove app certificate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apps_logoutCmd).Standalone()

	appsCmd.AddCommand(apps_logoutCmd)

	carapace.Gen(apps_logoutCmd).PositionalCompletion(
		tsh.ActionApps(),
	)
}
