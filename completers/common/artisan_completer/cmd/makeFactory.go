package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var makeFactoryCmd = &cobra.Command{
	Use:   "make:factory [-m|--model [MODEL]] [--] <name>",
	Short: "Create a new model factory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(makeFactoryCmd).Standalone()

	makeFactoryCmd.Flags().String("model", "", "The name of the model")
	rootCmd.AddCommand(makeFactoryCmd)
}
