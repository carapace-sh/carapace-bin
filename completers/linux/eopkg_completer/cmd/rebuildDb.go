package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rebuildDbCmd = &cobra.Command{
	Use:     "rebuild-db",
	Short:   "Rebuild the eopkg database",
	Aliases: []string{"rdb"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rebuildDbCmd).Standalone()

	rebuildDbCmd.Flags().BoolP("files", "f", false, "Rebuild files database")

	rootCmd.AddCommand(rebuildDbCmd)
}
