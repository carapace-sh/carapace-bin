package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var alias_importCmd = &cobra.Command{
	Use:   "import [<filename> | -]",
	Short: "Import aliases from a YAML file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(alias_importCmd).Standalone()

	alias_importCmd.Flags().Bool("clobber", false, "Overwrite existing aliases of the same name")
	aliasCmd.AddCommand(alias_importCmd)

	carapace.Gen(alias_importCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
