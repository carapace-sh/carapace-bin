package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nixos-generate-config",
	Short: "generate a NixOS configuration",
	Long:  "https://nixos.org/manual/nixos/stable/configuration.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("dir", "", "Directory to write the configuration to")
	rootCmd.Flags().Bool("force", false, "Overwrite existing configuration files")
	rootCmd.Flags().Bool("help", false, "Show usage information")
	rootCmd.Flags().Bool("no-filesystems", false, "Do not include filesystem information")
	rootCmd.Flags().String("root", "", "Root directory to generate configuration for")
	rootCmd.Flags().Bool("show-hardware-config", false, "Print the hardware configuration to stdout")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"dir":  carapace.ActionDirectories(),
		"root": carapace.ActionDirectories(),
	})
}