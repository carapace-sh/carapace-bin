package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var novaFieldCmd = &cobra.Command{
	Use:   "nova:field <name>",
	Short: "Create a new field",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(novaFieldCmd).Standalone()

	rootCmd.AddCommand(novaFieldCmd)
}
