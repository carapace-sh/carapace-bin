package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "continue",
	Short: "Resume the next iteration of a loop",
	Long:  "https://zsh.sourceforge.io/Doc/Release/Shell-Builtin-Commands.html#index-continue",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("i", "i", false, "resume from loops running in interactive mode")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValues("1", "2", "3", "4", "5", "6", "7", "8", "9").Usage("loop count to continue"),
	)
}
