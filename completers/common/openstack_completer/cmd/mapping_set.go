package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mapping_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set mapping properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mapping_setCmd).Standalone()

	mapping_setCmd.Flags().String("rules", "", "Filename that contains a new set of mapping rules")
	mapping_setCmd.Flags().String("schema-version", "", "The federated attribute mapping schema version.")
	mappingCmd.AddCommand(mapping_setCmd)
}
