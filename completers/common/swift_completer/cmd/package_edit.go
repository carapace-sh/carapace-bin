package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var package_editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Put a package in editable mode",
	Args:  cobra.ExactArgs(1),
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(package_editCmd).Standalone()
	package_editCmd.Flags().SetInterspersed(false)

	package_editCmd.Flags().String("branch", "", "The branch to create")
	package_editCmd.Flags().BoolP("help", "h", false, "Show help information")
	package_editCmd.Flags().String("path", "", "Create or use the checkout at this path")
	package_editCmd.Flags().String("revision", "", "The revision to edit")
	package_editCmd.Flags().Bool("version", false, "Show the version")

	packageCmd.AddCommand(package_editCmd)

	carapace.Gen(package_editCmd).FlagCompletion(carapace.ActionMap{
		"path": carapace.ActionDirectories(),
	})
}
