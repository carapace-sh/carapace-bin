package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nixops",
	Short: "NixOps deployment tool",
	Long:  "https://nixos.org/nixops/manual/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().Bool("confirm", false, "Confirm deployment")
	rootCmd.PersistentFlags().String("cores", "", "Set NIX_BUILD_CORES environment variable in builders")
	rootCmd.PersistentFlags().Bool("debug", false, "Turn on debugging output")
	rootCmd.PersistentFlags().StringP("deployment", "d", "", "Deployment to use")
	rootCmd.PersistentFlags().Bool("fallback", false, "Fall back on installation from source")
	rootCmd.PersistentFlags().Bool("help", false, "Show usage information")
	rootCmd.PersistentFlags().StringP("include", "I", "", "Append a directory to the Nix search path")
	rootCmd.PersistentFlags().BoolP("keep-failed", "K", false, "Keep temporary directories of failed builds")
	rootCmd.PersistentFlags().Bool("keep-going", false, "Keep going after failed builds")
	rootCmd.PersistentFlags().StringP("max-jobs", "j", "", "Set maximum number of concurrent Nix builds")
	rootCmd.PersistentFlags().String("network", "", "Path to a directory containing either nixops.nix or flake.nix")
	rootCmd.PersistentFlags().Bool("no-build-output", false, "Suppress output written by builders")
	rootCmd.PersistentFlags().StringSlice("option", nil, "Set a Nix option")
	rootCmd.PersistentFlags().Bool("read-only-mode", false, "Run Nix evaluations in read-only mode")
	rootCmd.PersistentFlags().Bool("show-trace", false, "Print a Nix stack trace if evaluation fails")

	rootCmd.Flags().Bool("pdb", false, "Invoke pdb on unhandled exception")
	rootCmd.Flags().Bool("version", false, "Print NixOps's version number")

	rootCmd.Flag("option").Nargs = 2

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"network": carapace.ActionDirectories(),
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
	})
}
