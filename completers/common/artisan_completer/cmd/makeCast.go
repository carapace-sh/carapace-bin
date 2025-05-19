package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var makeCastCmd = &cobra.Command{
	Use:   "make:cast [-f|--force] [--inbound] [--] <name>",
	Short: "Create a new custom Eloquent cast class",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(makeCastCmd).Standalone()

	makeCastCmd.Flags().Bool("force", false, "Create the class even if the cast already exists")
	makeCastCmd.Flags().Bool("inbound", false, "Generate an inbound cast class")
	rootCmd.AddCommand(makeCastCmd)
}
