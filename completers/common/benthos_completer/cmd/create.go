package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new Benthos config",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(createCmd).Standalone()

	createCmd.Flags().BoolP("small", "s", false, "Print only the main components of a Benthos config (input, pipeline, output) and omit all fields marked as advanced.")
	rootCmd.AddCommand(createCmd)
}
