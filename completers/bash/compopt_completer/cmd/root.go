package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "compopt",
	Short: "Modify or display completion options",
	Long:  "https://www.gnu.org/software/bash/manual/html_node/Bash-Builtins.html#index-compopt",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("D", "D", false, "change options for the default command completion")
	rootCmd.Flags().BoolS("E", "E", false, "change options for the empty command completion")
	rootCmd.Flags().BoolS("I", "I", false, "change options for completion on the initial word")
	rootCmd.Flags().StringS("o", "o", "", "set completion option")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"o": carapace.ActionValuesDescribed(
			"bashdefault", "perform default bash completions",
			"default", "use readline's default filename completion",
			"dirnames", "perform directory name completion",
			"filenames", "completions are filenames",
			"noargs", "do not complete arguments",
			"nobashdefault", "do not perform default bash completions",
			"nodefault", "do not use readline's default filename completion",
			"nodirnames", "do not perform directory name completion",
			"nofilenames", "completions are not filenames",
			"noquote", "do not quote completions",
			"nosort", "do not sort completions",
			"nospace", "do not append a space after completion",
			"plusdirs", "also perform directory name completion",
			"quote", "quote completions",
			"sort", "sort completions",
		),
	})
}
