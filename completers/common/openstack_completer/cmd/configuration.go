package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configurationCmd = &cobra.Command{
	Use:   "configuration",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configurationCmd).Standalone()

	rootCmd.AddCommand(configurationCmd)
}
