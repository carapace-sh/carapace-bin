package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var db_importCmd = &cobra.Command{
	Use:   "import FILE",
	Short: "import a vulnerability database archive",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(db_importCmd).Standalone()

	dbCmd.AddCommand(db_importCmd)

	carapace.Gen(db_importCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
