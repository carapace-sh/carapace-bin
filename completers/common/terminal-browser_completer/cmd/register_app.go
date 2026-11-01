package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var registerAppCmd = &cobra.Command{
	Use:   "register-app",
	Short: "Register a terminal-browser application",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(registerAppCmd).Standalone()

	registerAppCmd.Flags().String("args", "", "Arguments passed to the binary")
	registerAppCmd.Flags().String("bin", "", "Path to the application binary")
	registerAppCmd.Flags().String("id", "", "Application id")
	registerAppCmd.Flags().String("name", "", "Application name")
	rootCmd.AddCommand(registerAppCmd)

	carapace.Gen(registerAppCmd).FlagCompletion(carapace.ActionMap{
		"bin": carapace.ActionFiles(),
	})
}
