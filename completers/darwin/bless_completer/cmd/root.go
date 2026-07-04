package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bless",
	Short: "set volume bootable and startup options",
	Long:  "https://keith.github.io/xcode-manpages/bless.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("bootBlocks", "bootBlocks", false, "Bless with boot blocks")
	rootCmd.Flags().BoolP("bootinfo", "bootinfo", false, "Create boot info files")
	rootCmd.Flags().StringP("file", "file", "", "Specify the file to boot")
	rootCmd.Flags().StringP("folder", "folder", "", "Specify the folder to bless")
	rootCmd.Flags().BoolP("help", "h", false, "Display usage information")
	rootCmd.Flags().BoolP("mount", "mount", false, "Mount the volume")
	rootCmd.Flags().BoolP("setBoot", "setBoot", false, "Set the boot volume")
	rootCmd.Flags().BoolP("unmount", "unmount", false, "Unmount the volume")
	rootCmd.Flags().BoolP("verbose", "v", false, "Verbose mode")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"file":   carapace.ActionFiles(),
		"folder": carapace.ActionDirectories(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionDirectories())
}
