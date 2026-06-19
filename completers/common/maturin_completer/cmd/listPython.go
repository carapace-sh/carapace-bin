package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/rust"
	"github.com/spf13/cobra"
)

var listPythonCmd = &cobra.Command{
	Use:   "list-python",
	Short: "Search and list the available python installations",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listPythonCmd).Standalone()

	listPythonCmd.Flags().String("target", "", "Target triple to list interpreters for")
	listPythonCmd.Flags().CountP("verbose", "v", "Use verbose output")
	rootCmd.AddCommand(listPythonCmd)

	carapace.Gen(listPythonCmd).FlagCompletion(carapace.ActionMap{
		"target": rust.ActionTargets(),
	})
}
