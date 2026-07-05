package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var csListLvFamilyCmd = &cobra.Command{
	Use:   "listLvFamily",
	Short: "List logical volume families",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(csListLvFamilyCmd).Standalone()
	coreStorageCmd.AddCommand(csListLvFamilyCmd)
	carapace.Gen(csListLvFamilyCmd).PositionalAnyCompletion(fs.ActionBlockDevices())
}
