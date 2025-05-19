package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var constraints_sourceCmd = &cobra.Command{
	Use:   "source",
	Short: "Print the source code for the constraints",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(constraints_sourceCmd).Standalone()

	constraints_sourceCmd.Flags().BoolP("verbose", "v", false, "Also print the fact database automatically compiled from the workspace manifests")
	constraintsCmd.AddCommand(constraints_sourceCmd)
}
