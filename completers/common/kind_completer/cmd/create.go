package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates one of [cluster]",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(createCmd).Standalone()

	rootCmd.AddCommand(createCmd)
}
