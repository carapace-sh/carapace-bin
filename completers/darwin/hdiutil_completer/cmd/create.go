package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a disk image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(createCmd).Standalone()

	createCmd.Flags().String("encryption", "", "Specify encryption: AES-128 or AES-256")
	createCmd.Flags().String("filesystem", "", "Filesystem type: HFS+, JHFS+, APFS, UFS, MS-DOS")
	createCmd.Flags().String("format", "", "Image format: UDZO, UDCO, UDRO, UDRW, UDSP, UDBZ")
	createCmd.Flags().Bool("plist", false, "Provide result output in plist format")
	createCmd.Flags().Bool("quiet", false, "Close stdout and stderr")
	createCmd.Flags().String("size", "", "Specify the size of the image")
	createCmd.Flags().String("srcfolder", "", "Source folder to create the image from")
	createCmd.Flags().Bool("verbose", false, "Produce extra progress output")
	createCmd.Flags().String("volname", "", "Volume name")
	rootCmd.AddCommand(createCmd)

	carapace.Gen(createCmd).FlagCompletion(carapace.ActionMap{
		"encryption": carapace.ActionValues("AES-128", "AES-256"),
		"filesystem": carapace.ActionValues("HFS+", "JHFS+", "APFS", "UFS", "MS-DOS", "FAT32"),
		"format":     carapace.ActionValues("UDZO", "UDCO", "UDRO", "UDRW", "UDSP", "UDBZ"),
		"srcfolder":  carapace.ActionDirectories(),
	})

	carapace.Gen(createCmd).PositionalCompletion(carapace.ActionFiles())
}
