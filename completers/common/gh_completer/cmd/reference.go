package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var referenceCmd = &cobra.Command{
	Use:    "reference",
	Short:  "A comprehensive reference of all gh commands",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(referenceCmd).Standalone()

	rootCmd.AddCommand(referenceCmd)
}
