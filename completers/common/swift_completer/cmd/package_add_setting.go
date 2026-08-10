package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var package_addSettingCmd = &cobra.Command{
	Use:   "add-setting",
	Short: "Add a new setting to the manifest",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(package_addSettingCmd).Standalone()
	package_addSettingCmd.Flags().SetInterspersed(false)

	package_addSettingCmd.Flags().BoolP("help", "h", false, "Show help information")
	package_addSettingCmd.Flags().String("swift", "", "The Swift language setting(s) to add")
	package_addSettingCmd.Flags().String("target", "", "The target to add the setting to")
	package_addSettingCmd.Flags().Bool("version", false, "Show the version")

	packageCmd.AddCommand(package_addSettingCmd)
}
