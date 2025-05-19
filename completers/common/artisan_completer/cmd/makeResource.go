package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var makeResourceCmd = &cobra.Command{
	Use:   "make:resource [-f|--force] [-c|--collection] [--] <name>",
	Short: "Create a new resource",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(makeResourceCmd).Standalone()

	makeResourceCmd.Flags().Bool("collection", false, "Create a resource collection")
	makeResourceCmd.Flags().Bool("force", false, "Create the class even if the resource already exists")
	rootCmd.AddCommand(makeResourceCmd)
}
