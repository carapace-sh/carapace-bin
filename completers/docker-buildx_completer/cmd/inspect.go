package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Inspect current builder instance",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspectCmd).Standalone()
	inspectCmd.Flags().Bool("bootstrap", false, "Ensure builder has booted before inspecting")
	rootCmd.AddCommand(inspectCmd)

	// TODO positional completion
}
