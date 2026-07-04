package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fsutil",
	Short: "file system utility for performing tasks related to FAT and NTFS file systems",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/fsutil",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"8dot3name", "manage 8.3 short names",
			"behavior", "query and set volume behavior",
			"dirty", "query/set dirty bit",
			"file", "file specific operations",
			"fsinfo", "file system information",
			"hardlink", "hardlink management",
			"objectid", "object ID management",
			"quota", "quota management",
			"repair", "NTFS repair management",
			"reparsepoint", "reparse point management",
			"sparse", "sparse file control",
			"transaction", "transaction management",
			"usn", "USN journal management",
			"volume", "volume management",
		),
	)
}
