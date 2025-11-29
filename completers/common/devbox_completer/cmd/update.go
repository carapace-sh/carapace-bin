package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update [pkg]...",
	Short: "Update packages in your devbox",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(updateCmd).Standalone()

	updateCmd.Flags().Bool("all-projects", false, "update all projects in the working directory, recursively.")
	updateCmd.Flags().StringP("config", "c", "", "path to directory containing a devbox.json config file")
	updateCmd.Flags().String("environment", "", "environment to use, when supported (e.g.secrets support dev, prod, preview.)")
	updateCmd.Flags().Bool("no-install", false, "update lockfile but don't install anything")
	updateCmd.Flags().Bool("sync-lock", false, "sync all devbox.lock dependencies in multiple projects. Dependencies will sync to the latest local version.")
	rootCmd.AddCommand(updateCmd)

	// TODO environment
	carapace.Gen(updateCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionDirectories(),
	})

	// TODO positional completion
}
