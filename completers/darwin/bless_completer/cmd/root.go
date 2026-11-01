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

	rootCmd.Flags().Bool("bootBlocks", false, "Bless with boot blocks")
	rootCmd.Flags().Bool("bootinfo", false, "Create boot info files")
	rootCmd.Flags().String("file", "", "Specify the file to boot")
	rootCmd.Flags().String("folder", "", "Specify the folder to bless")
	rootCmd.Flags().BoolP("help", "h", false, "Display usage information")
	rootCmd.Flags().Bool("mount", false, "Mount the volume")
	rootCmd.Flags().Bool("setBoot", false, "Set the boot volume")
	rootCmd.Flags().Bool("unmount", false, "Unmount the volume")
	rootCmd.Flags().BoolP("verbose", "v", false, "Verbose mode")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"file":   carapace.ActionFiles(),
		"folder": carapace.ActionDirectories(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionDirectories())
}
