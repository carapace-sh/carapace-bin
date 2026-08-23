package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var verifychecksumsCmd = &cobra.Command{
	Use:   "verifychecksums",
	Short: "verify checksums of data in a backup",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(verifychecksumsCmd).Standalone()
	rootCmd.AddCommand(verifychecksumsCmd)

	carapace.Gen(verifychecksumsCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}