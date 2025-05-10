package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nixos-rebuild",
	Short: "reconfigure a NixOS machine",
	Long:  "https://nixos.org",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().String("build-host", "localhost", "Specify host to perform the rebuild with")
	rootCmd.PersistentFlags().StringArray("builders", nil, "Specify ad-hoc remote builders")
	rootCmd.PersistentFlags().String("commit-lockfile-summary", "", "Set the commit-lockfile-summary setting")
	rootCmd.PersistentFlags().Bool("fallback", false, "Enable the fallback setting")
	rootCmd.PersistentFlags().Bool("fast", false, "Do not rebuild the 'nixUnstable' nixpkgs attribute before rebuilding")
	rootCmd.PersistentFlags().String("flake", "", "Build the NixOS system from specified flake")
	rootCmd.PersistentFlags().Bool("help", false, "Show usage information")
	rootCmd.PersistentFlags().Bool("install-bootloader", false, "Causes the bootloader to be (re)installed")
	rootCmd.PersistentFlags().Bool("keep-failed", false, "Enable the keep-failed setting")
	rootCmd.PersistentFlags().Bool("keep-going", false, "Enable the keep-going setting")
	rootCmd.PersistentFlags().String("log-format", "", "Set the format of log output")
	rootCmd.PersistentFlags().String("max-jobs", "", "Set the max-jobs setting")
	rootCmd.PersistentFlags().Bool("no-build-nix", false, "Do not rebuild the 'nixUnstable' nixpkgs attribute before rebuilding")
	rootCmd.PersistentFlags().BoolP("no-build-output", "-Q", false, "Do not output on stdout or stderr")
	rootCmd.PersistentFlags().Bool("offline", false, "Disable substituters and consider all previously downloaded files up-to-date")
	rootCmd.PersistentFlags().StringSlice("option", nil, "Set the Nix configuration setting name to value")
	rootCmd.PersistentFlags().BoolP("print-build-logs", "L", false, "Print full build logs on standard error")
	rootCmd.PersistentFlags().StringP("profile-name", "p", "", "Specify Nix profile to place the new configuration in")
	rootCmd.PersistentFlags().Bool("quiet", false, "Decrease the logging verbosity level")
	rootCmd.PersistentFlags().Bool("refresh", false, "Consider all previously downloaded files out-of-date")
	rootCmd.PersistentFlags().Bool("rollback", false, "Rollback to previous configuration instead of rebuilding")
	rootCmd.PersistentFlags().Bool("show-trace", false, "Enable the show-trace setting")
	rootCmd.PersistentFlags().StringP("specialisation", "c", "", "Activates given specialisation")
	rootCmd.PersistentFlags().String("target-host", "localhost", "Specify host to apply rebuilt configuration to")
	rootCmd.PersistentFlags().Bool("upgrade", false, "Update the root user's channel named 'nixos' before rebuilding")
	rootCmd.PersistentFlags().Bool("upgrade-all", false, "Update all of the root user's channels before rebuilding")
	rootCmd.PersistentFlags().Bool("use-remote-sudo", false, "Prefix activation commands on the target host with `sudo`")
	rootCmd.PersistentFlags().Bool("use-substitutes", false, "Use substitue caches when running `nix-copy-closure`")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Increase the logging verbosity level")

	rootCmd.MarkFlagsMutuallyExclusive("fast", "no-build-nix") // aliases
	rootCmd.MarkFlagsMutuallyExclusive("no-build-output", "verbose", "quiet")

	rootCmd.Flag("option").Nargs = 2

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"build-host": net.ActionHosts(),
		"flake":      nix.ActionFlakes(),
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
		"target-host": net.ActionHosts(),
	})
}
