package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nix-hash",
	Short: "compute the cryptographic hash of a path",
	Long:  "https://nixos.org/manual/nix/stable/command-ref/nix-hash.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("base16", false, "Print hash in hexadecimal representation")
	rootCmd.Flags().Bool("base32", false, "Print hash in base-32 representation")
	rootCmd.Flags().Bool("base64", false, "Print hash in base-64 representation")
	rootCmd.Flags().String("cores", "", "Maximum number of CPU cores to use inside a build")
	rootCmd.Flags().Bool("debug", false, "Set the logging verbosity level to 'debug'")
	rootCmd.Flags().Bool("fallback", false, "Build from source if substitution fails")
	rootCmd.Flags().Bool("flat", false, "Hash each regular file directly")
	rootCmd.Flags().Bool("help", false, "Show usage information")
	rootCmd.Flags().BoolP("keep-failed", "K", false, "Keep temporary directories of failed builds")
	rootCmd.Flags().BoolP("keep-going", "k", false, "Keep going after a build fails")
	rootCmd.Flags().String("log-format", "", "Set the format of log output")
	rootCmd.Flags().StringP("max-jobs", "j", "", "Maximum number of parallel builds")
	rootCmd.Flags().String("max-silent-time", "", "Number of seconds of silence before a build is killed")
	rootCmd.Flags().BoolP("no-build-output", "Q", false, "Do not show build output")
	rootCmd.Flags().Bool("no-gc-warning", false, "Disable warnings about not using --add-root")
	rootCmd.Flags().StringSlice("option", nil, "Set the Nix configuration setting name to value")
	rootCmd.Flags().Bool("quiet", false, "Decrease the logging verbosity level")
	rootCmd.Flags().Bool("readonly-mode", false, "Do not write to the Nix store")
	rootCmd.Flags().Bool("sri", false, "Print hash in SRI format")
	rootCmd.Flags().String("store", "", "The URL of the Nix store to use")
	rootCmd.Flags().String("timeout", "", "Number of seconds before a build is killed")
	rootCmd.Flags().Bool("to-base16", false, "Convert base-32 hash to hexadecimal")
	rootCmd.Flags().Bool("to-base32", false, "Convert hexadecimal hash to base-32")
	rootCmd.Flags().Bool("to-base64", false, "Convert hexadecimal hash to base-64")
	rootCmd.Flags().Bool("to-sri", false, "Convert hexadecimal hash to SRI format")
	rootCmd.Flags().Bool("truncate", false, "Truncate hashes longer than 160 bits")
	rootCmd.Flags().String("type", "", "Hash algorithm to use")
	rootCmd.Flags().BoolP("verbose", "v", false, "Increase the logging verbosity level")
	rootCmd.Flags().Bool("version", false, "Show version information")

	rootCmd.Flag("option").Nargs = 2

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"log-format": carapace.ActionValues("bar", "bar-with-logs", "internal-json", "raw"),
		"option": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return nix.ActionConfigKeys()
			case 1:
				return nix.ActionConfigValues(c.Parts[0])
			default:
				return carapace.ActionValues()
			}
		}),
		"type": carapace.ActionValues("blake3", "md5", "sha1", "sha256", "sha512"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		nix.ActionPaths(),
	)
}
