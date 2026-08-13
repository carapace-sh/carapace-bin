package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var git_disconnectCmd = &cobra.Command{
	Use:   "disconnect",
	Short: "Disconnect the Git repository from your Vercel Project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(git_disconnectCmd).Standalone()

	git_disconnectCmd.Flags().Bool("confirm", false, "(deprecated)")
	git_disconnectCmd.Flags().String("project", "", "Project name or ID")
	git_disconnectCmd.Flags().Bool("yes", false, "Skip confirmation")

	gitCmd.AddCommand(git_disconnectCmd)
}
