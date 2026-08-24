package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:    "import",
	Short:  "Import metadata from another format",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(importCmd).Standalone()

	importCmd.Flags().Bool("dry-run", false, "Show what would be imported without writing")
	importCmd.Flags().String("format", "", "Legacy source format: \"entire\" or \"git-ai\"")
	importCmd.Flags().BoolP("help", "h", false, "Print help")
	importCmd.Flags().String("since", "", "Only import metadata for commits on or after this date (YYYY-MM-DD)")
	importCmd.Flag("dry-run").Hidden = true
	importCmd.Flag("format").Hidden = true
	importCmd.Flag("since").Hidden = true
	rootCmd.AddCommand(importCmd)
}
