package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get info about the given branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_infoCmd).Standalone()

	branch_infoCmd.Flags().BoolP("help", "h", false, "Print help")
	branchCmd.AddCommand(branch_infoCmd)
}
