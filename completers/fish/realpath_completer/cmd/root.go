package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "realpath",
	Short: "Resolve path to absolute",
	Long:  "https://fishshell.com/docs/current/cmds/realpath.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("h", "h", false, "display help")
	rootCmd.Flags().BoolS("s", "s", false, "don't resolve symlinks")
	rootCmd.Flags().Bool("no-symlinks", false, "don't resolve symlinks")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
