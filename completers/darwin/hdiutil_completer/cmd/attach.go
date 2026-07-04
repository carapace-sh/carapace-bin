package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var attachCmd = &cobra.Command{
	Use:   "attach",
	Short: "Attach a disk image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(attachCmd).Standalone()

	attachCmd.Flags().Bool("autoopen", false, "Auto-open volumes in the Finder")
	attachCmd.Flags().Bool("autoopenro", false, "Auto-open read-only volumes in the Finder")
	attachCmd.Flags().Bool("autoopenrw", false, "Auto-open read/write volumes in the Finder")
	attachCmd.Flags().Bool("debug", false, "Be very verbose")
	attachCmd.Flags().String("drivekey", "", "Specify a drive key=value pair")
	attachCmd.Flags().String("encryption", "", "Specify encryption: AES-128 or AES-256")
	attachCmd.Flags().Bool("ignorebadchecksums", false, "Ignore bad checksums")
	attachCmd.Flags().Bool("insecurehttp", false, "Ignore SSL host validation failures")
	attachCmd.Flags().Bool("kernel", false, "Attach without a helper process")
	attachCmd.Flags().String("mount", "", "Mount behavior: required, optional, or suppressed")
	attachCmd.Flags().String("mountpoint", "", "Mount the volume at the given path")
	attachCmd.Flags().Bool("noautoopen", false, "Do not auto-open volumes")
	attachCmd.Flags().Bool("nobrowse", false, "Render volumes invisible in Finder")
	attachCmd.Flags().Bool("nokernel", false, "Attach with a helper process")
	attachCmd.Flags().Bool("notremovable", false, "Prevent the image from being detached")
	attachCmd.Flags().Bool("noverify", false, "Do not verify the image")
	attachCmd.Flags().String("owners", "", "Specify owners on filesystems: on or off")
	attachCmd.Flags().Bool("plist", false, "Provide result output in plist format")
	attachCmd.Flags().Bool("quiet", false, "Close stdout and stderr")
	attachCmd.Flags().Bool("readonly", false, "Force the resulting device to be read-only")
	attachCmd.Flags().Bool("readwrite", false, "Attempt to override read-only flag")
	attachCmd.Flags().String("shadow", "", "Use a shadow file for read-write access")
	attachCmd.Flags().Bool("verbose", false, "Produce extra progress output")
	attachCmd.Flags().Bool("verify", false, "Verify the image before attaching")
	rootCmd.AddCommand(attachCmd)

	carapace.Gen(attachCmd).FlagCompletion(carapace.ActionMap{
		"encryption": carapace.ActionValues("AES-128", "AES-256"),
		"mount":      carapace.ActionValues("required", "optional", "suppressed"),
		"owners":     carapace.ActionValues("on", "off"),
		"shadow":     carapace.ActionFiles(),
	})

	carapace.Gen(attachCmd).PositionalCompletion(carapace.ActionFiles())
}
