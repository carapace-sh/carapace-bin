package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var package_showExecutablesCmd = &cobra.Command{
	Use:   "show-executables",
	Short: "List the available executables from this package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(package_showExecutablesCmd).Standalone()
	package_showExecutablesCmd.Flags().SetInterspersed(false)

	package_showExecutablesCmd.Flags().String("format", "", "Set the output format")
	package_showExecutablesCmd.Flags().BoolP("help", "h", false, "Show help information")
	package_showExecutablesCmd.Flags().Bool("version", false, "Show the version")

	packageCmd.AddCommand(package_showExecutablesCmd)

	carapace.Gen(package_showExecutablesCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("flatlist", "json"),
	})
}
