package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Creates a new workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(initCmd).Standalone()

	initCmd.Flags().StringSliceP("channel", "c", nil, "Channel to use in the workspace")
	initCmd.Flags().StringSlice("conda-pypi-map", nil, "Set a mapping between conda channels and pypi channels")
	initCmd.Flags().String("format", "", "The manifest format to create")
	initCmd.Flags().StringP("import", "i", "", "Environment.yml file to bootstrap the workspace")
	initCmd.Flags().StringSliceP("platform", "p", nil, "Platforms that the workspace supports")
	initCmd.Flags().StringP("scm", "s", "", "Source Control Management used for this workspace")
	rootCmd.AddCommand(initCmd)

	carapace.Gen(initCmd).FlagCompletion(carapace.ActionMap{
		"format":   carapace.ActionValues("pixi", "pyproject", "mojoproject"),
		"import":   carapace.ActionFiles(),
		"platform": pixi.ActionPlatforms(),
		"scm":      carapace.ActionValues("github", "gitlab", "codeberg"),
	})
}
