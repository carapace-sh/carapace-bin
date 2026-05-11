package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Modify the Baloo configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configCmd).Standalone()

	rootCmd.AddCommand(configCmd)

	carapace.Gen(configCmd).PositionalCompletion(
		carapace.ActionValues("add", "help", "list", "ls", "remove", "rm", "set", "show").StyleF(style.ForKeyword),
	)
}

func actionConfigListParameters() carapace.Action {
	return carapace.ActionValues(
		"contentIndexing",
		"excludeFilters",
		"excludeFolders",
		"excludeMimetypes",
		"hidden",
		"includeFolders",
	).StyleF(style.ForKeyword)
}

func actionConfigModifyParameters() carapace.Action {
	return carapace.ActionValues(
		"excludeFilters",
		"excludeFolders",
		"excludeMimetypes",
		"includeFolders",
	).StyleF(style.ForKeyword)
}

func actionConfigSetParameters() carapace.Action {
	return carapace.ActionValues("contentIndexing", "hidden").StyleF(style.ForKeyword)
}

func actionBooleanValues() carapace.Action {
	return carapace.ActionValues("true", "false", "yes", "no", "y", "n", "1", "0").StyleF(style.ForKeyword)
}
