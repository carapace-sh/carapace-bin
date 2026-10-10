package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/mise_completer/cmd/action"
	"github.com/spf13/cobra"
)

var useCmd = &cobra.Command{
	Use:     "use",
	Short:   "Install tool version and add it to config file",
	Aliases: []string{"u"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(useCmd).Standalone()

	useCmd.Flags().StringP("env", "e", "", "Modify an environment-specific config file")
	useCmd.Flags().BoolP("force", "f", false, "Force reinstall even if already installed")
	useCmd.Flags().BoolP("global", "g", false, "Use the global config file")
	useCmd.Flags().Bool("pin", false, "Pin tool version to current full version")
	useCmd.Flags().BoolP("remove", "r", false, "Remove tool from config file")
	rootCmd.AddCommand(useCmd)

	carapace.Gen(useCmd).FlagCompletion(carapace.ActionMap{
		"env": carapace.ActionValues("development", "production", "staging", "test"),
	})

	carapace.Gen(useCmd).PositionalAnyCompletion(
		action.ActionTools(),
	)
}
