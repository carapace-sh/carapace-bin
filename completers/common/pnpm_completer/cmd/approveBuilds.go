package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var approveBuildsCmd = &cobra.Command{
	Use:   "approve-builds",
	Short: "Approve dependencies for running scripts during installation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(approveBuildsCmd).Standalone()

	approveBuildsCmd.Flags().Bool("all", false, "Approve all pending dependencies without interactive prompts")
	approveBuildsCmd.Flags().BoolP("global", "g", false, "Approve builds for globally installed packages")
	approveBuildsCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(approveBuildsCmd)
}
