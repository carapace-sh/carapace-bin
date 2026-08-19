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

	rootCmd.PersistentFlags().Bool("help", false, "Show usage information")
	rootCmd.PersistentFlags().Bool("impure", false, "Allow impure builds")
	rootCmd.PersistentFlags().String("log-format", "", "Set the log format")
	rootCmd.PersistentFlags().StringSlice("option", nil, "Set a Nix configuration option")
	rootCmd.PersistentFlags().Bool("refresh", false, "Refresh the Nix cache")

	rootCmd.Flag("option").Nargs = 2

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
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
	})
}
