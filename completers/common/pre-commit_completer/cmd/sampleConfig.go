package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/pre-commit_completer/cmd/action"
	"github.com/spf13/cobra"
)

var sampleConfigCmd = &cobra.Command{
	Use:   "sample-config",
	Short: "Produce a sample .pre-commit-config.yaml file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sampleConfigCmd).Standalone()

	sampleConfigCmd.Flags().String("color", "", "Whether to use color in output.")

	carapace.Gen(sampleConfigCmd).FlagCompletion(carapace.ActionMap{
		"color": action.ActionColors(),
	})

	rootCmd.AddCommand(sampleConfigCmd)
}
