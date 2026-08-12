package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var file_hashCmd = &cobra.Command{
	Use:   "hash",
	Short: "Hash a local file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(file_hashCmd).Standalone()

	file_hashCmd.Flags().BoolP("help", "h", false, "Print help")
	file_hashCmd.Flags().String("targets", "", "Path to a targets file")
	fileCmd.AddCommand(file_hashCmd)

	carapace.Gen(file_hashCmd).FlagCompletion(carapace.ActionMap{
		"targets": carapace.ActionFiles(),
	})

	carapace.Gen(file_hashCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
