package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net/ssh"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nix-copy-closure",
	Short: "copy store objects to or from a remote machine via SSH",
	Long:  "https://nixos.org/manual/nix/stable/command-ref/nix-copy-closure.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("bzip2", false, "Compress the SSH connection using bzip2 (not implemented, falls back to gzip)")
	rootCmd.Flags().String("cores", "", "Maximum number of CPU cores to use inside a build")
	rootCmd.Flags().Bool("debug", false, "Set the logging verbosity level to 'debug'")
	rootCmd.Flags().Bool("dry-run", false, "Perform a dry run without actually copying")
	rootCmd.Flags().Bool("fallback", false, "Build from source if substitution fails")
	rootCmd.Flags().Bool("from", false, "Copy closure from remote machine to local machine")
	rootCmd.Flags().Bool("gzip", false, "Enable compression of the SSH connection")
	rootCmd.Flags().Bool("help", false, "Show usage information")
	rootCmd.Flags().Bool("include-outputs", false, "Also copy outputs of store derivations included in the closure")
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
	rootCmd.Flags().Bool("show-progress", false, "Show progress (not implemented)")
	rootCmd.Flags().String("store", "", "The URL of the Nix store to use")
	rootCmd.Flags().String("timeout", "", "Number of seconds before a build is killed")
	rootCmd.Flags().Bool("to", false, "Copy closure from local to remote machine")
	rootCmd.Flags().BoolP("use-substitutes", "s", false, "Attempt to download missing store objects on target from substituters")
	rootCmd.Flags().BoolP("verbose", "v", false, "Increase verbosity level")
	rootCmd.Flags().Bool("version", false, "Print the Nix version number")
	rootCmd.Flags().Bool("xz", false, "Compress the SSH connection using xz (not implemented, falls back to gzip)")

	rootCmd.Flag("option").Nargs = 2

	rootCmd.MarkFlagsMutuallyExclusive("to", "from")

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
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		ssh.ActionHosts(""),
	)
	carapace.Gen(rootCmd).PositionalAnyCompletion(
		nix.ActionPaths(),
	)
}
