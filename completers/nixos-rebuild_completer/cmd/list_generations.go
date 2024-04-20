package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var list_generationsCmd = &cobra.Command{
	Use:   "list-generations",
	Short: "List the available generations",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(list_generationsCmd).Standalone()

	list_generationsCmd.Flags().Bool("json", false, "Output in JSON")

	rootCmd.AddCommand(list_generationsCmd)
}
