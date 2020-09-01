package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var scaleCmd = &cobra.Command{
	Use:   "scale",
	Short: "Set number of containers for a service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scaleCmd).Standalone()

	scaleCmd.Flags().StringP("timeout", "t", "", "Specify a shutdown timeout in seconds.")
	rootCmd.AddCommand(scaleCmd)

	carapace.Gen(scaleCmd).PositionalAnyCompletion(
		ActionServices(), // TODO Multiparts Action/ or simply add = suffix
	)
}
