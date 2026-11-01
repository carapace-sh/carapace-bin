package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listFeaturesCmd = &cobra.Command{
	Use:   "features",
	Short: "Print all features of the system",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listFeaturesCmd).Standalone()

	listCmd.AddCommand(listFeaturesCmd)

}
