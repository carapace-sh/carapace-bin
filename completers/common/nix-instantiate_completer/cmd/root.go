package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nix-instantiate",
	Short: "instantiate store derivations from Nix expression",
	Long:  "https://nixos.org/manual/nix/stable/command-ref/nix-instantiate.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("add-root", "", "Create a symlink to the output path")
	rootCmd.Flags().StringSlice("arg", nil, "Pass [name] [expr] to Nix functions.")
	rootCmd.Flags().StringSlice("arg-from-file", nil, "Pass the contents of file as the argument name to Nix functions")
	rootCmd.Flags().StringSlice("arg-from-stdin", nil, "Pass the contents of stdin as the argument name to Nix functions")
	rootCmd.Flags().StringSlice("argstr", nil, "Pass [name] [string] to Nix functions")
	rootCmd.Flags().StringP("attr", "A", "", "Attribute to build")
	rootCmd.Flags().String("cores", "", "Maximum number of CPU cores to use inside a build")
	rootCmd.Flags().Bool("debug", false, "Set the logging verbosity level to 'debug'")
	rootCmd.Flags().Bool("dry-run", false, "Show what store paths would be built or downloaded")
	rootCmd.Flags().Bool("eval", false, "Evaluate the input files, and print the resulting values on standard output")
	rootCmd.Flags().Bool("eval-only", false, "Evaluate the input files, and print the resulting values on standard output")
	rootCmd.Flags().String("eval-store", "", "The Nix store to use for evaluation")
	rootCmd.Flags().BoolP("expr", "E", false, "Interpret command line arguments as Nix expressions")
	rootCmd.Flags().Bool("fallback", false, "Build from source if substitution fails")
	rootCmd.Flags().Bool("find-file", false, "Lookup files in the nix search path")
	rootCmd.Flags().Bool("help", false, "Show usage information")
	rootCmd.Flags().Bool("impure", false, "Allow access to mutable paths and repositories")
	rootCmd.Flags().StringP("include", "I", "", "Include paths")
	rootCmd.Flags().Bool("indirect", false, "Create the symlink indirectly (obsolete)")
	rootCmd.Flags().Bool("json", false, "Print the result of --eval as JSON")
	rootCmd.Flags().BoolP("keep-failed", "K", false, "Keep temporary directories of failed builds")
	rootCmd.Flags().BoolP("keep-going", "k", false, "Keep going after a build fails")
	rootCmd.Flags().String("log-format", "", "Set the format of log output")
	rootCmd.Flags().StringP("max-jobs", "j", "", "Maximum number of parallel builds")
	rootCmd.Flags().String("max-silent-time", "", "Number of seconds of silence before a build is killed")
	rootCmd.Flags().BoolP("no-build-output", "Q", false, "Do not show build output")
	rootCmd.Flags().Bool("no-gc-warning", false, "Disable warnings about not using --add-root")
	rootCmd.Flags().Bool("no-location", false, "Do not print source location in XML output")
	rootCmd.Flags().StringSlice("option", nil, "Set the Nix configuration setting name to value")
	rootCmd.Flags().String("override-flake", "", "Override the flake registries, redirecting original-ref to resolved-ref")
	rootCmd.Flags().Bool("parse", false, "Just parse the input files, and print their abstract syntax trees on standard output in ATerm format.")
	rootCmd.Flags().Bool("parse-only", false, "Just parse the input files, and print their abstract syntax trees on standard output in ATerm format.")
	rootCmd.Flags().Bool("quiet", false, "Decrease the logging verbosity level")
	rootCmd.Flags().Bool("raw", false, "Print strings without quotes or escaping")
	rootCmd.Flags().Bool("read-write-mode", false, "Perform evaluation in read/write mode")
	rootCmd.Flags().Bool("readonly-mode", false, "Do not write to the Nix store")
	rootCmd.Flags().Bool("repair", false, "During evaluation, rewrite missing or corrupted files in the Nix store")
	rootCmd.Flags().String("store", "", "The URL of the Nix store to use")
	rootCmd.Flags().Bool("strict", false, "Recursively evaluate list elements and attributes.")
	rootCmd.Flags().String("timeout", "", "Number of seconds before a build is killed")
	rootCmd.Flags().BoolP("verbose", "v", false, "Increase the logging verbosity level")
	rootCmd.Flags().Bool("version", false, "Show version information")
	rootCmd.Flags().Bool("xml", false, "Print the result of --eval as XML")

	rootCmd.Flag("arg").Nargs = 2
	rootCmd.Flag("arg-from-file").Nargs = 2
	rootCmd.Flag("arg-from-stdin").Nargs = 1
	rootCmd.Flag("argstr").Nargs = 2
	rootCmd.Flag("option").Nargs = 2
	rootCmd.Flag("override-flake").Nargs = 2

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"add-root": carapace.ActionFiles(),
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
		"log-format": carapace.ActionValues("raw", "internal-json", "bar", "bar-with-logs"),
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
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		nix.ActionPaths(),
	)

}
