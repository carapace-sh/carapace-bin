package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var pkg_cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Removes unnecessary packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pkg_cleanCmd).Standalone()
	pkgCmd.AddCommand(pkg_cleanCmd)
}
