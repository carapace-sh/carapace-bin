package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var package_addProductCmd = &cobra.Command{
	Use:   "add-product",
	Short: "Add a new product to the manifest",
	Args:  cobra.ExactArgs(1),
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(package_addProductCmd).Standalone()
	package_addProductCmd.Flags().SetInterspersed(false)

	package_addProductCmd.Flags().BoolP("help", "h", false, "Show help information")
	package_addProductCmd.Flags().StringArray("targets", nil, "A list of targets that are part of this product")
	package_addProductCmd.Flags().String("type", "", "The type of target to add")
	package_addProductCmd.Flags().Bool("version", false, "Show the version")

	packageCmd.AddCommand(package_addProductCmd)

	carapace.Gen(package_addProductCmd).FlagCompletion(carapace.ActionMap{
		"type": carapace.ActionValues("executable", "library", "static-library", "dynamic-library", "plugin"),
	})
}
