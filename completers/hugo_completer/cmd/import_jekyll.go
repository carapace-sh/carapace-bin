package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var import_jekyllCmd = &cobra.Command{
	Use:   "jekyll",
	Short: "hugo import from Jekyll",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(import_jekyllCmd).Standalone()
	import_jekyllCmd.Flags().Bool("force", false, "allow import into non-empty target directory")
	importCmd.AddCommand(import_jekyllCmd)

	carapace.Gen(import_jekyllCmd).PositionalCompletion(
		carapace.ActionDirectories(),
		carapace.ActionDirectories(),
	)
}
