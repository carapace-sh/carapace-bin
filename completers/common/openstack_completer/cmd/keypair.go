package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var keypairCmd = &cobra.Command{
	Use:   "keypair",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keypairCmd).Standalone()

	rootCmd.AddCommand(keypairCmd)
}
