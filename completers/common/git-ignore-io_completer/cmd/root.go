package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-ignore-io",
	Short: "Generate sample .gitignore from gitignore.io",
	Long:  "https://github.com/tj/git-extras/blob/master/Commands.md#git-ignore-io",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringArrayP("append", "a", nil, "Append types to .gitignore")
	rootCmd.Flags().Bool("help", false, "show help")
	rootCmd.Flags().BoolP("list-alphabetically", "L", false, "List available types alphabetically")
	rootCmd.Flags().BoolP("list-in-table", "l", false, "List available types in table format")
	rootCmd.Flags().StringArrayP("replace", "r", nil, "Replace .gitignore with types")
	rootCmd.Flags().StringP("search", "s", "", "Search for a word")
	rootCmd.Flags().BoolP("show-update-time", "t", false, "Show update time")
	rootCmd.Flags().BoolP("update-list", "u", false, "Update the list of types")
}
