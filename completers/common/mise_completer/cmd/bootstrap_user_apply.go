package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_user_applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Set your login shell from `[bootstrap.user]`",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_user_applyCmd).Standalone()

	bootstrap_user_applyCmd.Flags().BoolP("dry-run", "n", false, "Show what would change without changing anything")
	bootstrap_user_applyCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_user_applyCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
	bootstrap_userCmd.AddCommand(bootstrap_user_applyCmd)
}
