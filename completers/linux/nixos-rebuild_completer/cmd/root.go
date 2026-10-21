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

	rootCmd.PersistentFlags().Bool("accept-flake-config", false, "Accept flake nixConfig")
	rootCmd.PersistentFlags().BoolP("ask-elevate-password", "S", false, "Prompt locally for the password to feed to the elevation method")
	rootCmd.PersistentFlags().Bool("ask-sudo-password", false, "Alias for --elevate=sudo --ask-elevate-password")
	rootCmd.PersistentFlags().StringP("attr", "A", "", "Use the specified attribute path from the file specified by --file")
	rootCmd.PersistentFlags().String("build-host", "", "Specify host to perform the rebuild with")
	rootCmd.PersistentFlags().StringArray("builders", nil, "Specify ad-hoc remote builders")
	rootCmd.PersistentFlags().Bool("commit-lock-file", false, "Commit changes to the flake's lock file")
	rootCmd.PersistentFlags().String("cores", "", "Set the cores setting")
	rootCmd.PersistentFlags().Bool("debug", false, "Enable debug logging")
	rootCmd.PersistentFlags().Bool("diff", false, "Show diff between current and new system closures")
	rootCmd.PersistentFlags().String("elevate", "", "Privilege-elevation method for activation commands")
	rootCmd.PersistentFlags().Bool("fallback", false, "Enable the fallback setting")
	rootCmd.PersistentFlags().Bool("fast", false, "Deprecated, use --no-reexec instead")
	rootCmd.PersistentFlags().StringP("file", "f", "", "Build the NixOS system from the specified file")
	rootCmd.PersistentFlags().StringP("flake", "F", "", "Build the NixOS system from specified flake")
	rootCmd.PersistentFlags().BoolP("help", "h", false, "Show usage information")
	rootCmd.PersistentFlags().String("image-variant", "", "Selects an image variant to build")
	rootCmd.PersistentFlags().Bool("impure", false, "Use impure nix evaluation")
	rootCmd.PersistentFlags().StringP("include", "I", "", "Add path to the list of locations used to look up <...> file names")
	rootCmd.PersistentFlags().Bool("install-bootloader", false, "Causes the bootloader to be (re)installed")
	rootCmd.PersistentFlags().Bool("install-grub", false, "Deprecated, use --install-bootloader instead")
	rootCmd.PersistentFlags().Bool("json", false, "Produce JSON output")
	rootCmd.PersistentFlags().BoolP("keep-failed", "K", false, "Enable the keep-failed setting")
	rootCmd.PersistentFlags().BoolP("keep-going", "k", false, "Enable the keep-going setting")
	rootCmd.PersistentFlags().String("log-format", "", "Set the format of log output")
	rootCmd.PersistentFlags().StringP("max-jobs", "j", "", "Set the max-jobs setting")
	rootCmd.PersistentFlags().Bool("no-build-nix", false, "Deprecated, use --no-reexec instead")
	rootCmd.PersistentFlags().BoolP("no-build-output", "Q", false, "Do not output on stdout or stderr")
	rootCmd.PersistentFlags().Bool("no-flake", false, "Do not imply --flake if /etc/nixos/flake.nix exists")
	rootCmd.PersistentFlags().Bool("no-link", false, "Do not show build output")
	rootCmd.PersistentFlags().Bool("no-net", false, "Disable substituters and consider all previously downloaded files up-to-date")
	rootCmd.PersistentFlags().Bool("no-reexec", false, "Do not update nixos-rebuild in-place before rebuilding")
	rootCmd.PersistentFlags().Bool("no-registries", false, "Don't allow lookups in the flake registries")
	rootCmd.PersistentFlags().Bool("no-ssh-tty", false, "Deprecated")
	rootCmd.PersistentFlags().Bool("no-update-lock-file", false, "Do not allow any updates to the flake's lock file")
	rootCmd.PersistentFlags().Bool("no-write-lock-file", false, "Do not write the flake's newly generated lock file")
	rootCmd.PersistentFlags().Bool("offline", false, "Disable substituters and consider all previously downloaded files up-to-date")
	rootCmd.PersistentFlags().StringSlice("option", nil, "Set the Nix configuration setting name to value")
	rootCmd.PersistentFlags().StringSlice("override-input", nil, "Override a specific flake input")
	rootCmd.PersistentFlags().BoolP("print-build-logs", "L", false, "Print full build logs on standard error")
	rootCmd.PersistentFlags().StringP("profile-name", "p", "system", "Specify Nix profile to place the new configuration in")
	rootCmd.PersistentFlags().Bool("quiet", false, "Decrease the logging verbosity level")
	rootCmd.PersistentFlags().Bool("recreate-lock-file", false, "Recreate the flake's lock file from scratch")
	rootCmd.PersistentFlags().Bool("refresh", false, "Consider all previously downloaded files out-of-date")
	rootCmd.PersistentFlags().Bool("repair", false, "Fix corrupted or missing store paths")
	rootCmd.PersistentFlags().Bool("rollback", false, "Rollback to previous configuration instead of rebuilding")
	rootCmd.PersistentFlags().Bool("show-trace", false, "Enable the show-trace setting")
	rootCmd.PersistentFlags().StringP("specialisation", "c", "", "Activates given specialisation")
	rootCmd.PersistentFlags().String("store-path", "", "Use a pre-built NixOS system store path instead of building")
	rootCmd.PersistentFlags().Bool("substitute-on-destination", false, "Use substitute caches when copying to the destination")
	rootCmd.PersistentFlags().Bool("sudo", false, "Alias for --elevate=sudo")
	rootCmd.PersistentFlags().String("target-host", "", "Specify host to apply rebuilt configuration to")
	rootCmd.PersistentFlags().StringArray("update-input", nil, "Update a specific flake input")
	rootCmd.PersistentFlags().Bool("upgrade", false, "Update the root user's channel named 'nixos' before rebuilding")
	rootCmd.PersistentFlags().Bool("upgrade-all", false, "Update all of the root user's channels before rebuilding")
	rootCmd.PersistentFlags().Bool("use-remote-sudo", false, "Deprecated, use --elevate=sudo instead")
	rootCmd.PersistentFlags().BoolP("use-substitutes", "s", false, "Use substitute caches when running nix-copy-closure")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Increase the logging verbosity level")

	rootCmd.MarkFlagsMutuallyExclusive("no-reexec", "no-build-nix") // aliases
	rootCmd.MarkFlagsMutuallyExclusive("offline", "no-net")         // aliases
	rootCmd.MarkFlagsMutuallyExclusive("no-build-output", "verbose", "quiet")

	rootCmd.Flag("option").Nargs = 2
	rootCmd.Flag("override-input").Nargs = 2

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"build-host": net.ActionHosts(),
		"elevate":    carapace.ActionValues("none", "sudo", "run0"),
		"file":       carapace.ActionFiles(".nix"),
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
		"store-path":  carapace.ActionFiles(),
		"target-host": net.ActionHosts(),
	})
}
