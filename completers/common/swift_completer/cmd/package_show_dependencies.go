package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var package_showDependenciesCmd = &cobra.Command{
	Use:   "show-dependencies",
	Short: "Print the resolved dependency graph",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(package_showDependenciesCmd).Standalone()
	package_showDependenciesCmd.Flags().SetInterspersed(false)

	package_showDependenciesCmd.Flags().String("format", "", "Set the output format")
	package_showDependenciesCmd.Flags().BoolP("help", "h", false, "Show help information")
	package_showDependenciesCmd.Flags().StringP("output-path", "o", "", "The absolute or relative path to output the resolved dependency graph")
	package_showDependenciesCmd.Flags().Bool("version", false, "Show the version")

	packageCmd.AddCommand(package_showDependenciesCmd)

	carapace.Gen(package_showDependenciesCmd).FlagCompletion(carapace.ActionMap{
		"format":      carapace.ActionValues("text", "dot", "json", "flatlist"),
		"output-path": carapace.ActionFiles(),
	})
}
