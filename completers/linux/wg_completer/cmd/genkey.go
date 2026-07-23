package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var genkeyCmd = &cobra.Command{
	Use:   "genkey",
	Short: "Generates a new private key and writes it to stdout",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(genkeyCmd).Standalone()

	rootCmd.AddCommand(genkeyCmd)
}
