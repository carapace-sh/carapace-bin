package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listGenerationsCmd = &cobra.Command{
	Use:   "list-generations",
	Short: "List the available generations",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listGenerationsCmd).Standalone()

	listGenerationsCmd.Flags().Bool("json", false, "Output in JSON")
	rootCmd.AddCommand(listGenerationsCmd)
}
