package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var addonsCmd = &cobra.Command{
	Use:     "addons",
	Short:   "Enable or disable a minikube addon",
	GroupID: "configuration",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addonsCmd).Standalone()
	rootCmd.AddCommand(addonsCmd)
}
