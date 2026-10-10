package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var orbitCmd = &cobra.Command{
	Use:   "orbit [<command>] [flags]",
	Short: "Run the GitLab Orbit CLI. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(orbitCmd).Standalone()

	orbitCmd.Flags().BoolP("help", "h", false, "Show the Orbit binary's help, or this text until the binary is installed.")
	orbitCmd.Flags().Bool("install", false, "Install the Orbit binary without running it.")
	orbitCmd.Flags().Bool("update", false, "Check for and install updates to the binary. Same as the update command.")
	orbitCmd.Flags().BoolP("yes", "y", false, "Skip confirmation prompts.")
	rootCmd.AddCommand(orbitCmd)
}
