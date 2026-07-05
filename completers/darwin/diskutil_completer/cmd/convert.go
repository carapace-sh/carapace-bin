package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var apfsConvertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Non-destructively convert an HFS volume to APFS",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apfsConvertCmd).Standalone()
	apfsCmd.AddCommand(apfsConvertCmd)
	carapace.Gen(apfsConvertCmd).PositionalAnyCompletion(fs.ActionBlockDevices())
}
