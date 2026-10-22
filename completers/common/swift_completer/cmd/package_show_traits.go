package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var package_showTraitsCmd = &cobra.Command{
	Use:   "show-traits",
	Short: "List the available traits for a package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(package_showTraitsCmd).Standalone()
	package_showTraitsCmd.Flags().SetInterspersed(false)

	package_showTraitsCmd.Flags().String("format", "", "Set the output format")
	package_showTraitsCmd.Flags().BoolP("help", "h", false, "Show help information")
	package_showTraitsCmd.Flags().String("package-id", "", "Show traits for any package id in the transitive dependencies")
	package_showTraitsCmd.Flags().Bool("version", false, "Show the version")

	packageCmd.AddCommand(package_showTraitsCmd)

	carapace.Gen(package_showTraitsCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("text", "json"),
	})
}
