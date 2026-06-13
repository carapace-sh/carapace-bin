package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/shell"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "builtin",
	Short: "Run a builtin command",
	Long:  "https://fishshell.com/docs/current/cmds/builtin.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("n", "n", false, "list all defined builtins")
	rootCmd.Flags().Bool("names", false, "list all defined builtins")
	rootCmd.Flags().BoolS("q", "q", false, "query builtins")
	rootCmd.Flags().Bool("query", false, "query builtins")
	rootCmd.Flags().BoolS("h", "h", false, "display help")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		shell.ActionBuiltins(),
	)
}
