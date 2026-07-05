package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imCmd = &cobra.Command{
	Use:   "im",
	Short: "install event publisher manifest",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imCmd).Standalone()
	rootCmd.AddCommand(imCmd)
}
