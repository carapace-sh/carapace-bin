package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var recordCmd = &cobra.Command{
	Use:   "record",
	Short: "manage resource records",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(recordCmd).Standalone()
	rootCmd.AddCommand(recordCmd)
}
