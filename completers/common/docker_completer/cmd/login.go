package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:     "login [OPTIONS] [SERVER]",
	Short:   "Log in to a registry",
	GroupID: "common",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(loginCmd).Standalone()

	loginCmd.Flags().StringP("password", "p", "", "Password")
	loginCmd.Flags().Bool("password-stdin", false, "Take the password from stdin")
	loginCmd.Flags().StringP("username", "u", "", "Username")
	rootCmd.AddCommand(loginCmd)
}
