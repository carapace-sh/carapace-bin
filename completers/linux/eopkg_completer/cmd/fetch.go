package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fetchCmd = &cobra.Command{
	Use:     "fetch",
	Short:   "Fetch a package",
	Aliases: []string{"fc"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fetchCmd).Standalone()

	fetchCmd.Flags().StringP("output-dir", "O", "", "Output directory for fetched packages")

	rootCmd.AddCommand(fetchCmd)
}
