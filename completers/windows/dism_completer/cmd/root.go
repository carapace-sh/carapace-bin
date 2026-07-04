package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dism",
	Short: "deployment image servicing and management",
	Long:  "https://learn.microsoft.com/en-us/windows-hardware/manufacture/desktop/what-is-dism",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"Get-WimInfo", "display information about a Windows image file",
			"Get-Features", "display all features in a package",
			"Get-Packages", "display all packages in an image",
			"Add-Package", "add a package to an image",
			"Remove-Package", "remove a package from an image",
			"Enable-Feature", "enable a feature in an image",
			"Disable-Feature", "disable a feature in an image",
			"Cleanup-Image", "perform cleanup and recovery on an image",
			"Mount-Image", "mount an image to a directory",
			"Unmount-Image", "unmount a mounted image",
			"Apply-Image", "apply an image to a directory",
			"Capture-Image", "capture an image from a directory",
			"Split-Image", "split an image file into smaller files",
			"Export-Image", "export an image to a new file",
		),
	)
}
