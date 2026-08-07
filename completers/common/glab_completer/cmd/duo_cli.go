package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var duo_cliCmd = &cobra.Command{
	Use:   "cli [command]",
	Short: "Run the GitLab Duo CLI.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(duo_cliCmd).Standalone()

	duo_cliCmd.Flags().Bool("install", false, "Install the GitLab Duo CLI binary without running it.")
	duo_cliCmd.Flags().Bool("update", false, "Check for and install updates to the binary.")
	duo_cliCmd.Flags().BoolP("yes", "y", false, "Skip confirmation prompts.")
	duoCmd.AddCommand(duo_cliCmd)
}
