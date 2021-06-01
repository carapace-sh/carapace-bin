package cmd

import (
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in to the Pulumi service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	loginCmd.PersistentFlags().StringP("cloud-url", "c", "", "A cloud URL to log in to")
	loginCmd.PersistentFlags().BoolP("local", "l", false, "Use Pulumi in local-only mode")
	rootCmd.AddCommand(loginCmd)
}
