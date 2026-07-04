package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "createinstallmedia",
	Short: "Create a bootable installer for macOS",
	Long:  "https://support.apple.com/en-us/101578",
	Run:   func(*cobra.Command, []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("applicationpath", "", "path to the installer app (deprecated in macOS 10.14+)")
	rootCmd.Flags().Bool("nointeraction", false, "erase the disk and create the installer without prompting")
	rootCmd.Flags().String("volume", "", "path to the target volume to erase and make bootable")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"applicationpath": carapace.ActionDirectories(),
		"volume":          carapace.ActionDirectories(),
	})
}
