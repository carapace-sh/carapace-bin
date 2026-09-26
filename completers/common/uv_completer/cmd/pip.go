package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pipCmd = &cobra.Command{
	Use:   "pip",
	Short: "Manage Python packages with a pip-compatible interface",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pipCmd).Standalone()

	pipCmd.PersistentFlags().String("cert", "", "Path to a PEM-encoded CA certificate bundle")
	rootCmd.AddCommand(pipCmd)
	carapace.Gen(pipCmd).FlagCompletion(carapace.ActionMap{
		"cert": carapace.ActionFiles(),
	})
}
