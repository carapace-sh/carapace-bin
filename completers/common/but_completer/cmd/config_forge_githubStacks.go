package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var config_forge_githubStacksCmd = &cobra.Command{
	Use:   "github-stacks",
	Short: "View or configure native GitHub stacked pull requests for this repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_forge_githubStacksCmd).Standalone()

	config_forge_githubStacksCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_forgeCmd.AddCommand(config_forge_githubStacksCmd)

	carapace.Gen(config_forge_githubStacksCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"auto", "Use native stacks when the repository supports them (the default)",
			"enable", "",
			"disable", "",
		).StyleF(style.ForKeyword),
	)
}
