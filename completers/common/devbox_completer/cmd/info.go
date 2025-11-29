package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info <pkg>",
	Short: "Display package info",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(infoCmd).Standalone()

	infoCmd.Flags().StringP("config", "c", "", "path to directory containing a devbox.json config file")
	infoCmd.Flags().String("environment", "", "environment to use, when supported (e.g.secrets support dev, prod, preview.)")
	infoCmd.Flags().Bool("markdown", false, "output in markdown format")
	rootCmd.AddCommand(infoCmd)

	// TODO environment
	carapace.Gen(infoCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionDirectories(),
	})

	// TODO positional completion
}
