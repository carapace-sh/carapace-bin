package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var makeObserverCmd = &cobra.Command{
	Use:   "make:observer [-f|--force] [-m|--model [MODEL]] [--] <name>",
	Short: "Create a new observer class",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(makeObserverCmd).Standalone()

	makeObserverCmd.Flags().Bool("force", false, "Create the class even if the observer already exists")
	makeObserverCmd.Flags().String("model", "", "The model that the observer applies to")
	rootCmd.AddCommand(makeObserverCmd)
}
