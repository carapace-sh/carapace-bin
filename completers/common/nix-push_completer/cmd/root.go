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

	rootCmd.Flags().Bool("bzip2", false, "Compress NARs using bzip2 instead of xz -9")
	rootCmd.Flags().String("dest", "", "Destination directory")
	rootCmd.Flags().Bool("force", false, "Overwrite .narinfo files if they already exist")
	rootCmd.Flags().Bool("help", false, "Show usage information")
	rootCmd.Flags().Bool("link", false, "Hard link files into the destination directory rather than copying")
	rootCmd.Flags().Bool("manifest", false, "Force the generation of a manifest suitable for use by nix-pull")
	rootCmd.Flags().String("manifest-path", "", "Like --manifest, but specify the manifest filename")
	rootCmd.Flags().Bool("none", false, "Do not compress")
	rootCmd.Flags().String("url-prefix", "", "Specify the prefix URL used in the Manifest")

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
