package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nix-env",
	Short: "manage user environments",
	Long:  "https://nixos.org/manual/nix/stable/command-ref/nix-env.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("always", false, "Always upgrade a package, even if the version does not change")
	rootCmd.Flags().StringSlice("arg", nil, "Pass [name] [expr] to Nix functions")
	rootCmd.Flags().StringSlice("arg-from-file", nil, "Pass the contents of file as the argument name to Nix functions")
	rootCmd.Flags().StringSlice("arg-from-stdin", nil, "Pass the contents of stdin as the argument name to Nix functions")
	rootCmd.Flags().StringSlice("argstr", nil, "Pass [name] [string] to Nix functions")
	rootCmd.Flags().StringP("attr", "A", "", "Interpret args as attribute paths")
	rootCmd.Flags().Bool("attr-path", false, "Print attribute path")
	rootCmd.Flags().BoolP("available", "a", false, "Query available packages")
	rootCmd.Flags().BoolP("compare-versions", "c", false, "Compare installed vs available versions")
	rootCmd.Flags().String("cores", "", "Maximum number of CPU cores to use inside a build")
	rootCmd.Flags().Bool("debug", false, "Set the logging verbosity level to 'debug'")
	rootCmd.Flags().Bool("delete-generations", false, "Delete profile generations")
	rootCmd.Flags().Bool("description", false, "Print one-line description")
	rootCmd.Flags().Bool("drv-path", false, "Print store derivation path")
	rootCmd.Flags().Bool("dry-run", false, "Print what would be done without doing it")
	rootCmd.Flags().Bool("eq", false, "Upgrade only if the version is equal")
	rootCmd.Flags().String("eval-store", "", "The Nix store to use for evaluation")
	rootCmd.Flags().Bool("fallback", false, "Build from source if substitution fails")
	rootCmd.Flags().StringP("file", "f", "", "Path to Nix expression")
	rootCmd.Flags().String("force-name", "", "Force the name of the derivation")
	rootCmd.Flags().BoolP("from-expression", "E", false, "Interpret args as Nix expressions")
	rootCmd.Flags().String("from-profile", "", "Copy from another profile")
	rootCmd.Flags().Bool("help", false, "Show usage information")
	rootCmd.Flags().Bool("impure", false, "Allow access to mutable paths and repositories")
	rootCmd.Flags().StringP("include", "I", "", "Include paths")
	rootCmd.Flags().BoolP("install", "i", false, "Add packages to user environment")
	rootCmd.Flags().Bool("installed", false, "Operate on installed packages (default)")
	rootCmd.Flags().Bool("json", false, "JSON output")
	rootCmd.Flags().BoolP("keep-failed", "K", false, "Keep temporary directories of failed builds")
	rootCmd.Flags().BoolP("keep-going", "k", false, "Keep going after a build fails")
	rootCmd.Flags().Bool("leq", false, "Upgrade only if the version is less than or equal")
	rootCmd.Flags().Bool("list-generations", false, "List profile generations")
	rootCmd.Flags().String("log-format", "", "Set the format of log output")
	rootCmd.Flags().Bool("lt", false, "Upgrade only if the version is less than")
	rootCmd.Flags().StringP("max-jobs", "j", "", "Maximum number of parallel builds")
	rootCmd.Flags().String("max-silent-time", "", "Number of seconds of silence before a build is killed")
	rootCmd.Flags().Bool("meta", false, "Print all meta-attributes")
	rootCmd.Flags().BoolP("no-build-output", "Q", false, "Do not show build output")
	rootCmd.Flags().Bool("no-gc-warning", false, "Disable warnings about not using --add-root")
	rootCmd.Flags().Bool("no-name", false, "Suppress printing of name attribute")
	rootCmd.Flags().StringSlice("option", nil, "Set the Nix configuration setting name to value")
	rootCmd.Flags().Bool("out-path", false, "Print output path")
	rootCmd.Flags().String("override-flake", "", "Override the flake registries, redirecting original-ref to resolved-ref")
	rootCmd.Flags().BoolP("prebuilt-only", "b", false, "Only use derivations with pre-built binaries")
	rootCmd.Flags().BoolP("preserve-installed", "P", false, "Don't remove derivations with matching names")
	rootCmd.Flags().Int("priority", 0, "Set priority of derivations being installed")
	rootCmd.Flags().StringP("profile", "p", "", "Path to profile")
	rootCmd.Flags().BoolP("query", "q", false, "Display information about packages")
	rootCmd.Flags().Bool("quiet", false, "Decrease the logging verbosity level")
	rootCmd.Flags().Bool("readonly-mode", false, "Do not write to the Nix store")
	rootCmd.Flags().BoolP("remove-all", "r", false, "Remove all previously installed packages first")
	rootCmd.Flags().Bool("repair", false, "During evaluation, rewrite missing or corrupted files in the Nix store")
	rootCmd.Flags().Bool("rollback", false, "Switch to previous generation")
	rootCmd.Flags().Bool("set", false, "Set profile to contain exactly one derivation")
	rootCmd.Flags().Bool("set-flag", false, "Modify meta attributes of installed packages")
	rootCmd.Flags().BoolP("status", "s", false, "Print status")
	rootCmd.Flags().String("store", "", "The URL of the Nix store to use")
	rootCmd.Flags().BoolP("switch-generation", "G", false, "Switch to a specific generation")
	rootCmd.Flags().BoolP("switch-profile", "S", false, "Set user environment to given profile")
	rootCmd.Flags().Bool("system", false, "Print system attribute")
	rootCmd.Flags().String("system-filter", "", "Filter derivations by platform")
	rootCmd.Flags().String("timeout", "", "Number of seconds before a build is killed")
	rootCmd.Flags().BoolP("uninstall", "e", false, "Remove packages from user environment")
	rootCmd.Flags().BoolP("upgrade", "u", false, "Upgrade packages in user environment")
	rootCmd.Flags().BoolP("verbose", "v", false, "Increase the logging verbosity level")
	rootCmd.Flags().Bool("version", false, "Show version information")
	rootCmd.Flags().Bool("xml", false, "XML output")

	rootCmd.Flag("arg").Nargs = 2
	rootCmd.Flag("arg-from-file").Nargs = 2
	rootCmd.Flag("arg-from-stdin").Nargs = 1
	rootCmd.Flag("argstr").Nargs = 2
	rootCmd.Flag("option").Nargs = 2
	rootCmd.Flag("override-flake").Nargs = 2
	rootCmd.Flag("set-flag").Nargs = 2

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"file":         carapace.ActionFiles(".nix"),
		"from-profile": carapace.ActionFiles(),
		"include":      carapace.ActionFiles(),
		"log-format":   carapace.ActionValues("raw", "internal-json", "bar", "bar-with-logs"),
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
		"profile":        carapace.ActionFiles(),
		"system-filter":  nix.ActionSystems(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		nix.ActionChannelPackages(),
	)
}
