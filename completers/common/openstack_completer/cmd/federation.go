package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var federationCmd = &cobra.Command{
	Use:   "federation",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(federationCmd).Standalone()

	rootCmd.AddCommand(federationCmd)
}
