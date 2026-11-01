package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rebuildDbCmd = &cobra.Command{
	Use:     "rebuild-db",
	Aliases: []string{"rdb"},
	Short:   "rebuild all eopkg databases",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rebuildDbCmd).Standalone()

	rebuildDbCmd.Flags().BoolP("files", "f", false, "only rebuild the files database")

	rootCmd.AddCommand(rebuildDbCmd)
}
