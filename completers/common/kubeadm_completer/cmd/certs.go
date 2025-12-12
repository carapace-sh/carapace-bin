package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var certsCmd = &cobra.Command{
	Use:     "certs",
	Short:   "Commands related to handling Kubernetes certificates",
	Aliases: []string{"certificates"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(certsCmd).Standalone()

	rootCmd.AddCommand(certsCmd)
}
