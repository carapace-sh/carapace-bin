package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nix-prefetch-url",
	Short: "download a file and print its hash into the Nix store",
	Long:  "https://nixos.org/manual/nix/stable/command-ref/nix-prefetch-url.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringSlice("arg", nil, "Pass [name] [expr] to Nix functions")
	rootCmd.Flags().StringSlice("arg-from-file", nil, "Pass the contents of file as the argument name to Nix functions")
	rootCmd.Flags().StringSlice("arg-from-stdin", nil, "Pass the contents of stdin as the argument name to Nix functions")
	rootCmd.Flags().StringSlice("argstr", nil, "Pass [name] [string] to Nix functions")
	rootCmd.Flags().StringP("attr", "A", "", "Fetch from a Nix expression attribute path")
	rootCmd.Flags().String("cores", "", "Maximum number of CPU cores to use inside a build")
	rootCmd.Flags().Bool("debug", false, "Set the logging verbosity level to 'debug'")
	rootCmd.Flags().String("eval-store", "", "The Nix store to use for evaluation")
	rootCmd.Flags().Bool("executable", false, "Set the executable bit on the downloaded file")
	rootCmd.Flags().Bool("fallback", false, "Build from source if substitution fails")
	rootCmd.Flags().Bool("help", false, "Show usage information")
	rootCmd.Flags().Bool("impure", false, "Allow access to mutable paths and repositories")
	rootCmd.Flags().StringP("include", "I", "", "Include paths")
	rootCmd.Flags().BoolP("keep-failed", "K", false, "Keep temporary directories of failed builds")
	rootCmd.Flags().BoolP("keep-going", "k", false, "Keep going after a build fails")
	rootCmd.Flags().String("log-format", "", "Set the format of log output")
	rootCmd.Flags().StringP("max-jobs", "j", "", "Maximum number of parallel builds")
	rootCmd.Flags().String("max-silent-time", "", "Number of seconds of silence before a build is killed")
	rootCmd.Flags().String("name", "", "Override the name of the file in the Nix store")
	rootCmd.Flags().BoolP("no-build-output", "Q", false, "Do not show build output")
	rootCmd.Flags().Bool("no-gc-warning", false, "Disable warnings about not using --add-root")
	rootCmd.Flags().StringSlice("option", nil, "Set the Nix configuration setting name to value")
	rootCmd.Flags().String("override-flake", "", "Override the flake registries, redirecting original-ref to resolved-ref")
	rootCmd.Flags().Bool("print-path", false, "Print the store path of the downloaded file")
	rootCmd.Flags().Bool("quiet", false, "Decrease the logging verbosity level")
	rootCmd.Flags().Bool("readonly-mode", false, "Do not write to the Nix store")
	rootCmd.Flags().Bool("repair", false, "During evaluation, rewrite missing or corrupted files in the Nix store")
	rootCmd.Flags().String("store", "", "The URL of the Nix store to use")
	rootCmd.Flags().String("timeout", "", "Number of seconds before a build is killed")
	rootCmd.Flags().String("type", "", "Use the specified cryptographic hash algorithm")
	rootCmd.Flags().Bool("unpack", false, "Unpack the archive and add the result to the Nix store")
	rootCmd.Flags().BoolP("verbose", "v", false, "Increase the logging verbosity level")
	rootCmd.Flags().Bool("version", false, "Show version information")

	rootCmd.Flag("arg").Nargs = 2
	rootCmd.Flag("arg-from-file").Nargs = 2
	rootCmd.Flag("arg-from-stdin").Nargs = 1
	rootCmd.Flag("argstr").Nargs = 2
	rootCmd.Flag("option").Nargs = 2
	rootCmd.Flag("override-flake").Nargs = 2

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"attr": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			arg, _ := rootCmd.Flags().GetStringSlice("arg")
			argstr, _ := rootCmd.Flags().GetStringSlice("argstr")
			opts := nix.AttributeOpts{
				Source:  "default.nix",
				Include: rootCmd.Flag("include").Value.String(),
				Arg:     arg,
				ArgStr:  argstr,
			}
			if len(c.Args) > 0 {
				opts.Source = c.Args[0]
			}
			return nix.ActionAttributes(opts)
		}),
		"include":    carapace.ActionFiles(),
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
		"override-flake": nix.ActionFlakeRefs(),
		"type":           carapace.ActionValues("blake3", "md5", "sha1", "sha256", "sha512"),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValues(), // url
		carapace.ActionValues(), // expected hash
	)
}
