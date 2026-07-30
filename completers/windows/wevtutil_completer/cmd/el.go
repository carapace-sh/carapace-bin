package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elCmd = &cobra.Command{
	Use:   "el",
	Short: "enumerate event log names",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elCmd).Standalone()
	rootCmd.AddCommand(elCmd)
}
