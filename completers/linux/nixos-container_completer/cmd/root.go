package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nixos-container",
	Short: "manage NixOS containers",
	Long:  "https://nixos.org/manual/nixos/stable/containers.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().Bool("commit-lock-file", false, "Commit changes to the flake's lock file")
	rootCmd.PersistentFlags().Bool("help", false, "Show usage information")
	rootCmd.PersistentFlags().Bool("impure", false, "Allow impure builds")
	rootCmd.PersistentFlags().String("log-format", "", "Set the log format")
	rootCmd.PersistentFlags().Bool("no-allow-dirty", false, "Do not allow dirty Git/Mercurial trees")
	rootCmd.PersistentFlags().Bool("no-registries", false, "Don't allow lookups in the flake registries")
	rootCmd.PersistentFlags().Bool("no-update-lock-file", false, "Do not allow any updates to the flake's lock file")
	rootCmd.PersistentFlags().Bool("no-write-lock-file", false, "Do not write the flake's newly generated lock file")
	rootCmd.PersistentFlags().StringSlice("option", nil, "Set a Nix configuration option")
	rootCmd.PersistentFlags().StringSlice("override-input", nil, "Override a specific flake input")
	rootCmd.PersistentFlags().Bool("recreate-lock-file", false, "Recreate the flake's lock file from scratch")
	rootCmd.PersistentFlags().Bool("refresh", false, "Refresh the Nix cache")
	rootCmd.PersistentFlags().StringArray("update-input", nil, "Update a specific flake input")

	rootCmd.Flag("option").Nargs = 2
	rootCmd.Flag("override-input").Nargs = 2

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
		"update-input": nix.ActionFlakeRefs(),
	})
}
