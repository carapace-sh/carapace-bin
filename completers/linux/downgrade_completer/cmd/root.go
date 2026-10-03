package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pacman"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "downgrade",
	Short: "Downgrade Arch Linux packages",
	Long:  "https://github.com/pbrisbin/downgrade",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("ala-only", false, "only use ALA server")
	rootCmd.Flags().String("ala-url", "", "location of ALA server")
	rootCmd.Flags().Bool("cached-only", false, "only use cached packages")
	rootCmd.Flags().BoolP("help", "h", false, "show help script")
	rootCmd.Flags().String("ignore", "", "whether to add packages to IgnorePkg")
	rootCmd.Flags().Bool("latest", false, "pick latest matching version")
	rootCmd.Flags().String("maxdepth", "", "maximum depth to search for cached packages")
	rootCmd.Flags().Bool("oldest", false, "pick oldest matching version")
	rootCmd.Flags().String("pacman", "", "pacman command to use")
	rootCmd.Flags().String("pacman-cache", "", "pacman cache directory")
	rootCmd.Flags().String("pacman-conf", "", "pacman configuration file")
	rootCmd.Flags().String("pacman-log", "", "pacman log file")
	rootCmd.Flags().Bool("prefer-cache", false, "do not query ALA if a matching package was found in cache")
	rootCmd.Flags().Bool("unignore", false, "remove packages from IgnorePkg")
	rootCmd.Flags().Bool("version", false, "show downgrade version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"ignore":       carapace.ActionValues("prompt", "always", "never"),
		"pacman-cache": carapace.ActionDirectories(),
		"pacman-conf":  carapace.ActionFiles(),
		"pacman-log":   carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		pacman.ActionPackages().FilterArgs(),
	)
}
