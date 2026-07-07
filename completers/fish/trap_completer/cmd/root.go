package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "trap",
	Short: "Handle signals",
	Long:  "https://fishshell.com/docs/current/cmds/trap.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("h", "h", false, "display help")
	rootCmd.Flags().BoolS("l", "l", false, "print signal names")
	rootCmd.Flags().Bool("list-signals", false, "print signal names")
	rootCmd.Flags().BoolS("p", "p", false, "print handlers")
	rootCmd.Flags().Bool("print", false, "print handlers")
}
