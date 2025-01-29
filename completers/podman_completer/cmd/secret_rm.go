package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var secret_rmCmd = &cobra.Command{
	Use:   "rm [options] SECRET [SECRET...]",
	Short: "Remove one or more secrets",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secret_rmCmd).Standalone()

	secret_rmCmd.Flags().BoolP("all", "a", false, "Remove all secrets")
	secret_rmCmd.Flags().BoolP("ignore", "i", false, "Ignore errors when a specified secret is missing")
	secretCmd.AddCommand(secret_rmCmd)
}
