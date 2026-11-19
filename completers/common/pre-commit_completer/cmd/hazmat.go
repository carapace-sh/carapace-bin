package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/pre-commit_completer/cmd/action"
	"github.com/spf13/cobra"
)

var hazmatCmd = &cobra.Command{
	Use:   "hazmat",
	Short: "Composable tools for rare use in hook `entry`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hazmatCmd).Standalone()

	hazmatCmd.Flags().String("color", "", "Whether to use color in output.")

	carapace.Gen(hazmatCmd).FlagCompletion(carapace.ActionMap{
		"color": action.ActionColors(),
	})

	rootCmd.AddCommand(hazmatCmd)
}
