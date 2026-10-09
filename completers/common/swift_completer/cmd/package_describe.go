package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var package_describeCmd = &cobra.Command{
	Use:   "describe",
	Short: "Describe the current package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(package_describeCmd).Standalone()
	package_describeCmd.Flags().SetInterspersed(false)

	package_describeCmd.Flags().BoolP("help", "h", false, "Show help information")
	package_describeCmd.Flags().String("type", "", "Set the output format")
	package_describeCmd.Flags().Bool("version", false, "Show the version")

	packageCmd.AddCommand(package_describeCmd)

	carapace.Gen(package_describeCmd).FlagCompletion(carapace.ActionMap{
		"type": carapace.ActionValues("json", "text", "mermaid"),
	})
}
