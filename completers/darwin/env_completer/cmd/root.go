package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "env",
	Short: "set environment and execute command, or print environment",
	Long:  "https://keith.github.io/xcode-manpages/env.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("0", "0", false, "End each output line with NUL")
	rootCmd.Flags().BoolS("i", "i", false, "Start with empty environment")

	rootCmd.Flags().StringS("C", "C", "", "Change to directory")
	rootCmd.Flags().StringS("P", "P", "", "Set PATH")
	rootCmd.Flags().StringS("S", "S", "", "Process string")
	rootCmd.Flags().StringS("u", "u", "", "Unset variable")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
