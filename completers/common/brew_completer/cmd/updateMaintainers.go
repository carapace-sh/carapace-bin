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

	updateMaintainersCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	updateMaintainersCmd.Flags().BoolP("help", "h", false, "Show this message.")
	updateMaintainersCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	updateMaintainersCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(updateMaintainersCmd)
}
