package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "contains",
	Short: "Test if a word is present in a list",
	Long:  "https://fishshell.com/docs/current/cmds/contains.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("h", "h", false, "display help")
	rootCmd.Flags().BoolS("i", "i", false, "print index")
	rootCmd.Flags().Bool("index", false, "print index")
}
