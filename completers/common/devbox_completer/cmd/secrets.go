package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var secretsCmd = &cobra.Command{
	Use:     "secrets",
	Short:   "Interact with devbox secrets in jetify cloud.",
	Aliases: []string{"envsec"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secretsCmd).Standalone()

	secretsCmd.PersistentFlags().StringP("config", "c", "", "path to directory containing a devbox.json config file")
	secretsCmd.PersistentFlags().String("environment", "", "environment to use, when supported (e.g. secrets support dev, prod, preview.)")
	rootCmd.AddCommand(secretsCmd)

	// TODO environment
	carapace.Gen(secretsCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionDirectories(),
	})
}
