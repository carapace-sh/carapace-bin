package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nh",
	Short: "Yet another nix helper",
	Long:  "https://github.com/nix-community/nh",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().CountP("verbose", "v", "Increase logging verbosity")
	rootCmd.PersistentFlags().StringP("elevation-strategy", "e", "", "Choose the privilege elevation strategy")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"elevation-strategy": carapace.ActionValues("none", "auto", "passwordless"),
	})
}
