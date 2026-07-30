package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apikeyCmd = &cobra.Command{
	Use:   "apikey",
	Short: "set API key for a source",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apikeyCmd).Standalone()
	rootCmd.AddCommand(apikeyCmd)
}
