package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dependsCmd = &cobra.Command{
	Use:   "depends",
	Short: "visualize tox environment dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dependsCmd).Standalone()

	add_common_flags(dependsCmd)

	dependsCmd.Flags().String("result-json", "", "write a JSON file with detailed information about all commands and results involved (default: None)")
	dependsCmd.Flags().String("hashseed", "", "set PYTHONHASHSEED to SEED before running commands. Defaults to a random integer in the range [1, 4294967295] ([1, 1024] on Windows). Passing 'notset' suppresses this behavior. (default: 264197440)")
	dependsCmd.Flags().StringArray("discover", []string{}, "for Python discovery first try these Python executables (default: [])")
	dependsCmd.Flags().Bool("list-dependencies", false, "list the dependencies installed during environment setup (default: False)")
	dependsCmd.Flags().Bool("no-list-dependencies", false, "never list the dependencies installed during environment setup (default: True)")

	rootCmd.AddCommand(dependsCmd)
}
