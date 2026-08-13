package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var certsCmd = &cobra.Command{
	Use:     "certs",
	Aliases: []string{"cert"},
	Short:   "Manages your SSL certificates",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(certsCmd).Standalone()

	rootCmd.AddCommand(certsCmd)
}
