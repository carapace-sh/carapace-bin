package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "echo",
	Short: "Display a line of text",
	Long:  "https://fishshell.com/docs/current/cmds/echo.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("E", "E", false, "disable backslash escapes")
	rootCmd.Flags().BoolS("e", "e", false, "enable backslash escapes")
	rootCmd.Flags().BoolS("n", "n", false, "do not output a trailing newline")
	rootCmd.Flags().BoolS("s", "s", false, "do not separate arguments with spaces")
}
