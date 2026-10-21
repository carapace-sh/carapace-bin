package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nixos-install",
	Short: "install NixOS",
	Long:  "https://nixos.org/manual/nixos/stable/installation.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("attr", "A", "", "Attribute name to build from the Nix file")
	rootCmd.Flags().String("channel", "", "Path to the nixos channel to copy")
	rootCmd.Flags().String("closure", "", "Pre-built system path to use (same as --system)")
	rootCmd.Flags().Bool("commit-lock-file", false, "Commit changes to the flake's lock file")
	rootCmd.Flags().String("cores", "", "Set the cores setting")
	rootCmd.Flags().Bool("debug", false, "Enable shell debug mode")
	rootCmd.Flags().StringP("file", "f", "", "Path to Nix file")
	rootCmd.Flags().String("flake", "", "Flake URI")
	rootCmd.Flags().Bool("help", false, "Show usage information")
	rootCmd.Flags().Bool("impure", false, "Allow impure builds")
	rootCmd.Flags().StringP("include", "I", "", "Add path to the list of locations used to look up <...> file names")
	rootCmd.Flags().Bool("keep-going", false, "Keep building on failure")
	rootCmd.Flags().String("log-format", "", "Set the format of log output")
	rootCmd.Flags().StringP("max-jobs", "j", "", "Set the max-jobs setting")
	rootCmd.Flags().Bool("no-bootloader", false, "Skip bootloader installation")
	rootCmd.Flags().Bool("no-channel-copy", false, "Skip copying the nixos channel to target")
	rootCmd.Flags().Bool("no-registries", false, "Don't allow lookups in the flake registries")
	rootCmd.Flags().Bool("no-root-passwd", false, "Skip setting root password")
	rootCmd.Flags().Bool("no-root-password", false, "Skip setting root password")
	rootCmd.Flags().Bool("no-update-lock-file", false, "Do not allow any updates to the flake's lock file")
	rootCmd.Flags().Bool("no-write-lock-file", false, "Do not write the flake's newly generated lock file")
	rootCmd.Flags().StringSlice("option", nil, "Set Nix configuration option")
	rootCmd.Flags().StringSlice("override-input", nil, "Override a specific flake input")
	rootCmd.Flags().Bool("recreate-lock-file", false, "Recreate the flake's lock file from scratch")
	rootCmd.Flags().String("root", "", "Root directory to install to")
	rootCmd.Flags().Bool("show-trace", false, "Show the stack trace on evaluation errors")
	rootCmd.Flags().String("store-path", "", "Pre-built system path to use (same as --system)")
	rootCmd.Flags().String("substituters", "", "Set the substituters setting")
	rootCmd.Flags().String("system", "", "Pre-built system path to use")
	rootCmd.Flags().StringArray("update-input", nil, "Update a specific flake input")
	rootCmd.Flags().BoolP("verbose", "v", false, "Increase verbosity")

	rootCmd.Flag("option").Nargs = 2
	rootCmd.Flag("override-input").Nargs = 2

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"channel":    carapace.ActionDirectories(),
		"closure":    carapace.ActionFiles(),
		"file":       carapace.ActionFiles(".nix"),
		"include":    carapace.ActionDirectories(),
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
		"override-input": nix.ActionFlakeRefs(),
		"root":           carapace.ActionDirectories(),
		"store-path":     carapace.ActionFiles(),
		"system":         carapace.ActionFiles(),
	})
}
