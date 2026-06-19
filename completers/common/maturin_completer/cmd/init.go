package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/maturin_completer/cmd/action"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a new cargo project in an existing directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(initCmd).Standalone()

	initCmd.Flags().StringP("bindings", "b", "", "Which kind of bindings to use")
	initCmd.Flags().Bool("mixed", false, "Use mixed Rust/Python project layout")
	initCmd.Flags().String("name", "", "Set the resulting package name, defaults to the directory name")
	initCmd.Flags().Bool("src", false, "Use Python first src layout for mixed Rust/Python project")
	initCmd.Flags().CountP("verbose", "v", "Use verbose output")
	rootCmd.AddCommand(initCmd)

	carapace.Gen(initCmd).FlagCompletion(carapace.ActionMap{
		"bindings": action.ActionProjectBindings(),
	})

	carapace.Gen(initCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
