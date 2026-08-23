package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var launchCmd = &cobra.Command{
	Use:   "launch",
	Short: "Launch an application with CoreApplicationServices",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(launchCmd).Standalone()
	launchCmd.Flags().String("arg", "", "Argument to pass to the application")
	rootCmd.AddCommand(launchCmd)
}
