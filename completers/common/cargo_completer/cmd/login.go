package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in to a registry.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(loginCmd).Standalone()

	loginCmd.Flags().BoolP("help", "h", false, "Print help")
	loginCmd.Flags().String("registry", "", "Registry to use")
	rootCmd.AddCommand(loginCmd)
}
