package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var package_archiveSourceCmd = &cobra.Command{
	Use:   "archive-source",
	Short: "Create a source archive for the package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(package_archiveSourceCmd).Standalone()
	package_archiveSourceCmd.Flags().SetInterspersed(false)

	package_archiveSourceCmd.Flags().BoolP("help", "h", false, "Show help information")
	package_archiveSourceCmd.Flags().StringP("output", "o", "", "The absolute or relative path for the generated source archive")
	package_archiveSourceCmd.Flags().Bool("version", false, "Show the version")

	packageCmd.AddCommand(package_archiveSourceCmd)

	carapace.Gen(package_archiveSourceCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionFiles(),
	})
}
