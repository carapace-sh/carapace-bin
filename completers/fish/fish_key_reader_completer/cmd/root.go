package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fish_key_reader",
	Short: "Explore keyboard input",
	Long:  "https://fishshell.com/docs/current/cmds/fish_key_reader.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("c", "c", false, "continuous mode")
	rootCmd.Flags().Bool("continuous", false, "continuous mode")
	rootCmd.Flags().BoolS("h", "h", false, "display help")
	rootCmd.Flags().BoolS("V", "V", false, "explain sequence")
	rootCmd.Flags().Bool("verbose", false, "explain sequence")
	rootCmd.Flags().BoolS("v", "v", false, "display version")
}
