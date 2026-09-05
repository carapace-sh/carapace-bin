package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var completionServerCmd = &cobra.Command{
	Use:    "completion-server",
	Short:  "Dynamic completion endpoint used by generated shell scripts",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(completionServerCmd).Standalone()

	completionServerCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(completionServerCmd)
}
