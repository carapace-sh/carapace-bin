package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fc",
	Short: "Display or execute commands from the history list",
	Long:  "https://www.gnu.org/software/bash/manual/html_node/Bash-Builtins.html#index-fc",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("e", "e", "", "select which editor to use")
	rootCmd.Flags().BoolS("l", "l", false, "list lines instead of editing")
	rootCmd.Flags().BoolS("n", "n", false, "omit line numbers when listing")
	rootCmd.Flags().BoolS("r", "r", false, "reverse the order of the lines")
	rootCmd.Flags().BoolS("s", "s", false, "re-execute command after substitution")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"e": carapace.ActionValuesDescribed(
			"vi", "vi editor",
			"emacs", "emacs editor",
			"nano", "nano editor",
		),
	})
}
