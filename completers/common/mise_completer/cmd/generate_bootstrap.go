package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var generate_bootstrapCmd = &cobra.Command{
	Use:   "bootstrap",
	Short: "Generate bootstrap file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generate_bootstrapCmd).Standalone()

	generateCmd.AddCommand(generate_bootstrapCmd)
}
