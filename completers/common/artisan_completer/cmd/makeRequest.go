package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var makeRequestCmd = &cobra.Command{
	Use:   "make:request [-f|--force] [--] <name>",
	Short: "Create a new form request class",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(makeRequestCmd).Standalone()

	makeRequestCmd.Flags().Bool("force", false, "Create the class even if the request already exists")
	rootCmd.AddCommand(makeRequestCmd)
}
