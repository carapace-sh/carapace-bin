package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var updateMaintainersCmd = &cobra.Command{
	Use:     "update-maintainers",
	Short:   "Update the list of maintainers in the `Homebrew/brew` README",
	GroupID: "developer",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(updateMaintainersCmd).Standalone()

	updateMaintainersCmd.Flags().Bool("debug", false, "Display any debugging information.")
	updateMaintainersCmd.Flags().Bool("help", false, "Show this message.")
	updateMaintainersCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	updateMaintainersCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(updateMaintainersCmd)
}
