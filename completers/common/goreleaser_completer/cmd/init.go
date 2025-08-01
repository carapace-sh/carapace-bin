package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/goreleaser"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:     "init",
	Short:   "Generates a .goreleaser.yaml file",
	Aliases: []string{"i"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(initCmd).Standalone()

	initCmd.Flags().StringP("config", "f", "", "Load configuration from file")
	initCmd.Flags().StringP("language", "l", "", "Which language will be used")
	rootCmd.AddCommand(initCmd)

	carapace.Gen(initCmd).FlagCompletion(carapace.ActionMap{
		"config":   carapace.ActionFiles(),
		"language": goreleaser.ActionLanguages(),
	})
}
