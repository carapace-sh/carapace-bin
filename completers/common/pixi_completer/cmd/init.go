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

	initCmd.Flags().StringP("channel", "c", "", "The channels to use")
	initCmd.Flags().String("conda-pypi-map", "", "The conda-pypi mapping to use")
	initCmd.Flags().String("format", "", "The format of the workspace to create")
	initCmd.Flags().StringP("import", "i", "", "The path to an existing conda environment file")
	initCmd.Flags().StringP("platform", "p", "", "The platforms to use")
	initCmd.Flags().Bool("pyproject-toml", false, "Create a pyproject.toml manifest")
	initCmd.Flags().StringP("scm", "s", "", "The SCM to use")
	rootCmd.AddCommand(initCmd)

	carapace.Gen(initCmd).FlagCompletion(carapace.ActionMap{
		"format":   carapace.ActionValues("pixi", "pyproject", "mojoproject"),
		"import":   carapace.ActionFiles(),
		"platform": pixi.ActionPlatforms(),
		"scm":      carapace.ActionValues("github", "gitlab", "codeberg"),
	})
}
