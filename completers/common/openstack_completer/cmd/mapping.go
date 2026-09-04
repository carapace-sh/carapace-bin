package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mappingCmd = &cobra.Command{
	Use:   "mapping",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mappingCmd).Standalone()

	rootCmd.AddCommand(mappingCmd)
}
