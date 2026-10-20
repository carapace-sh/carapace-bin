package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_userCmd = &cobra.Command{
	Use:   "user",
	Short: "Manage your login shell from `[bootstrap.user]`",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_userCmd).Standalone()

	bootstrap_userCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrapCmd.AddCommand(bootstrap_userCmd)
}
