package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register with the configured server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(registerCmd).Standalone()

	registerCmd.Flags().StringP("email", "e", "", "")
	registerCmd.Flags().BoolP("help", "h", false, "Print help")
	registerCmd.Flags().StringP("password", "p", "", "")
	registerCmd.Flags().StringP("username", "u", "", "")
	rootCmd.AddCommand(registerCmd)
}
