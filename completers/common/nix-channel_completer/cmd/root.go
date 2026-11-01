package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nix-channel",
	Short: "manage Nix channels",
	Long:  "https://nixos.org/manual/nix/stable/command-ref/nix-channel.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("add", false, "Adds a channel named name with URL url to the list of subscribed channels")
	rootCmd.Flags().String("cores", "", "Maximum number of CPU cores to use inside a build")
	rootCmd.Flags().Bool("debug", false, "Set the logging verbosity level to 'debug'")
	rootCmd.Flags().Bool("fallback", false, "Build from source if substitution fails")
	rootCmd.Flags().Bool("help", false, "Show usage information")
	rootCmd.Flags().BoolP("keep-failed", "K", false, "Keep temporary directories of failed builds")
	rootCmd.Flags().BoolP("keep-going", "k", false, "Keep going after a build fails")
	rootCmd.Flags().Bool("list", false, "Prints the names and URLs of all subscribed channels on standard output")
	rootCmd.Flags().Bool("list-generations", false, "List the generations of the channels profile")
	rootCmd.Flags().String("log-format", "", "Set the format of log output")
	rootCmd.Flags().StringP("max-jobs", "j", "", "Maximum number of parallel builds")
	rootCmd.Flags().String("max-silent-time", "", "Number of seconds of silence before a build is killed")
	rootCmd.Flags().BoolP("no-build-output", "Q", false, "Do not show build output")
	rootCmd.Flags().Bool("no-gc-warning", false, "Disable warnings about not using --add-root")
	rootCmd.Flags().StringSlice("option", nil, "Set the Nix configuration setting name to value")
	rootCmd.Flags().Bool("quiet", false, "Decrease the logging verbosity level")
	rootCmd.Flags().Bool("readonly-mode", false, "Do not write to the Nix store")
	rootCmd.Flags().Bool("remove", false, "Removes the channel named name from the list of subscribed channels")
	rootCmd.Flags().Bool("rollback", false, "Reverts the previous call to nix-channel --update")
	rootCmd.Flags().String("store", "", "The URL of the Nix store to use")
	rootCmd.Flags().String("timeout", "", "Number of seconds before a build is killed")
	rootCmd.Flags().Bool("update", false, "Downloads the Nix expressions of all subscribed channels")
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
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if (rootCmd.Flag("remove").Changed && len(c.Args) == 0) ||
				rootCmd.Flag("update").Changed {
				return nix.ActionLocalChannels().FilterArgs()
			} else if rootCmd.Flag("add").Changed {
				if len(c.Args) == 0 {
					return nix.ActionRemoteChannels().Prefix("https://nixos.org/channels/")
				}
			}
			return carapace.ActionValues()
		}),
	)
}
