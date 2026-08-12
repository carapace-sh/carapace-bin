package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_latestCmd = &cobra.Command{
	Use:   "latest",
	Short: "Branch latest related commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_latestCmd).Standalone()

	branch_latestCmd.Flags().BoolP("help", "h", false, "Print help")
	branchCmd.AddCommand(branch_latestCmd)
}
