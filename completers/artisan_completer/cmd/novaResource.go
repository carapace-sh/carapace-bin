package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var novaResourceCmd = &cobra.Command{
	Use:   "nova:resource [-m|--model MODEL] [--] <name>",
	Short: "Create a new resource class",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(novaResourceCmd).Standalone()

	novaResourceCmd.Flags().String("model", "", "The model class being represented.")
	rootCmd.AddCommand(novaResourceCmd)
}
