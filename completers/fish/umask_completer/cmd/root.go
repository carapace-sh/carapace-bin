package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "umask",
	Short: "Set file creation mask",
	Long:  "https://fishshell.com/docs/current/cmds/umask.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("S", "S", false, "symbolic mode")
	rootCmd.Flags().Bool("symbolic", false, "symbolic mode")
	rootCmd.Flags().BoolS("h", "h", false, "display help")
	rootCmd.Flags().BoolS("p", "p", false, "output reusable form")
	rootCmd.Flags().Bool("as-command", false, "output reusable form")
}
