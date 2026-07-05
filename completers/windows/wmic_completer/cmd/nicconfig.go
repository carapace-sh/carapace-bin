package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var nicconfigCmd = &cobra.Command{
	Use:   "nicconfig",
	Short: "network adapter configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(nicconfigCmd).Standalone()
	rootCmd.AddCommand(nicconfigCmd)
}
