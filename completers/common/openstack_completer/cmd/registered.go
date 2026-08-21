package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var registeredCmd = &cobra.Command{
	Use:   "registered",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(registeredCmd).Standalone()

	rootCmd.AddCommand(registeredCmd)
}
