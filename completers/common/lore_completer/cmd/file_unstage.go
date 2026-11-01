package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var file_unstageCmd = &cobra.Command{
	Use:   "unstage",
	Short: "Unstage changes to a file or directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(file_unstageCmd).Standalone()

	file_unstageCmd.Flags().BoolP("help", "h", false, "Print help")
	file_unstageCmd.Flags().String("targets", "", "Path to a targets file")
	fileCmd.AddCommand(file_unstageCmd)

	carapace.Gen(file_unstageCmd).FlagCompletion(carapace.ActionMap{
		"targets": carapace.ActionFiles(),
	})

	carapace.Gen(file_unstageCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
