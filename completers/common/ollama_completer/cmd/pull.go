package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pullCmd = &cobra.Command{
	Use:   "pull MODEL",
	Short: "Pull a model from a registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pullCmd).Standalone()

	pullCmd.Flags().Bool("insecure", false, "Use an insecure registry")
	rootCmd.AddCommand(pullCmd)

	// TODO complete remote models
}
