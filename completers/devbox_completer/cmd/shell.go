package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "Start a new shell or run a command with access to your packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shellCmd).Standalone()
	rootCmd.AddCommand(shellCmd)

	carapace.Gen(shellCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)

	// TODO dash completion - if possible?
}
