package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var global_updateCmd = &cobra.Command{
	Use:   "update [pkg]...",
	Short: "Update packages in your devbox",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(global_updateCmd).Standalone()

	global_updateCmd.Flags().Bool("all-projects", false, "update all projects in the working directory, recursively.")
	global_updateCmd.Flags().StringP("config", "c", "", "path to directory containing a devbox.json config file")
	global_updateCmd.Flags().String("environment", "", "environment to use, when supported (e.g.secrets support dev, prod, preview.)")
	global_updateCmd.Flags().Bool("no-install", false, "update lockfile but don't install anything")
	global_updateCmd.Flags().Bool("sync-lock", false, "sync all devbox.lock dependencies in multiple projects. Dependencies will sync to the latest local version.")
	global_updateCmd.Flag("config").Hidden = true
	globalCmd.AddCommand(global_updateCmd)

	// TODO environment
	carapace.Gen(global_updateCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionDirectories(),
	})

	// TODO positional completion
}
