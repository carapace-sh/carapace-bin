package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var git_connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect your Vercel Project to your Git repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(git_connectCmd).Standalone()

	git_connectCmd.Flags().Bool("confirm", false, "(deprecated)")
	git_connectCmd.Flags().String("project", "", "Project name or ID")
	git_connectCmd.Flags().Bool("yes", false, "Skip confirmation")

	gitCmd.AddCommand(git_connectCmd)
}
