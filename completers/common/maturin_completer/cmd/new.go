package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/maturin_completer/cmd/action"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new cargo project",
	Args:  cobra.ExactArgs(1),
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(newCmd).Standalone()

	newCmd.Flags().StringP("bindings", "b", "", "Which kind of bindings to use")
	newCmd.Flags().Bool("mixed", false, "Use mixed Rust/Python project layout")
	newCmd.Flags().String("name", "", "Set the resulting package name, defaults to the directory name")
	newCmd.Flags().Bool("src", false, "Use Python first src layout for mixed Rust/Python project")
	newCmd.Flags().CountP("verbose", "v", "Use verbose output")
	rootCmd.AddCommand(newCmd)

	carapace.Gen(newCmd).FlagCompletion(carapace.ActionMap{
		"bindings": action.ActionProjectBindings(),
	})

	carapace.Gen(newCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
