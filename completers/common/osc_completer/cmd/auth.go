package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Cloud authentication operations",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var authLoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to the cloud",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var authShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show current authentication information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(authCmd).Standalone()
	carapace.Gen(authLoginCmd).Standalone()
	carapace.Gen(authShowCmd).Standalone()

	rootCmd.AddCommand(authCmd)
	authCmd.AddCommand(authLoginCmd)
	authCmd.AddCommand(authShowCmd)
}
