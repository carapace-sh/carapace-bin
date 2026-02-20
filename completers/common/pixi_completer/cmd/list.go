package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List the packages of the current workspace",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listCmd).Standalone()

	listCmd.Flags().StringP("environment", "e", "", "The environment to list packages for. Defaults to the default environment")
	listCmd.Flags().BoolP("explicit", "x", false, "Only list packages that are explicitly defined in the workspace")
	listCmd.Flags().StringSlice("fields", nil, "Select which fields to display and in what order (comma-separated)")
	listCmd.Flags().Bool("frozen", false, "Install the environment as defined in the lockfile, doesn't update lockfile if it isn't up-to-date with the manifest file")
	listCmd.Flags().Bool("json", false, "Whether to output in json format")
	listCmd.Flags().Bool("json-pretty", false, "Whether to output in json format")
	listCmd.Flags().Bool("locked", false, "Check if lockfile is up-to-date before installing the environment, aborts when lockfile isn't up-to-date with the manifest file")
	listCmd.Flags().Bool("no-install", false, "Don't modify the environment, only modify the lock-file")
	listCmd.Flags().String("platform", "", "The platform to list packages for. Defaults to the current platform")
	listCmd.Flags().String("sort-by", "", "Sorting strategy")
	listCmd.Flag("json-pretty").Hidden = true
	rootCmd.AddCommand(listCmd)

	carapace.Gen(listCmd).FlagCompletion(carapace.ActionMap{
		"environment": pixi.ActionEnvironments(),
		"platform":    pixi.ActionPlatforms(),
		"sort-by":     carapace.ActionValues("size", "name", "kind"),
	})
}
