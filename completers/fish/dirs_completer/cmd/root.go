package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dirs",
	Short: "Print directory stack",
	Long:  "https://fishshell.com/docs/current/cmds/dirs.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("c", "c", false, "clear the directory stack")
	rootCmd.Flags().BoolS("h", "h", false, "display help")
}
