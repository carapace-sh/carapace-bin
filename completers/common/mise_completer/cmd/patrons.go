package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var patronsCmd = &cobra.Command{
	Use:   "patrons",
	Short: "Show the individuals supporting mise as Patron-tier members",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(patronsCmd).Standalone()

	patronsCmd.Flags().BoolP("help", "h", false, "Print help")
	patronsCmd.Flags().BoolP("json", "J", false, "Output in JSON format")
	patronsCmd.Flags().Bool("refresh", false, "Fetch the list now instead of using the cached copy")
	rootCmd.AddCommand(patronsCmd)
}
