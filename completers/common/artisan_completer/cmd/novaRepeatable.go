package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var novaRepeatableCmd = &cobra.Command{
	Use:   "nova:repeatable [-m|--model MODEL] [--] <name>",
	Short: "Create a new repeatable class",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(novaRepeatableCmd).Standalone()

	novaRepeatableCmd.Flags().String("model", "", "The model class being represented.")
	rootCmd.AddCommand(novaRepeatableCmd)
}
