package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qdescriptionCmd = &cobra.Command{
	Use:   "qdescription",
	Short: "query the description of a service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qdescriptionCmd).Standalone()
	rootCmd.AddCommand(qdescriptionCmd)
}
