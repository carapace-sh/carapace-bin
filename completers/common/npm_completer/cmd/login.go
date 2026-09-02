package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Add a registry user account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(loginCmd).Standalone()
	loginCmd.Flags().String("auth-type", "", "authentication strategy")
	loginCmd.Flags().String("scope", "", "associate with scope")

	rootCmd.AddCommand(loginCmd)

	carapace.Gen(loginCmd).FlagCompletion(carapace.ActionMap{
		"auth-type": carapace.ActionValues("legacy", "web"),
	})
}
