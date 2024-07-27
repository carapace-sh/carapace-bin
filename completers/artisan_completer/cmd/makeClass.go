package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var makeClassCmd = &cobra.Command{
	Use:   "make:class [-i|--invokable] [-f|--force] [--] <name>",
	Short: "Create a new class",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(makeClassCmd).Standalone()

	makeClassCmd.Flags().Bool("force", false, "Create the class even if the class already exists")
	makeClassCmd.Flags().Bool("invokable", false, "Generate a single method, invokable class")
	rootCmd.AddCommand(makeClassCmd)
}
