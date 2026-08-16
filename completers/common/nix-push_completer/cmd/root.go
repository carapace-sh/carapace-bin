package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nix-push",
	Short: "push store paths to a binary cache",
	Long:  "https://nixos.org/manual/nix/stable/command-ref/nix-push.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("bzip2", false, "Use bzip2 compression")
	rootCmd.Flags().String("dest", "", "Destination directory")
	rootCmd.Flags().Bool("force", false, "Force overwriting of existing files")
	rootCmd.Flags().Bool("help", false, "Show usage information")
	rootCmd.Flags().Bool("link", false, "Create hard links instead of copying")
	rootCmd.Flags().String("manifest", "", "Path to manifest file")
	rootCmd.Flags().String("manifest-path", "", "Path to manifest file")
	rootCmd.Flags().Bool("none", false, "Use no compression")
	rootCmd.Flags().String("url-prefix", "", "URL prefix for the cache")

	rootCmd.MarkFlagsMutuallyExclusive("bzip2", "none")
	rootCmd.MarkFlagsMutuallyExclusive("manifest", "manifest-path")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"dest":          carapace.ActionDirectories(),
		"manifest":      carapace.ActionFiles(),
		"manifest-path": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		nix.ActionPaths(),
	)
}
