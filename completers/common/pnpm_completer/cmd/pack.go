package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var packCmd = &cobra.Command{
	Use:   "pack",
	Short: "Create a tarball from a package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(packCmd).Standalone()

	packCmd.Flags().Bool("dry-run", false, "Do everything `pack` would do except writing the tarball to disk")
	packCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	packCmd.Flags().Bool("json", false, "Print the packed tarball and its contents in JSON")
	packCmd.Flags().Bool("no-skip-manifest-obfuscation", false, "Apply pnpm's normal packed-manifest filtering")
	packCmd.Flags().String("out", "", "Customize the output path. `%s` expands to the package name and `%v` to the version, e.g. `%s.tgz` or `some-dir/%s-%v.tgz`")
	packCmd.Flags().String("pack-destination", "", "Directory in which to save the tarball. Defaults to the current working directory")
	packCmd.Flags().String("pack-gzip-level", "", "gzip compression level (`0`–`9`) for the tarball")
	packCmd.Flags().Bool("skip-manifest-obfuscation", false, "Keep the original `packageManager` field and publish-lifecycle scripts in the packed manifest instead of stripping them")
	packCmd.Flag("no-skip-manifest-obfuscation").Hidden = true
	rootCmd.AddCommand(packCmd)
}
