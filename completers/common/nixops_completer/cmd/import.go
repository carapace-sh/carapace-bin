package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ImportCmd = &cobra.Command{
	Use:   "import",
	Short: "Import",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ImportCmd).Standalone()
	rootCmd.AddCommand(ImportCmd)
}
