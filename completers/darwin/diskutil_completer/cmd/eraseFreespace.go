package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var eraseFreespaceCmd = &cobra.Command{
	Use:   "eraseFreespace",
	Short: "Erase the free space on a volume",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eraseFreespaceCmd).Standalone()
	rootCmd.AddCommand(eraseFreespaceCmd)
	carapace.Gen(eraseFreespaceCmd).PositionalCompletion(fs.ActionBlockDevices())
}
