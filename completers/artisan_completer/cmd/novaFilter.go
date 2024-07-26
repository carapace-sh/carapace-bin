package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var novaFilterCmd = &cobra.Command{
	Use:   "nova:filter [--boolean] [--date] [--] <name>",
	Short: "Create a new filter class",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(novaFilterCmd).Standalone()

	novaFilterCmd.Flags().Bool("boolean", false, "Indicates if the generated filter should be a boolean filter")
	novaFilterCmd.Flags().Bool("date", false, "Indicates if the generated filter should be a date filter")
	rootCmd.AddCommand(novaFilterCmd)
}
