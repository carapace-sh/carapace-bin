package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "history",
	Short: "Display or manipulate the history list",
	Long:  "https://www.gnu.org/software/bash/manual/html_node/Bash-Builtins.html#index-history",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("a", "a", false, "append history lines from this session to the history file")
	rootCmd.Flags().BoolS("c", "c", false, "clear the history list by erasing the entries")
	rootCmd.Flags().BoolS("d", "d", false, "delete the history entry at offset offset")
	rootCmd.Flags().BoolS("n", "n", false, "read all history lines not already read from the history file")
	rootCmd.Flags().StringS("p", "p", "", "perform history expansion on each arg and display the result")
	rootCmd.Flags().BoolS("r", "r", false, "read the history file and append the contents to the history list")
	rootCmd.Flags().StringS("s", "s", "", "append the args to the history list as a single entry")
	rootCmd.Flags().BoolS("w", "w", false, "write the current history to the history file")
}
