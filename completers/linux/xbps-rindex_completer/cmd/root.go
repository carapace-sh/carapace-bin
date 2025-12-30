package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xbps-rindex [OPTIONS] MODE ARGUMENTS",
	Short: "XBPS utility to manage local binary package repositories",
	Long:  "https://man.voidlinux.org/xbps-remove",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringSliceP("add", "a", nil, "Add package(s) to repository index")
	rootCmd.Flags().StringP("clean", "c", "", "Clean repository index")
	rootCmd.Flags().String("compression", "zstd", "Compression format: none, gzip, bzip2, lz4, xz, zstd (default)")
	rootCmd.Flags().BoolP("debug", "d", false, "Debug mode shown to stderr")
	rootCmd.Flags().BoolP("force", "", false, "Force mode to overwrite entry in add mode")
	rootCmd.Flags().BoolP("hashcheck", "C", false, "Consider file hashes for cleaning up packages")
	rootCmd.Flags().BoolP("help", "h", false, "Show usage")
	rootCmd.Flags().String("privkey", "", "Path to the private key for signing")
	rootCmd.Flags().StringP("remove-obsoletes", "r", "", "Removes obsolete packages from repository")
	rootCmd.Flags().StringP("sign", "s", "", "Initialize repository metadata signature")
	rootCmd.Flags().StringSliceP("sign-pkg", "S", nil, "Sign binary package archive")
	rootCmd.Flags().String("signedby", "", "Signature details, i.e \"name <email>\"")
	rootCmd.Flags().BoolP("verbose", "v", false, "Show XBPS version")

	rootCmd.MarkFlagsMutuallyExclusive("add", "clean", "remove-obsoletes", "sign", "sign-pkg")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"add":              carapace.ActionFiles(".xbps"),
		"compression":      carapace.ActionValues("none", "gzip", "bzip2", "lz4", "xz", "zstd"),
		"privkey":          carapace.ActionFiles(),
		"remove-obsoletes": carapace.ActionDirectories(),
		"sign":             carapace.ActionDirectories(),
		"sign-pkg":         carapace.ActionFiles(".xbps"),
	})
}
